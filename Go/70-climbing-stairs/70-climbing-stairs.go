package leetcode

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	ways := []int{1, 1}

	for i := 2; i < n+1; i++ {
		ways = append(ways, ways[i-1]+ways[i-2])
	}
	return ways[len(ways)-1]

}
