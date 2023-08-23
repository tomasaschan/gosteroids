package graphics

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/text"
)

func TextAt(target pixel.Target, location pixel.Vec, content string, scale float64) {
	txt := text.New(location, text.Atlas7x13)
	transform := pixel.IM.Scaled(location, scale)
	bbx := txt.BoundsOf(content)
	transform = transform.Moved(pixel.V(-scale*(bbx.Max.X-bbx.Min.X)/2, 0))
	txt.WriteString(content)
	txt.Draw(target, transform)
}
