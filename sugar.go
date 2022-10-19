package utils

// 三元表达式的写法
// 例如： max := If(a > b, a, b).(int)
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
