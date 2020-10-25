package main

func main() {

}

func subarraysDivByK(A []int, K int) int {
	count := [10000]int{}
	count[0] = 1
	prefix, res := 0, 0
	for _, a := range A {
		prefix = (prefix + a%K + K) % K
		res += count[prefix]
		count[prefix]++
	}
	return res
}

func subarraysDivByK2(A []int, K int) int {
	preSum := map[int]int{0: 1}
	count := 0
	sum := 0
	for _, a := range A {
		sum = (sum + a) % K
		if sum < 0 {
			sum += K
		}
		count += preSum[sum]
		preSum[sum]++
	}

	return count
}
