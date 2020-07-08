func threeSum(nums []int) [][]int {
    res := [][]int{}
    sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
    for i:= range nums {
        if i > 0 && nums[i-1] == nums[i] {
            continue
        }
        left := i+1
        right := len(nums)-1
        for left < right {
            if left > i+1 && nums[left-1] == nums[left] {
                left++
                continue
            }
            if right < len(nums)-1 && nums[right+1] == nums[right] {
                right--
                continue
            }
            if nums[left] + nums[right] == -1 * nums[i] {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                left++
                right--
            }else if nums[left] + nums[right] < -1 * nums[i] {
                left++
            }else {
                right--
            }
        }
    }
    return res
}