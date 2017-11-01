package storyeng

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

//let credits = ...
const Credits = `
The story engine behind this game was created by ...
           [0;31m        __   .__.__            __          __
           [0;31m  _____|  | _|__|  |   _______/  |______  |  | __
           [0;31m /  ___/  |/ /  |  |  /  ___/\\   __\\__  \\ |  |/ /
           [0;31m \\___ \\|    <|  |  |__\\___ \\  |  |  / __ \\|    <
           [0;31m/____  >__|_ \\__|____/____  > |__| (____  /__|_ \\[1;37m_______[0m
           [0;31m     \\/     \\/            \\/            \\/     \\/[1;37m______/[0m
                                        [1;37mCoding Arts[0m
`

// let copyright = ...
const Copyright = `
Copyright (c) 2017 SkilStak, Inc.
Redistribution and use in source and binary forms, with or without modification,
are permitted provided that the following conditions are met:
1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation and/or
   other materials provided with the distribution.
THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON
ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
`

//not import checked
//let parts = {}
type part map[string]interface{}
type parts map[string]part

var Parts parts

//let current = ""
var current string

//let previous = ""
var previous string

//let laste = ''
var last string

/* let actions = {}
let actions.also = {}*/
type Actions struct {
	Also struct {
	}
}

var actions Actions

//let data = ...
type data struct {
	Parts    parts  `json:"parts"`
	Current  string `json:"current"`
	Previous string `json:"previous"`
	Name     string `json:"name"`
}

var Data = data{Name: "friend"}

//let iswindows,ismac,islinux,isbsd,isposix = ...
//called in init
var isWindows bool
var isMac bool
var isLinux bool
var isBSD bool
var isPosix bool
var isDebug = true

func OsTest() {
	var os = runtime.GOOS
	debug(os)
	switch os {
	case "windows":
		isWindows = true
		isPosix = false
	case "mac":
		isMac = true
	case "linux":
		isLinux = true
	case "freebsd", "netbsd", "openbsd":
		isBSD = true
	}
}

//class cache...
type cache struct {
	Name        string
	Path        string
	Modpath     string
	Id          string
	Data        string
	Partspath   string
	Actionspath string
}

//constructor...
func NewCache() *cache {
	//this.name = ...
	_, filename, _, ok := runtime.Caller(1)
	if ok != true {
		panic(ok)
	}
	name := "." + filepath.Base(filename)
	debug(name)
	//this.path =
	dir, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(dir, name)

	debug(path)
	//this.modpath = ...
	//for now set to string
	modpath := "/home/liam/go/src/github.com/liamnaddell/storygame-sample-go"
	//this.id = ..
	id := filepath.Base(modpath)
	//this.data = ...
	data := filepath.Join(path, id+".json")
	//this.partspath
	partspath := filepath.Join(modpath, "parts")
	//this.actionspath
	actionspath := filepath.Join(modpath, "actions")
	ourcache := &cache{name, path, modpath, id, data, partspath, actionspath}
	ourcache.make()
	return ourcache
}

func (c *cache) make() {
	_, err := os.Stat(c.Path)
	if err != nil {
		_ = os.Mkdir(c.Path, 755)
	}
}
func (c *cache) save(data data) {
	dat, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(c.Data, dat, 755)
	if err != nil {
		panic(err)
	}

}
func (c *cache) load() map[string]interface{} {
	_, err := os.Stat(c.Data)
	if err != nil {
		return nil
	}
	var plain, err2 = ioutil.ReadFile(c.Data)
	if err2 != nil {
		panic(err2)
	}
	var dat map[string]interface{}
	err = json.Unmarshal(plain, &dat)
	if err != nil {
		panic(err)
	}
	return dat

}
func (c *cache) remove() {
	err := os.RemoveAll(c.Data)
	if err != nil {
		panic(err)
	}
}

var Cache = NewCache()

func crunch(text string) string {
	var regex = regexp.MustCompile(`[\W]+`)
	return string(regex.ReplaceAll([]byte(text), []byte("")))
}

type event struct {
	Data data
}

func NewEvent() *event {
	return &event{Data: Data}
}

type inputEvent struct {
	Line     string
	Empty    bool
	Lower    string
	Crunched string
	Yes      bool
	No       bool
}

func NewInputEvent() *inputEvent {
	//what is line?
	var line string
	line = strings.TrimSpace(line)
	var empty bool
	if line == "" {
		empty = true
	}
	lower := strings.ToLower(line)
	crunched := crunch(line)
	var Yes, No bool
	//I have no idea bout yes/no/action
	for i := 0; i < len(yes); i++ {
		s := strings.Split(lower, yes[i])
		fmt.Println("DEBUG:", s)
		if len(s) > 2 {
			Yes = true
		}
	}
	for i := 0; i < len(no); i++ {
		s := strings.Split(lower, no[i])
		fmt.Println("DEBUG:", s)
		if len(s) > 2 {
			No = true
		}
	}

	return &inputEvent{line, empty, lower, crunched, Yes, No}
}
func (i *inputEvent) actionFor() {
	fmt.Println("TODO")
}

