package src

import (
	"bufio"
	//	"encoding/json"
	"fmt"
	//	"os/exec"
	//	"strconv"
	"net"
	"os"
)

//type Workspaces struct {
//	Num     int    `json:"num"`
//	Name    string `json:"name"`
//	Visible bool   `json:"visible"`
//	Focused bool   `json:"focused"`
//}

func Hyprws() {
	fmt.Println("hyprland")

}

func Hyprsub() {

	Hyprws()
	Readline()

	conn, _ := net.Dial("unix", "/tmp/hypr/"+os.Getenv("HYPRLAND_INSTANCE_SIGNATURE")+"/.socket2.sock")

	scanner := bufio.NewScanner(conn)
	for {
		if ok := scanner.Scan(); !ok {
			break
		}
		Hyprws()
		Readline()
	}

}
