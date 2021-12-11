package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
)

func main() {
	// a := 0.48
	// fmt.Printf("%T \n", a) //不指定类型则浮点型默认为float64

	// b := float32(0.32)
	// fmt.Printf("%T \n", b) //强制指定类型为float32

	// //将32位的b转换为64位,会失真，不能直接转
	// b := float32(0.42)
	// fmt.Printf("b的值为%v,类型为%T\n",b,b)  //0.42	float32
	// fmt.Println(b == 0.42)	//true
	// a := float64(b)
	// fmt.Printf("a的值为%v,类型为%T\n",a,a)	//0.42	float64
	// fmt.Println(a == 0.42)	//false

	//将64位的b转换位32位
	b := 0.4254111515151
	fmt.Printf("b的值为%v,类型为%T\n", b, b) //0.4254111515151	float64
	fmt.Println(b == 0.4254111515151)  //true
	a := float32(b)
	fmt.Printf("a的值为%v,类型为%T\n", a, a) //0.42541116	float32
	fmt.Println(a == 0.4254111515151)  //true 0.42541116 == 0.4254111515151 ?

}

//
//  RetainValidNumber
//  @Description: 保留有效数字
//  @param f
//  @param n 要保留的位数
//  @return string
//
func RetainValidNumber(f float64, n int) string {
	//获取整数
	trunc := int(math.Trunc(f))

	//计算要保留的小数位数
	length := len(strconv.Itoa(trunc))

	//四舍五入保留有效数字
	var rtValue float64
	if n <= length {
		return strconv.Itoa(trunc)[:n]
	} else {
		rtValue, _ = decimal.NewFromFloat(f).Round(int32(n - length)).Float64()

		//float64 -> string
		return strconv.FormatFloat(rtValue, 'f', n-length, 64)
	}
}
