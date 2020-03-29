package removeDuplicates

//func main() {
//
//	var a = []int{00, 10, 20, 30, 40, 40, 40, 50, 60, 60, 70, 80, 80}
//	fmt.Println(a)
//	fmt.Println(removeDuplicates(a))
//
//	fmt.Println(a)
//}

func removeDuplicates(nums []int) int {
	unique := 0

	for cur := 1; cur < len(nums); cur++ {
		if nums[cur] != nums[unique] {
			//nums[cur] is not a duplicate, increment unique index by 1
			unique++
			//the following operation will copy the value of nums[cur] to nums[unique]
			nums[unique] = nums[cur]
		}
	}

	return unique + 1
}
