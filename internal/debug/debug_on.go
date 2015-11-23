//+build debug

package debug

import (
	"log"
	"os"
)

var logger = log.New(os.Stdout, "|DEBUG| ", 0)

func Printf(f string, args ...interface{}) {
	logger.Printf(f, args...)
}
