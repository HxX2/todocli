package pprint

import (
	"fmt"

	"github.com/fatih/color"
)

func Print(s string, colors ...color.Attribute) {
	color.Set(colors...)
	fmt.Print(s)
	color.Set(color.Reset)
}

func Error(msg string) {
	Print("Error: ", color.FgRed, color.Bold)
	Print(msg)
}
