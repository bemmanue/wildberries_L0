package logger

import (
	"fmt"
	"time"
)

type Writer struct{}

func (w Writer) Write(bytes []byte) (int, error) {
	output := fmt.Sprintf(
		"[LOG] %s | %s",
		time.Now().Format("2006/01/02 - 15:04:05"),
		bytes,
	)
	return fmt.Print(output)
}
