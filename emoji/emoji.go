package emoji

import (
	"regexp"
	"strconv"
	"strings"
)

type EmojiTool struct{}

// Emoji表情解码
func (s *EmojiTool) UnicodeEmojiDecode(str string) string {
	//emoji表情的数据表达式
	re := regexp.MustCompile("\\[[\\\\u0-9a-zA-Z]+\\]")
	//提取emoji数据表达式
	reg := regexp.MustCompile("\\[\\\\u|]")
	src := re.FindAllString(str, -1)
	for i := 0; i < len(src); i++ {
		e := reg.ReplaceAllString(src[i], "")
		p, err := strconv.ParseInt(e, 16, 32)
		if err == nil {
			str = strings.Replace(str, src[i], string(rune(p)), -1)
		}
	}
	return str
}

// Emoji表情转换
func (s *EmojiTool) UnicodeEmojiCode(str string) string {
	ret := ""
	rs := []rune(str)
	for i := 0; i < len(rs); i++ {
		if len(string(rs[i])) == 4 {
			u := `[\u` + strconv.FormatInt(int64(rs[i]), 16) + `]`
			ret += u

		} else {
			ret += string(rs[i])
		}
	}
	return ret
}
