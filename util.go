package editor

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func callNum(function func(), num int) {
	for i := 0; i < num; i++ {
		function()
	}
}
