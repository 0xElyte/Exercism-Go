package clock

import "fmt"

// Clock represents a time of day without a date.
// The minutes field is unexported to enforce the invariant that it always 
// holds a normalized value in the range [0, 1439].
type Clock struct {
	minutes int
}

// New creates a new Clock from the given hours and minutes.
func New(h, m int) Clock {
	return Clock{minutes: normalize(h, m)}
}

// Add adds minutes to the Clock and returns a new Clock.
func (c Clock) Add(m int) Clock {
	return Clock{minutes: normalize(0, c.minutes+m)}
}

// Subtract subtracts minutes from the Clock and returns a new Clock.
func (c Clock) Subtract(m int) Clock {
	return Clock{minutes: normalize(0, c.minutes-m)}
}

// String returns the formatted time as HH:MM.
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// normalize converts arbitrary hours and minutes into a total minute count
// strictly bounded within a single 24-hour day [0, 1439].
func normalize(h, m int) int {
	total := h*60 + m
	// The double modulo operation correctly handles negative values in Go,
	// mapping them to the positive range [0, 1439].
	return ((total % 1440) + 1440) % 1440
}