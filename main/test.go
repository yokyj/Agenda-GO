package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"user"
)

const tempFilePath = "temp.txt"

func main() {
	//input := "123321"
	//fmt.Println("111")
	fmt.Print(user.IsLogin())
}

func hashFunc(hashString string) string {
	h := md5.New()
	h.Write([]byte(hashString)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
