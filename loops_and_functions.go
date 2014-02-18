package GoInda13

/*
#include <stdio.h>

int epsilon() {

float epsilon = 1.0f;

  do{
    epsilon /= 2.0f;
  } while ((float)(1.0 + (epsilon/2.0)) != 1.0);

return epsilon;
}
*/
import "C"

import (
	"math"
)

var (
	epsilon = float64(C.epsilon())
)

func Sqrt(x, guess float64) float64 {
	var end float64
	for {
		end = guess - (math.Pow(guess, 2)-x)/(2*guess)
		if guess-end < math.Abs(epsilon) {
			break
		}
		guess = end
	}
	return end
}
