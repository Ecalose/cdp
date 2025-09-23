package cdp

import (
	"github.com/gospider007/tools"
)

type Rect struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

func (obj Rect) Center() Point {
	return Point{
		X: obj.X + obj.Width/2,
		Y: obj.Y + obj.Height/2,
	}
}
func (obj Rect) RandCenter() Point {
	return Point{
		X: obj.X + tools.RanFloat64(0, int64(obj.Width)),
		Y: obj.Y + tools.RanFloat64(0, int64(obj.Height)),
	}
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
