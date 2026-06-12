func topKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	numsAndFrequencys := make([]numberAndFrequency, 0)
	for i := 0; i < len(nums);  i++ {
		num := nums[i]
		mp[num]++
	}
	
	for key, val := range mp {
		numsAndFrequencys = append(numsAndFrequencys, numberAndFrequency{number: key, frequency: val})
	}
	numsAndFrequencys = sortNumberAndFrequencyList(numsAndFrequencys)

	res := make([]int, 0)
	for i := 0 ; i < len(numsAndFrequencys); i++ {
		if k == 0 {
			break
		}
		res = append(res, numsAndFrequencys[i].number)
		k--
	}

	return res
}

func sortNumberAndFrequencyList (list []numberAndFrequency) []numberAndFrequency {
	sort.Slice(list, func(i, j int) bool {
        return list[i].frequency > list[j].frequency
    })
    return list
}

type numberAndFrequency struct {
	number int
	frequency int
}
