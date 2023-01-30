package enigma

import (
	/* "fmt" */
	"math/rand"
	"time"
)

func NewRotor() ([162]int, bool) {

	shuffleArr()

	settingmirrorNumbers()

	return mirror, checkErrorsInArr()
}

func shuffleArr() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(mirror),
		func(i, j int) { mirror[i], mirror[j] = mirror[j], mirror[i] })
}

func settingmirrorNumbers() {
	for i := 0; i < len(mirror); i++ { // отзеркаливание значения и индекса
		x := mirror[i]
		y := mirror[x]

		if i != y { // если еще не зеркальны
			mirror[x] = i
			for l := 0; l < len(mirror); l++ {
				if l != x && mirror[l] == i { // удаление лишнего значения, с заменой на остаточное чилсо после отзеркаливания первых двух
					mirror[l] = y
					break
				}
			}
		}
	}
}

func checkErrorsInArr() bool {
	ok := true
	for i := 0; i < len(mirror); i++ {
		x := mirror[i]
		if mirror[x] != i {
			ok = false
			/* fmt.Println("!", i, x) */
		}
	}
	return ok
}
