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

    command := fields[0]
    arguments := fields[1:]

    switch command {
      case "exit":
        if len(fields) > 1 {
          code, errConv := strconv.Atoi(fields[1])
          if errConv != nil {
            fmt.Fprintln(os.Stderr, "Invalid error code", errConv)
            continue
          }
          os.Exit(code)
        } else {
          os.Exit(0)
        }  
      case "echo":
        if len(arguments) < 1 {
          fmt.Fprintln(os.Stderr, "Nothing to echo")
        } else {
          fmt.Println(strings.Join(arguments, " "))
        }
      default:
        fmt.Println(strings.TrimSuffix(input, "\n") + ": command not found")
    }
  }
}
