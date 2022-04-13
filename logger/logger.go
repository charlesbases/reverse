package logger

import (
	"fmt"
	"os"

	"github.com/charlesbases/colors"
)

// Warnf .
func Warnf(format string, a ...interface{}) {
	fmt.Println(colors.WhiteSprintf(format, a...))
}

// Infor .
func Infor(a ...interface{}) {
	fmt.Println(colors.GreenSprint(a...))
}

// Inforf .
func Inforf(format string, a ...interface{}) {
	fmt.Println(colors.GreenSprintf(format, a...))
}

// Debugf .
func Debugf(format string, a ...interface{}) {
	fmt.Println(colors.PurpleSprintf(format, a...))
}

// Errorf .
func Errorf(format string, a ...interface{}) {
	fmt.Println(colors.RedSprint(a...))
}

// Fatal .
func Fatal(a ...interface{}) {
	fmt.Println(colors.RedSprint(a...))
	os.Exit(1)
}

// Fatalf .
func Fatalf(format string, a ...interface{}) {
	fmt.Println(colors.RedSprintf(format, a...))
	os.Exit(1)
}
