func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)

	for i := 0; i < len(nums);  i++ {
		num := nums[i]
		mp[num]++
	}
	res := make([]int, 0)
	for key, val := range mp {
		if val >= k {
			res = append(res, key)
		}
	}

	return res
}
