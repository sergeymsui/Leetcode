// 1140. Stone Game II

// Alice and Bob continue their games with piles of stones.  There are a number
// of piles arranged in a row, and each pile has a positive integer number of stones
// piles[i].  The objective of the game is to end with the most stones.

// Alice and Bob take turns, with Alice starting first.  Initially, M = 1.
// On each player's turn, that player can take all the stones in the
// first X remaining piles, where 1 <= X <= 2M.  Then, we set M = max(M, X).

// The game continues until all the stones have been taken.
// Assuming Alice and Bob play optimally, return the maximum number of stones Alice can get.

package main

import "fmt"

func stoneGameII(piles []int) int {
	n := len(piles)

	dp := [][]int{}
	for i := 0; i < n; i++ {

		arr := []int{}
		for j := 0; j <= n+1; j++ {
			arr = append(arr, 0)
		}

		dp = append(dp, arr)
	}

	suffixSum := make([]int, n)

	suffixSum[n-1] = piles[n-1]

	for i := n - 2; i >= 0; i-- {
		suffixSum[i] = suffixSum[i+1] + piles[i]
	}

	for i := n - 1; i >= 0; i-- {
		for m := 1; m <= n; m++ {
			if i+2*m >= n {
				dp[i][m] = suffixSum[i]
			} else {
				for x := 1; x <= 2*m; x++ {
					dp[i][m] = max(dp[i][m], suffixSum[i]-dp[i+x][max(m, x)])
				}
			}
		}
	}

	return dp[0][1]
}

func main() {
	arr := []int{2, 7, 9, 4, 4}
	fmt.Println(stoneGameII(arr))
}
