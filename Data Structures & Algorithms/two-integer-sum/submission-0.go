func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		guessNumber := nums[i]
		val, ok := mp[target - guessNumber] 
		if ok {
			res = append(res, val)
			res = append(res, i)
			break
		} else {
			mp[guessNumber] = i
		}
	}
	return res
}
