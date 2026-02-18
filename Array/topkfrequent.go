func topKFrequent(nums []int, k int) []int {

    mp:=make(map[int]int)

    for _ , num := range nums{
        mp[num]++
    }
   
    result:=make([]int,k)

    for i:=0; i<k; i++{
        maxFreq:=0
        maxKey:=0
        for key ,val := range mp{
            if val > maxFreq {
                maxFreq=val
                maxKey=key
            }
        }
        result[i] = maxKey
        delete(mp , maxKey)
    }

    return result
}