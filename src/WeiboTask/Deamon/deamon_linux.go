// @Title        deamon_linux
// @Description  启动守护进程的实现
// @Author       星辰
// @Update
package Deamon

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

// @title         InitDeamon
// @description   启动子进程运行
// @auth          星辰
// @param
// @return
func InitDeamon() {
	STAGE := os.Getenv("__DAEMON_STAGE")
	if STAGE == "" {
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Env = []string{"__DAEMON_STAGE=1"}
		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil
		cmd.ExtraFiles = nil
		cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
		cmd.Start()
		os.Exit(0)
	}else if STAGE == "1" {
		syscall.Setsid()
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Env = []string{"__DAEMON_STAGE=2"}
		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil
		cmd.ExtraFiles = nil
		cmd.Start()
		os.Exit(0)
	}else if STAGE == "2" {
		syscall.Umask(0)
		syscall.Chdir("/tmp")
	}
}