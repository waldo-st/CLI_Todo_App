package main

import (
	"bufio"
	"fmt"
	"os"

	todo "github.com/waldo-st/CLI_Todo_App.git/internal"
)

const (
	todoFile = "todos.json"
)

var Read = bufio.NewReader(os.Stdin)

func main() {
	todos := &todo.Todos{}
	id := 0

	if err := todos.Laod_file(todoFile); err != nil {
		fmt.Fprintln(os.Stdout, err.Error())
	}

	for {
		fmt.Fprintln(os.Stdout, "\n\x1b[1m\x1b[34m📑 MANUEL D'UTILISATION DU TODO:\x1b[0m\n\n 👉 \x1b[1m\x1b[34mAdd\x1b[0m: \x1b[32mAjoute une nouvelle tache sur le todo\x1b[0m\n\n 👉 \x1b[1m\x1b[34mComplete\x1b[0m: \x1b[32mMarquer la tache du todo comme fait\x1b[0m\n\n 👉 \x1b[1m\x1b[34mDelete\x1b[0m: \x1b[32mSupprimer la tache du todo\x1b[0m\n\n 👉 \x1b[1m\x1b[34mListe\x1b[0m: \x1b[32mListe des taches du todo\x1b[0m\x1b[0m")

		saisie, err := todo.Collect_input("Entrer une commande: ")
		if err != nil {
			fmt.Fprintln(os.Stdout, err.Error())
		}
		switch saisie {
		case "Add":
			id++
			if err := todos.Add(id); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				os.Exit(1)
			}
			if err := todos.Store_in_file(todoFile); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				os.Exit(1)
			}
		case "Complete":
			id, err := todo.Collect_input("Entrer l'id de la tache exécutée: ")
			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				os.Exit(1)
			}
			if err = todos.Complete(id); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				os.Exit(1)
			}
			if err := todos.Store_in_file(todoFile); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				os.Exit(1)
			}
		case "Delete":
			id, err := todo.Collect_input("Entrer l'id de la tache: ")
			if err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				os.Exit(1)
			}
			if err = todos.Delete(id); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				os.Exit(1)
			}
			if err := todos.Store_in_file(todoFile); err != nil {
				fmt.Fprintln(os.Stdout, err.Error())
				os.Exit(1)
			}
		}
	}

	// add := flag.Bool("add", false, "ajoute une nouvelle tache sur le todo")
	// complete := flag.String("complete", "", "marquer la tache du todo comme fait")
	// del := flag.String("del", "", "supprimer la tache du todo")
	// list := flag.Bool("list", false, "liste des taches du todo")
	// flag.Parse()

	// todos := &todo.Todos{}

	// if err := todos.Laod_file(todoFile); err != nil {
	// 	fmt.Fprintln(os.Stdout, err.Error())
	// }

	// switch {
	// case *add:
	// 	todos.Add("sample todo plus", "apprendre a creer un todo sur la console (CLI_TODo_APP)...")
	// 	if err := todos.Store_in_file(todoFile); err != nil {
	// 		fmt.Fprintln(os.Stdout, err.Error())
	// 		os.Exit(1)
	// 	}
	// case *complete != "":
	// 	if err := todos.Complete(*complete); err != nil {
	// 		fmt.Fprintln(os.Stdout, err.Error())
	// 		os.Exit(1)
	// 	}
	// 	if err := todos.Store_in_file(todoFile); err != nil {
	// 		fmt.Fprintln(os.Stdout, err.Error())
	// 		os.Exit(1)
	// 	}
	// case *del != "":
	// 	if err := todos.Delete(*del); err != nil {
	// 		fmt.Fprintln(os.Stdout, err.Error())
	// 		os.Exit(1)
	// 	}
	// 	if err := todos.Store_in_file(todoFile); err != nil {
	// 		fmt.Fprintln(os.Stdout, err.Error())
	// 		os.Exit(1)
	// 	}
	// case *list:
	// 	todos.Display_list()
	// default:
	// 	fmt.Fprintln(os.Stdout, "invalid command")
	// 	os.Exit(0)
	// }
}
