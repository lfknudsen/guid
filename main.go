package main

import (
	"fmt"
	"os"
	"uuid"

	"github.com/atotto/clipboard"
)

func main() {
	generatedUuid := uuid.New().String()
	err := clipboard.WriteAll(generatedUuid)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	} else {
		_, _ = fmt.Fprintln(os.Stdout, generatedUuid)
	}
}
