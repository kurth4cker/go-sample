package sample

func SortedMerge(s1, s2 []int) []int {
	slice := make([]int, 0, len(s1)+len(s2))

	var i1, i2 int
	for {
		switch {
		case i1 == len(s1):
			slice = append(slice, s2[i2:]...)
			return slice
		case i2 == len(s2):
			slice = append(slice, s1[i1:]...)
			return slice
		default:
			if s1[i1] < s2[i2] {
				slice = append(slice, s1[i1])
				i1++
			} else {
				slice = append(slice, s2[i2])
				i2++
			}
		}
	}
}
