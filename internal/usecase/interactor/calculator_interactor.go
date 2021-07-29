package interactor

type CalculatorInteractor interface {
	Add(a1, a2 int) int
}

type calculatorInteractor struct {
}

func NewCalculatorInteractor() CalculatorInteractor {
	return &calculatorInteractor{}
}

func (ci *calculatorInteractor) Add(a1, a2 int) int {
	return a1 + a2
}
