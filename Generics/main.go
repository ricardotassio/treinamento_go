package main

// essa interface pertence a terceira parte
type Number interface {
	~int | ~float64
}

// essa interface pertence a quarta parte
type MyNumber int

// Primeira parte
func SumInt(m map[string]int) int {
	var sum int
	for _, v := range m {
		sum += v
	}
	return sum
}

func SumFloat(m map[string]float64) float64 {
	var sum float64
	for _, v := range m {
		sum += v
	}
	return sum
}

// Segunda parte
func Sum[T int | float64](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

// Terceira parte usando Constraint
func SumWithConstraint[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

//quarta parte

func Compare[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Ricardo": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Ricardo": 100.20, "João": 2000.3, "Maria": 3000.0}
	println(SumInt(m))
	println(SumFloat(m2))
	println(Sum(m))
	println(Sum(m2))
	println(SumWithConstraint(m))
	println(SumWithConstraint(m2))
	m3 := map[string]MyNumber{"Ricardo": 100, "João": 200, "Maria": 300}
	println(SumWithConstraint(m3))
	println(Compare(10, 10))
}
