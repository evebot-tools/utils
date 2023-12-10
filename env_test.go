package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetEnv(t *testing.T) {
	Convey("TestGetEnv", t, func() {
		t.Setenv("FOUND", "found")
		Convey("Should return fallback value if key is not found", func() {
			So(GetEnv("NOT_FOUND", "fallback"), ShouldEqual, "fallback")
		})
		Convey("Should return value if key is found", func() {
			So(GetEnv("FOUND", "fallback"), ShouldEqual, "found")
		})
	})
}
