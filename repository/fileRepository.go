package repository

import (
	"fmt"
	"os"
	"strings"
)

type TextFileRepository struct {
	FilePath string
}

func NewTextFileRepository(filePath string) *TextFileRepository {
	return &TextFileRepository{FilePath: filePath}
}

func (t *TextFileRepository) Create(data string) {
	file, err := os.OpenFile(t.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	if _, err := file.WriteString(data + "\n"); err != nil {
		fmt.Println(err)
	}

}

func (t *TextFileRepository) ReadAll() []string {
	data, err := os.ReadFile(t.FilePath)

	if err != nil {
		fmt.Println(err)

		return []string{}
	}

	return strings.Split(string(data), "\n")
}

func (t *TextFileRepository) FindById(targetID int) string {
	lines := t.ReadAll()

	if targetID >= 0 && targetID < len(lines) {
		return lines[targetID]
	}

	return ""
}

func (t *TextFileRepository) Update(index int, newData string) {
	lines := t.ReadAll()

	if index >= 0 && index < len(lines) {
		lines[index] = newData + "\n"
		updatedData := []byte(strings.Join(lines, "\n"))

		if err := os.WriteFile(t.FilePath, updatedData, 0644); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Index out of range")
	}
}

func (t *TextFileRepository) Delete(index int) {
	lines := t.ReadAll()

	if index >= 0 && index < len(lines) {
		lines = append(lines[:index], lines[index+1:]...)

		updatedData := []byte(strings.Join(lines, "\n"))

		if err := os.WriteFile(t.FilePath, updatedData, 0644); err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Index out of range.")
	}
}
