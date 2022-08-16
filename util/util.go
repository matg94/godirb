package util

func ListContains(val int, arr []int) bool {
	for _, v := range arr {
		if val == v {
			return true
		}
	}
	return false
}
