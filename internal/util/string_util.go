package util

import (
	mathRand "math/rand"
)

const lowLetters = "abcdefghijklmnopqrstuvwxyz"

// RandLowStr 获取随机小写字母
func RandLowStr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = lowLetters[mathRand.Intn(len(lowLetters))]
	}
	return string(b)
}
