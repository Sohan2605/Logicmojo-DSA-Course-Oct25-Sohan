func fib(n int) int {
     
     dp:=make([]int , n+1)
    
     for i := range dp{
       dp[i] = -1
     }
    
    return dfs(n,dp)
}

func dfs(n int , dp []int) int{
        if n==0 || n==1 {
            return n
        }

        if dp[n]!=-1{
            return dp[n]
        }

        dp[n] = dfs(n-1 , dp) + dfs(n-2 , dp)

        return dp[n]
    }