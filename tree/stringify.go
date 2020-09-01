package tree

import (
	"fmt"
	"log"
	"math"
	"sort"

	"github.com/m1gwings/treedrawer/drawer"
)

// stringify takes a pointer to a node and draws all the tree below in a drawer.
// Returns the drawn drawer.
// This function is called recursively
func stringify(t *Tree) *drawer.Drawer {
	// Getting drawer and dimensions of this NodeValue
	dVal := t.val.Draw()
	dValW, dValH := dVal.Dimens()

	// No children
	if len(t.Children()) == 0 {
		// Allocating new drawer to return
		// Ensuring that width is odd
		d, err := drawer.NewDrawer(dValW+2+1-dValW%2, dValH+2)
		if err != nil {
			log.Fatal(fmt.Errorf("error while allocating new drawer with no children: %v", err))
		}

		// Drawing dVal drawer onto the drawer to return
		err = d.DrawDrawer(dVal, 1, 1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing val with no children: %v", err))
		}

		// Adding a box in the drawer to return, around where the dVal drawer has been drawn
		err = addBoxAround(d, 0, 0, dValW+1, dValH+1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while adding box with no children: %v", err))
		}
		return d
	}

	// One child
	if len(t.Children()) == 1 {
		// Drawer of the child
		var dChild *drawer.Drawer

		// Recursively calling stringify of the child and initializing dChild drawer
		tChild, err := t.Child(0)
		if err != nil {
			log.Fatal(fmt.Errorf("error while getting child 0 with one child: %v", err))
		}
		dChild = stringify(tChild)
		// Getting dimensions of dChild drawer
		dChildW, dChildH := dChild.Dimens()

		// w and h represent respectively width and height of the drawer to return
		// w is the max between the width of dVal + 2 (considering the box) and the width of the one child
		// h is equal to the height of dVal + 2 (considering the box) + 1 (considering the "pipe") + the height of dChild
		w := int(math.Max(float64(dValW+2), float64(dChildW)))
		// Ensuring that w is odd
		w += 1 - w%2
		h := dValH + 3 + dChildH

		// Allocating new drawer to return
		d, err := drawer.NewDrawer(w, h)
		if err != nil {
			log.Fatal(fmt.Errorf("error while allocating new drawer with one child: %v", err))
		}

		// Drawing dVal onto the drawer to return with x in (w-dValW)/2 to put dVal in the middle
		// (remember that drawer.DrawDrawer takes coordinates of top left corner)
		// and y in 1 (considering the box)
		err = d.DrawDrawer(dVal, (w-dValW)/2, 1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing val with one child: %v", err))
		}

		// Adding a box in the drawer to return, around where the dVal drawer has been drawn
		// start coordinates are taken considering d.DrawDrawer above - 1 in order to not overwrite
		// end coordinates are just start coordinates plus respectively dValW+1 and dValH+1 in order to not overwrite
		err = addBoxAround(d, (w-dValW)/2-1, 0, (w-dValW)/2+dValW, dValH+1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while adding box with one child: %v", err))
		}

		// Drawing the upper-link onto the drawer to return with x in the middle
		// and y just above the pipe
		err = d.DrawRune('┬', w/2, dValH+1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ┬ with one child: %v", err))
		}

		// Drawing the pipe onto the drawer to return with x in the middle
		// and y in dValH + 2 (considering the box)
		err = d.DrawRune('│', w/2, dValH+2)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing | with one child: %v", err))
		}

		// Drawing dChild onto the drawer to return with x in (w-dChildW)/2 to put dChild in the middle
		// (remember that drawer.DrawDrawer takes coordinates of top left corner)
		// and y in dValH + 3 (considering the box and pipe)
		err = d.DrawDrawer(dChild, (w-dChildW)/2, dValH+3)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing child drawer with one child: %v", err))
		}

		// Drawing the lower-link onto the drawer to return with x in the middle
		// and y just below the pipe
		// this drawing must be the latest because it has to overwrite dChild
		err = d.DrawRune('┴', w/2, dValH+3)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ┴ with one child: %v", err))
		}

		return d
	}

	// More children

	// nChildren is the number of children of t
	nChildren := len(t.Children())
	// dChildren stores the result of recursively call of stringify for each child
	dChildren := make([]*drawer.Drawer, 0, nChildren)
	// childrenLeft is a slice with the x coordinate of the upper-left corner of each child drawer to draw onto d
	childrenLeft := make([]int, 0, nChildren)
	// childrenMiddle is a slice with the x coordinate of the middle of each child drawer to draw onto d
	childrenMiddle := make([]int, 0, nChildren)
	// childrenW is the width required to draw children
	// it is incremented child by child to obtain the x coordinate of the upper-left corner for each child
	childrenW := 0
	// maxChildH is the maximum height of a child
	maxChildH := 0

	// Iterates over children to calculate maxChildH, childrenLeft and childrenMiddle
	for i, tChild := range t.Children() {
		dChild := stringify(tChild)
		dChildren = append(dChildren, dChild)
		dChildW, dChildH := dChild.Dimens()
		maxChildH = int(math.Max(float64(maxChildH), float64(dChildH)))

		if i == nChildren-1 {
			// When the child is the last
			if (childrenW+dChildW)%2 == 1 {
				// If final childrenW (notice that childrenW gets incremented at the end) is odd than we just have to add dChildW
				childrenLeft = append(childrenLeft, childrenW)
				childrenMiddle = append(childrenMiddle, childrenW+dChildW/2)
				childrenW += dChildW
			} else {
				// Otherwise we add one more space to make childrenW odd
				childrenLeft = append(childrenLeft, childrenW+1)
				childrenMiddle = append(childrenMiddle, childrenW+1+dChildW/2)
				childrenW += dChildW + 1
			}
		} else {
			// When the child isn't the last just add it to the left of the child before with a space in between
			childrenLeft = append(childrenLeft, childrenW)
			childrenMiddle = append(childrenMiddle, childrenW+dChildW/2)
			childrenW += dChildW + 1
		}
	}

	// Assert that childrenLeft and childrenMiddle are sorted, this is required because we are going to use binary search later
	sorted := sort.SliceIsSorted(childrenLeft, func(i, j int) bool { return childrenLeft[i] < childrenLeft[j] })
	if !sorted {
		log.Fatal(fmt.Errorf("childrenLeft is not sorted"))
	}
	sorted = sort.SliceIsSorted(childrenMiddle, func(i, j int) bool { return childrenMiddle[i] < childrenMiddle[j] })
	if !sorted {
		log.Fatal(fmt.Errorf("childrenMiddle is not sorted"))
	}

	// w is the width of the final drawer and is equal to the maximum between dValW+2 and childrenW
	var w int
	if dValW+2 > childrenW {
		w = dValW + 2
		// If parent width is greater than children width, children get centered by shifting each child
		for i := 0; i < nChildren; i++ {
			childrenLeft[i] += (w - childrenW) / 2
			childrenMiddle[i] += (w - childrenW) / 2
		}
	} else {
		w = childrenW
	}
	h := dValH + 3 + maxChildH

	// Allocating new drawer to return
	d, err := drawer.NewDrawer(w, h)
	if err != nil {
		log.Fatal(fmt.Errorf("error while allocating new drawer with more children: %v", err))
	}

	// Drawing dVal onto the drawer to return with x in (w-dValW)/2 to put dVal in the middle
	// (remember that drawer.DrawDrawer takes coordinates of top left corner)
	// and y in 1 (considering the box)
	err = d.DrawDrawer(dVal, (w-dValW)/2, 1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing val with more children: %v", err))
	}

	// Adding a box in the drawer to return, around where the dVal drawer has been drawn
	// start coordinates are taken considering d.DrawDrawer above - 1 in order to not overwrite
	// end coordinates are just start coordinates plus respectively dValW+1 and dValH+1 in order to not overwrite
	err = addBoxAround(d, (w-dValW)/2-1, 0, (w-dValW)/2+dValW, dValH+1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while adding box with more children: %v", err))
	}

	// Drawing children onto the drawer to return
	for i := 0; i < nChildren; i++ {
		err = d.DrawDrawer(dChildren[i], childrenLeft[i], dValH+3)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing %d child: %v", i, err))
		}
	}

	// Drawing upper-link ┬ under the parent
	err = d.DrawRune('┬', w/2, dValH+1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing upper-link ┬ under the parent: %v", err))
	}

	// Drawing lower-link ┴ above the children
	for i, x := range childrenMiddle {
		err = d.DrawRune('┴', x, dValH+3)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing lower-link ┴ above the %dth child: %v", i, err))
		}
	}

	// Drawing left-corner ╭ above the left most child
	err = d.DrawRune('╭', childrenMiddle[0], dValH+2)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing left-corner ╭ above the left most child: %v", err))
	}
	// Drawing right-corner ╮ above the right most child
	err = d.DrawRune('╮', childrenMiddle[len(childrenMiddle)-1], dValH+2)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing right-corner ╮ above the right most child: %v", err))
	}

	// Finish to connect the pipe
	for x := childrenMiddle[0] + 1; x < childrenMiddle[len(childrenMiddle)-1]; x++ {
		underParent := x == w/2
		shouldBeAt := sort.SearchInts(childrenMiddle, x)
		aboveChild := shouldBeAt < len(childrenMiddle) && childrenMiddle[shouldBeAt] == x
		var connection rune
		switch {
		case underParent && aboveChild:
			connection = '┼'
		case underParent:
			connection = '┴'
		case aboveChild:
			connection = '┬'
		default:
			connection = '─'
		}
		err = d.DrawRune(connection, x, dValH+2)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing %c at position %d to finish connection: %v", connection, x, err))
		}
	}

	return d
}

