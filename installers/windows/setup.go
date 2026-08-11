//go:build ignore

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

//go:embed dry.exe
var dryBinary []byte

func main() {
	fmt.Println("=====================================")
	fmt.Println("    dryLang Setup (dryLang)     ")
	fmt.Println("=====================================")
	fmt.Println()

	userProfile, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error: Could not find user home directory: %v\n", err)
		waitForExit()
		return
	}

	installDir := filepath.Join(userProfile, ".drylang", "bin")
	binPath := filepath.Join(installDir, "dry.exe")

	fmt.Printf("[*] Creating directory: %s\n", installDir)
	err = os.MkdirAll(installDir, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		waitForExit()
		return
	}

	fmt.Println("[*] Extracting dry.exe...")
	err = os.WriteFile(binPath, dryBinary, 0755)
	if err != nil {
		fmt.Printf("Error extracting binary: %v\n", err)
		waitForExit()
		return
	}

	fmt.Println("[*] Updating Windows User PATH...")
	err = addToPath(installDir)
	if err != nil {
		fmt.Printf("Error updating PATH: %v\n", err)
		fmt.Println("You may need to add the path manually.")
	} else {
		fmt.Println("[+] PATH updated successfully.")
	}

	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("       INSTALLATION SUCCESSFUL       ")
	fmt.Println("=====================================")
	fmt.Println("Please restart your terminal/command prompt.")
	fmt.Println("Then type 'dry' to start using dryLang.")
	fmt.Println()
	
	waitForExit()
}

func addToPath(newPath string) error {
	k, err := registry.OpenKey(registry.CURRENT_USER, `Environment`, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		return err
	}
	defer k.Close()

	val, _, err := k.GetStringValue("Path")
	if err != nil && err != registry.ErrNotExist {
		return err
	}

	// Check if already exists
	paths := strings.Split(val, ";")
	for _, p := range paths {
		if strings.EqualFold(strings.TrimSpace(p), newPath) {
			return nil // Already in PATH
		}
	}

	// Append
	var updatedPath string
	if val == "" {
		updatedPath = newPath
	} else if strings.HasSuffix(val, ";") {
		updatedPath = val + newPath
	} else {
		updatedPath = val + ";" + newPath
	}

	err = k.SetStringValue("Path", updatedPath)
	if err != nil {
		return err
	}

	// Broadcast environment change to system
	broadcastEnvironmentChange()

	return nil
}

func broadcastEnvironmentChange() {
	// SendMessageTimeout to notify explorer.exe of env change
	user32 := syscall.NewLazyDLL("user32.dll")
	sendMessageTimeout := user32.NewProc("SendMessageTimeoutW")
	
	// HWND_BROADCAST = 0xFFFF
	// WM_SETTINGCHANGE = 0x001A
	// SMTO_ABORTIFHUNG = 0x0002
	
	envStr, _ := syscall.UTF16PtrFromString("Environment")
	sendMessageTimeout.Call(
		uintptr(0xFFFF),
		uintptr(0x001A),
		0,
		uintptr(unsafe.Pointer(envStr)),
		uintptr(0x0002),
		uintptr(5000), // 5 seconds timeout
		0,
	)
}

func waitForExit() {
	fmt.Print("Press Enter to exit...")
	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)
}
