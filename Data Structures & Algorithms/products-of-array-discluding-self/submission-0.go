func productExceptSelf(nums []int) []int {
	productNums := make([]int, len(nums))
	product := 1
	numberOfZeros := 0
	noneZeroProduct := 1
	positionOfZeroNumber := -1
	for i := 0; i < len(nums); i++ {
		product *= nums[i]
		if nums[i] == 0 {
			numberOfZeros++
			positionOfZeroNumber = i
		} else {
			noneZeroProduct *= nums[i]
		}
	}
	if numberOfZeros == 0 {
		for i := 0; i < len(nums); i++ {
			productNums[i] = product / nums[i]
		}
	} else if numberOfZeros == 1 {
		for i := 0; i < len(nums); i++ {
			if i == positionOfZeroNumber {
				productNums[i] = noneZeroProduct
			}
		} 
	} else {
		for i := 0; i < len(nums); i++ {
			productNums[i] = 0
		}
	}


	return productNums
}
