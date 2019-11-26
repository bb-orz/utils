package xos

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

/**
执行系统命令
*/
func ExecCmd(name string, params ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(name, params...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stdout.String(), stderr.String(), err
	}

	fmt.Println("execute command finish:", name, params)
	return stdout.String(), stderr.String(), err
}

/**
执行系统命令,且标准输出
*/
func ExecStdCmd(name string, params ...string) error {
	cmd := exec.Command(name, params...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("execute command finish:", name, params)
	return err
}
