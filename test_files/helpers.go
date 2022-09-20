package helpers

import (
	"fmt"
	"os"
)

func PipePath(username string) string {
	return fmt.Sprintf("/tmp/ezquake_fifo_%s", username)
}
func ResetPipe(username string) {
	os.Truncate(PipePath(username), 0)
}

func ReadPipe(username string) string {
	contentAsBytes, _ := os.ReadFile(PipePath(username))
	return string(contentAsBytes)
}
