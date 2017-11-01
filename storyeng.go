package storyeng

import (
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
	Parts    parts
	Current  string
	Previous string
	Name     string
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
func (c *cache) save(data string) {
	err := ioutil.WriteFile(c.Data, []byte(data), 755)
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
	//I have no idea bout yes/no/action
	return &inputEvent{line, empty, lower, crunched}
}
func (i *inputEvent) actionFor() {
	fmt.Println("TODO")
}

var yes = []string{"yes"}
var no = []string{"no"}

//todo colors

type themePrompt struct {
	text  string
	color string
}

func newthemePrompt() *themePrompt {
	return &themePrompt{"--->", "TODO"}
}

type themeInput struct {
	color string
}

func newthemeInput() *themeInput {
	return &themeInput{"TODO"}
}

type themeTell struct {
	color string
	wrap  bool
}

func newthemeTell() *themeTell {
	var color = "TODO"
	var wrap bool
	return &themeTell{color, wrap}
}

type themeMessages struct {
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

func init() {
	OsTest()
}

func debug(msg string) {
	if isDebug {
		fmt.Println(msg)
	}
}
