package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/dwivedi-ritik/menv/pkg/menv"
)

type ArgumentAction[T any] struct {
	Action string
	Flag   string
	Value  T
}

func (argAction *ArgumentAction[T]) validateAction() bool {
	validCommands := map[string]bool{
		"init":     true,
		"generate": true,
		"update":   true,
	}

	return validCommands[argAction.Action]

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
		performAction[string](custom_init_action)
	case "update":
		update_arg := menv.UpdateArgument[string](os.Args[2:])
		custom_init_action := &ArgumentAction[string]{
			Action: "update",
			Flag:   update_arg.Flag,
			Value:  update_arg.Value,
		}
		performAction[string](custom_init_action)
	default:
		fmt.Println("nothing to perform")
	}
}

func performAction[T any](actionArgument *ArgumentAction[T]) error {
	if !actionArgument.validateAction() {
		return &menv.InvalidAction{}
	}
	if actionArgument.Action == "update" {
		conf_path := menv.FetchConfigPath()
		_, err := os.Stat(conf_path)

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Couldn't find secret key, generating new key")
			err = menv.InitPathConfig()
			if err != nil {
				panic(err)
			}
			fmt.Println("New key generated")
		}

		err = menv.CreateEnv(actionArgument.Flag)
		if err != nil {
			panic(err)
		}

	} else if actionArgument.Action == "init" {
		conf_path := menv.FetchConfigPath()
		_, err := os.Stat(conf_path)

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Couldn't find secret key, generating new key")
			err = menv.InitPathConfig()
			if err != nil {
				panic(err)
			}
			fmt.Println("New key generated")
		}

		err = menv.CreateMenv(actionArgument.Flag)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
