package Xargs

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	CheckArgsLen()

	exec_cmd := MakeCommand()
	exec_cmd.Stdin = os.Stdin
	exec_cmd.Stdout = os.Stdout
	exec_cmd.Stderr = os.Stderr

	err := exec_cmd.Run()
	CheckError(err)
}

func MakeCommand() *exec.Cmd {
	cmd := os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		arg := strings.TrimSpace(scanner.Text())
		cmd = append(cmd, arg)
	}

	err := scanner.Err()
	CheckError(err)

	return exec.Command(cmd[0], cmd[1:]...)
}

func CheckArgsLen() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: myXargs <command>")
	}
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
