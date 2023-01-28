package enigma

func replaceSignArray(n int) [162]string {
	var newSignArr [162]string

	for i := 0; i < len(SignsArray); i++ {
		sign := SignsArray[i]
		if i < n {
			j := i
			j += len(SignsArray)
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

	for i := 0; i < len(Mirror); i++ {
		sign := Mirror[i]
		if i < n {
			j := i
			j += len(Mirror)
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
