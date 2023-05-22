package main

import (
	"flag"
	"fmt"

	"github.com/RenanAlmeida225/todo-cli/src/internal"
)

func main() {
	save := flag.Bool("save", false, "save task")
	complete := flag.Int("complete", 0, "complete task")
	delete := flag.Int("delete", 0, "delete an task")
	list := flag.Bool("list", false, "list all tasks")
	todos := &internal.Todos{}
	flag.Parse()
	switch {
	case *save:
		todos.Save(flag.Args()[0])
	case *complete > 0:
		todos.Complete(*complete)
	case *delete > 0:
		todos.Delete(*delete)
	case *list:
		todos.List()
	default:
		fmt.Println("Help")
	}
}
