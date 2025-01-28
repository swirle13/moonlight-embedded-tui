package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"sync"
)

// Always update this when new options are added
const configSize = 4

// TODO: write defaults to file for use on subsequent calls
const defaultWidth = 1400
const defaultHeight = 1200
const defaultBitrate = 17500
const defaultHostIP = "192.168.1.117"

// TODO: Write templates and template engine population logic

type moonlightConfig struct {
	appName     string
	description string
	// TODO: add config file name attribute once templates are working
	width   uint16
	height  uint16
	bitrate uint16
}

func (i moonlightConfig) GenerateCommand() *exec.Cmd {
	return exec.Command("moonlight", "stream", "-width", strconv.Itoa(int(i.width)), "-height", strconv.Itoa(int(i.height)), "-app", i.appName, "-bitrate", strconv.Itoa(int(i.bitrate)))
}

func (i moonlightConfig) RunCmd(cmd *exec.Cmd) {
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println(string(output))
		return
	}

	fmt.Println(string(output))
}

type moonlightConfigItemGenerator struct {
	moonlightConfigs [configSize]moonlightConfig
	index            int
	mtx              *sync.Mutex
}

// TODO: Add choice between multiple hosts, if found. Currently assumes one host.
// Fetch values from moonlight cli and populate instead of hardcode
// func getMoonlightApps() []string {
// 	cmd := exec.Command("moonlight", "list")
// 	cmdOutput := &bytes.Buffer{}
// 	cmd.Stdout = cmdOutput
// 	err := cmd.Run()
// 	if err != nil {
// 		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
// 	}

// 	output := cmdOutput.Bytes()
// 	if len(output) > 0 {
// 		fmt.Printf("==> Output: %s\n", string(output))
// 	}

// 	var appNames []string

// 	// parse output from `moonlight list`
// 	var appNameRegex = regexp.MustCompile(`^\d\.\s(\w+\s*\w*)$`)
// 	result := appNameRegex.FindAllStringSubmatch(string(output), -1)

// 	if len(result) > 0 {

// 	}
// }

// TODO: can make a child model for a sub menu if I want to choose app then resolution. Will be returned in a switch case
var configs = [configSize]moonlightConfig{
	// appName, description, configName, width, height, bitrate
	{"Desktop", "The big tamale", defaultWidth, defaultHeight, defaultBitrate},
	{"SteamBigPicture", "The world is your oyster", defaultWidth, defaultHeight, defaultBitrate},
	{"SquareGolf", "Meat and potatoes", defaultWidth, defaultHeight, defaultBitrate},
	{"GSPro", "The cream of the crop", defaultWidth, defaultHeight, defaultBitrate},
}

func (r *moonlightConfigItemGenerator) reset() {
	r.mtx = &sync.Mutex{}

	// set again when reset, unsure if needed; the example set values here
	r.moonlightConfigs = configs
}

func (r *moonlightConfigItemGenerator) next() item {
	if r.mtx == nil {
		r.reset()
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	i := item{
		title:       r.moonlightConfigs[r.index].appName,
		description: r.moonlightConfigs[r.index].description,
		config:      r.moonlightConfigs[r.index],
	}

	r.index++
	if r.index >= len(r.moonlightConfigs) {
		r.index = 0
	}

	return i
}
