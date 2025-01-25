package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "strconv"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
  for {
    fmt.Fprint(os.Stdout, "$ ")

	  // Wait for user input
	  var input, err = bufio.NewReader(os.Stdin).ReadString('\n')

    if err != nil {
      fmt.Fprintln(os.Stderr, "Error reading input: ", err)
      os.Exit(1)
      break
    }

    fields := strings.Fields(input)

    if fields[0] == "exit" {
      if len(fields) > 1 {
        num, errConv := strconv.Atoi(fields[1])
        if errConv != nil {
          fmt.Fprintln(os.Stderr, "Invalid error code", errConv)
          continue
        }
        os.Exit(num)
      } else {
        os.Exit(0)
      }
      break
    }

    fmt.Println(strings.TrimSuffix(input, "\n") + ": command not found")
  }
}
