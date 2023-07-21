package src

import (
	"bufio"
	"encoding/json"
	//"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
)

type HyprWorkspaces struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type HyprActive struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Hyprws() {

	cmd1 := exec.Command("hyprctl", "-j", "workspaces")
	cmd2 := exec.Command("hyprctl", "-j", "activeworkspace")
	out1, _ := cmd1.CombinedOutput()
	out2, _ := cmd2.CombinedOutput()

	s := []HyprWorkspaces{}
	err := json.Unmarshal([]byte(out1), &s)
	if err != nil {
		panic(err)
	}

	var ha HyprActive
	json.Unmarshal([]byte(out2), &ha)

	stop := NumWorkspaces + 1
	start := 1

	Line = ""

	for start < stop {

		unoccupied := true

		for _, v := range s {
			if v.Id == start {
				if ha.Id == start {
					if start == stop {
						Line = Line + "F" + ha.Name
					} else {
						Line = Line + "F" + ha.Name + ":"
					}
				} else {
					if start == stop {
						Line = Line + "O" + v.Name
					} else {
						Line = Line + "O" + v.Name + ":"
					}
				}
				unoccupied = false
			}
		}

		if unoccupied {
			t := strconv.Itoa(start)
			if start == stop {
				Line = Line + "U" + t
			} else {
				Line = Line + "U" + t + ":"
			}
		}
		start++
	}
	Line = Line
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
