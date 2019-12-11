package calculate

import "time"

import "encoding/json"

type Response struct {
	Product int
}

func Calculate(a, b int) string {
	valueA := make(chan int, 0)
	valueB := make(chan int, 0)

	go func() {
		time.Sleep(time.Second * 5)
		valueA <- calculateA(a)
	}()

	go func() {
		time.Sleep(time.Second * 5)
		valueB <- calculateB(b)
	}()

	x := <-valueA
	y := <-valueB
	resp, _ := json.Marshal(Response{x * y})
	return string(resp)
}

func calculateA(a int) int {
	var factorialA = a
	n := 1
	if a == 0 {
		return 1
	}

	for n < a {
		factorialA = factorialA * (a - n)
		n++
	}
	return factorialA
}

func calculateB(b int) int {
	var factorialB = b
	n := 1

	if b == 0 {
		return 1
	}

	for n < b {
		factorialB = factorialB * (b - n)
		n++
	}
	return factorialB
}
