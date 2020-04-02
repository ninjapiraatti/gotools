package gotools

// Fprime IS NOT FINISHED
func Fprime(nbr int64) []int64 {
	var nbrs []int64
	var i int64
	var prime int64 = 1
	for i = 0; i < nbr; i++ {
		nbrs = append(nbrs, prime+i)
	}
	return nbrs
}
