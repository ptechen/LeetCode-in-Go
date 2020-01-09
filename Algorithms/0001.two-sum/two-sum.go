package problem0001


func twoSum(nums []int, target int) []int {
	if nums == nil {
		return []int{}
	}
	if target == 0 {
		return []int{}
	}
	left :=0
	right := len(nums) - 1
	for {
		if nums[left] < target {
			if left == right {
				break
			}
			if nums[left] + nums[right] == target {
				return []int{left, right}
			}else if nums[left] + nums[right] > target && nums[left + 1] < nums[right - 1]{
				right -= 1
			} else {
				left += 1
			}
		}
	}

	return []int{}
}
