package kr

import (
	"fmt"
	"net"
	"os/exec"
	"time"
)

func DaemonDial(unixFile string) (conn net.Conn, err error) {
	conn, err = net.Dial("unix", unixFile)
	if err != nil {
		//	restart then try again
		exec.Command("killall", "krd").Run()
		exec.Command("nohup", "/usr/bin/krd", "&").Run()
		<-time.After(time.Second)
		conn, err = net.Dial("unix", unixFile)
	}
	if err != nil {
		err = fmt.Errorf("Failed to connect to Kryptonite daemon. Please make sure it is running by typing \"kr restart\".")
	}
	return
}
