package canvas

import (
	iotmaker_types "github.com/helmutkemper/iotmaker.types"
	"syscall/js"
)

// en: Draws text on the canvas (no fill)
//     text:     Specifies the text that will be written on the canvas
//     x:        The x coordinate where to start painting the text (relative to the canvas)
//     y:        The y coordinate where to start painting the text (relative to the canvas)
//     maxWidth: Optional. The maximum allowed width of the text, in pixels
//
//     The strokeText() method draws text (with no fill) on the canvas. The default color of the text is black.
//     Tip: Use the font property to specify font and font size, and use the strokeStyle property to render the text in another color/gradient.
//     JavaScript syntax: context.strokeText(text, x, y, maxWidth);
//
//     var c = document.getElementById("myCanvas");
//     var ctx = c.getContext("2d");
//     ctx.font = "20px Georgia";
//     ctx.strokeText("Hello World!", 10, 50);
//     ctx.font = "30px Verdana";
//     // Create gradient
//     var gradient = ctx.createLinearGradient(0, 0, c.width, 0);
//     gradient.addColorStop("0", "magenta");
//     gradient.addColorStop("0.5", "blue");
//     gradient.addColorStop("1.0", "red");
//     // Fill with gradient
//     ctx.strokeStyle = gradient;
//     ctx.strokeText("Big smile!", 10, 90);
func (el *Canvas) StrokeText(text string, x, y, maxWidth iotmaker_types.Coordinate) {
	el.SelfContext.Call("strokeText", text, x, y, maxWidth)
}
