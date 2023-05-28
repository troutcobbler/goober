package src

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
)

func Dksub() {

	cmd := exec.Command("dkcmd", "status", "type=ws")

	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error starting Cmd: ", err)
		return
	}

	_ = cmd.Start()

	listener := bufio.NewScanner(out)

	for listener.Scan() {

		payload := listener.Text()

		// change lettering to goobers default UFO
		re := regexp.MustCompile(`i`)
		u := re.ReplaceAllString(payload, "U")

		re = regexp.MustCompile(`I`)
		f := re.ReplaceAllString(u, "F")

		re = regexp.MustCompile(`A`)
		f = re.ReplaceAllString(f, "F")

		re = regexp.MustCompile(`a`)
		o := re.ReplaceAllString(f, "O")

		Line = o
		Readline()
	}

}
