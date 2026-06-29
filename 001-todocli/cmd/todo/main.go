package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"todocli"

	"github.com/alexeyco/simpletable"
)

const filename = "todos.json"

func save(todos *todocli.Todos) {
	err := todos.SaveToFile(filename)
	if err != nil {
		fmt.Printf("Error saving todos: %v\n", err)
		os.Exit(1)
	}
}

func printList(todos *todocli.Todos) {

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Created At"},
			{Align: simpletable.AlignCenter, Text: "Completed At"},
		},
	}

	for i, todo := range todos.List() {
		task := blue(todo.Task)
		status := "No"
		if todo.Done {
			task = green(fmt.Sprintf("\u2705 %s", todo.Task))
			status = "Yes"
		}
		row := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", i+1)},
			{Align: simpletable.AlignLeft, Text: task},
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%v", status)},
			{Align: simpletable.AlignCenter, Text: todo.CreatedAt.Format(time.RFC822)},
			{Align: simpletable.AlignCenter, Text: todo.CompletedAt.Format(time.RFC822)},
		}
		table.SetStyle(simpletable.StyleUnicode)
		table.Body.Cells = append(table.Body.Cells, row)
	}
	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You got %d pending todo tasks", getPendingTaskCount(todos)))},
		},
	}
	table.Println()
}

func addTask(task string, todos *todocli.Todos) {
	if task == "" {
		fmt.Println("Please provide a task description")
		os.Exit(1)
	}
	todos.Add(task)
	save(todos)
}

func completeTask(complete *int, todos *todocli.Todos) {
	err := todos.Complete(*complete)
	if err != nil {
		fmt.Printf("Error completing todo: %v\n", err)
		os.Exit(1)
	}
	save(todos)
}
func deleteTask(delete *int, todos *todocli.Todos) {
	err := todos.Delete(*delete)
	if err != nil {
		fmt.Printf("Error deleting todo: %v\n", err)
		os.Exit(1)
	}
	save(todos)
}

func getInput(r io.Reader, args []string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", err
	}
	input := scanner.Text()
	if len(input) == 0 {
		return "", fmt.Errorf("task description cannot be empty")
	}
	return input, nil
}

func getPendingTaskCount(todos *todocli.Todos) int {
	count := 0
	tasks := *todos
	for _, task := range tasks {
		if task.Done == false {
			count++
		}
	}

	return count
}

func main() {

	add := flag.Bool("add", false, "Add a new todo")
	complete := flag.Int("complete", 0, "Mark a todo as completed by index")
	delete := flag.Int("delete", 0, "Delete a todo by index")
	list := flag.Bool("list", false, "List all todos")

	flag.Parse()

	todos := &todocli.Todos{}

	err := todos.LoadFromFile(filename)
	if err != nil {
		fmt.Printf("Error loading todos: %v\n", err)
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getInput(os.Stdin, flag.Args())
		if err != nil {
			fmt.Printf("Error getting input: %v\n", err)
			os.Exit(1)
		}
		addTask(task, todos)
	case *complete > 0:
		completeTask(complete, todos)
	case *delete > 0:
		deleteTask(delete, todos)
	case *list:
		printList(todos)
	default:
		fmt.Println("Please provide a valid flag: -add, -complete, -delete, or -list")
		os.Exit(1)
	}
}
