package clock

import "fmt"

type Clock struct {
	minutes int
}

// normalize ensures minutes is in range [0, 1439]
func normalize(m int) int {
	m = m % 1440
	if m < 0 {
		m += 1440
	}
	return m
}

func New(h, m int) Clock {
	return Clock{minutes: normalize(h*60 + m)}
}

func (c Clock) Add(m int) Clock {
	return Clock{minutes: normalize(c.minutes + m)}
}

func (c Clock) hour() int {
	return c.minutes / 60
}

func (c Clock) minute() int {
	return c.minutes % 60
}

func (c Clock) Subtract(m int) Clock {
	return Clock{minutes: normalize(c.minutes - m)}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour(), c.minute())
}
