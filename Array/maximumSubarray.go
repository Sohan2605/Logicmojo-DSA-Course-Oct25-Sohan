func maxSubArray(nums []int) int {
     currentSum:=nums[0]
     maxSum:=nums[0]

     for i:=1;i<len(nums);i++{
         if currentSum < 0 {
            currentSum = nums[i]
         }else{
            currentSum+=nums[i]
         }

         if currentSum>maxSum{
            maxSum = currentSum
         }
     }

     return maxSum
}