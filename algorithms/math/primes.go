package math

import (
	"math"
)

var smallPrimes = [100]int{2, 3, 5, 7, 11, 13, 17, 19,
	23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73,
	79, 83, 89, 97, 101, 103, 107, 109, 113,
	127, 131, 137, 139, 149, 151, 157, 163, 167, 173,
	179, 181, 191, 193, 197, 199, 211, 223, 227, 229,
	233, 239, 241, 251, 257, 263, 269, 271, 277, 281,
	283, 283, 293, 307, 311, 313, 317, 331, 337, 347,
	349, 353, 359, 367, 373, 379, 383, 389, 397, 401, 409, 419,
	421, 431, 433, 439, 443, 449, 457, 461, 463, 467, 479, 487, 491, 499, 503, 509, 521, 523}

func NthPrime(n int) int {

	if n < 1 {
		panic("n has to be at least 1")
	}

	if n <= 100 {
		return smallPrimes[n-1]
	}

	limit := float64(n) * (math.Log(float64(n)) + math.Log(math.Log(float64(n))))
	sqrtlimit := int(math.Sqrt(limit))
	sieve := make([]bool, int(limit/2))

	for p := 3; p <= sqrtlimit; p += 2 {
		if sieve[((p+1)/2)-2] == false {
			for i := p * p; i <= int(limit); i += 2 * p {
				sieve[(i/2)-1] = true
			}
		}
	}

	cnt := 1
	for p := 0; p < len(sieve); p++ {
		if sieve[p] == false {
			cnt += 1
		}
		if cnt == n {
			return p*2 + 3
		}
	}

	panic("unr")
	return 0
}
