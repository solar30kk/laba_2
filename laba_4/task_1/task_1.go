package main

import (
	"fmt"
	"math"
)

func f(x float64) float64 {
	if x <= 0 {
		return math.NaN()
	}
	return math.Pow(x-6, 2) + math.Log(x) - 5
}

func df(x float64) float64 {
	if x <= 0 {
		return math.NaN()
	}
	return 2*(x-6) + 1.0/x
}

func phi(x float64, usePlus bool) float64 {
	if x <= 0 {
		return math.NaN()
	}
	val := 5 - math.Log(x)
	if val < 0 {
		return math.NaN()
	}
	if usePlus {
		return 6 + math.Sqrt(val)
	}
	return 6 - math.Sqrt(val)
}

func bisectionMethod(a, b, eps float64) {
	var aVec, bVec, diffVec []float64

	for math.Abs(b-a) > eps {
		aVec = append(aVec, a)
		bVec = append(bVec, b)
		diffVec = append(diffVec, math.Abs(b-a))

		c := (a + b) / 2.0
		if f(a)*f(c) < 0 {
			b = c
		} else {
			a = c
		}
	}

	fmt.Printf("%-3s %-12s %-12s %-12s\n", "N", "an", "bn", "bn-an")
	for i, an := range aVec {
		fmt.Printf("%-3d %-12.6f %-12.6f %-12.6f\n", i, an, bVec[i], diffVec[i])
	}
	fmt.Printf("Корень: %.6f\n\n", (a+b)/2.0)
}

func newtonMethod(x0, eps float64) {
	var xnVec, xnextVec, diffVec []float64
	x := x0

	for i := 0; i < 100; i++ {
		fx := f(x)
		dfx := df(x)
		if math.Abs(dfx) < 1e-10 {
			break
		}

		xNext := x - fx/dfx
		diff := math.Abs(xNext - x)

		xnVec = append(xnVec, x)
		xnextVec = append(xnextVec, xNext)
		diffVec = append(diffVec, diff)

		x = xNext
		if diff < eps {
			break
		}
	}

	fmt.Printf("%-3s %-12s %-12s %-12s\n", "N", "xn", "xn+1", "|xn+1-xn|")
	for i, xn := range xnVec {
		fmt.Printf("%-3d %-12.6f %-12.6f %-12.6f\n", i, xn, xnextVec[i], diffVec[i])
	}
	fmt.Printf("Корень: %.6f\n\n", x)
}

func simpleIterationMethod(x0, eps float64, usePlus bool) {
	var xnVec, xnextVec, diffVec []float64
	x := x0

	for i := 0; i < 100; i++ {
		xNext := phi(x, usePlus)
		if math.IsNaN(xNext) {
			break
		}

		diff := math.Abs(xNext - x)
		xnVec = append(xnVec, x)
		xnextVec = append(xnextVec, xNext)
		diffVec = append(diffVec, diff)

		x = xNext
		if diff < eps {
			break
		}
	}

	fmt.Printf("%-3s %-12s %-12s %-12s\n", "N", "xn", "xn+1", "|xn+1-xn|")
	for i, xn := range xnVec {
		fmt.Printf("%-3d %-12.6f %-12.6f %-12.6f\n", i, xn, xnextVec[i], diffVec[i])
	}
	fmt.Printf("Корень: %.6f\n\n", x)
}

func main() {
	eps := 1e-4

	fmt.Println("ПЕРВЫЙ КОРЕНЬ (интервал [4, 5])")

	fmt.Println("Метод половинного деления")
	bisectionMethod(4.0, 5.0, eps)

	fmt.Println("Метод Ньютона")
	newtonMethod(4.5, eps)

	fmt.Println("Метод простых итераций (ветвь с минусом)")
	simpleIterationMethod(4.5, eps, false)

	fmt.Println("ВТОРОЙ КОРЕНЬ (интервал [7, 8])")

	fmt.Println("Метод половинного деления")
	bisectionMethod(7.0, 8.0, eps)

	fmt.Println("Метод Ньютона")
	newtonMethod(7.5, eps)

	fmt.Println("Метод простых итераций (ветвь с плюсом)")
	simpleIterationMethod(7.5, eps, true)
}
