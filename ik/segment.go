package ik

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Point struct {
	X, Y float64
}
type Segment struct {
	angle      float64
	length     float64
	start      Point
	color      color.RGBA
	width      float64
	adjustment float64
}

func segmentNew(start Point, length, width, adjustment float64, color color.RGBA) *Segment {
	return &Segment{
		length:     length,
		start:      start,
		width:      width,
		color:      color,
		adjustment: adjustment,
	}
}

func (s *Segment) setPoint(p Point) {
	s.start = p
}

// Updates the angle of the segment, (to make the endpoint align with the target position)
func (s *Segment) updateAngle(end, target Point) {
	angle := math.Atan2(end.Y-s.start.Y, end.X-s.start.X)
	targetAngle := math.Atan2(target.Y-s.start.Y, target.X-s.start.X)
	delta := targetAngle - angle
	for delta < -math.Pi {
		delta += 2 * math.Pi
	}
	for delta > math.Pi {
		delta -= 2 * math.Pi
	}
	s.angle += delta * s.adjustment * 0.1 // Tip: Don't change if "end arm" is close to target
}

func (s *Segment) end() Point {
	x1 := s.start.X + math.Cos(s.angle)*s.length
	y1 := s.start.Y + math.Sin(s.angle)*s.length
	return Point{x1, y1}
}

func (s *Segment) draw(screen *ebiten.Image) {
	end := s.end()
	vector.StrokeLine(screen,
		float32(s.start.X), float32(s.start.Y),
		float32(end.X), float32(end.Y),
		float32(s.width), s.color, true,
	)
}
