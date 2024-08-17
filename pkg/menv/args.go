package menv

import (
	"flag"
)

type ArgumentType[T any] struct {
	Argument string
	Flag     string
	Value    T
}

func (arg *ArgumentType[T]) isValid() bool {
	validArguments := [3]string{"update", "init", "generate"}
	for _, argument := range validArguments {
		if argument == arg.Argument {
			return true
		}
	}
	return false
}

func InitArgument[T string](args []string) *ArgumentType[T] {
	initArg := flag.NewFlagSet("init", flag.ExitOnError)
	initArgFile := initArg.String("f", "", "Pass environment file name")
	initArg.Parse(args)
	argumentType := ArgumentType[T]{
		Argument: "init",
		Flag:     "f",
		Value:    T(*initArgFile),
	}
	return &argumentType
}

func UpdateArgument[T string](args []string) *ArgumentType[T] {
	initArg := flag.NewFlagSet("update", flag.ExitOnError)
	initArg.Parse(args)
	argumentType := ArgumentType[T]{
		Argument: "update",
	}
	return &argumentType
}

func GenerateArgument[T bool](args []string) *ArgumentType[T] {
	initArg := flag.NewFlagSet("generate", flag.ExitOnError)
	overrideArg := initArg.Bool("y", false, "Override environment file")
	initArg.Parse(args)
	argumentType := &ArgumentType[T]{
		Argument: "generate",
		Flag:     "y",
		Value:    T(*overrideArg),
	}
	return argumentType
}
