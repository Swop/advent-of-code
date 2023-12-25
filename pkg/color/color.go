package color

import "fmt"

const escape = "\x1b"

const (
	None = iota
	Red
	Green
	Yellow
	Blue
	Purple
)

func color(c int) string {
	if c == None {
		return fmt.Sprintf("%s[%dm", escape, c)
	}

	return fmt.Sprintf("%s[3%dm", escape, c)
}

func Format(c int, text string) string {
	return color(c) + text + color(None)
}
