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
		valueA <- calculateFactorial(a)
	}()

	go func() {
		time.Sleep(time.Second * 5)
		valueB <- calculateFactorial(b)
	}()

	x := <-valueA
	y := <-valueB
	resp, _ := json.Marshal(Response{x * y})
	return string(resp)
}

func calculateFactorial(a int) int {
	var factorialA = a
	n := 1
	if a == 0 || a == 1 {
		return 1
	}

	for n < a {
		factorialA = factorialA * (a - n)
		n++
	}
	return factorialA
}
