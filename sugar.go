package utils

// If 三元表达式的写法
// 例如： max := If(a > b, a, b).(int)
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}
