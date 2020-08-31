package drawer

import (
	"fmt"
	"testing"
)

func TestNewDrawer(t *testing.T) {
	_, err := NewDrawer(-1, 1)
	if err == nil {
		t.Errorf("you shouldn't be able to create a drawer wiht negative width")
	}
	_, err = NewDrawer(1, -1)
	if err == nil {
		t.Errorf("you shouldn't be able to create a drawer wiht negative height")
	}
	_, err = NewDrawer(1, 5)
	if err != nil {
		t.Errorf("you should be able to create a drawer with positive height and width: %v", err)
	}
}

func TestDrawRune(t *testing.T) {
	d, err := NewDrawer(10, 1)
	if err != nil {
		t.Errorf("you should be able to create a drawer with positive height and width: %v", err)
	}
	err = d.DrawRune('🤨', 0, 0)
	if err != nil {
		t.Errorf("0 0 should be a valid place for drawing a rune in this drawer: %v", err)
	}
	err = d.DrawRune('🤪', 2, -1)
	if err == nil {
		t.Errorf("2 -1 shouldn't be a valid place for drawing a rune in this drawer: %v", err)
	}
	err = d.DrawRune('🤪', 10, 0)
	if err == nil {
		t.Errorf("10 0 shouldn't be a valid place for drawing a rune in this drawer: %v", err)
	}
	err = d.DrawRune('🤪', 2, 1)
	if err == nil {
		t.Errorf("2 1 shouldn't be a valid place for drawing a rune in this drawer: %v", err)
	}
	err = d.DrawRune('🥶', 9, 0)
	if err != nil {
		t.Errorf("9 0 should be a valid place for drawing a rune in this drawer: %v", err)
	}

	fmt.Println(d)

	// Output:
	// 🤨        🥶
}

func TestDrawDrawer(t *testing.T) {
	d, err := NewDrawer(10, 10)
	if err != nil {
		t.Errorf("you should be able to create a drawer with positive height and width: %v", err)
	}
	dToDraw, err := NewDrawer(10, 11)
	if err != nil {
		t.Errorf("you should be able to create a drawer to draw with 10 and 11 as height and width: %v", err)
	}
	err = d.DrawDrawer(dToDraw, 0, 0)
	if err == nil {
		t.Errorf("you shouldn't be able to draw a drawer onto a smaller drawer")
	}
	dToDraw, err = NewDrawer(4, 4)
	if err != nil {
		t.Errorf("you should be able to create a drawer to draw with 4 and 4 as height and width: %v", err)
	}
	err = d.DrawDrawer(dToDraw, -1, 0)
	if err == nil {
		t.Errorf("you shouldn't be able to draw a drawer with negative x coordinate in upper left corner")
	}
	err = d.DrawDrawer(dToDraw, 0, -1)
	if err == nil {
		t.Errorf("you shouldn't be able to draw a drawer with negative y coordinate in upper left corner")
	}
	err = d.DrawDrawer(dToDraw, 7, 0)
	if err == nil {
		t.Errorf("there should be no place to draw dToDraw at coordinate 7 0")
	}
	err = d.DrawDrawer(dToDraw, 0, 7)
	if err == nil {
		t.Errorf("there should be no place to draw dToDraw at coordinate 0 7")
	}
	err = dToDraw.DrawRune('🖖', 2, 2)
	if err != nil {
		t.Errorf("2 2 should be a valid place to draw a rune in dToDraw drawer: %v", err)
	}
	err = d.DrawDrawer(dToDraw, 0, 1)
	if err != nil {
		t.Errorf("0 1 should be a valid place to draw dToDraw onto d: %v", err)
	}

	fmt.Println(d)

	// Output:
	//
	//
	//
	//  🖖
	//
	//
	//
	//
	//
	//
}
