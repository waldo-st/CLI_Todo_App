package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/google/uuid"
)

type task struct {
	Name       string
	Describe   string
	Done       bool
	CreatAt    time.Time
	CompletaAt time.Time
}

type Todos map[string]task

func (t *Todos) Add(name, describe string) {
	id := uuid.New()
	todo := task{
		Name:       name,
		Describe:   describe,
		Done:       false,
		CreatAt:    time.Now(),
		CompletaAt: time.Time{},
	}

	(*t)[id.String()] = todo
}

func (t *Todos) Complete(id string) error {
	list_todo := *t

	if _, exist := list_todo[id]; !exist {
		return errors.New("invalid index")
	}

	updateTask := list_todo[id] //copier la tache dans une variable local

	// modifier les champs CompletaAt et Done
	updateTask.CompletaAt = time.Now()
	updateTask.Done = true

	//et ensuite mettre a jour la tache dans le map
	list_todo[id] = updateTask
	return nil
}

func (t *Todos) Delete(id string) error {
	list_todo := *t

	if _, exist := list_todo[id]; !exist {
		return errors.New("invalid index")
	}

	delete(list_todo, id)
	return nil
}

func (t *Todos) Laod_file(filename string) error {
	// ouvre le fichier et s'il n'existe pas il le cree
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// lis le contenu du fichier une fois ouvert
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// verifier la taille du fichier
	if len(data) == 0 {
		return errors.New("le fichier est vide")
	}

	// Désérialisation du type t
	return json.Unmarshal(data, t)
}

func (t *Todos) Store_in_file(nameFile string) error {
	// serialisation du type t
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(nameFile, data, 0644)
}

func (t *Todos) Display_list() {
	num := 0

	for _, val := range *t {
		num++
		fmt.Fprintln(os.Stdout, num, "-", val.Name)
	}
}
