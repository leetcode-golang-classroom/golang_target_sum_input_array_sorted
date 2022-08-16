package sol

func twoSum(numbers []int, target int) []int {
	lp, rp := 0, len(numbers)-1
	for lp < rp {
		if numbers[lp]+numbers[rp] == target {
			return []int{lp + 1, rp + 1}
		}
		if numbers[lp]+numbers[rp] < target {
			lp++
		}
		if numbers[lp]+numbers[rp] > target {
			rp--
		}
	}
	return nil
}
