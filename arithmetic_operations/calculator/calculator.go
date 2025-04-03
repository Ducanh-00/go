package calculator

type Calculator struct {
	A int
	B int
}

func (c *Calculator) Add() int {
	return c.A + c.B
}

func (c *Calculator) Subtract() int {
	return c.A - c.B
}

func (C *Calculator) Multiply() int {
	return C.A * C.B
}

func (c *Calculator) Divide() int {
	return c.A / c.B
}

func (c *Calculator) SetValues(a, b int) {
	c.A = a
	c.B = b
}
