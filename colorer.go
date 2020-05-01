package bashColor

type Colorer interface {
	Coder
	Writer
	FormatEventMessenger
	FormatColorMessenger
}

type Writer interface {
	Format(color string) func(...interface{}) string
	Print(f func(...interface{}) string, s string)
}

type FormatEventMessenger interface {
	Info(message ...interface{}) string
	Success(message ...interface{}) string
	Warning(message ...interface{}) string
	Fatal(message ...interface{}) string
}

type FormatColorMessenger interface {
	Black(message ...interface{}) string
	Red(message ...interface{}) string
	Green(message ...interface{}) string
	Yellow(message ...interface{}) string
	Purple(message ...interface{}) string
	Magenta(message ...interface{}) string
	Teal(message ...interface{}) string
	White(message ...interface{}) string
}
