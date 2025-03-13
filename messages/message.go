package messages

import (
	"fmt"
	"time"
)

func DisplayCustomMessage(author, message string) string {
	pubDate := time.Now()
	return fmt.Sprintf("%v : %v at %v ", author, message, pubDate)
}
