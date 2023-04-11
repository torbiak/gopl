// ex3.12 determines if strings are anagrams of each other.
package anagram

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	bitset := 0
	for _, v := range a {
		bitset = bitset ^ int(v)
	}
	for _, v := range b {
		bitset = bitset ^ int(v)
	}
	return bitset == 0
}
