package Loops

/*
	Author: Christopher Lillthors
	Version 1.0
*/

/*
#include <stdio.h>

float epsilon() {

float machine_epsilon = 1.0f;

  do{
    machine_epsilon /= 2.0f;
  } while ((float)(1.0 + (machine_epsilon/2.0)) != 1.0);

return (float)machine_epsilon;
}
*/
import "C"

import (
	"fmt"
	"math"
)

var (
	epsilon = float64(C.epsilon())
)

func Sqrt(x, guess float64) float64 {
	counter := 0
	var end float64
	for {
		end = guess - (math.Pow(guess, 2)-x)/(2*guess)
		counter += 1
		fmt.Println(...)
		if math.Abs(guess-end) < epsilon {
			break
		}
		guess = end
	}
	return end
}
