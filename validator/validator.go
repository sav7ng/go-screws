package validator

import "regexp"

type ValidatorTool struct{}

// 手机号码检测
func (s *ValidatorTool) IsMobile(mobileNum string) bool {
	var regular = "^1[345789]{1}\\d{9}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

// 判断是否是18或15位身份证
func (s *ValidatorTool) IsIdCard(cardNo string) bool {
	//18位身份证 ^(\d{17})([0-9]|X)$
	if m, _ := regexp.MatchString(`(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)`, cardNo); !m {
		return false
	}
	return true
}

// 判断是否是True
func (s *ValidatorTool) IsTrue(value bool) bool {
	return value
}

// 判断是否是False
func (s *ValidatorTool) IsFalse(value bool) bool {
	return false == value
}
