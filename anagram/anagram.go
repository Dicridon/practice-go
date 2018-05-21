package anagram

import "strings"

func FindAnagrams(dictionary []string, word string) (result []string) {
	if len(word) == 0 {
		return nil
	}

	word = normalize(word)
	charHash := make(map[rune]int)
	for _, c := range word {
		charHash[c]++
	}
	for _, v := range dictionary {
		refv := normalize(v)
		if refv == word || len(refv) == 0 {
			continue
		}
		if compareDict(toHash(refv), charHash) {
			result = append(result, v)
		}
	}
	return
}

func normalize(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "", -1)
}

func toHash(s string) (res map[rune]int) {
	res = make(map[rune]int)
	for _, c := range s {
		res[c]++
	}
	return
}

func compareDict(d1 map[rune]int, d2 map[rune]int) bool {
	if len(d1) != len(d2) {
		return false
	}
	for k, v := range d1 {
		if v != d2[k] {
			return false
		}
	}
	return true
}
