package models

type Operation struct {
	firstNumber  int
	secondNumber int
}

func New(firstNumber, secondNumber int) Operation {
	n := Operation{firstNumber, secondNumber}

	return n
}

func (n Operation) Multiplication() int{
	operationM := n.firstNumber * n.secondNumber

	return operationM
}

func (n Operation) Division() int{
	operationD := n.firstNumber / n.secondNumber

	return operationD
}

func (n Operation) Addition() int{
	operationA := n.firstNumber + n.secondNumber

	return operationA
}

func (n Operation) Subtraction() int{
	operationS := n.firstNumber - n.secondNumber

	return operationS
}