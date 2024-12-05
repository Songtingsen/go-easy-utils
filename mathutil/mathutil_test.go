package mathutil

import (
	"fmt"
	"testing"
)

// 测试Exponent函数
func TestExponent(t *testing.T) {
	result := Exponent(2, 3)
	expected := int64(8)
	fmt.Printf("Exponent(2, 3) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("Exponent(2, 3) = %d; want %d", result, expected)
	}
}

// 测试Fibonacci函数
func TestFibonacci(t *testing.T) {
	result := Fibonacci(1, 1, 5)
	expected := 5
	fmt.Printf("Fibonacci(1, 1, 5) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("Fibonacci(1, 1, 5) = %d; want %d", result, expected)
	}
}

// 测试Factorial函数
func TestFactorial(t *testing.T) {
	result := Factorial(5)
	expected := uint(120)
	fmt.Printf("Factorial(5) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("Factorial(5) = %d; want %d", result, expected)
	}
}

// 测试Percent函数
func TestPercent(t *testing.T) {
	result := Percent(20, 100, 2)
	expected := 20.00
	fmt.Printf("Percent(20, 100, 2) 实际结果: %f, 预期结果: %f\n", result, expected)
	if result != expected {
		t.Errorf("Percent(20, 100, 2) = %f; want %f", result, expected)
	}
}

// 测试RoundToFloat函数
func TestRoundToFloat(t *testing.T) {
	result := RoundToFloat(3.1415926, 2)
	expected := 3.14
	fmt.Printf("RoundToFloat(3.1415926, 2) 实际结果: %f, 预期结果: %f\n", result, expected)
	if result != expected {
		t.Errorf("RoundToFloat(3.1415926, 2) = %f; want %f", result, expected)
	}
}

// 测试Max函数
func TestMax(t *testing.T) {
	numbers := []int{1, 5, 3}
	result := Max(numbers...)
	expected := 5
	fmt.Printf("Max([1, 5, 3]) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("Max([1, 5, 3]) = %d; want %d", result, expected)
	}
}

// 测试Min函数
func TestMin(t *testing.T) {
	numbers := []int{1, 5, 3}
	result := Min(numbers...)
	expected := 1
	fmt.Printf("Min([1, 5, 3]) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("Min([1, 5, 3]) = %d; want %d", result, expected)
	}
}

// 测试Sum函数
func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}
	result := Sum(numbers...)
	expected := 6
	fmt.Printf("Sum([1, 2, 3]) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("Sum([1, 2, 3]) = %d; want %d", result, expected)
	}
}

// 测试Average函数
func TestAverage(t *testing.T) {
	numbers := []int{1, 2, 3}
	result := Average(numbers...)
	expected := 2.0
	fmt.Printf("Average([1, 2, 3]) 实际结果: %f, 预期结果: %f\n", result, expected)
	if result != expected {
		t.Errorf("Average([1, 2, 3]) = %f; want %f", result, expected)
	}
}

// 测试IsPrime函数
func TestIsPrime(t *testing.T) {
	result := IsPrime(7)
	expected := true
	fmt.Printf("IsPrime(7) 实际结果: %t, 预期结果: %t\n", result, expected)
	if result != expected {
		t.Errorf("IsPrime(7) = %t; want %t", result, expected)
	}
	result = IsPrime(4)
	expected = false
	fmt.Printf("IsPrime(4) 实际结果: %t, 预期结果: %t\n", result, expected)
	if result != expected {
		t.Errorf("IsPrime(4) = %t; want %t", result, expected)
	}
}

// 测试GCD函数
func TestGCD(t *testing.T) {
	numbers := []int{12, 18}
	result := GCD(numbers...)
	expected := 6
	fmt.Printf("GCD([12, 18]) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("GCD([12, 18]) = %d; want %d", result, expected)
	}
}

// 测试LCM函数
func TestLCM(t *testing.T) {
	numbers := []int{4, 6}
	result := LCM(numbers...)
	expected := 12
	fmt.Printf("LCM([4, 6]) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("LCM([4, 6]) = %d; want %d", result, expected)
	}
}

// 测试Cos函数
func TestCos(t *testing.T) {
	result := Cos(0)
	expected := 1.0
	fmt.Printf("Cos(0) 实际结果: %f, 预期结果: %f\n", result, expected)
	if result != expected {
		t.Errorf("Cos(0) = %f; want %f", result, expected)
	}
}

// 测试Sin函数
func TestSin(t *testing.T) {
	result := Sin(0)
	expected := 0.0
	fmt.Printf("Sin(0) 实际结果: %f, 预期结果: %f\n", result, expected)
	if result != expected {
		t.Errorf("Sin(0) = %f; want %f", result, expected)
	}
}

// 测试Log函数
func TestLog(t *testing.T) {
	result := Log(100, 10)
	expected := 2.0
	fmt.Printf("Log(100, 10) 实际结果: %f, 预期结果: %f\n", result, expected)
	if result != expected {
		t.Errorf("Log(100, 10) = %f; want %f", result, expected)
	}
}

// 测试Abs函数
func TestAbs(t *testing.T) {
	var num int = -5
	result := Abs(num)
	expected := 5
	fmt.Printf("Abs(-5) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("Abs(-5) = %d; want %d", result, expected)
	}
}

// 测试Div函数
func TestDiv(t *testing.T) {
	var x int = 10
	var y int = 2
	result := Div(x, y)
	expected := 5.0
	fmt.Printf("Div(10, 2) 实际结果: %f, 预期结果: %f\n", result, expected)
	if result != expected {
		t.Errorf("Div(10, 2) = %f; want %f", result, expected)
	}
}

// 测试Variance函数
func TestVariance(t *testing.T) {
	numbers := []int{1, 2, 3}
	result := Variance(numbers)
	expected := 0.6666666666666666
	fmt.Printf("Variance([1, 2, 3]) 实际结果: %f, 预期结果: %f\n", result, expected)
	if result != expected {
		t.Errorf("Variance([1, 2, 3]) = %f; want %f", result, expected)
	}
}

// 测试StdDev函数
func TestStdDev(t *testing.T) {
	numbers := []int{1, 2, 3}
	result := StdDev(numbers)
	expected := 0.816496580927726
	fmt.Printf("StdDev([1, 2, 3]) 实际结果: %f, 预期结果: %f\n", result, expected)
	if result != expected {
		t.Errorf("StdDev([1, 2, 3]) = %f; want %f", result, expected)
	}
}

// 测试Permutation函数
func TestPermutation(t *testing.T) {
	result := Permutation(5, 3)
	expected := uint(60)
	fmt.Printf("Permutation(5, 3) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("Permutation(5, 3) = %d; want %d", result, expected)
	}
}

// 测试Combination函数
func TestCombination(t *testing.T) {
	result := Combination(5, 3)
	expected := uint(10)
	fmt.Printf("Combination(5, 3) 实际结果: %d, 预期结果: %d\n", result, expected)
	if result != expected {
		t.Errorf("Combination(5, 3) = %d; want %d", result, expected)
	}
}
