package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"github.com/fatih/color"
)

var todoFile string

func printList() {
  i := 1
  file, err := os.Open(todoFile)

  if err != nil {
    fmt.Println("error")
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  
  color.Set(color.FgMagenta)
  fmt.Print("")
  color.Set(color.BgMagenta, color.FgBlack, color.Bold)
  fmt.Print("      ToDo List      ")
  color.Set(color.Reset)
  color.Set(color.FgMagenta)
  fmt.Print("")
  fmt.Print("\n\n")
  color.Set(color.Reset)

  for scanner.Scan() {
    line := scanner.Text()
    color.Set(color.FgWhite, color.Faint)
    fmt.Printf("%2d  ", i)
    color.Set(color.Reset)
    if (strings.Contains(line, "X")) {
      color.Set(color.FgGreen)
      fmt.Print("  ")
      color.Set(color.Reset)
      color.Set(color.Bold, color.CrossedOut)
      fmt.Println(line[4:])
      color.Set(color.Reset)
    } else {
      fmt.Print("  ")
      color.Set(color.Bold)
      fmt.Println(line[3:])
      color.Set(color.Reset)
    }
    i++
  }
  
  if err := scanner.Err(); err != nil {
    fmt.Println("error")
  }
}

func remTask(taskId int) {
  i := 1
  file, err := os.Open(todoFile)

  if err != nil {
    fmt.Println("error")
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  newLines := make([]string, 0)
  
  for scanner.Scan() {
    line := scanner.Text()
    if (i != taskId) {
      newLines = append(newLines, line)
   }
    i++
  }
  
  if err := scanner.Err(); err != nil {
    fmt.Println("error")
  }

  file, err = os.Create(todoFile)
  if err != nil {
    fmt.Println("Error reading file", err)
    return
  } 
  defer file.Close()

  for _, line := range newLines {
    fmt.Fprintln(file, line)
  }

  printList()
}

func addTask(task string) {
  file, err := os.OpenFile(todoFile, os.O_APPEND | os.O_WRONLY, 0644)

  if err != nil {
    fmt.Println("Error opening file:", err)
  }
  defer file.Close()

  ftask := fmt.Sprintf("[] %s\n", task)

  _, err = file.WriteString(ftask)
  if err != nil {
    fmt.Println("Error wiriting to file", err)
    return
  }

 printList()
}

func togleTask(taskId int) {
  i := 1
  file, err := os.Open(todoFile)

  if err != nil {
    fmt.Println("error")
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  newLines := make([]string, 0)
  
  for scanner.Scan() {
    line := scanner.Text()
    if (i == taskId) {
      if (strings.Contains(line, "X")) {
        line = strings.Replace(line, "X", "", 1)
      } else {
        line = strings.Replace(line, "[", "[X", 1)
      }
    }
    newLines = append(newLines, line)
    i++
  }
  
  if err := scanner.Err(); err != nil {
    fmt.Println("error")
  }

  file, err = os.Create(todoFile)
  if err != nil {
    fmt.Println("Error reading file", err)
    return
  } 
  defer file.Close()

  for _, line := range newLines {
    fmt.Fprintln(file, line)
  }

  printList()
}

func openEditor() {
  editor := os.Getenv("EDITOR")

  if editor == ""{
    fmt.Println("Error opening editor", "must set EDITOR environnement variable")
    return
  }

  err := exec.Command(editor, todoFile).Run()
  
  if err != nil {
    fmt.Println("Error running the command:", err)
    return
  }
}

func todoInit()  {
  home := os.Getenv("HOME")
  configDir := filepath.Join(home, ".config", "todo")
  filePath := filepath.Join(configDir, "todo.txt")

  _, err := os.Stat(filePath)
  
  if os.IsNotExist(err) {
    err = os.MkdirAll(configDir, 0755)
    if err != nil {
        fmt.Println("Error creating directory:", err)
        return
    }

    file, err := os.Create(filePath)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }

    file.Close()
  } else if err != nil {
      fmt.Println("Error checking file:", err)
      return
  }

  todoFile = filePath
}


func main()  {
  todoInit()

  addPtr := flag.String("a", "", "add a task")
  remPtr := flag.Int("r", 0, "remove a task")
  donePtr := flag.Int("t", 0, "toggle done for a task")
  editPtr := flag.Bool("e", false, "edite todo file")

  flag.Parse()

  remTaskNum := *remPtr
  newTask := *addPtr
  togleTaskNum := *donePtr
  editFlag := *editPtr  
  
  switch {
    case remTaskNum != 0:
      remTask(remTaskNum)
    case newTask != "":
      addTask(newTask)
    case togleTaskNum != 0:
      togleTask(togleTaskNum)
    case editFlag:
      openEditor()
    default:
      printList()
  }
}

