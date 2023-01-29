package enigma

func replaceSignArray(n int) [162]string {
	var newSignArr [162]string

	for i := 0; i < len(signsArray); i++ {
		sign := signsArray[i]
		if i < n {
			j := i
			j += len(signsArray)
			j -= n
			newSignArr[j] = sign
		} else {
			j := i
			j -= n
			newSignArr[j] = sign
		}
	}

	return newSignArr
}

func replaceRotor(n int) [162]int {
	var newSignArr [162]int

	for i := 0; i < len(mirror); i++ {
		sign := mirror[i]
		if i < n {
			j := i
			j += len(mirror)
			j -= n
			newSignArr[j] = sign
		} else {
			j := i
			j -= n
			newSignArr[j] = sign
		}
	}

	return newSignArr
}
