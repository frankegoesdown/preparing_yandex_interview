package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
}

func countPrimes(n int) int {
	primes := make([]bool, n)
	for i := 2; i*i < len(primes); i++ {
		if !primes[i] {
			for j := i; j*i < len(primes); j++ {
				primes[i*j] = true
			}
		}
	}
	count := 0
	for i := 2; i < len(primes); i++ {
		if !primes[i] {
			count++
		}
	}
	return count
}
