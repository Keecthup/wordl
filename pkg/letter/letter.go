package letter

import "fmt"

type Letter struct {
	Char  rune
	Color string
}

func NewLetter(char rune, color string) *Letter {
	return &Letter{Char: char, Color: color}
}

func (l *Letter) String() string {
	switch l.Color {
	case "yellow":
		return fmt.Sprintf("<span style=\"color: yellow;\">%s</span>", string(l.Char))
	case "green":
		return fmt.Sprintf("<span style=\"color: green;\">%s</span>", string(l.Char))
	case "gray":
		return fmt.Sprintf("<span style=\"color: gray;\">%s</span>", string(l.Char))
	default:
		return string(l.Char)
	}
}
