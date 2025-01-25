package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
  fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	var input, err = bufio.NewReader(os.Stdin).ReadString('\n')

  if err != nil {
    fmt.Fprintln(os.Stderr, "Error reading input: ", err)
  }

  fmt.Println(strings.TrimSuffix(input, "\n") + ": command not found")
}
