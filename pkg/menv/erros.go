package menv

import "fmt"

type FileNotExists struct {
	Err error
}

func (f *FileNotExists) Error() string {
	return fmt.Sprintf("Couldn't able to locate file\n")
}

type InvalidAction struct {
	Err error
}

func (i *InvalidAction) Error() string {
	return fmt.Sprintf("Invalid action, Please check help command\n")
}

type MenvFileNotExists struct {
	Err error
}

func (err *MenvFileNotExists) Error() string {
	return fmt.Sprintf("Menvfile isn't present, please execute menv init to create.\n")
}

type ConfigNotExists struct {
	Err error
}

func (err *ConfigNotExists) Error() string {
	return fmt.Sprintf("Couldn't found secret key, please execute menv init to create.\n")

}
