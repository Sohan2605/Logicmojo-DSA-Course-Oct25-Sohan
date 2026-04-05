func removeDuplicates(nums []int) int {

   left := 0
   right := 1

   for right < len(nums){
    
    if nums[left]!=nums[right]{
        left++;
        nums[left] , nums[right] = nums[right] , nums[left]
    }
    right++
   } 
   
   return left+1
}