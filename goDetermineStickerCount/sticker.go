package main

import (
	"fmt"
	"strings"
)

// simple test cases as proof of concept.

func stickerCount(input string) int {
	sticker := "wpengine"
	if input == sticker {
		return 1
	}
	for _, char := range input {
		if !strings.Contains(sticker, string(char)) {
			return -1
		}
	}
	stickerCount := 0
	inputLength := len(input)

	for inputLength > 0 {
		for _, char := range sticker {
			if strings.Contains(input, string(char)) {
				inputLength--
			}
		}
		stickerCount++
	}

	return stickerCount

}

func main() {
	fmt.Println(stickerCount("winning"))
	fmt.Println(stickerCount("wpengine"))
	fmt.Println(stickerCount("zebra"))
}
