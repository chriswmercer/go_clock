//Package clock contains types and functions for the Clock assignment
package clock

import "fmt"

//Clock provides a sructure for defining a simple clock
type Clock struct {
	Hour int
	Minute int
}

//New creates a new instance of the Clock type
func New(hour, minute int) Clock {
	c := new(Clock)
	c.Hour = hour
	c.Minute = minute
	c.rationalise()
	return *c
}

//String stringifies a clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hour, c.Minute)
}

//Add adds the given number of minutes to a clock
func (c Clock) Add(minutes int) Clock {
	c.Minute += minutes
	c.rationalise()
	return c
}

//Subtract subtracts the given number of minutes from a clock
func (c Clock) Subtract(minutes int) Clock {
	c.Minute -= minutes
	c.rationalise()
	return c
}

//Equal will allow equality comparison
func (c Clock) Equal(d Clock) bool {
	return c.Hour == d.Hour && c.Minute == d.Minute
}

func (c *Clock) rationalise() {
	for c.Minute < 0 {
		c.Hour--
		c.Minute += 60
	}
	for c.Minute > 59 {
		c.Hour++
		c.Minute -= 60
	}
	for c.Hour < 0 {
		c.Hour += 24
	}
	for c.Hour >= 24 {
		c.Hour -= 24
	}
}
