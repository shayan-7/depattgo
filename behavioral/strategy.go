package behavioral

type Operation interface {
	Operate(n1, n2 int) int
}

type Calculator struct {
	Operation

	Num1 int
	Num2 int
}

func NewCalculator(n1, n2 int, operation Operation) *Calculator {
	return &Calculator{operation, n1, n2}
}

func (c *Calculator) SetOperation(operation Operation) {
	c.Operation = operation
}

func (c *Calculator) Calculate() int {
	return c.Operate(c.Num1, c.Num2)
}

type Addition struct{}

func (a *Addition) Operate(n1, n2 int) int {
	return n1 + n2
}

type Substraction struct{}

func (s *Substraction) Operate(n1, n2 int) int {
	return n1 - n2
}
