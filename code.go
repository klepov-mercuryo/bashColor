package bashColor

const Default = "Default"
const Black = "Black"
const Red = "Red"
const Green = "Green"
const Yellow = "Yellow"
const Purple = "Purple"
const Magenta = "Magenta"
const Teal = "Teal"
const White = "White"

const DefaultCode = "\033[0m"
const BlackCode = "\033[1;30m"
const RedCode = "\033[1;31m"
const GreenCode = "\033[1;32m"
const YellowCode = "\033[1;33m"
const PurpleCode = "\033[1;34m"
const MagentaCode = "\033[1;35m"
const TealCode = "\033[1;36m"
const WhiteCode = "\033[1;37m"

type Code struct {
	codes map[string]string
}

func (c *Code) GetColor(name string) string {
	return c.codes[name]
}
func (c *Code) SetColors(codes map[string]string) Coder {
	c.codes = codes
	return c
}

type Coder interface {
	SetColors(map[string]string) Coder
	GetColor(string) string
}

func NewCode() Coder {
	var code Coder
	code = &Code{}

	return code.SetColors(map[string]string{
		Default: DefaultCode,
		Black:   BlackCode,
		Red:     RedCode,
		Green:   GreenCode,
		Yellow:  YellowCode,
		Purple:  PurpleCode,
		Magenta: MagentaCode,
		Teal:    TealCode,
		White:   WhiteCode,
	})
}
