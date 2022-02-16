package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	s1 := "19981025Qweasd@"
	//  计算16字节的校验码，校验码使用16进制
	checkSum := md5.Sum([]byte(s1))
	fmt.Println(hex.EncodeToString(checkSum[:])) //  719fef03d0453f6f799fbaa18e83c85d
	//  获取对应的二进制
	fmt.Printf("%b", checkSum) //  [1110001 10011111 11101111 11 11010000 1000101 111111 1101111 1111001 10011111 10111010 10100001 10001110 10000011 11001000 1011101]

}
