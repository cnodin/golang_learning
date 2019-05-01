package words

import "strings"

/**
统计字符数量
 */
func CountWords(text string) (count int)  {
	count = len(strings.Fields(text))
	return
}