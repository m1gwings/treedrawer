package tree

import (
	"fmt"
	"log"
	"math"

	"github.com/m1gwings/treedrawer/drawer"
)

func stringify(t *Tree) *drawer.Drawer {
	dVal := t.val.Draw()
	dValW, dValH := dVal.Dimens()
	if t.left == nil && t.right == nil {
		d, err := drawer.NewDrawer(dValW+2, dValH+2)
		if err != nil {
			log.Fatal(fmt.Errorf("error while allocating new drawer with no children: %v", err))
		}
		err = d.DrawDrawer(dVal, 1, 1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing val with no children: %v", err))
		}
		err = d.DrawRune('╭', 0, 0)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ╭ with no children: %v", err))
		}
		err = d.DrawRune('╮', dValW+1, 0)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ╮ with no children: %v", err))
		}
		err = d.DrawRune('╰', 0, dValH+1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ╰ with no children: %v", err))
		}
		err = d.DrawRune('╯', dValW+1, dValH+1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ╯ with no children: %v", err))
		}
		for x := 1; x <= dValW; x++ {
			for yMul := 0; yMul <= 1; yMul++ {
				err = d.DrawRune('─', x, yMul*(dValH+1))
				if err != nil {
					log.Fatal(fmt.Errorf("error while drawing ─ with no children: %v", err))
				}
			}
		}
		for y := 1; y <= dValH; y++ {
			for xMul := 0; xMul <= 1; xMul++ {
				err = d.DrawRune('│', xMul*(dValW+1), y)
				if err != nil {
					log.Fatal(fmt.Errorf("error while drawing │ with no children: %v", err))
				}
			}
		}
		return d
	}

	if (t.left != nil && t.right == nil) || (t.left == nil && t.right != nil) {
		var dChild *drawer.Drawer
		if t.left != nil {
			dChild = stringify(t.left)
		} else {
			dChild = stringify(t.right)
		}
		dChildW, dChildH := dChild.Dimens()
		w := int(math.Max(float64(dValW+2), float64(dChildW)))
		h := dValH + 1 + dChildH
		d, err := drawer.NewDrawer(w, h)
		if err != nil {
			log.Fatal(fmt.Errorf("error while allocating new drawer with one child: %v", err))
		}
		err = d.DrawDrawer(dVal, (w-dValW)/2, 0)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing val with one child: %v", err))
		}
		err = d.DrawRune('│', w/2, dValH)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing | with one child: %v", err))
		}
		err = d.DrawDrawer(dChild, (w-dChildW)/2, dValH+1)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing child drawer with one child: %v", err))
		}
		return d
	}

	dLeft, dRight := stringify(t.left), stringify(t.right)
	dLeftW, dLeftH := dLeft.Dimens()
	dRightW, dRightH := dRight.Dimens()
	maxChildW := int(math.Max(float64(dLeftW), float64(dRightW)))
	w := maxChildW*2 + 1
	maxChildH := int(math.Max(float64(dLeftH), float64(dRightH)))
	h := dValH + 1 + maxChildH
	d, err := drawer.NewDrawer(w, h)
	if err != nil {
		log.Fatal(fmt.Errorf("error while allocating new drawer with two children: %v", err))
	}
	err = d.DrawDrawer(dVal, (w-dValW)/2, 0)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing val with two children: %v", err))
	}
	err = d.DrawRune('┴', w/2, dValH)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ┴ rune with two childern: %v", err))
	}
	for i := 1; i <= maxChildW/2; i++ {
		err = d.DrawRune('─', w/2-i, dValH)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ─ rune with two childern with negative i: %v", err))
		}
		err = d.DrawRune('─', w/2+i, dValH)
		if err != nil {
			log.Fatal(fmt.Errorf("error while drawing ─ rune with two childern with positive i: %v", err))
		}
	}
	err = d.DrawRune('╭', w/2-maxChildW/2-1, dValH)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ╭ rune with two children: %v", err))
	}
	err = d.DrawRune('╮', w/2+maxChildW/2+1, dValH)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing ╮ rune with two children: %v", err))
	}
	err = d.DrawDrawer(dLeft, (maxChildW-dLeftW)/2, dValH+1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing left child: %v", err))
	}
	err = d.DrawDrawer(dRight, maxChildW+1+(maxChildW-dRightW)/2, dValH+1)
	if err != nil {
		log.Fatal(fmt.Errorf("error while drawing right child: %v", err))
	}
	return d
}
