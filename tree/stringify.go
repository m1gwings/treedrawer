package tree

import (
	"fmt"
	"log"
	"math"

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
	if t.left == nil && t.right == nil {
		// Allocating new drawer to return
		d, err := drawer.NewDrawer(dValW+2, dValH+2)
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
	if (t.left != nil && t.right == nil) || (t.left == nil && t.right != nil) {
		// Drawer of the child
		var dChild *drawer.Drawer

		// Recursively calling stringify of the child which is not nil and initializing dChild drawer
		if t.left != nil {
			dChild = stringify(t.left)
		} else {
			dChild = stringify(t.right)
		}
		// Getting dimensions of dChild drawer
		dChildW, dChildH := dChild.Dimens()

		// w and h represent respectively width and height of the drawer to return
		// w is the max between the width of dVal + 2 (considering the box) and the width of the one child
		// h is equal to the height of dVal + 2 (considering the box) + 1 (considreing the "pipe") + the height of dChild
		w := int(math.Max(float64(dValW+2), float64(dChildW)))
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

	// Two children

	// Recursively calling stringify of the children and initializing dLeft and dRight drawer
	dLeft, dRight := stringify(t.left), stringify(t.right)
	// Getting dimensions of dLeft and dRight
	dLeftW, dLeftH := dLeft.Dimens()
	dRightW, dRightH := dRight.Dimens()

	// maxChildW is the max between the width of left child, the width of right child and half the width of the parent + 1
	maxNodeW := int(math.Max(float64(dLeftW), math.Max(float64(dRightW), float64(dValW/2+1))))
	// w represents the width of the drawer to return and is equal to maxChildW*2+1 to keep dVal in the center
	// and give the same space to each child
	w := maxNodeW*2 + 1
	// maxChildH is the max between the height of the left child and the height of the right child
	maxChildH := int(math.Max(float64(dLeftH), float64(dRightH)))
	// h represents the height of the drawer to return and is equal to dValH + 3 (considering the box around and the pipe)
	// + maxChildH
	h := dValH + 3 + maxChildH

	// Allocating new drawer to return
	d, err := drawer.NewDrawer(w, h)
	if err != nil {
		log.Fatal(fmt.Errorf("error while allocating new drawer with two children: %v", err))
	}

	// Drawing dVal onto the drawer to return with x in (w-dValW)/2 to put dVal in the middle
	// (remember that drawer.DrawDrawer takes coordinates of top left corner)
	// and y in 1 (considering the box)
	err = d.DrawDrawer(dVal, (w-dValW)/2, 1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing val with two children: %v", err))
	}

	// Adding a box in the drawer to return, around where the dVal drawer has been drawn
	// start coordinates are taken considering d.DrawDrawer above - 1 in order to not overwrite
	// end coordinates are just start coordinates plus respectively dValW+1 and dValH+1 in order to not overwrite
	err = addBoxAround(d, (w-dValW)/2-1, 0, (w-dValW)/2+dValW, dValH+1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while adding box with two children: %v", err))
	}

	// Drawing the pipe onto the drawer to return with x in the middle
	// and y in dValH + 2 (considering the box)
	err = d.DrawRune('┴', w/2, dValH+2)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ┴ rune with two childern: %v", err))
	}
	// Drawing the pipe onto the drawer with a loop
	for i := 1; i <= maxNodeW/2; i++ {
		err = d.DrawRune('─', w/2-i, dValH+2)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ─ rune with two childern with negative i: %v", err))
		}
		err = d.DrawRune('─', w/2+i, dValH+2)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ─ rune with two childern with positive i: %v", err))
		}
	}
	// Drawing left and right end of pipe, respectively at the middle of each child area
	err = d.DrawRune('╭', w/2-maxNodeW/2-1, dValH+2)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ╭ rune with two children: %v", err))
	}
	err = d.DrawRune('╮', w/2+maxNodeW/2+1, dValH+2)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ╮ rune with two children: %v", err))
	}

	// Drawing dLeft and dRight onto the drawer to return, each one is centered in its child area
	err = d.DrawDrawer(dLeft, (maxNodeW-dLeftW)/2, dValH+3)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing left child: %v", err))
	}
	// The weird expression that you see as x coordinate fixes an alignment problem
	// preferring the position one unit on the right where there is no perfect center
	err = d.DrawDrawer(dRight, maxNodeW*2+1-(maxNodeW-dRightW)/2-dRightW, dValH+3)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing right child: %v", err))
	}

	// Drawing links, this must be the latest things to be drawn because they overwrite dVal, dLeft and dRight
	err = d.DrawRune('┬', w/2, dValH+1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ┬ link with two children: %v", err))
	}
	err = d.DrawRune('┴', w/2-maxNodeW/2-1, dValH+3)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing left ┴ link with two children: %v", err))
	}
	err = d.DrawRune('┴', w/2+maxNodeW/2+1, dValH+3)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing right ┴ link with two children: %v", err))
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
