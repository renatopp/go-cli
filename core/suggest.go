package core

// suggest returns the candidate closest to input, or "" if none is close
// enough. Ties are resolved by the order of the candidates.
func suggest(input string, candidates []string) string {
	limit := min(max(2, len(input)/3), len(input)-1)
	best, bestDist := "", limit+1
	for _, c := range candidates {
		if d := levenshtein(input, c); d < bestDist {
			best, bestDist = c, d
		}
	}
	return best
}

// levenshtein returns the edit distance between a and b.
func levenshtein(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	prev := make([]int, len(rb)+1)
	curr := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		curr[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			curr[j] = min(prev[j]+1, curr[j-1]+1, prev[j-1]+cost)
		}
		prev, curr = curr, prev
	}
	return prev[len(rb)]
}
