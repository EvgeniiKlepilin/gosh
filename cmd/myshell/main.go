package main

import (
"bufio"
	"fmt"
	"os"
  "os/exec"
  "strings"
  "strconv"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

var BUILTINS = map[string]bool{
  "echo": true,
  "exit": true,
  "type": true,
  "pwd": true,
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
        ExitCommand(arguments)
      case "echo":
        fmt.Println(strings.Join(arguments, " "))
      case "type":
        TypeCommand(arguments)
      case "pwd":
        PwdCommand()
      default:
        paths := strings.Split(PATH, ":")
        isFound := false
        for _, path := range paths {
          executables, errReadDir := os.ReadDir(path)
          if errReadDir != nil {
            // skipping directories that we can't read
            continue
          }
          for _, executable := range executables {
            if executable.Name() == command {
              ExecutableCommand(command, arguments)
              isFound = true
              break
            }
          }
          if isFound {
            break
          }
        }
        if !isFound {
          fmt.Println(strings.TrimSuffix(input, "\n") + ": command not found")
        }
    }
  }
}

func ExitCommand(arguments []string) {
  if len(arguments) > 0 {
    code, errConv := strconv.Atoi(arguments[0])
    if errConv != nil {
      fmt.Fprintln(os.Stderr, "Invalid error code", errConv)
      return
    }
    os.Exit(code)
  } else {
    os.Exit(0)
  }  
}

func TypeCommand(arguments []string) {
  if len(arguments) != 1 {
    fmt.Fprintln(os.Stderr, "Invalid arguments")
    return 
  }
  typeCommand := arguments[0]
  _, present := BUILTINS[typeCommand]
  if present {
    fmt.Println(typeCommand + " is a shell builtin")
    return
  }
  paths := strings.Split(PATH, ":")
  for _, path := range paths {
    executables, errReadDir := os.ReadDir(path)
    if errReadDir != nil {
      // skipping directories that we can't read
     continue 
    }
    for _, executable := range executables {
      if executable.Name() == typeCommand {
        fmt.Println(typeCommand + " is " + path + "/" + executable.Name())
        return
      }
    }
  }
  fmt.Println(typeCommand + ": not found")
}

func ExecutableCommand(command string, arguments []string) {
  cmd := exec.Command(command, arguments...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func PwdCommand() {
  dir, err := os.Getwd()
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error getting working directory", err)
    return
  }
  fmt.Println(dir)
}
