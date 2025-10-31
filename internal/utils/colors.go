package utils

import (
	"fmt"
	"log"
	"os"
)

const (
	ColorRed   = "\033[0;31m"
	ColorGreen = "\033[0;32m"
	ColorBlue  = "\033[0;34m"
	ColorCyan  = "\033[0;36m"
	ColorBold  = "\033[1m"
	ColorReset = "\033[0m"
)

var (
	logger = log.New(os.Stderr, "", 0)
)

func Success(f string, args ...any) { success(fmt.Sprintf(f, args...)) }
func Error(f string, args ...any)   { error(fmt.Sprintf(f, args...)) }
func Info(f string, args ...any)    { info(fmt.Sprintf(f, args...)) }

func success(msg string) { logger.Printf("%s✓ %s%s\n", ColorGreen, msg, ColorReset) }
func error(msg string)   { logger.Printf("%s✗ %s%s\n", ColorRed, msg, ColorReset) }
func info(msg string)    { logger.Printf("%sℹ %s%s\n", ColorCyan, msg, ColorReset) }
func Header(msg string) {
	logger.Printf("%s%s===== %s =====%s\n", ColorBold, ColorBlue, msg, ColorReset)
}
