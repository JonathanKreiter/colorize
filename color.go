package colorize

import "fmt"

type Color string

const (
	black  string = "\u001b[30m"
	red    string = "\u001b[31m"
	green  string = "\u001b[32m"
	yellow string = "\u001b[33m"
	blue   string = "\u001b[34m"
	reset  string = "\u001b[0m"
	Black  Color  = Color(black)
	Red           = Color(red)
	Green         = Color(green)
	Yellow        = Color(yellow)
	Blue          = Color(blue)
	Reset         = Color(reset)
)

func Colorize(color Color, message string) string {
	return fmt.Sprint(string(color), message, reset)
}

func PrintRed(message string) {
	fmt.Println(red, message, reset)
}

func PrintGreen(message string) {
	fmt.Println(green, reset)
}

func PrintYellow(message string) {
	fmt.Println(yellow, message, reset)
}

func PrintBlue(message string) {
	fmt.Println(blue, message, reset)
}
