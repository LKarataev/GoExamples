package Knapsack

type Present struct {
	Value int
	Size  int
}

func grabPresents(presents []Present, capacity int) []Present {
	n := len(presents)
	maxVal := make([][]int, n+1)
	for i := range maxVal {
		maxVal[i] = make([]int, capacity+1)
	}

	for i := 1; i <= n; i++ {
		for j := 0; j <= capacity; j++ {
			if presents[i-1].Size > j {
				maxVal[i][j] = maxVal[i-1][j]
			} else {
				maxVal[i][j] = max(maxVal[i-1][j], presents[i-1].Value+maxVal[i-1][j-presents[i-1].Size])
			}
		}
	}

	selectedPresents := make([]Present, 0)
	for i, j := n, capacity; i > 0; i-- {
		if maxVal[i][j] != maxVal[i-1][j] {
			selectedPresents = append(selectedPresents, presents[i-1])
			j -= presents[i-1].Size
		}
	}
	return selectedPresents
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
