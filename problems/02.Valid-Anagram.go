package leetcode

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var m [26]int
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
