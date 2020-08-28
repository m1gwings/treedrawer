package drawer

import "fmt"

// Drawer is a canvas on which you can draw ascii bytes.
type Drawer struct {
	canvas [][]byte
}

// NewDrawer returns a new Drawer with width w and height h.
func NewDrawer(w, h int) *Drawer {
	d := new(Drawer)
	d.canvas = make([][]byte, h)
	for i := range d.canvas {
		d.canvas[i] = make([]byte, w)
	}
	return d
}

// NewDrawerFromString return a new  Drawer which contains the string s.
func NewDrawerFromString(s string) *Drawer {
	d := new(Drawer)
	d.canvas = make([][]byte, 1)
	d.canvas[0] = []byte(s)
	return d
}

// DrawByte draws a byte in position x, y in the drawer canvas.
// Returns an error if the x, y position in input is outside the canvas.
func (d *Drawer) DrawByte(b byte, x, y int) error {
	w, h := d.Dimens()
	if x >= w || y >= h {
		return fmt.Errorf("position (%d, %d) is outside the canvas of dimension (%d, %d)", x, y, w, h)
	}
	d.canvas[y][x] = b
	return nil
}

// DrawDrawer draws the canvas inside e onto d with the up left corner in position x, y.
// Returns an error if the canvas inside e, drawn in position x, y, overflows the canvas in d.
func (d *Drawer) DrawDrawer(e *Drawer, x, y int) error {
	w, h := d.Dimens()
	eW, eH := e.Dimens()
	if x+eW-1 >= w || y+eH-1 >= h {
		return fmt.Errorf("canvas e of dimension (%d, %d) drawn in position (%d, %d) overflows canvas d of dimension (%d, %d)", eW, eH, x, y, w, h)
	}
	for i, row := range e.canvas {
		for j, b := range row {
			d.canvas[i+y][j+x] = b
		}
	}
	return nil
}

// Dimens returns width and height of the canvas.
func (d *Drawer) Dimens() (w, h int) {
	h, w = len(d.canvas), len(d.canvas[0])
	return
}

// String returns the string representation of the canvas.
func (d *Drawer) String() string {
	var s string
	for _, row := range d.canvas {
		for _, b := range row {
			if b == 0 {
				s += " "
			} else {
				s += string(b)
			}
		}
		s += "\n"
	}
	return s
}
