package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// 初始化日志组件
func InitLogger() *zap.SugaredLogger {

	// 定义日志级别
	logMode := zapcore.DebugLevel
	if !viper.GetBool("mode.develop") {
		logMode = zapcore.InfoLevel
	}
	// 构建日志输出的核心配置对象,
	// 1.配置日志的输出格式
	// 2.配置日志以什么方式输出（可以配置多个输出端，文件/strin )
	// 3.配置日志输出的级别
	core := zapcore.NewCore(
		getEncoder(),
		zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)),
		logMode,
	)

	return zap.New(core).Sugar()
}

// 指定日志编码器（日志格式）
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		// time.DateTime 是 1.20 之后定义的常量，方便做时间格式化
		encoder.AppendString(t.Local().Format(time.DateTime))
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 指定日志保存在哪里
func getWriteSyncer() zapcore.WriteSyncer {
	sepator := string(filepath.Separator)
	rootDir, _ := os.Getwd()
	logFilePath := rootDir + sepator + "log" + sepator + time.Now().Format(time.DateOnly) + ".log"
	fmt.Println("create log file =>> " + logFilePath)

	// 日志文件分割, 需要安装第三方的依赖包 lumberjack
	lumberjackSyncer := &lumberjack.Logger{
		Filename:   logFilePath,                    // file position
		MaxSize:    viper.GetInt("log.MaxSize"),    // max file size
		MaxBackups: viper.GetInt("log.MaxBackups"), // backup pies
		MaxAge:     viper.GetInt("log.MaxAge"),     // store days
		Compress:   false,                          // 是否对文件进行压缩
	}
	return zapcore.AddSync(lumberjackSyncer)
}
