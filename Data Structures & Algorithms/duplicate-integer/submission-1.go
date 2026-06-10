func hasDuplicate(nums []int) bool {
    len := len(nums)
    mp := make(map[string]bool)
    for i := 0 ; i < len ; i++ {
        val := strconv.Itoa(nums[i])
        if mp[val] {
            return true
        }
        mp[val] = true
    }
    
    return false
}
