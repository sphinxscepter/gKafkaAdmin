package zlog

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

var logConfig *Log

func LogConfiguration() {
	configFile := "conf/log.yaml"

	path, _ := os.Getwd()

	if configEnv := os.Getenv("VIPER_CONFIG"); configEnv != "" {
		configFile = configEnv
	}

	v := *viper.New()
	fmt.Println(path)
	v.AddConfigPath(path)
	v.SetConfigFile(configFile)
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		panic(fmt.Errorf("read config file error: %s\n", err))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed: %s\n", in.Name)

		if err := v.Unmarshal(&logConfig); err != nil {
			fmt.Println(err)
			panic(err)
		}
	})

	if err := v.Unmarshal(&logConfig); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(logConfig.Log.Level)
}

func init() {
	LogConfiguration()
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置日志中的日志记录时间的格式
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	// encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if logConfig.Log.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	logFile, _ := os.OpenFile(
		filepath.Join(logConfig.Log.RootDir, logConfig.Log.Filename),
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)
	fileWriteSyncer := zapcore.AddSync(logFile)

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zap.DebugLevel),
		zapcore.NewCore(encoder, fileWriteSyncer, zap.DebugLevel),
		zapcore.NewCore(encoder, getLogFileWriter(), zap.InfoLevel),
	)

	logger = zap.New(core)
}

func getLogFileWriter() (writeSyncer zapcore.WriteSyncer) {
	lumberjackLogger := &lumberjack.Logger{
		Filename:   filepath.Join(logConfig.Log.RootDir, logConfig.Log.Filename),
		MaxSize:    logConfig.Log.MaxSize,
		MaxBackups: logConfig.Log.MaxBackups,
		MaxAge:     logConfig.Log.MaxAge,
		Compress:   logConfig.Log.Compress,
	}

	return zapcore.AddSync(lumberjackLogger)
}

func getCallerInfoForLog() (callerFields []zap.Field) {
	// 回溯两层，拿到写日志的调用方的函数信息
	pc, file, line, ok := runtime.Caller(2)

	if !ok {
		return
	}

	funcName := runtime.FuncForPC(pc).Name()
	funcName = path.Base(funcName) //Base函数返回路径的最后一个元素，只保留函数名

	callerFields = append(callerFields, zap.String("file", file), zap.String("func", funcName), zap.Int("line", line))
	return
}

func Info(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Warn(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	callerFields := getCallerInfoForLog()
	fields = append(fields, callerFields...)
	logger.Error(message, fields...)
}
