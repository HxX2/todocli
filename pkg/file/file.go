package file

import (
	"fmt"
	"os"

	"github.com/HxX2/todo/pkg/pprint"
)

func Open(filePath string) *os.File {
	file, err := os.Open(filePath)

	if err != nil {
		pprint.Error(fmt.Sprintf("Can't open %s", filePath))
	}

	return file
}

func Write(filePath string, content string, flag int) {
	file, err := os.OpenFile(filePath, flag|os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()
	if err != nil {
		pprint.Error(fmt.Sprintf("Can't open %s", filePath))
		return
	}

	_, err = file.WriteString(content)
	if err != nil {
		pprint.Error(fmt.Sprintf("Can't write in %s", filePath))
		return
	}
}
