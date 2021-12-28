package tiny_proc

import (
	"errors"
	"os/exec"
)

func Proc(args []string, dir *string) error {
	if len(args) == 0 {
		return errors.New("Proc needs atleast 1 arg")
	}

	proc := exec.Command(args[0], args[1:]...)
	if dir != nil {
		cwd := *dir
		proc.Dir = cwd
	}

	out, err := proc.Output()
	println(string(out))
	return err
}
