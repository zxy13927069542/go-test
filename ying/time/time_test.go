package time

import (
	"fmt"
	"testing"
	"time"
)

//
// TestBeforeAndAfter
//  @Description: 时间前后比较
//  @param t
//
func TestBeforeAndAfter(t *testing.T) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}

	strTime, err := time.ParseInLocation("2006-01-02 15:04", "2020-12-21 11:33", location)
	if err != nil {
		fmt.Println(err)
		return
	}

	after := time.Now().After(strTime)
	fmt.Println(after)

	before := time.Now().Before(strTime)
	fmt.Println(before)
}
