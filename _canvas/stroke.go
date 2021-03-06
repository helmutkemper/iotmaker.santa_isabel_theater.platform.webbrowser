package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Actually draws the path you have defined
//     The stroke() method actually draws the path you have defined with all those moveTo() and lineTo() methods. The
//     default color is black.
//     Tip: Use the strokeStyle property to draw with another color/gradient.
func (el *Canvas) Stroke() {
	el.SelfContext.Call("stroke")
}
