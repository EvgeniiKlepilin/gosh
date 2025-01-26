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

var BUILTINS = map[string]bool{
  "echo": true,
  "exit": true,
}

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
        fmt.Println(strings.Join(arguments, " "))
      case "type":
        if len(arguments) != 1 {
          fmt.Fprintln(os.Stderr, "Invalid arguments")
          continue
        }
        _, present := BUILTINS[arguments[0]]
        if present {
          fmt.Println(arguments[0] + " is a shell builtin")
        } else {
          fmt.Println(arguments[0] + ": not found")
        }
      default:
        fmt.Println(strings.TrimSuffix(input, "\n") + ": command not found")
    }
  }
}
