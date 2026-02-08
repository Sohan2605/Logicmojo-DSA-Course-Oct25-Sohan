func sortColors(nums []int)  {
    red:=0
    white:=0
    blue:=0

    for _ , num := range nums{
        if num == 0 {
            red++
        }else if num == 1{
            white++;
        }else{
            blue++;
        }
    }
    index:=0
    for red > 0 {
        nums[index]=0
        red--
        index++
    }
    for white > 0{
        nums[index]=1
        white--
        index++
    }
    for blue > 0{
        nums[index]=2
        blue--
        index++
    }
}