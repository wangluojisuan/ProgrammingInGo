package main

import "fmt"
import "errors"

type binaryOperation func(operand1 int, operand2 int) (result int, err error)

func operate(op1 int, op2 int, bop binaryOperation) (result int, err error) {
	if bop == nil {
		err = errors.New("invalid binary operation function")
		return
	}
	
	return bop(op1, op2)
}

func divide(dividend int, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("division by zero")
		return
	}
	
	result = dividend / divisor
	return
}

func add(op1 int, op2 int) (result int, err error) {
	result = op1+op2
	return
}

func add1(op int) (result int, err error) {
	result = op
	return
}

func add2(o1 int, o2 int) (int, error) {
	return o1+o2, nil
}

func main() {
	result, err := operate(10, 0, divide)	
	if err == nil {
		fmt.Println("result=", result)
	} else {
		fmt.Println(err)
	}
	
	result, err = operate(10, 2, add)
	if err == nil {
		fmt.Println("result=", result)
	} else {
		fmt.Println(err)
	}
}