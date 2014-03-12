package Bugs

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestForCorrectness(t *testing.T) {
	Convey("Output should be ok", t, func() {
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		So(expected, ShouldResemble, WriteNumbers())
	})

	Convey("Should be able to reverse", t, func() {
		input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
		expected := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		Reverse(input)
		So(expected, ShouldResemble, input)
	})
}
