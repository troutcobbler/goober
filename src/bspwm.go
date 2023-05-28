package src

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func Bspwmsub() {

	cmd := exec.Command("bspc", "subscribe", "report")

	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error starting Cmd: ", err)
		return
	}

	_ = cmd.Start()

	listener := bufio.NewScanner(out)

	for listener.Scan() {

		payload := listener.Text()

		// beginning of string
		regex := payload[:strings.Index(payload, ":")] + ":"
		s := strings.TrimPrefix(payload, regex)

		// end of string
		result := strings.LastIndex(s, ":L")

		// change lettering to goobers default UFO
		re := regexp.MustCompile(`f`)
		u := re.ReplaceAllString(s[:result], "U")

		re = regexp.MustCompile(`O`)
		f := re.ReplaceAllString(u, "F")

		re = regexp.MustCompile(`o`)
		o := re.ReplaceAllString(f, "O")

		Line = o
		Readline()
	}

}
