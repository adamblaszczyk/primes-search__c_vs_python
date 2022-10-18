/*
 *  -----------------------------------------
 *  Primes Search - Go version
 *  -----------------------------------------
 *  Author: Adam Blaszczyk
 *          http://wyciekpamieci.blogspot.com
 *  Date:   2022-10-07
 *  -----------------------------------------
 *  
 *  Compilation (Windows):
 *         go build -o primes-search.exe primes-search.go
 *  Compilation (FreeBSD, Linux):
 *         go build -o primes-search primes-search.go
 *
 *  Usage:
 *         primes-search
 *
 */

package main

import "fmt"
import "math"
import "time"

func csqrt(n int) int {
	return int(math.Ceil(math.Sqrt(float64(n))))
}

func is_prime(n int) bool {
	var i int
	if n<2 {
		return false
	} else if n==2 {
		return true
	} else {
		for i=2; i<=csqrt(n); i++ {
			if n % i == 0 {
				return false
			}
		}
		return true
	}
}

func main() {
	var max int = 8000000
	var i int
	var t1, t2 int64
	
	fmt.Printf("Checking %d numbers. Please wait...\n", max)
	
	t1 = time.Now().Unix()
	
	for i=1; i<=max; i++ {
		is_prime(i)
		//fmt.Printf("%d   %t\n", i, is_prime(i))
	}
	
	t2 = time.Now().Unix()
	
	fmt.Printf("Done in %d seconds.\n", t2-t1)

}