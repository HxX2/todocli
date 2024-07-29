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
	editPtr := flag.Bool("e", false, "edite todo file")
	listDonePtr := flag.Bool("ld", false, "list done tasks")
	listUndonePtr := flag.Bool("lu", false, "list undone tasks")

	flag.Parse()

	remTaskNum := *remPtr
	newTask := *addPtr
	toggleTaskNum := *togglePtr
	editFlag := *editPtr

	t.ListDone = !*listUndonePtr
	t.ListUndone = !*listDonePtr

	switch {
	case remTaskNum != 0:
		t.RemTask(remTaskNum)
	case newTask != "":
		t.AddTask(newTask)
	case toggleTaskNum != 0:
		t.ToggleTask(toggleTaskNum)
	case editFlag:
		t.OpenEditor()
	case t.ListUndone:
		t.PrintList()
	case t.ListDone:
		t.PrintList()
	default:
		t.PrintList()
	}
}
