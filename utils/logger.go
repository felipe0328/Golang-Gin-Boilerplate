package utils

import "fmt"

const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorBlue   = "\033[34m"
	ColorPurple = "\033[35m"
	ColorCyan   = "\033[36m"
	ColorWhite  = "\033[37m"
)

func LogError(err error) {
	fmt.Printf("%sERROR:%s %s/n", ColorRed, ColorReset, err.Error())
}

func LogWarning(warning string) {
	fmt.Printf("%sWARNING:%s %s /n", ColorYellow, ColorReset, warning)
}

func LogDebug(message string){
	fmt.Printf("%sDEBUG:%s %s/n", ColorGreen, ColorReset, message)
}