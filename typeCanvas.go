package iotmaker_platform_webbrowser

import (
	"syscall/js"
)

const (
	kCanvasNotSet int = iota
	kCanvas2DContext
	kCanvas3DContext
)

// todo: selfContextType deve ser um enum

// en: The Canvas API provides a means for drawing graphics via JavaScript and the HTML <canvas> element. Among other
// things, it can be used for animation, game graphics, data visualization, photo manipulation, and real-time video
// processing.
//
// The Canvas API largely focuses on 2D graphics. The WebGL API, which also uses the <canvas> element, draws
// hardware-accelerated 2D and 3D graphics.
type Canvas struct {
	selfCanvas      js.Value
	selfContext     js.Value
	selfContextType int
	Element
}

func (el *Canvas) GetCanvas() js.Value {
	return el.selfCanvas
}

func (el *Canvas) GetContext() js.Value {
	return el.selfContext
}

func (el *Canvas) GetElement() js.Value {
	return el.selfElement
}

func (el *Canvas) Call(jsFunction string, value interface{}) js.Value {
	return el.selfDocument.Call(jsFunction, value)
}

func (el *Canvas) Set(jsParam string, value ...interface{}) {
	el.selfDocument.Set(jsParam, value)
}

/*func (el *Canvas) CreateNewWith3DContext(width, height float64) {
	el.selfCanvas = el.Call("getElementsById", "myCanvas")
	el.Set("width", width)
	el.selfCanvas.Set("height", height)
	el.selfCanvas.Call("getContext", "3d")
}*/

func (el *Canvas) GlobalAlpha(value float64) {
	el.selfContext.Call("globalAlpha", value)
}

func (el *Canvas) MoveTo(x, y float64) {
	el.selfContext.Call("moveTo", x, y)
}

// en: Adds a new point and creates a line from that point to the last specified point in the canvas
func (el *Canvas) LineTo(x, y float64) {
	el.selfContext.Call("lineTo", x, y)
}

func (el *Canvas) StrokeStyle(value string) {
	el.selfContext.Set("strokeStyle", value)
}

func (el *Canvas) Stroke() {
	el.selfContext.Call("stroke")
}

func (el *Canvas) BeginPath() {
	el.selfContext.Call("beginPath")
}

// todo: tem que saber que id é um canvas
func (el *Canvas) InitializeContext2DById(id string) {
	el.Document.Initialize()
	el.selfCanvas = el.Element.NewCanvas(id)
	el.selfContextType = 1
	el.selfContext = el.selfCanvas.Call("getContext", "2d")
}

// todo: tem que saber que id é um canvas
func (el *Canvas) InitializeContext3DById(id string) {
	el.Element.NewCanvas(id)
	el.selfContextType = 2
	el.selfContext = el.selfCanvas.Call("getContext", "3d")
}

func (el *Canvas) AppendToDocumentBody() {
	el.selfDocument.Get("body").Call("appendChild", el.selfElement)
}
