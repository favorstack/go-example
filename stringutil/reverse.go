// stringutil包包含字符串处理相关的函数
package stringutil

// Reverse函数返回其参数字符串从左到右的反转形式
func Reverse(s string) string {
	r := []rune(s)
	length := len(r)
	for i, j := 0, length - 1; i < length/2; i, j = i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}
