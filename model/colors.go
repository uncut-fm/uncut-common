package model

import (
	"fmt"
	"github.com/PerformLine/go-stockutil/mathutil"
	"math"
)

type RGB struct {
	R, G, B uint32
}

type HSL struct {
	H, S, L float64
}

func (c HSL) GetAccent() HSL {
	accent := HSL{
		H: c.H,
		S: c.S,
		L: c.L,
	}
	if c.S < 0.55 {
		accent.S = 0.55
	}

	if c.L < 0.6 {
		if 180 < c.H && c.H < 290 {
			accent.L = 0.7
		} else {
			accent.L = 0.6
		}
	}

	if c.L > 0.8 {
		c.L = 0.8
	}

	return accent
}

func (c HSL) GetBackground() HSL {
	background := HSL{
		H: c.H,
		S: c.S,
		L: 0.08,
	}

	if c.S < 0.4 {
		c.S = 0.4
	}

	return background
}

func (c HSL) String() string {
	return fmt.Sprintf("H%f S%f L%f", c.H, c.S, c.L)
}

func (c RGB) String() string {
	return fmt.Sprintf("rgb(%d,%d,%d)", c.R, c.G, c.B)
}

const precision = 2

func (c RGB) ToHSL() HSL {
	var h, s, lvi float64

	r := float64(c.R) / 255
	g := float64(c.G) / 255
	b := float64(c.B) / 255

	var huePrime float64

	// hue
	// ---------------------------------------------------------------------------------------------
	max := math.Max(math.Max(r, g), b)
	min := math.Min(math.Min(r, g), b)
	chroma := (max - min)

	if chroma == 0 {
		h = 0
	} else {
		if r == max {
			huePrime = math.Mod(((g - b) / chroma), 6)
		} else if g == max {
			huePrime = ((b - r) / chroma) + 2

		} else if b == max {
			huePrime = ((r - g) / chroma) + 4

		}

		h = huePrime * 60
	}

	// lightness
	// ---------------------------------------------------------------------------------------------
	if r == g && g == b {
		lvi = r
	} else {

		lvi = (max + min) / 2
	}

	// saturation
	// ---------------------------------------------------------------------------------------------
	if lvi == 1 {
		s = 0
	} else {
		s = (chroma / (1 - math.Abs(2*lvi-1)))
	}

	if math.IsNaN(s) {
		s = 0
	}

	h = mathutil.RoundPlaces(h, precision)
	s = mathutil.RoundPlaces(s, precision)
	lvi = mathutil.RoundPlaces(lvi, precision)

	if h < 0 {
		h = 360 + h
	}

	return HSL{h, s, lvi}
}

func (c HSL) ToRgb() RGB {
	var r, g, b float64

	hueDegrees := math.Mod(c.H, 360)

	if c.S == 0 {
		r = c.L
		g = c.L
		b = c.L
	} else {
		var chroma float64
		var m float64

		chroma = (1 - math.Abs((2*c.L)-1)) * c.S

		hueSector := hueDegrees / 60

		intermediate := chroma * (1 - math.Abs(
			math.Mod(hueSector, 2)-1,
		))

		switch {
		case hueSector >= 0 && hueSector <= 1:
			r = chroma
			g = intermediate
			b = 0

		case hueSector > 1 && hueSector <= 2:
			r = intermediate
			g = chroma
			b = 0

		case hueSector > 2 && hueSector <= 3:
			r = 0
			g = chroma
			b = intermediate

		case hueSector > 3 && hueSector <= 4:
			r = 0
			g = intermediate
			b = chroma
		case hueSector > 4 && hueSector <= 5:
			r = intermediate
			g = 0
			b = chroma

		case hueSector > 5 && hueSector <= 6:
			r = chroma
			g = 0
			b = intermediate

		default:
			panic(fmt.Errorf("hue input %v yielded sector %v", hueDegrees, hueSector))
		}

		m = c.L - (chroma / 2)

		r += m
		g += m
		b += m

	}

	r = mathutil.RoundPlaces(r, precision)
	g = mathutil.RoundPlaces(g, precision)
	b = mathutil.RoundPlaces(b, precision)

	return RGB{
		R: uint32(r * 255.0),
		G: uint32(g * 255.0),
		B: uint32(b * 255.),
	}
}
