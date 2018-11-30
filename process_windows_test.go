package gowindows

import (
	"os"
	"os/exec"
	"testing"

	"golang.org/x/sys/windows"
)

func TestGetProcessParameters(t *testing.T) {
	cmd := exec.Command(`C:\WINDOWS\system32\notepad.exe`, `test.txt`)
	err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer cmd.Process.Kill()

	// time.Sleep(1 * time.Second)

	processHandle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|PROCESS_VM_READ, false, uint32(cmd.Process.Pid))
	if err != nil {
		t.Fatal(err)
	}
	defer windows.CloseHandle(processHandle)

	imagePathName, commandLine, err := GetProcessParameters(Handle(processHandle))
	if err != nil {
		t.Fatal(err)
	}

	// 32位程序@64位系统， C:\WINDOWS\system32\notepad.exe 会被系统替换为 C:\WINDOWS\SysWOW64\notepad.exe
	if imagePathName != `C:\WINDOWS\system32\notepad.exe` && imagePathName != `C:\WINDOWS\SysWOW64\notepad.exe` {
		t.Fatalf(`%v!=C:\WINDOWS\system32\notepad.exe`, imagePathName)
	}
	if commandLine != `C:\WINDOWS\system32\notepad.exe test.txt` {
		t.Fatalf(`%v!=C:\WINDOWS\system32\notepad.exe test.txt`, commandLine)
	}
}

func TestGetProcessParameters2(t *testing.T) {
	cmd := exec.Command(`C:\WINDOWS\notepad.exe`, `test.txt`)
	err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer cmd.Process.Kill()

	// time.Sleep(1 * time.Second)

	processHandle, err := windows.OpenProcess(windows.PROCESS_QUERY_INFORMATION|PROCESS_VM_READ, false, uint32(cmd.Process.Pid))
	if err != nil {
		t.Fatal(err)
	}
	defer windows.CloseHandle(processHandle)

	imagePathName, commandLine, err := GetProcessParameters(Handle(processHandle))
	if err != nil {
		t.Fatal(err)
	}

	if imagePathName != `C:\WINDOWS\notepad.exe` {
		t.Fatalf(`%v!=C:\WINDOWS\notepad.exe`, imagePathName)
	}
	if commandLine != `C:\WINDOWS\notepad.exe test.txt` {
		t.Fatalf(`%v!=C:\WINDOWS\notepad.exe test.txt`, commandLine)
	}
}

func TestGetProcessParametersWid(t *testing.T) {
	cmd := exec.Command(`C:\WINDOWS\notepad.exe`, `test.txt`)
	err := cmd.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer cmd.Process.Kill()

	// time.Sleep(1 * time.Second)

	imagePathName, commandLine, err := GetProcessParametersWPid(uint32(cmd.Process.Pid))
	if err != nil {
		t.Fatal(err)
	}

	if imagePathName != `C:\WINDOWS\notepad.exe` {
		t.Fatalf(`%v!=C:\WINDOWS\notepad.exe`, imagePathName)
	}
	if commandLine != `C:\WINDOWS\notepad.exe test.txt` {
		t.Fatalf(`%v!=C:\WINDOWS\notepad.exe test.txt`, commandLine)
	}
}

func TestMyCreateProcess(t *testing.T) {

	info, err := MyCreateProcess(`C:\WINDOWS\system32\notepad.exe`, false, `d:\\aaa.txt`)
	if err != nil {
		t.Fatal(err)
	}

	if info.ProcessId == 0 {
		t.Fatal(err)
	}

	p, err := os.FindProcess(int(info.ProcessId))
	if err != nil {
		t.Fatal(err)
	}

	err = p.Kill()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMyCreateProcess2(t *testing.T) {

	info, err := MyCreateProcess(`C:\WINDOWS\system32\notepad.exe`, true, `d:\\aaa.txt`)
	if err != nil {
		t.Fatal(err)
	}

	if info.ProcessId == 0 {
		t.Fatal(err)
	}

	p, err := os.FindProcess(int(info.ProcessId))
	if err != nil {
		t.Fatal(err)
	}

	err = p.Kill()
	if err != nil {
		t.Fatal(err)
	}
}
