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

func error(msg string)  {
  color.Set(color.FgRed, color.Bold)
  fmt.Print("Error: ")
  color.Set(color.Reset)
  println(msg)
}

func progress(nd float32, nu float32)  {
  lineWidth := float32(50)
  nt := nu + nd
  doneWidth := lineWidth * (nd / nt)
  undoneWidth := (lineWidth - doneWidth) - 1
  percent := (nd / nt) * 100

  fmt.Print("\n")
  color.Set(color.FgGreen)
  for i := 0; i < int(doneWidth); i++ {
    fmt.Print("━")
  }

  color.Set(color.FgWhite, color.Faint)
  if nu != 0 {
    fmt.Print("╺")
  }
  for i := 0; i < int(undoneWidth); i++ {
    fmt.Print("━")
  }
  color.Set(color.Reset)

  fmt.Printf(" %d%% Done\n", int(percent))
}

func printList(listDone bool, listUndone bool) {
  i := 1
  file, err := os.Open(todoFile)
  var nd float32 = 0
  var nu float32 = 0

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
    if strings.Contains(line, "[X]") && listDone {
      color.Set(color.FgWhite, color.Faint)
      fmt.Printf("%2d  ", i)
      color.Set(color.Reset)
      color.Set(color.FgGreen)
      fmt.Print("  ")
      color.Set(color.Reset)
      color.Set(color.Bold, color.CrossedOut)
      fmt.Println(line[4:])
      color.Set(color.Reset)
      nd++
    } else if strings.Contains(line, "[]") && listUndone {
      color.Set(color.FgWhite, color.Faint)
      fmt.Printf("%2d  ", i)
      color.Set(color.Reset)
      fmt.Print("  ")
      color.Set(color.Bold)
      fmt.Println(line[3:])
      color.Set(color.Reset)
      nu++
    }
    i++
  }

  if listDone && listUndone && (nu != 0 || nd != 0) {
    progress(nd, nu)
  }
  
  if err := scanner.Err(); err != nil {
    error("Can't read file")
  }
}

func remTask(taskId int) {
  i := 1
  file, err := os.Open(todoFile)

  if err != nil {
    error("Can't open file")
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  newLines := make([]string, 0)
  
  for scanner.Scan() {
    line := scanner.Text()
    if i != taskId {
      newLines = append(newLines, line)
   }
    i++
  }
  
  if err := scanner.Err(); err != nil {
    error("Can't read file")
  }

  file, err = os.Create(todoFile)
  if err != nil {
    error("Can't create file")
    return
  } 
  defer file.Close()

  for _, line := range newLines {
    fmt.Fprintln(file, line)
  }

  printList(true, true)
}

func addTask(task string) {
  file, err := os.OpenFile(todoFile, os.O_APPEND | os.O_WRONLY, 0644)

  if err != nil {
    error("Can't open file")
  }
  defer file.Close()

  ftask := fmt.Sprintf("[] %s\n", task)

  _, err = file.WriteString(ftask)
  if err != nil {
    error("Can't write in file")
    return
  }

 printList(true, true)
}

func togleTask(taskId int) {
  i := 1
  file, err := os.Open(todoFile)

  if err != nil {
    error("Can't open file")
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  newLines := make([]string, 0)
  
  for scanner.Scan() {
    line := scanner.Text()
    if i == taskId {
      if strings.Contains(line, "[X]") {
        line = strings.Replace(line, "[X", "[", 1)
      } else {
        line = strings.Replace(line, "[", "[X", 1)
      }
    }
    newLines = append(newLines, line)
    i++
  }
  
  if err := scanner.Err(); err != nil {
    error("Can't read file")
  }

  file, err = os.Create(todoFile)
  if err != nil {
    error("Can't read file")
    return
  } 
  defer file.Close()

  for _, line := range newLines {
    fmt.Fprintln(file, line)
  }

  printList(true, true)
}

func openEditor() {
  editor := os.Getenv("EDITOR")

  if editor == ""{
    error("Can't open editor [no $EDITOR env]")
    return
  }

  err := exec.Command(editor, todoFile).Run()
  
  if err != nil {
    error("Failed to open editor")
    fmt.Println(err)
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
        error("Can't create config directory")
        fmt.Println(err)
        return
    }

    file, err := os.Create(filePath)
    if err != nil {
        error("Can't create todo.txt file")
        fmt.Println(err)
        return
    }

    file.Close()
  } else if err != nil {
      error("Can't check file")
      fmt.Println(err)
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
  listDonePtr := flag.Bool("ld", false, "list done tasks")
  listUndonePtr := flag.Bool("lu", false, "list undone tasks")
 
  flag.Parse()

  remTaskNum := *remPtr
  newTask := *addPtr
  togleTaskNum := *donePtr
  editFlag := *editPtr
  listUndoneFlag := *listUndonePtr 
  listDoneFlag := *listDonePtr 
  
  switch {
    case remTaskNum != 0:
      remTask(remTaskNum)
    case newTask != "":
      addTask(newTask)
    case togleTaskNum != 0:
      togleTask(togleTaskNum)
    case editFlag:
      openEditor()
    case listDoneFlag:
      printList(true, false)
    case listUndoneFlag:
      printList(false, true)
    default:
      printList(true, true)
  }
}

