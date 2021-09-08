package myint

import (
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAdd(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Calculate with Add function on MyInt", t, func() {
		var (
			i   MyInt = 0
			res MyInt
			err error
		)

		Convey("1+1 should be 2", func() {
			i = 1
			res, err = i.Add(1)
			So(res, ShouldEqual, 2)
			So(err, ShouldBeNil)
		})

		Convey("2-1 should be 1", func() {
			i = 2
			res, err = i.Add(-1)
			So(res, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("0+math.MaxInt32 should be math.MaxInt32", func() {
			i = 0
			res, err = i.Add(math.MaxInt32)
			So(res, ShouldEqual, math.MaxInt32)
			So(err, ShouldBeNil)
		})

		Convey("0+math.MinInt32 should be math.MinInt32", func() {
			i = 0
			res, err = i.Add(math.MinInt32)
			So(res, ShouldEqual, math.MinInt32)
			So(err, ShouldBeNil)
		})

		Convey("1+math.MaxInt32 should be -1", func() {
			i = 1
			res, err = i.Add(math.MaxInt32)
			So(res, ShouldEqual, -1)
			So(err, ShouldEqual, ErrOutOfRange)
		})

		Convey("math.MinInt32+math.MaxInt32 should be -1", func() {
			i = math.MinInt32
			res, err = i.Add(math.MaxInt32)
			So(res, ShouldEqual, -1)
			So(err, ShouldBeNil)
		})

		Convey("should be zero value", func() {
			var a int = 0
			So(a, ShouldBeZeroValue)
		})

	})
}
