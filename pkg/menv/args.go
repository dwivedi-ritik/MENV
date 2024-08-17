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
	valid_arguments := [3]string{"update", "init", "generate"}
	for _, argument := range valid_arguments {
		if argument == (*arg).Argument {
			return true
		}
	}
	return false
}

func InitArgument[T string | int](args []string) *ArgumentType[T] {
	init_arg := flag.NewFlagSet("init", flag.ExitOnError)
	init_arg_file := init_arg.String("f", "", "Pass file name to be process")
	init_arg.Parse(args)
	new_argument_type := ArgumentType[T]{
		Argument: "init",
		Flag:     *init_arg_file,
	}
	return &new_argument_type
}

func UpdateArgument[T string | int](args []string) *ArgumentType[T] {
	init_arg := flag.NewFlagSet("update", flag.ExitOnError)
	init_arg_file := init_arg.String("f", "", "Pass your .env file")
	init_arg.Parse(args)
	new_argument_type := ArgumentType[T]{
		Argument: "update",
		Flag:     *init_arg_file,
	}
	return &new_argument_type
}

// deprecated
func GenerateArgument[T string | int](args []string) *ArgumentType[T] {
	init_arg := flag.NewFlagSet("generate", flag.ExitOnError)
	var init_arg_file string
	init_arg.StringVar(&init_arg_file, "f", "", "Pass your .env file that need to be menv(it search for .env file)")
	init_arg.Parse(args)
	new_argument_type := &ArgumentType[T]{
		Argument: "generate",
		Flag:     init_arg_file,
	}
	return new_argument_type
}
