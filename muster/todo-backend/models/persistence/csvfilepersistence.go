// Package persistence contains the todo backend models and corresponding business logic
package persistence

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"todo-rest-backend/models"
)

const FileName = "data.csv"

type CsvFileTodoPersistence struct {
}

// ReadTodos returns todo's stored in csv file
func (c CsvFileTodoPersistence) ReadTodos() ([]models.Todo, error) {
	return readDataFromFile()
}

func readDataFromFile() ([]models.Todo, error) {
	file, err := os.Open(FileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var readTodos = []models.Todo{}
	csvReader := csv.NewReader(file)
	for {
		records, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		readTodos = append(readTodos, parseTodoData(records))
	}

	if err != nil {
		return nil, err
	}

	return readTodos, nil
}

// ReadTodoById returns todo with passed id when existing
func (c CsvFileTodoPersistence) ReadTodoById(id string) (models.Todo, error) {
	todos, err := models.ReadTodos()
	if err != nil {
		return models.Todo{}, err
	}

	for _, todo := range todos {
		if id == todo.Id {
			return todo, nil
		}
	}

	return models.Todo{}, errors.New("id not found")
}

// CreateTodo stores the passed todo in the csv file and returns the stored todo
func (c CsvFileTodoPersistence) CreateTodo(todo models.Todo) (models.Todo, error) {
	todoCount, err := getLineCount(FileName)
	if err != nil {
		return models.Todo{}, err
	}

	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return models.Todo{}, err
	}
	defer file.Close()

	todo.Id = strconv.Itoa(todoCount)

	writer := csv.NewWriter(file)
	err = writer.Write(todo.Serialize())
	if err != nil {
		return models.Todo{}, err
	}

	writer.Flush()
	err = file.Close()
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func (c CsvFileTodoPersistence) UpdateTodo(id string, todo models.Todo) (models.Todo, error) {
	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return models.Todo{}, err
	}
	defer file.Close()

	todos, err := readDataFromFile()
	if err != nil {
		return models.Todo{}, err
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return models.Todo{}, err
	}

	todos[intID] = todo

	writer := csv.NewWriter(file)
	err = writer.Write(todo.Serialize())
	if err != nil {
		return models.Todo{}, err
	}

	writer.Flush()
	err = file.Close()
	if err != nil {
		return models.Todo{}, err
	}

	return todo, nil
}

func getLineCount(fileName string) (int, error) {
	if fileExists(fileName) == false {
		return 0, nil
	}
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	rowCount := 0
	for {
		_, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		rowCount++
	}

	return rowCount, nil
}

func fileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func parseTodoData(rec []string) models.Todo {
	id := rec[0]
	title := rec[1]
	description := rec[2]
	terminated := ToBool(rec[3])

	// Create new todo based on parsed values
	//
	todo := models.Todo{Id: id, Title: title, Description: description, Terminated: terminated}
	return todo
}

// ToBool converts a string to a boolean value
func ToBool(info string) bool {
	aBool, _ := strconv.ParseBool(info)
	return aBool
}
