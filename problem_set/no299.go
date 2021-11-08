package problemset

import "fmt"

func getHint(secret string, guess string) string {
	bulls := 0
	cows := 0
	secretMap := make(map[uint8]int)
	guessMap := make(map[uint8]int)
	n := len(secret)
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretMap[secret[i]]++
			guessMap[guess[i]]++
		}
	}
	for k, v1 := range secretMap {
		v2 := guessMap[k]
		if v2 > v1 {
			cows += v1
		} else {
			cows += v2
		}
	}
	return fmt.Sprintf("%dA%dB", bulls, cows)
}
