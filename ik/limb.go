package ik

// Procedurally animate limbs with Inverse Kinamatics
// Using:	Cyclic coordinate descent algorithms
//			One of the simplier and faster approach to Inverse Kinematics
//			1) Loop backwards, updating the each element in an array
//			2) AND, run another inner loop forward, updating every follwing element, (of the current element in 1.)

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Limb struct {
	segments []*Segment
}

func lerp(t, lo, hi float64) float64 {
	return float64((1-t)*lo + t*hi)
}

func LimbNew(x, y, nbrSegments int) *Limb {
	p := Point{float64(x), float64(y)}
	segments := make([]*Segment, nbrSegments)
	for i := 0; i < nbrSegments; i++ {
		t := float64(i) / float64(nbrSegments)
		segments[i] = segmentNew(p, lerp(t, 30, 100), lerp(t, 10, 1), lerp(t, 1.0, 0.1), color.RGBA{255, 255, 255, 255})
	}
	return &Limb{segments: segments}
}

// Update all segments of the limb with respect to the target point
func (l *Limb) Update(target Point) {
	lastIdx := len(l.segments) - 1
	lastSegment := l.segments[lastIdx]

	// Outer loop: Iterate from the last segment to the first, and update each.
	for i := lastIdx; i >= 0; i-- {
		end := lastSegment.end()
		l.segments[i].updateAngle(end, target)

		// Inner loop: Update all the following segments start positions, based on the previous segments
		for j := i; j < lastIdx; j++ {
			curr := l.segments[j]
			next := l.segments[j+1]
			next.start = curr.end()
		}
	}
}

func (l *Limb) Draw(screen *ebiten.Image) {
	for _, s := range l.segments {
		s.draw(screen)
	}
}
