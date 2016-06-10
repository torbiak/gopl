// ex4.5 dedupes a slice of strings.
package unique

func unique(strs []string) []string {
	w := 0 // index of last written string
	for _, s := range strs {
		if strs[w] == s {
			continue
		}
		w++
		strs[w] = s
	}
	return strs[:w+1]
}
