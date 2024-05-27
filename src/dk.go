package src

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os/exec"
)

type Dk struct {
	Workspaces []struct {
		Name    string
		Number  int
		Focused bool
		Active  bool
	}
}

func Dksub() {

	cmd := exec.Command("dkcmd", "status")

	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error starting Cmd: ", err)
		return
	}

	_ = cmd.Start()

	listener := bufio.NewScanner(out)

	for listener.Scan() {

		payload := listener.Text()

		var dk Dk
		err := json.Unmarshal([]byte(payload), &dk)
		if err != nil {
			panic(err)
		}

		Line = ""

		for _, status := range dk.Workspaces {

			if status.Number == NumWorkspaces {
				if status.Focused {
					Line = Line + "F" + status.Name
				} else if status.Active {
					Line = Line + "O" + status.Name
				} else {
					Line = Line + "U" + status.Name
				}
			} else {
				if status.Focused {
					Line = Line + "F" + status.Name + ":"
				} else if status.Active {
					Line = Line + "O" + status.Name + ":"
				} else {
					Line = Line + "U" + status.Name + ":"
				}
			}

		}

		Readline()
	}

}
