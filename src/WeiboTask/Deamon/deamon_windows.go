// @Title        deamon_windows
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
		/*
		os.Setenv("__DAEMON_STAGE","0")
		WinExec(strings.Join(os.Args, " "), 1) */
		filePath, _ := filepath.Abs(os.Args[0])
		cmd := exec.Command(filePath, os.Args[1:]...)
		cmd.Env = []string{"__DAEMON_STAGE=1"}
		cmd.Stdin = nil
		cmd.Stdout = nil
		cmd.Stderr = nil
		cmd.ExtraFiles = nil
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
		cmd.Start()
		os.Exit(0)
	}else if STAGE == "1" {
		dir, err := os.Executable()
		if err == nil {
			os.Chdir(filepath.Dir(dir))
		}
	}
}

/*
func WinExec(lpCmdLine string, uCmdShow uint32) uint32 {
	lpCmdLineByte, err := syscall.BytePtrFromString(lpCmdLine)
	if err != nil {
		panic(err)
	}
	lib, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		panic(err)
	}
	addr, err := syscall.GetProcAddress(lib, "WinExec")
	if err != nil {
		panic(err)
	}
	ret, _, _ := syscall.Syscall(addr,
		2,
		uintptr(unsafe.Pointer(lpCmdLineByte)),
		uintptr(uCmdShow),
		0)
	return uint32(ret)
}
*/