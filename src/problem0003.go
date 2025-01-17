/**
 * Largest prime factor
 *
 * https://projecteuler.net/problem=3
 *
 * The prime factors of 13195 are 5, 7, 13 and 29.
 *
 * What is the largest prime factor of the number 600851475143 ?
 */

package projecteuler

import (
	"gon.cl/projecteuler.net/src/helpers"
	log "gon.cl/projecteuler.net/src/lib"
)

func Problem0003(_top int) int {
	var maxPrimeFactor int

	divs := helpers.Divisors(_top)

	log.Info("Divisors(%d) = %v \n", _top, divs)

	var i = len(divs) - 1

	for ok := true; ok; ok = i >= 0 && maxPrimeFactor == 0 {
		prime := helpers.IsPrime(divs[i])

		log.Debug("%d is Prime? => %t \n", divs[i], prime)

		if prime {
			maxPrimeFactor = divs[i]
		}

		i -= 1
	}

	log.Info("maxPrimeFactor is %d \n", maxPrimeFactor)

	return maxPrimeFactor
}
