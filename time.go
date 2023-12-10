package utils

import (
	"github.com/golang-module/carbon/v2"
	"time"
)

func CheckTTL(t time.Time, ttl int) bool {
	c := carbon.CreateFromStdTime(t)
	c.AddSeconds(ttl)
	if c.IsPast() {
		return false
	}
	return true
}
