package utility

import (
	"awesomeProject/entity"
	"encoding/csv"
	"github.com/joho/godotenv"
	"io"
	"os"
	"strconv"
)

var filePath string

func filepath() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	filePath = os.Getenv("FILEPATH")
}

func OpenRead() ([]entity.CSVData, error) {
	filepath()
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	readTodos := []entity.CSVData{}

	reader := csv.NewReader(file)

	for {
		todos, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		flags, _ := strconv.ParseBool(todos[3])
		todo := entity.Todo{
			Id:          todos[0],
			Title:       todos[1],
			Description: todos[2],
			Terminated:  flags,
		}

		readTodos = append(readTodos, entity.CSVData{
			Todo: []entity.Todo{todo},
		})
	}

	return readTodos, nil
}

func WriteSaveFile(todos entity.CSVData) error {
	filepath()
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, todo := range todos.Todo {
		err = writer.Write(todo.Serialize())
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteUpdateFile(todos entity.CSVData) error {
	filepath()
	file, err := os.OpenFile(filePath, os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, todo := range todos.Todo {
		err = writer.Write(todo.Serialize())
		if err != nil {
			return err
		}
	}

	return nil
}

func NextID(data []entity.Todo) int {
	fileTodos, err := OpenRead()
	if err != nil {
		panic(err)
	}

	var todos []entity.Todo
	for _, todo := range fileTodos {
		todos = append(todos, todo.Todo...)
	}

	todos = append(todos, data...)

	// Find the maximum ID in the todos
	maxID := 0
	for _, todo := range todos {
		id, err := strconv.Atoi(todo.Id)
		if err == nil && id > maxID {
			maxID = id
		}
	}

	// Generate the next ID by incrementing the maximum ID by 1
	nextID := maxID + 1

	return nextID
}

/*func IsNumeric(input string) bool {
	return regexp.MustCompile(`^[0-9]+$`).MatchString(input)
}*/
