package utils

import "github.com/fatih/color"

func Error(fmt string, a ... interface{}) {
	color.Red(fmt, a...)
}
