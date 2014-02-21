package Sum

func Add(a []int, res chan<- int) {
	var sum int = 0
	for _, value := range a {
		sum += value
	}
	res <- sum
}
