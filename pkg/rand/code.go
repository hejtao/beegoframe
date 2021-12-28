package rand

import (
	"math/rand"
	"time"
)

const (
	digitCase = "01234567890123456789"
	lowerCase = "abcdefghijklmnopqrstuvwxyz"
	upperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// 生成n位随机码, 可选择是否包含数字、大小写字母
func GenCode(n int, digit, lower, upper bool) (code string) {
	strRange := ""
	if digit {
		strRange += digitCase
	}

	if lower {
		strRange += lowerCase
	}

	if upper {
		strRange += upperCase
	}

	m := len(strRange)

	for i := 0; i < n; i++ {
		code += string(strRange[rand.Intn(m)])
	}

	return
}
