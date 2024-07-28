package main

import (
	"flag"

	"github.com/HxX2/todo/pkg/todo"
)

func main() {
	t := todo.Init()

	// addPtr := flag.String("a", "", "add a task")
	remPtr := flag.Int("r", 0, "remove a task")
	// donePtr := flag.Int("t", 0, "toggle done for a task")
	// editPtr := flag.Bool("e", false, "edite todo file")
	listDonePtr := flag.Bool("ld", false, "list done tasks")
	listUndonePtr := flag.Bool("lu", false, "list undone tasks")

	flag.Parse()

	remTaskNum := *remPtr
	// newTask := *addPtr
	// togleTaskNum := *donePtr
	// editFlag := *editPtr
	listUndoneFlag := *listUndonePtr
	listDoneFlag := *listDonePtr

	switch {
	case remTaskNum != 0:
		t.RemTask(remTaskNum)
	// case newTask != "":
	//   addTask(newTask)
	// case togleTaskNum != 0:
	//   togleTask(togleTaskNum)
	// case editFlag:
	// openEditor()
	case listDoneFlag:
		t.ListUndone = false
		t.PrintList()
	case listUndoneFlag:
		t.ListDone = false
		t.PrintList()
	default:
		t.PrintList()
	}
}
