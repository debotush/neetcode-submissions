func maxArea(heights []int) int {
	left, right := 0, len(heights) - 1
	maxArea := 0

	for left < right {
		area := min(heights[left], heights[right]) * (right - left)

		if maxArea < area {
			maxArea = area
		}
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	} 

	return b
}
