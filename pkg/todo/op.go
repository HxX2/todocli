package todo

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"

	"github.com/HxX2/todo/pkg/file"
	"github.com/HxX2/todo/pkg/pprint"
)

func (t Todo) PrintList() {
	file := file.Open(t.filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	pprint.Print("\n", color.FgMagenta)
	pprint.Print("", color.FgMagenta)
	pprint.Print("      ToDo List      ", color.BgMagenta, color.FgBlack, color.Bold)
	pprint.Print("\n\n", color.FgMagenta)

	i := 1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[X]") && t.ListDone {
			pprint.Print(fmt.Sprintf("%2d  ", i), color.FgWhite, color.Faint)
			pprint.Print("  ", color.FgGreen)
			pprint.Print(fmt.Sprintf("%s\n", line[4:]), color.Bold, color.CrossedOut)
			t.doneCount++
		} else if strings.Contains(line, "[]") && t.ListUndone {
			pprint.Print(fmt.Sprintf("%2d  ", i), color.FgWhite, color.Faint)
			pprint.Print("  ")
			pprint.Print(fmt.Sprintf("%s\n", line[3:]), color.Bold)
			t.undoneCount++
		}
		i++
	}

	t.printProgress()

	if err := scanner.Err(); err != nil {
		pprint.Error("Can't read file")
	}
}

func (t Todo) RemTask(taskId int) {
	newLines := make([]string, 0)
	todoFile := file.Open(t.filePath)
	defer todoFile.Close()

	i := 1
	scanner := bufio.NewScanner(todoFile)
	for scanner.Scan() {
		line := scanner.Text()
		if i != taskId {
			newLines = append(newLines, line)
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		pprint.Error(fmt.Sprintf("Can't read %s", t.filePath))
	}

	file.Write(t.filePath, strings.Join(newLines, "\n"), os.O_TRUNC)

	t.PrintList()
}

func (t Todo) AddTask(task string) {
	ftask := fmt.Sprintf("\n[] %s", task)

	file.Write(t.filePath, ftask, os.O_APPEND)

	t.PrintList()
}

func (t Todo) ToggleTask(taskId int) {
	newLines := make([]string, 0)
	todoFile := file.Open(t.filePath)
	defer todoFile.Close()

	i := 1
	scanner := bufio.NewScanner(todoFile)
	for scanner.Scan() {
		line := scanner.Text()
		if i == taskId {
			if strings.Contains(line, "[X]") {
				line = strings.Replace(line, "[X", "[", 1)
			} else if strings.Contains(line, "[]") {
				line = strings.Replace(line, "[", "[X", 1)
			}
		}
		newLines = append(newLines, line)
		i++
	}

	if err := scanner.Err(); err != nil {
		pprint.Error(fmt.Sprintf("Can't read %s", t.filePath))
	}

	file.Write(t.filePath, strings.Join(newLines, "\n"), os.O_TRUNC)

	t.PrintList()
}

func (t Todo) OpenEditor() {
	editor := os.Getenv("EDITOR")

	if editor == "" {
		pprint.Error("Can't open editor [no $EDITOR env]")
		return
	}
	// Set up the command to use standard input/output/error
	cmd := exec.Command(editor, t.filePath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Run the command
	err := cmd.Run()
	if err != nil {
		pprint.Error(fmt.Sprintf("Failed to open editor\n%s\n", err))
		return
	}
}

func (t Todo) printProgress() {
	if !t.ShowProgress || !t.ListDone || !t.ListUndone || (t.doneCount == 0 && t.undoneCount == 0) {
		return
	}

	lineWidth := 50.0
	total := t.doneCount + t.undoneCount
	doneWidth := lineWidth * (t.doneCount / total)
	undoneWidth := (lineWidth - doneWidth) - 1
	percentage := (t.doneCount / total) * 100

	fmt.Print("\n")
	for i := 0; i < int(doneWidth); i++ {
		pprint.Print("━", color.FgGreen)
	}

	if t.undoneCount != 0 {
		pprint.Print("╺", color.FgWhite, color.Faint)
	}

	for i := 0; i < int(undoneWidth); i++ {
		pprint.Print("━", color.FgWhite, color.Faint)
	}

	fmt.Printf(" %d%% Done\n", int(percentage))
}
