package validate

import "regexp"

// 数字
func IsNumber(s string) bool {
	pattern := `^\d+(\.\d+)?$`
	match, err := regexp.MatchString(pattern, s)
	if err != nil {
		return false
	}
	return match
}

// 罗马数字
func IsRomanNumber(s string) bool {
	re := regexp.MustCompile("^[IVXLCDM]*$")
	return re.MatchString(s)
}

// url
func IsURL(s string) bool {
	regex := regexp.MustCompile(`^(http://|https://)?[\w-]+(\.[\w-]+)+([\w.,@?^=%&:/~+#-]*[\w@?^=%&/~+#-])?$`)
	return regex.MatchString(s)
}

// 邮箱
func IsEmail(email string) bool {
	res, _ := regexp.MatchString(`^([\w\.\_\-]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	return res
}

// func IsEmail(s string) bool {
// 	emailRegex := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)
// 	return emailRegex.MatchString(s)
// }

func IsNotUnicode(s string) bool {
	regex := `^[^\p{L}\p{N}]*$`
	match, _ := regexp.MatchString(regex, s)
	return match
}
