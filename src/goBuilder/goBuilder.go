package main

import (
	"os/exec"
	"os"
	"syscall"
	"strings"
	"fmt"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("use: goBuilder projectPath")
		return
	}

	path := os.Args[1]

	binary, lookErr := exec.LookPath("go")
	if lookErr != nil {
		panic(lookErr)
	}

	env := os.Environ()
	isset := false
	for i, e := range env{
		if strings.Index(e, "GOPATH") == 0 {
			env[i] = "GOPATH=" + path
			isset = true
			break
		}
	}
	if !isset {
		env = append(env, "GOPATH=" + path)
	}

	args := []string{"go", "install"}

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
