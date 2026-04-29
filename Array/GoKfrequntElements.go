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


//optimal solution

func topKFrequent(nums []int, k int) []int {
    
    mp := make(map[int]int)

    for _ , val := range nums{
        mp[val]++
    }

    bucket := make([][]int ,  len(nums)+1)
    
    for key , val := range mp{
        bucket[val] = append(bucket[val],key)
    }
    result:=[]int{}
    for i:=len(bucket)-1; i>=0 && len(result) < k ; i--{
        if len(bucket[i]) > 0{
            result = append(result,bucket[i]...)
        }
    }

    return result[:k]
}