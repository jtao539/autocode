package zlog

import (
	"github.com/jtao539/autocode/template/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

var logger *zap.Logger
var SugarLogger *zap.SugaredLogger

func InitLogger() {
	encoder := getEncoder()
	writerSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writerSyncer, zapcore.DebugLevel)
	logger = zap.New(core, zap.AddCaller())
	SugarLogger = logger.Sugar()
}

// core 三个参数之  编码
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// core 三大核心之  路径
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   config.Conf.App.LogPath + string(filepath.Separator) + "card.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	if !config.Conf.App.ZapLog {
		return zapcore.AddSync(os.Stdout)
	}
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),        //日志同时输出到控制台
		zapcore.AddSync(lumberJackLogger), //配置的hook
	)
}
