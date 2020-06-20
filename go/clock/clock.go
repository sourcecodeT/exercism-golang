package clock

import "fmt"

const MINUSTES_IN_A_DAY = 24 * 60

func New(h, m int) Clock {
	m = (h*60 + m) % MINUSTES_IN_A_DAY
	if m < 0 {
		m += MINUSTES_IN_A_DAY
	}

	return Clock{m}
}

type Clock struct {
	m int
}

func (c Clock) Add(m int) Clock {
	return New(0, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, c.m-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.m/60, c.m%60)
}
