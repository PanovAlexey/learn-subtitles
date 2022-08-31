package logging

type Logger interface {
	Panic(args ...interface{})
	Error(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
}
