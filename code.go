package bashColor

// key for color ascii code default
const Default = "Default"

// key for color ascii code black
const Black = "Black"

// key for color ascii code red
const Red = "Red"

// key for color ascii code green
const Green = "Green"

// key for color ascii code yellow
const Yellow = "Yellow"

// key for color ascii code purple
const Purple = "Purple"

// key for color ascii code magenta
const Magenta = "Magenta"

// key for color ascii code teal
const Teal = "Teal"

// key for color ascii code white
const White = "White"

// ascii code default
const defaultCode = "\033[0m"

// ascii code black
const blackCode = "\033[1;30m"

// ascii code red
const redCode = "\033[1;31m"

// ascii code green
const greenCode = "\033[1;32m"

// ascii code yellow
const yellowCode = "\033[1;33m"

// ascii code purple
const purpleCode = "\033[1;34m"

// ascii code magenta
const magentaCode = "\033[1;35m"

// ascii code teal
const tealCode = "\033[1;36m"

// ascii code white
const whiteCode = "\033[1;37m"

type code struct {
	codes map[string]string
}

// Getting ascii code color
func (c *code) GetColor(name string) string {
	return c.codes[name]
}

// Set ascii colors codes
func (c *code) SetColors(codes map[string]string) coder {
	c.codes = codes
	return c
}

// Get ascii colors codes
func (c *code) GetColors() map[string]string {
	return c.codes
}

// Interface codes
type coder interface {
	SetColors(map[string]string) coder
	GetColors() map[string]string
	GetColor(string) string
}

// Constructor Codes, is created and populated with default values.
func newCode() coder {
	var c coder
	c = &code{}

	return c.SetColors(map[string]string{
		Default: defaultCode,
		Black:   blackCode,
		Red:     redCode,
		Green:   greenCode,
		Yellow:  yellowCode,
		Purple:  purpleCode,
		Magenta: magentaCode,
		Teal:    tealCode,
		White:   whiteCode,
	})
}
