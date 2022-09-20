package helpers

import (
	"fmt"
	"os"
	"time"
)

func PipePath(username string) string {
	return fmt.Sprintf("/tmp/ezquake_fifo_%s", username)
}
func ResetPipe(username string) {
	os.Truncate(PipePath(username), 0)
}

func ReadPipe(username string) string {
	time.Sleep(10 * time.Millisecond)
	contentAsBytes, _ := os.ReadFile(PipePath(username))
	return string(contentAsBytes)
}

func WriteToPipe(path string, value string) error {
	file, errOpen := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if errOpen != nil {
		return errOpen
	}

	_, errWrite := file.WriteString(value)
	return errWrite
}
