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

const helpText = `
Usage: menv COMMAND [options]

A tool to manage your enviroment files

Commands:
  init     Initialize the menv file
  update   Update your environment file

update Options:
	-f		Name of environment file

Examples:
  menv init
  menv init -f config.json
  menv update
`

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
		break
	case "update":
		updateArg := menv.UpdateArgument[string](os.Args[2:])
		customInitAction := &ArgumentAction[string]{
			Action: "update",
			Flag:   updateArg.Flag,
			Value:  updateArg.Value,
		}
		performAction[string](customInitAction)
		break
	default:
		fmt.Printf("%v", helpText)
		break
	}
}

func performAction[T any](actionArgument *ArgumentAction[T]) error {
	if !actionArgument.validateAction() {
		return &menv.InvalidAction{}
	}

	isConfExist, err := menv.IsConfigExists()
	if err != nil {
		panic(err)
	}
	if !isConfExist {
		err = menv.InitConfig()
		if err != nil {
			panic(err)
		}
	}

	switch actionArgument.Action {
	case "update":
		err = menv.CreateEnv(actionArgument.Flag)
		if err != nil {
			panic(err)
		}
		break
	case "init":
		err = menv.CreateMenv(actionArgument.Flag)
		if err != nil {
			panic(err)
		}
		break
	default:
		break
	}
	return nil
}
