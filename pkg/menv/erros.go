package menv

import "fmt"

type FileNotExists struct {
	Err error
}

func (f *FileNotExists) Error() string {
	return fmt.Sprintf("%v", f.Err.Error())
}
