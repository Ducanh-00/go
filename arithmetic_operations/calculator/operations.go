package calculator

type Operation interface {
	PerformOperation() int
}

func (c *Calculator) PerformOperation() int {
	return c.Add()
}

func PerformCalculation(c *Calculator, operation func() int) int {
	return operation()
}

