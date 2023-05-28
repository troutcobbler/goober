package src

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var Line string

var Icons []string
var NumWorkspaces int
var HideUnoccupied bool

var WindowManager string

func Readline() {

	output := ""

	s := strings.Split(Line, ":")

	for i, v := range s {

		name := strconv.Itoa(i + 1)
		class := ""
		command := ""
		tag := strconv.Itoa(i)

		if WindowManager == "sway" {
			command = "swaymsg workspace"
		} else if WindowManager == "bspwm" {
			command = "bspc desktop -f"
		} else if WindowManager == "dk" {
			command = "dkcmd ws"
		} else if WindowManager == "herbstluftwm" {
			command = "herbstclient use_index"
			name = tag
		} else if WindowManager == "i3" {
			command = "i3-msg workspace"
		}

		unoccupied, err := regexp.MatchString(`^U`, v)
		focused, err := regexp.MatchString(`^F`, v)
		occupied, err := regexp.MatchString(`^O`, v)

		if err != nil {
			fmt.Println("Error parsing: ", err)
			return
		}

		if HideUnoccupied {
			if unoccupied {
				output = output + ""
			} else if focused {
				class = "active"
				output = output + " (button :class \"" + class + "\" :onclick \"" + command + " " + name + "\" \"" + Icons[i] + "\")"
			} else if occupied {
				class = "inactive"
				output = output + " (button :class \"" + class + "\" :onclick \"" + command + " " + name + "\" \"" + Icons[i] + "\")"
			}
		} else {
			if unoccupied {
				class = "inactive"
				output = output + " (button :class \"" + class + "\" :onclick \"" + command + " " + name + "\" \"" + Icons[i] + "\")"
			} else if focused {
				class = "active"
				output = output + " (button :class \"" + class + "\" :onclick \"" + command + " " + name + "\" \"" + Icons[i] + "\")"
			} else if occupied {
				class = "empty"
				output = output + " (button :class \"" + class + "\" :onclick \"" + command + " " + name + "\" \"" + Icons[i] + "\")"
			}
		}

	}

	goober := "(box :orientation \"v\" :class \"workspaces\" :space-evenly true :halign \"center\" :valign \"center\" :vexpand true " + output + ")"

	fmt.Println(goober)

}
