/*
Package jdn is the tool of Julian Day Number.

Example:
	package main

	import (
		"fmt"
		"time"

		"github.com/DestinyLab/go-jdn"
	)

	func main() {
		t1 := time.Date(2018, 4, 13, 0, 0, 0, 0, time.UTC)
		fmt.Printf("%v", jdn.ToNumber(t1))
		// Output: 2458222

		t2 := JDN(2458222)
		loc, _ := time.LoadLocation()
		fmt.Printf("%s", t2.ToTime(loc))
		// Output: 2018-04-13 00:00:00 +0000 UTC
	}
*/
package jdn

import (
	"time"
)

// JDN is Julian Day Number
type JDN int64

var (
	startGregorian    = time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC)
	startGregorianJDN = JDN(2299161)
)

// ToNumber from time to JDN
func ToNumber(t time.Time) JDN {
	if t.Sub(time.Unix(0, 0)) >= 0 {
		return dnUnix(t)
	}

	if t.Sub(startGregorian) < 0 {
		return dnJulian(t)
	}

	return dnGregorian(t)
}

func dnGregorian(t time.Time) JDN {
	y := int64(t.Year())
	m := int64(t.Month())
	d := int64(t.Day())

	return JDN((1461*(y+4800+(m-14)/12))/4 + (367*(m-2-12*((m-14)/12)))/12 - (3*((y+4900+(m-14)/12)/100))/4 + d - 32075)
}

func dnJulian(t time.Time) JDN {
	y := int64(t.Year())
	m := int64(t.Month())
	d := int64(t.Day())

	return JDN(367*y - (7*(y+5001+(m-9)/7))/4 + (275*m)/9 + d + 1729777)
}

func dnUnix(t time.Time) JDN {
	return JDN(t.Unix()/86400 + 2440588)
}

// ToTime from JDN to time in UTC
func (jdn JDN) ToTime() time.Time {
	f := jdn + 1401

	if jdn >= startGregorianJDN {
		f += (((4*jdn+274277)/146097)*3)/4 - 38
	}

	e := 4*f + 3
	g := modInt(int64(e), 1461) / 4
	h := 5*g + 2
	d := int((modInt(int64(h), 153))/5 + 1)
	m := time.Month(modInt(int64(h/153+2), 12) + 1)
	y := int(e/1461 - 4716 + (12+2-JDN(m))/12)

	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
}

func modInt(c, m int64) int64 {
	r := c % m
	if r < 0 {
		r += m
	}

	return r
}
