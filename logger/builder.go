package logger

import "github.com/dhanar-kusuma/go-spark/environment"

type Builder struct {
	env        environment.Type
	loggerType Type
}

func NewBuilder() *Builder {
	return &Builder{
		env:        environment.Development,
		loggerType: Default,
	}
}

func (b *Builder) SetEnv(env environment.Type) *Builder {
	b.env = env
	return b
}

func (b *Builder) SetLoggerType(loggerType Type) *Builder {
	b.loggerType = loggerType
	return b
}
