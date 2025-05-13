package main

import (
	"flag"

	"github.com/HxX2/todo/pkg/todo"
)

func main() {
	t := todo.Init()

	addPtr := flag.String("a", "", "add a task")
	remPtr := flag.Int("r", 0, "remove a task")
	togglePtr := flag.Int("t", 0, "toggle done for a task")
	editPtr := flag.Bool("e", false, "edit todo file")
	listDonePtr := flag.Bool("ld", false, "list done tasks")
	listUndonePtr := flag.Bool("lu", false, "list undone tasks")
	hideProgressPtr := flag.Bool("hp", false, "hide progress bar")
	initPtr := flag.Bool("i", false, "initialize todo file in git project")

	flag.Parse()

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
