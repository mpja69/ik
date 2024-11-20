package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type point struct {
	x, y float64
}
type segment struct {
	angle  float64
	length float64
	start  point
	color  color.RGBA
	width  float64
}

func segmentNew(start point, length, width float64, color color.RGBA) *segment {
	return &segment{
		length: length,
		start:  start,
		width:  width,
		color:  color,
	}
}

// Updates the angle of the segment, (to make the endpoint align with the target position)
func (s *segment) updateAngle(end, target point) {
	angle := math.Atan2(end.y-s.start.y, end.x-s.start.x)
	targetAngle := math.Atan2(target.y-s.start.y, target.x-s.start.x)
	delta := targetAngle - angle
	for delta < -math.Pi {
		delta += 2 * math.Pi
	}
	for delta > math.Pi {
		delta -= 2 * math.Pi
	}
	s.angle += delta * 0.5 // Tip: Don't change if "end arm" is close to target
}

func (s *segment) end() point {
	x1 := s.start.x + math.Cos(s.angle)*s.length
	y1 := s.start.y + math.Sin(s.angle)*s.length
	return point{x1, y1}
}

func (s *segment) draw(screen *ebiten.Image) {
	end := s.end()
	vector.StrokeLine(screen,
		float32(s.start.x), float32(s.start.y),
		float32(end.x), float32(end.y),
		float32(s.width), s.color, true,
	)
}
