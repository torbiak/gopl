// ex3.12 determines if strings are anagrams of each other.
package anagram

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	freq := make(map[rune]int)
	for _, c := range a {
		freq[c]++
	}
	for _, c := range b {
		freq[c]--
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}
