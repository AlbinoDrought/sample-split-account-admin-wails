//go:build windows

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"github.com/fourcorelabs/wintoken"
)

func runWithFixedToken() {
	println("Re-launching self")
	token, err := wintoken.OpenProcessToken(0, wintoken.TokenPrimary) //pass 0 for own process
	if err != nil {
		panic(err)
	}
	defer token.Close()

	token.EnableTokenPrivileges([]string{
		"SeBackupPrivilege",
		"SeDebugPrivilege",
		"SeRestorePrivilege",
	})

	cmd := exec.Command(os.Args[0])
	cmd.Args = os.Args
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, fmt.Sprintf("%v=%v", fixedTokenKey, fixedTokenVal))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Token: syscall.Token(token.Token())}
	if err := cmd.Run(); err != nil {
		println("Error after launching self:", err)
		os.Exit(1)
	}
	println("Clean self launch :)")
	os.Exit(0)
}
