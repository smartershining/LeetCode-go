package main

func numJewelsInStones(J string, S string) int {
	count := 0
	for i := 0; i < len(S); i++ {
		for j := 0; j < len(J); j++ {
			if S[i] == J[j] {
				count++
				break
			}
		}
	}
	return count
}

func numJewelsInStones2(J string, S string) int {
	set := make(map[byte]bool)
	count := 0
	for i := 0; i < len(J); i++ {
		set[J[i]] = true
	}
	for i := 0; i < len(S); i++ {
		if _, ok := set[S[i]]; ok {
			count++
		}
	}
	return count
}
