package src

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
)

type I3workspaces struct {
	Num     int    `json:"num"`
	Name    string `json:"name"`
	Visible bool   `json:"visible"`
	Focused bool   `json:"focused"`
}

func I3ws() {

	cmd := exec.Command("i3-msg", "-t", "get_workspaces")
	out, _ := cmd.CombinedOutput()

	s := []I3workspaces{}
	err := json.Unmarshal([]byte(out), &s)
	if err != nil {
		panic(err)
	}

	stop := NumWorkspaces + 1
	start := 1

	Line = ""

	for start < stop {

		unoccupied := true

		for _, v := range s {
			if v.Num == start {
				if v.Focused {
					if start == stop {
						Line = Line + "F" + v.Name
					} else {
						Line = Line + "F" + v.Name + ":"
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

func I3sub() {

	I3ws()
	Readline()

	cmd := exec.Command("i3-msg", "-t", "subscribe", "-m", `[ "workspace" ]`)

	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error starting Cmd: ", err)
		return
	}

	_ = cmd.Start()

	listener := bufio.NewScanner(out)

	for listener.Scan() {
		I3ws()
		Readline()
	}

}
