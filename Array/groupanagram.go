func groupAnagrams(strs []string) [][]string {
    
    mp:=make(map[string][]string)

    for _ , str:= range strs {

        word:=[]rune(str)

        sort.Slice(word , func(i,j int)bool{
            return word[i] < word[j]
        })

         key:=string(word)
         mp[key] = append(mp[key],str)
    }
    result:=[][]string{}
    for _ , list := range mp {
      result = append(result,list)
    }
    return result
}