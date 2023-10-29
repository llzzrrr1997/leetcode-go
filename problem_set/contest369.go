package problemset

import "fmt"

func findKOr(nums []int, k int) int {
	bitCntMap := [32]int{}
	for _, val := range nums {
		i := 0
		v := 1
		for val > 0 {
			if val%2 == 1 {
				bitCntMap[i]++
			}
			val = val >> 1
			v = v << 1
			i++
		}
	}
	//fmt.Print(bitCntMap)
	ret := 0
	v := 1
	for i := 0; i < 32; i++ {
		if bitCntMap[i] >= k {
			ret += v
		}
		v = v << 1
	}
	return ret
}

func minSum(nums1 []int, nums2 []int) int64 {
	sum1 := 0
	sum2 := 0
	cnt1 := 0
	cnt2 := 0
	for _, val := range nums1 {
		sum1 += val
		if val == 0 {
			cnt1++
		}
	}
	for _, val := range nums2 {
		sum2 += val
		if val == 0 {
			cnt2++
		}
	}
	if cnt1 == 0 && cnt2 == 0 && sum1 != sum2 {
		return -1
	}
	if cnt1 == 0 && cnt2 != 0 && sum1 < sum2 {
		return -1
	}
	if cnt1 != 0 && cnt2 == 0 && sum1 > sum2 {
		return -1
	}
	if cnt1 == cnt2 && sum1 == sum2 {
		return int64(sum1 + cnt1)
	}
	if cnt1 < cnt2 {
		cnt1, cnt2 = cnt2, cnt1
		sum1, sum2 = sum2, sum1
	}
	fmt.Println(sum1, sum2)
	fmt.Println(cnt1, cnt2)
	//cnt1 >= cnt2
	//sum1 > sum2 ==> sum1+cnt1 ;
	//cnt1 > cnt2  sum1 == sum2 ==> -1 ;

	//sum1 < sum2 ==> sum2-sum1 比 cnt1-cnt2 大 ==> sum2+cnt1 ; 2 5 3 1  3>2 6
	//sum1 < sum2 ==> sum2-sum1 比 cnt1-cnt2 小 -1; 2 5 5 1   3 < 4
	if sum1 > sum2 {
		fmt.Println("sum1 > sum2")
		return int64(sum1 + cnt1)
	} else if sum1 == sum2 && cnt2 == 0 {
		fmt.Println(" sum1 == sum2 && cnt2 == 0")
		return -1
	} else {
		if sum2-sum1 >= cnt1-cnt2 {
			fmt.Println("sum2-sum1 >= cnt1-cnt2")
			return int64(sum2 + cnt2)
		} else if sum2-sum1 < cnt1-cnt2 && cnt2 != 0 {
			fmt.Println("sum2-sum1 < cnt1-cnt2 && cnt2 != 0")
			return int64(sum1 + cnt1)
		}
	}
	fmt.Println("return -1")
	return -1
}
