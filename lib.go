package storyeng

import (
	"fmt"
	"regexp"
	"runtime"
)

const credits = `
The story engine behind this game was created by ...
           [0;31m        __   .__.__            __          __
           [0;31m  _____|  | _|__|  |   _______/  |______  |  | __
           [0;31m /  ___/  |/ /  |  |  /  ___/\\   __\\__  \\ |  |/ /
           [0;31m \\___ \\|    <|  |  |__\\___ \\  |  |  / __ \\|    <
           [0;31m/____  >__|_ \\__|____/____  > |__| (____  /__|_ \\[1;37m_______[0m
           [0;31m     \\/     \\/            \\/            \\/     \\/[1;37m______/[0m
                                        [1;37mCoding Arts[0m
`
const copyright = `
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

//typedefs
type PartStruct struct {
}
type ActionsStruct struct {
	Also struct {
	}
}
type DataStruct struct {
	Parts struct {
	}
	Current  string
	Previous string
	Name     string
}
type CacheStruct struct {
	Name       string
	Path       string
	Modpath    string
	id         string
	data       string
	partspath  string
	actionpath string
}
type EventStruct struct {
	Data DataStruct
}

func NewEvent() *EventStruct {
	return &EventStruct{Data: data}
}

//vars
var (
	current   string
	previous  string
	laste     string
	isWindows bool
	isMac     bool
	isLinux   bool
	isPosix   bool
	isBSD     bool
	isDebug   = true
	parts     PartStruct
	actions   ActionsStruct
	data      = DataStruct{Name: "friend"}
	Cache     = NewCache()
	save      = Cache.Save
	load      = Cache.Load
	regex     = regexp.MustCompile(`[\W]+`)
)

//funcs

func init() {
	os := getOS()
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
func getOS() string {
	return runtime.GOOS
}

func debug(msg string) {
	if isDebug {
		fmt.Println(msg)
	}
}
func crunch(text string) string {
	return string(regex.ReplaceAll([]byte(text), []byte("")))
}

func (c *CacheStruct) Make() {
}
func (c *CacheStruct) Save(data string) {
}
func (c *CacheStruct) Load() string {
	return ""
}
func (c *CacheStruct) Remove() {
}

func NewCache() *CacheStruct {
	//TODO
	return new(CacheStruct)
}
