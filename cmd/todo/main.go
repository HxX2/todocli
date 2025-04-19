package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/HxX2/todo/pkg/todo"
)

var (
	version string
	commit  string
	date    string
	builtBy string
)

func main() {
	addPtr := flag.String("a", "", "add a task")
	remPtr := flag.Int("r", 0, "remove a task")
	togglePtr := flag.Int("t", 0, "toggle done for a task")
	editPtr := flag.Bool("e", false, "edite todo file")
	listDonePtr := flag.Bool("ld", false, "list done tasks")
	listUndonePtr := flag.Bool("lu", false, "list undone tasks")
	hideProgressPtr := flag.Bool("hp", false, "hide progress bar")
	initPtr := flag.Bool("i", false, "initialize todo file in git project")
	versionPtr := flag.Bool("v", false, "show version")

	flag.Parse()

	versionFlag := *versionPtr
	if versionFlag {
		fmt.Printf("todocli version %s\n", version)
		fmt.Printf("commit: %s\n", commit)
		fmt.Printf("built: %s\n", date)
		fmt.Printf("by: %s\n", builtBy)
		os.Exit(0)
	}

	t := todo.Init()

	remTaskNum := *remPtr
	newTask := *addPtr
	toggleTaskNum := *togglePtr
	editFlag := *editPtr
	initFlag := *initPtr

	t.ListDone = !*listUndonePtr
	t.ListUndone = !*listDonePtr
	t.ShowProgress = !*hideProgressPtr

	switch {
	case remTaskNum != 0:
		t.RemTask(remTaskNum)
	case newTask != "":
		t.AddTask(newTask)
	case toggleTaskNum != 0:
		t.ToggleTask(toggleTaskNum)
	case editFlag:
		t.OpenEditor()
	case initFlag:
		t.ProjectInit()
	case t.ListUndone:
		t.PrintList()
	case t.ListDone:
		t.PrintList()
	default:
		t.PrintList()
	}
}
