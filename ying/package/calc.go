package calc

//定义四种基本运算函数以供其他包调用，因此函数首字母大写
//加 a + b = rtValue
func Add(a, b int) (rtVale int) {
	rtVale = a + b
	return
}

//减 a - b = rtValue
func Reduce(a, b int) (rtValue int) {
	rtValue = a - b
	return
}

//乘 a * b = rtValue
func Multiply(a, b int) (rtValue int) {
	rtValue = a * b
	return
}

//除 a / b = rtValue
func divide(a, b int) (rtValue int) {
	rtValue = a / b
	return
}
