package bashColor

import "fmt"

// Color struct
type Color struct {
	coder
}

// Return interface Colorer
func NewColor() Colorer {
	var color Colorer
	color = &Color{
		newCode(),
	}
	return color
}

// Format with selected color and return string
func (c *Color) Format(color string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(color, fmt.Sprint(args...))
	}
	return sprint
}

// Format with selected color function and print string
func (c *Color) Print(f func(...interface{}) string, s string) {
	fmt.Println(f(s))
}

// Returns a string in information format
func (c *Color) Info(message ...interface{}) string {
	return c.Teal(message...)
}

// Returns a string in success format
func (c *Color) Success(message ...interface{}) string {
	return c.Green(message...)
}

// Returns a string in warning format
func (c *Color) Warning(message ...interface{}) string {
	return c.Yellow(message...)
}

// Returns a string in fatal format
func (c *Color) Fatal(message ...interface{}) string {
	return c.Red(message...)
}

// Returns a string in black format
func (c *Color) Black(message ...interface{}) string {
	return c.Format(c.GetColor(Black) + "%s" + c.GetColor(Default))(message...)
}

// Returns a string in red format
func (c *Color) Red(message ...interface{}) string {
	return c.Format(c.GetColor(Red) + "%s" + c.GetColor(Default))(message...)
}

// Returns a string in green format
func (c *Color) Green(message ...interface{}) string {
	return c.Format(c.GetColor(Green) + "%s" + c.GetColor(Default))(message...)
}

// Returns a string in yellow format
func (c *Color) Yellow(message ...interface{}) string {
	return c.Format(c.GetColor(Yellow) + "%s" + c.GetColor(Default))(message...)
}

// Returns a string in purple format
func (c *Color) Purple(message ...interface{}) string {
	return c.Format(c.GetColor(Purple) + "%s" + c.GetColor(Default))(message...)
}

// Returns a string in magenta format
func (c *Color) Magenta(message ...interface{}) string {
	return c.Format(c.GetColor(Magenta) + "%s" + c.GetColor(Default))(message...)
}

// Returns a string in teal format
func (c *Color) Teal(message ...interface{}) string {
	return c.Format(c.GetColor(Teal) + "%s" + c.GetColor(Default))(message...)
}

// Returns a string in white format
func (c *Color) White(message ...interface{}) string {
	return c.Format(c.GetColor(White) + "%s" + c.GetColor(Default))(message...)
}
