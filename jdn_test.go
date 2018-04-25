package jdn

import (
	"fmt"
	"testing"
	"time"
)

var jdntests = []struct {
	t   time.Time
	jdn JDN
}{
	{time.Date(1985, 11, 11, 0, 0, 0, 0, time.UTC), 2446381},
	{time.Date(1999, 9, 21, 0, 0, 0, 0, time.UTC), 2451443},
	{time.Date(2010, 12, 1, 0, 0, 0, 0, time.UTC), 2455532},
	{time.Date(2011, 3, 11, 0, 0, 0, 0, time.UTC), 2455632},
	{time.Date(2014, 3, 18, 0, 0, 0, 0, time.UTC), 2456735},
	{time.Date(2014, 9, 26, 0, 0, 0, 0, time.UTC), 2456927},
	{time.Date(2017, 3, 28, 0, 0, 0, 0, time.UTC), 2457841},
	{time.Date(2017, 8, 15, 0, 0, 0, 0, time.UTC), 2457981},
	{time.Date(2017, 10, 1, 0, 0, 0, 0, time.UTC), 2458028},
	{time.Date(2018, 4, 13, 0, 0, 0, 0, time.UTC), 2458222},
}

func TestToNumber(t *testing.T) {
	for _, tt := range jdntests {
		t.Run(tt.t.String(), func(t *testing.T) {
			got := ToNumber(tt.t)
			if got != tt.jdn {
				t.Errorf("got %q, want %q", got, tt.jdn)
			}
		})
	}
}

func TestToTime(t *testing.T) {
	for _, tt := range jdntests {
		t.Run(tt.t.String(), func(t *testing.T) {
			got := tt.jdn.ToTime()
			if got != tt.t {
				t.Errorf("got %q, want %q", got, tt.t)
			}
		})
	}
}

func BenchmarkDnGregorian(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dnGregorian(startGregorian)
	}
}

func BenchmarkDnJulian(b *testing.B) {
	t := time.Date(-4712, 1, 1, 0, 0, 0, 0, time.UTC)
	for n := 0; n < b.N; n++ {
		dnJulian(t)
	}
}

func BenchmarkDnUnix(b *testing.B) {
	t := time.Date(2018, 4, 18, 0, 0, 0, 0, time.UTC)
	for n := 0; n < b.N; n++ {
		dnGregorian(t)
	}
}

func BenchmarkToTime(b *testing.B) {
	jdn := JDN(2299161)
	for n := 0; n < b.N; n++ {
		jdn.ToTime()
	}
}

func ExampleToNumber() {
	tt := time.Date(1990, 6, 3, 0, 0, 0, 0, time.UTC)
	jdn := ToNumber(tt)
	fmt.Println(jdn)
	// Output: 2448046
}

func ExampleJDN_ToTime() {
	t := JDN(2448046).ToTime()
	fmt.Printf("%v %v %v\n", t.Year(), t.Month(), t.Day())
	// Output: 1990 June 3
}
