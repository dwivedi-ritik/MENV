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
  update   Update your Menvfile with environment file changes
  generate Generate your environment file from Menvfile


Options:
	init 	 -f	 Name of environment file
	generate -y  Yes for overridden message

Examples:
  menv init
  menv init -f config.json
  menv update
  menv generate
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
	case "generate":
		updateArg := menv.GenerateArgument[bool](os.Args[2:])
		customInitAction := &ArgumentAction[bool]{
			Action: "generate",
			Flag:   updateArg.Flag,
			Value:  updateArg.Value,
		}
		performAction[bool](customInitAction)
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

	isConfExist := menv.IsConfigExists()
	if !isConfExist {
		err := menv.InitConfig()
		if err != nil {
			fmt.Printf(err.Error())
		}
	}

	switch actionArgument.Action {
	case "init":
		err := menv.CreateMenv(any(actionArgument.Value).(string))
		if err != nil {
			fmt.Printf(err.Error())
		}
		break
	case "update":
		err := menv.UpdateMenvFile()
		if err != nil {
			fmt.Printf(err.Error())
		}
		break
	case "generate":
		err := menv.CreateEnv(any(actionArgument.Value).(bool))
		if err != nil {
			fmt.Printf(err.Error())
		}
		break
	default:
		break
	}
	return nil
}
