package main

import (
	"flag"
	"fmt"
	"os"

	todo "github.com/waldo-st/CLI_Todo_App.gitcd/internal"
)

const (
	todoFile = "todos.json"
)

func main() {

	add := flag.Bool("add", false, "ajoute une nouvelle tache sur le todo")
	complete := flag.String("complete", "", "marquer la tache du todo comme fait")
	del := flag.String("del", "", "supprimer la tache du todo")
	list := flag.Bool("list", false, "liste des taches du todo")
	flag.Parse()

	todos := &todo.Todos{}

	if err := todos.Laod_file(todoFile); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}

	switch {
	case *add:
		todos.Add("sample todo plus", "apprendre a creer un todo sur la console (CLI_TODo_APP)...")
		if err := todos.Store_in_file(todoFile); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(1)
		}
	case *complete != "":
		if err := todos.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(1)
		}
		if err := todos.Store_in_file(todoFile); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(1)
		}
	case *del != "":
		if err := todos.Delete(*del); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(1)
		}
		if err := todos.Store_in_file(todoFile); err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
			os.Exit(1)
		}
	case *list:
		todos.Display_list()
	default:
		fmt.Fprintln(os.Stdout, "invalid command")
		os.Exit(0)
	}
}
