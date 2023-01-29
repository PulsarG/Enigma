package enigma

import (
	"fmt"
)

func StartCrypt(text, key string) (string, bool) {
	// Проверка наличия ключа
	if key == "" {
		return "Key is empty", false
	}
	// Проверка на длину ключа
	short := checkLenKey(key)
	if !short {
		return "Sorry, but Key is Very Long", false
	}

	signArr := convertStringToArray(text)
	key1, key2, key3 := getKeyWordData(key)
	Rotor1, Rotor2, Rotor3 := replacing(key1, key2, key3)

	// Проверка на недопустимые символы
	var s string
	isWrong, wrongSing := checkWrongSing(signArr)
	if !isWrong {
		s = crypting(signArr, Rotor1, Rotor2, Rotor3)
	} else {
		return "Not valid character: " + wrongSing, false
	}

	return s, true
}

// Строка преобразуется в массив знаков
func convertStringToArray(s string) []string {
	var signArr []string
	for _, c := range s {
		signArr = append(signArr, fmt.Sprintf("%c", c))
	}
	return signArr
}

// Из ключа получаем три инта
func getKeyWordData(key string) (int, int, int) {
	var x, y, z int

	keyArr := convertStringToArray(key)

	x, y = findNumberFromKey(keyArr)
	z = len(keyArr)

	return x, y, z
}

func findNumberFromKey(arr []string) (int, int) {
	firstLileral := arr[0]
	lastLiteral := arr[(len(arr) - 1)]

	var numFirstLiteral, numLastLiteral int

	for i := 0; i < len(signsArray); i++ {
		if signsArray[i] == firstLileral {
			numFirstLiteral = i
		} else {
			continue
		}
	}

	for i := 0; i < len(signsArray); i++ {
		if signsArray[i] == lastLiteral {
			numLastLiteral = i
		} else {
			continue
		}
	}

	return numFirstLiteral, numLastLiteral
}

// Перемешивание:
func replacing(key1, key2, key3 int) ([162]string, [162]int, [162]int) {
	// Перемешиваем первым интом Базовый набор - ротор1
	newBaseSign := replaceSignArray(key1)
	// Перемешиваем вторым интом второй Ротор2
	newRotor2 := replaceRotor(key2)
	// Перемешиваем третьим интом третий Ротор3
	newRotor3 := replaceRotor(key3)

	return newBaseSign, newRotor2, newRotor3
}

// Для каждого знака из массива строки:
func crypting(signArr []string, Rotor1 [162]string, Rotor2, Rotor3 [162]int) string {
	var allResult string
	for i := 0; i < len(signArr); i++ {
		// Поиск Первого Индекса по знаку = х
		indexForRotor2 := findIndexInFirstRotor(signArr[i], Rotor1)
		// Поиск Второго Индекса по значению х = у
		indexForRotor3 := findIndexInRotor(indexForRotor2, Rotor2)
		// Поиск Третьего Индекса по значению у = z
		indexForMirror := findIndexInRotor(indexForRotor3, Rotor3)
		// Поиск Индекса в Отражателе по z = Хх
		indexFromMirror := mirror[indexForMirror]
		// Поиск значения в Роторе3 2 1 по индексу Хх = Уу
		result := Rotor1[Rotor2[Rotor3[indexFromMirror]]]
		allResult += result
	}
	return allResult
}

func findIndexInFirstRotor(s string, rotor [162]string) int {
	var index int
	for i := 0; i < len(rotor); i++ {
		if rotor[i] == s {
			index = i
			break
		} else {
			continue
		}
	}
	return index
}

func findIndexInRotor(s int, rotor [162]int) int {
	var index int
	for i := 0; i < len(rotor); i++ {
		if rotor[i] == s {
			index = i
			break
		} else {
			continue
		}
	}
	return index
}

func checkLenKey(key string) bool {
	keyArr := convertStringToArray(key)

	if len(keyArr) > len(signsArray) {
		return false
	} else {
		return true
	}
}

func checkWrongSing(arr []string) (bool, string) {
	check := false
	for indexOfEnterText := 0; indexOfEnterText < len(arr); indexOfEnterText++ {
		for indexOfBaseSingarr := 0; indexOfBaseSingarr < len(signsArray); indexOfBaseSingarr++ {
			if arr[indexOfEnterText] != signsArray[indexOfBaseSingarr] {
				check = true
			} else {
				check = false
				break
			}
		}
		if check {
			return check, arr[indexOfEnterText]
		}
	}
	return check, ""
}
