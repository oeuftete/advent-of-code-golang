package utils

func MakeRange(imin int, imax int) []int {
	a := make([]int, imax-imin+1)
	for i := range a {
		a[i] = imin + i
	}
	return a
}
