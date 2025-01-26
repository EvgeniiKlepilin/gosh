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
  "type": true,
}

var PATH = os.Getenv("PATH")

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
        typeCommand := arguments[0]
        _, present := BUILTINS[typeCommand]
        if present {
          fmt.Println(typeCommand + " is a shell builtin")
          break
        }
        paths := strings.Split(PATH, ":")
        isFound := false
        for _, path := range paths {
          executables, errReadDir := os.ReadDir(path)
          if errReadDir != nil {
            // skipping directories that we can't read
            continue 
          }
          for _, executable := range executables {
            if executable.Name() == typeCommand {
              fmt.Println(typeCommand + " is " + path + executable.Name())
              isFound = true
              break
            }
          }
          if isFound {
            break
          }
        }
        if !isFound {
          fmt.Println(typeCommand + ": not found")
        }
      default:
        fmt.Println(strings.TrimSuffix(input, "\n") + ": command not found")
    }
  }
}
