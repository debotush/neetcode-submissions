import (
	"slices"
)


func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	mp := make(map[[3]int]bool)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmp := [3]int{nums[i], nums[j], nums[k]}
					tmp = sortThreeElements(tmp)
					_, ok := mp[tmp]
					if !ok {
						mp[tmp] = true
						res = append(res, tmp[:])
					}
				}
			}
		}
	}
	return res
}

func sortThreeElements(nums [3]int) [3]int {
	slice := nums[:]      
	slices.Sort(slice)    
	return nums
}