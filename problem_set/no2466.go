package problemset

func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([][2]int, high+1) //dp[i][0] 以0结尾的数量,dp[i][1]以1结尾的数量； dp[i][0] = dp[i-zero][0]+dp[i-zero][1] dp[i][1]同理
	//base case dp[zero][0]=1 dp[one][1]=1
	for i := range dp {
		dp[i] = [2]int{0, 0}
	}
	mod := 1000000000 + 7
	dp[zero][0] = 1
	dp[one][1] = 1
	for i := 1; i <= high; i++ {
		if i > zero {
			dp[i][0] = (dp[i-zero][0] + dp[i-zero][1]) % mod
		}
		if i > one {
			dp[i][1] = (dp[i-one][0] + dp[i-one][1]) % mod
		}
	}
	//fmt.Println(dp)
	ret := 0
	for i := low; i <= high; i++ {
		ret = (ret + dp[i][0] + dp[i][1]) % mod
	}
	return ret
}
