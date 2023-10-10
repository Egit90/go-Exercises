package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	clock := Clock{h, m}
	clock.Adjust()
	return clock

}

func (c *Clock) Adjust() {
	min := c.h*60 + c.m
	min %= 24 * 60
	min += 24 * 60
	c.m = min % 60
	c.h = (min / 60) % 24
}

func (c Clock) Add(m int) Clock {
	c.m += m
	c.Adjust()
	return c
}

func (c Clock) Subtract(m int) Clock {
	c.m -= m
	c.Adjust()
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%.02d:%.02d", c.h, c.m)
}
