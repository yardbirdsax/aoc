package main

func find(input string) (index int) {
	found := false
	start := 0
	index = 4
	l := len(input)
	for !found {
		if index > l {
			break
		}
		found = true
		s := input[start:index]
		var bs int
		for ir := range s {
			b := 0x1 << (s[ir] % 32)
			if bs&b == b {
				found = false
				break
			}
			bs |= b
		}
		if !found {
			start++
			index++
		}
	}
	if !found {
		index = -1
	}
	return index
}
