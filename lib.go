package storyeng

type Part map[string]interface{}
type partst map[string]Part

type Event struct {}
type InputEvent struct {
	Event
	Line string
	Lower string
	Yes bool
	No bool
}

var Parts = partst{}
var Data = map[string]interface{}{}

func Start() {
	println("starting")
}
