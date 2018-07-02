package zapplug

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"go.uber.org/zap"
)

type CutConf struct {
	FileName string
	MaxSize int
	MaxBackups int
	MaxAge int
	Compress bool
}

func ZapCut(c CutConf,z zapcore.EncoderConfig,l zapcore.Level) zapcore.Core {
	lum := &lumberjack.Logger{Filename:c.FileName,MaxSize:c.MaxSize,MaxBackups:c.MaxBackups,MaxAge:c.MaxAge,Compress:c.Compress}
	zlvl := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= l
	})
	code := zapcore.NewJSONEncoder(z)
	return zapcore.NewCore(code,zapcore.AddSync(lum),zlvl)
}
