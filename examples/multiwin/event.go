package main

type appEventNewWindow struct {
	opt windowOption
	ret chan appEventNewWindowReturn
}

type appEventNewWindowReturn struct {
	w   *Window
	err error
}

type appEventDestroyWindow struct {
	w *Window
}

type appEventDestroyApplication struct {
}

type EventResize struct {
	Width  int
	Height int
}

type EventClose struct {
}

type EventMousePos struct {
	X, Y int
}

type EventChar struct {
	Char rune
}
