package src

import (
	"bufio"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

func Herbswm() {

	cmd := exec.Command("herbstclient", "tag_status")
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
		regex := payload[:strings.Index(payload, "	")] + "	"
		s := strings.TrimPrefix(payload, regex)

		// end of string
		result := strings.LastIndex(s, "	")

		// change lettering to goobers default UFO
		re := regexp.MustCompile(`\.`)
		u := re.ReplaceAllString(s[:result], "U")

		re = regexp.MustCompile(`#`)
		f := re.ReplaceAllString(u, "F")

		re = regexp.MustCompile(`:`)
		o := re.ReplaceAllString(f, "O")

		re = regexp.MustCompile(`!`)
		o = re.ReplaceAllString(o, "O")

		re = regexp.MustCompile(`	`)
		o = re.ReplaceAllString(o, ":")

		Line = o
		Readline()

	}

}

func Herbssub() {

	cmd := exec.Command("herbstclient", "-il")
	out, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error starting Cmd: ", err)
		return
	}

	_ = cmd.Start()

	listener := bufio.NewScanner(out)

	Herbswm()

	for listener.Scan() {
		Herbswm()
	}

}
