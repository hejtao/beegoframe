package json

import (
	"bytes"
	"regexp"
)

const timeZero = `"0001-01-01 00:00:00"`

var matchCamelKey = regexp.MustCompile(`"[A-Z][A-Za-z0-9]*":`)

var matchTime = regexp.MustCompile(`"[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9.]*[+][0-9]{2}:[0-9]{2}"`)

func FormatCamelKey(b []byte) []byte {
	return matchCamelKey.ReplaceAllFunc(b, func(match []byte) []byte {
		n := len(match)
		buffer := bytes.Buffer{}
		for k, v := range match {
			if k == 1 {
				buffer.Write([]byte{v + 'a' - 'A'})
				continue
			}

			// 处理大写字母
			if k > 1 && k <= n-3 && v >= 'A' && v <= 'Z' {
				v2 := match[k+1]
				if v2 >= 'a' && v2 <= 'z' {
					buffer.Write([]byte{'_', v + 'a' - 'A'})
					continue
				}
				buffer.Write([]byte{v + 'a' - 'A'})
				continue
			}

			// 处理数字
			if k > 1 && k <= n-3 && v >= '0' && v <= '9' {
				v2 := match[k-1]
				if v2 < '0' || v2 > '9' {
					buffer.Write([]byte{'_'})
				}
				buffer.Write([]byte{v})
				v2 = match[k+1]
				if k <= n-4 && (v2 < '0' || v2 > '9') {
					buffer.Write([]byte{'_'})
				}
				continue
			}
			buffer.Write([]byte{v})
		}
		return buffer.Bytes()
	})
}

func FormatTime(b []byte) []byte {
	return matchTime.ReplaceAllFunc(b, func(match []byte) []byte {
		buffer := bytes.Buffer{}
		buffer.Write(match[:11])
		buffer.Write([]byte(" "))
		buffer.Write(match[12:20])
		buffer.Write([]byte(`"`))
		if buffer.String() == timeZero {
			return []byte(`""`)
		}
		return buffer.Bytes()
	})
}
