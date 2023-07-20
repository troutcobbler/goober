package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"goober/src"
	"os"
)

type Config struct {
	Icons          []string
	Workspaces     int
	HideUnoccupied bool
}

func main() {

	configfile, _ := os.Open(os.Getenv("HOME") + "/.config/goober/goober.conf")
	defer configfile.Close()
	decoder := json.NewDecoder(configfile)
	config := Config{}
	err := decoder.Decode(&config)

	if err != nil {
		fmt.Println("error:", err)
	}

	src.Icons = config.Icons
	src.NumWorkspaces = config.Workspaces
	src.HideUnoccupied = config.HideUnoccupied

	wm := flag.String("wm", "wm", "wm")

	flag.Parse()

	if *wm == "sway" {
		src.WindowManager = *wm
		src.Swaysub()
	} else if *wm == "bspwm" {
		src.WindowManager = *wm
		src.Bspwmsub()
	} else if *wm == "dk" {
		src.WindowManager = *wm
		src.Dksub()
	} else if *wm == "herbstluftwm" {
		src.WindowManager = *wm
		src.Herbssub()
	} else if *wm == "i3" {
		src.WindowManager = *wm
		src.I3sub()
	} else if *wm == "hyprland" {
		src.WindowManager = *wm
		src.Hyprsub()
	} else {
		fmt.Println("no window manager specified")
		fmt.Println("please try again with: goober -wm DESIRED_WM")
	}

}
