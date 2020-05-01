package bashColor

type Colorer interface {
	coder
	writer
	formatEventMessenger
	formatColorMessenger
}

type writer interface {
	// Format with selected color and return string
	Format(color string) func(...interface{}) string
	// Format with selected color function and print string
	Print(f func(...interface{}) string, s string)
}

type formatEventMessenger interface {
	// Returns a string in information format
	Info(message ...interface{}) string
	// Returns a string in success format
	Success(message ...interface{}) string
	// Returns a string in warning format
	Warning(message ...interface{}) string
	// Returns a string in fatal format
	Fatal(message ...interface{}) string
}

type formatColorMessenger interface {
	// Returns a string in black format
	Black(message ...interface{}) string
	// Returns a string in red format
	Red(message ...interface{}) string
	// Returns a string in green format
	Green(message ...interface{}) string
	// Returns a string in yellow format
	Yellow(message ...interface{}) string
	// Returns a string in purple format
	Purple(message ...interface{}) string
	// Returns a string in magenta format
	Magenta(message ...interface{}) string
	// Returns a string in teal format
	Teal(message ...interface{}) string
	// Returns a string in white format
	White(message ...interface{}) string
}
