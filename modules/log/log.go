package log

import (
	"fmt"
	"os"
	"time"
)

func Debug(l ...interface{}) {
	fmt.Fprintln(os.Stdout, time.Now().String(), ":", l)
}

func Info(l ...interface{}) {
	fmt.Fprintln(os.Stdout, time.Now().String(), ":", l)
}

func Fatal(l ...interface{}) {
	fmt.Fprintln(os.Stdout, time.Now().String(), ":", l)
}