var yes = []string{"yes"}
var no = []string{"no"}

var c = make(map[string]string)

func PopulateColors() {
	if isPosix {
		c["b03"] = "\x1b[1;30m"
		c["b02"] = "\x1b[0;30m"
		c["b01"] = "\x1b[1;32m"
		c["b00"] = "\x1b[1;33m"
		c["b0"] = "\x1b[1;33m"
		c["b1"] = "\x1b[1;34m"
		c["b2"] = "\x1b[0;37m"
		c["b3"] = "\x1b[1;37m"
		c["y"] = "\x1b[0;33m"
		c["o"] = "\x1b[1;31m"
		c["r"] = "\x1b[0;31m"
		c["m"] = "\x1b[0;35m"
		c["v"] = "\x1b[1;35m"
		c["b"] = "\x1b[0;34m"
		c["c"] = "\x1b[0;36m"
		c["g"] = "\x1b[0;32m"
		c["x"] = "\x1b[0m"
		c["l"] = "\x1b[2K\x1b[G"
		c["c"] = "\x1b[2J\x1b[H"
	} else {
		c["b03"] = ""
		c["b02"] = ""
		c["b01"] = ""
		c["b00"] = ""
		c["b2"] = ""
		c["b0"] = ""
		c["b3"] = ""
		c["y"] = ""
		c["o"] = ""
		c["r"] = ""
		c["m"] = ""
		c["v"] = ""
		c["b"] = ""
		c["c"] = ""
		c["g"] = ""
		c["x"] = ""
		c["l"] = ""
		c["c"] = ""
	}
}

type themePrompt struct {
	text  string
	color string
}

func newthemePrompt() *themePrompt {
	return &themePrompt{"--->", c["b3"]}
}

type themeInput struct {
	color string
}

func newthemeInput() *themeInput {
	return &themeInput{c["y"]}
}

type themeTell struct {
	color string
	wrap  bool
}

func newthemeTell() *themeTell {
	var color = c["b0"]
	var wrap bool
	return &themeTell{color, wrap}
}

type themeMessages struct {
	nopart  string
	bye     string
	nostart string
	restart string
}

type theme struct {
	prompt *themePrompt
	input  *themeInput
	tell   *themeTell
	msgs   *themeMessages
}

func NewTheme() *theme {
	return &theme{newthemePrompt(), newthemeInput(), newthemeTell(), new(themeMessages)}
}

var Theme = NewTheme()

func SetThemeStuff() {
	Theme.msgs.nopart = "**I'm sorry. The author has not written the next part yet.**"
	Theme.msgs.bye = "Sorry to see you go, {name}. See you soon."
	Theme.msgs.nostart = "It appears the author has not added the required *Start* part."
	Theme.msgs.restart = "Do you really want to delete your data and restart?"
}

func clear() {
	fmt.Println(c["clear"])
}
func format(s string) {
	fmt.Println("TODOODODODO")
}
func tells(s string) {
	fmt.Println("TOOTOODO")
}
func tell(s string) {
	fmt.Println(s)
}
func show(s string) {
	fmt.Println(strings.TrimSpace(s))
}
func quit() {
	tell(Theme.msgs.bye)
	Cache.save(Data)
	os.Exit(0)
}
func leave(part part) {
	onleave := part["onleave"]
	var e = NewEvent()
	switch onleave.(type) {
	case func(*event) string:
		s := onleave.(func(*event) string)(e)
		tell(s)
	case string:
		tell(onleave.(string))
	}

}
func enter(part part) {
	onleave := part["onleave"]
	var e = NewEvent()
	switch onleave.(type) {
	case func(*event) string:
		rv := onleave.(func(*event) string)(e)
		tell(rv)

	case string:
		tell(onleave.(string))
	}

}

//don't know how to translate actions.X stuff
//or actions.also stuffz

//what is rl.on?
var rl = bufio.NewScanner(os.Stdin)

func update() {
	clear()
	if previous != "" {
		var prevpart = Parts[previous]
		Data.Previous = previous
		leave(prevpart)
	}
	var part = Parts[current]
	Data.Current = current
	enter(part)
	rl.Scan()
	fmt.Println(rl.Text())

}

func Go(name string) {
	if Parts[name] != nil {
		tell(Theme.msgs.nopart)
		return
	}
	previous = current
	current = name
	Cache.save(Data)
	update()
}
func init() {
	OsTest()
	PopulateColors()
}
func debug(msg string) {
	if isDebug {
		fmt.Println(msg)
	}
}