// addBoxAround draws a box onto d
// the box starts at startX and startY coordinates
// and ends at endX and endY
func addBoxAround(d *drawer.Drawer, startX, startY, endX, endY int) error {
	// Checking that start and end coordinates are valid
	if startX < 0 || startY < 0 || endX < 0 || endY < 0 {
		return fmt.Errorf("can't draw on negative coordinates %d %d %d %d", startX, startY, endX, endY)
	}
	if startX > endX || startY > endY {
		return fmt.Errorf("start should be before end %d %d %d %d", startX, startY, endX, endY)
	}
	dW, dH := d.Dimens()
	if endX >= dW || endY >= dH {
		return fmt.Errorf("end overflows the drawer with dimes %d %d, %d %d %d %d", dW, dH, startX, startY, endX, endY)
	}

	// Drawing corners
	err := d.DrawRune('╭', startX, startY)
	if err != nil {
		return fmt.Errorf("error while drawing ╭: %v", err)
	}
	err = d.DrawRune('╮', endX, startY)
	if err != nil {
		return fmt.Errorf("error while drawing ╮: %v", err)
	}
	err = d.DrawRune('╰', startX, endY)
	if err != nil {
		return fmt.Errorf("error while drawing ╰: %v", err)
	}
	err = d.DrawRune('╯', endX, endY)
	if err != nil {
		return fmt.Errorf("error while drawing ╯: %v", err)
	}

	// Drawing edges
	for x := startX + 1; x < endX; x++ {
		for yMul := 0; yMul <= 1; yMul++ {
			err = d.DrawRune('─', x, yMul*(endY-startY)+startY)
			if err != nil {
				return fmt.Errorf("error while drawing ─: %v", err)
			}
		}
	}
	for y := startY + 1; y < endY; y++ {
		for xMul := 0; xMul <= 1; xMul++ {
			err = d.DrawRune('│', xMul*(endX-startX)+startX, y)
			if err != nil {
				return fmt.Errorf("error while drawing │: %v", err)
			}
		}
	}
	return nil
}
