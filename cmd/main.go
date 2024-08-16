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
		initArg := menv.InitArgument[string](os.Args[2:])
		customInitAction := &ArgumentAction[string]{
			Action: "init",
			Flag:   initArg.Flag,
			Value:  initArg.Value,
		}
		performAction[string](customInitAction)
	case "update":
		updateArg := menv.UpdateArgument[string](os.Args[2:])
		customInitAction := &ArgumentAction[string]{
			Action: "update",
			Flag:   updateArg.Flag,
			Value:  updateArg.Value,
		}
		performAction[string](customInitAction)
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
