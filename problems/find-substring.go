package problems

func FindSubstring(str, sub string) bool {

	sl := len(str)
	ss := len(sub)

	if sl < ss {
		return false
	}

	for i := 0; i <= sl-ss; i++ {
		j := 0
		for ; j < ss; j++ {
			if str[i+j] != sub[j] {
				break
			}

		}

		if j == ss {
			return true
		}
	}

	return false

}
