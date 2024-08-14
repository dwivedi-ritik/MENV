package main

import (
	"fmt"
	"os"

	"github.com/dwivedi-ritik/menv/pkg/menv"
)

type ArgumentAction[T any] struct {
	Action string
	Flag   string
	Value  T
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	switch os.Args[1] {
	case "init":
		init_arg := menv.InitArgument[string](os.Args[2:])
		custom_init_action := &ArgumentAction[string]{
			Action: "init",
			Flag:   init_arg.Flag,
			Value:  init_arg.Value,
		}
		perform_action[string](custom_init_action)
	case "update":
		update_arg := menv.UpdateArgument[string](os.Args[2:])
		custom_init_action := &ArgumentAction[string]{
			Action: "init",
			Flag:   update_arg.Flag,
			Value:  update_arg.Value,
		}
		perform_action[string](custom_init_action)

	case "generate":
		generate_arg := menv.GenerateArgument[string](os.Args[2:])
		custom_init_action := &ArgumentAction[string]{
			Action: "init",
			Flag:   generate_arg.Flag,
			Value:  generate_arg.Value,
		}
		perform_action[string](custom_init_action)

	}

}

func perform_action[T any](actionArgument *ArgumentAction[T]) {
	fmt.Printf((*actionArgument).Action, (*actionArgument).Flag, (*actionArgument).Value)
}
