func twoSum(nums []int, target int) []int {
    m:=make(map[int]int)

    for i , val := range nums{
         complement := target - val

         if idx , found:= m[complement]; found{
            return []int{i,idx}
         }

         m[val] = i 
    }

    return []int{}
}