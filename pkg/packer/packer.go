package packer

// buckets is sorted MAX to MIN so linear seach gets the first fitting one
func maxFittingBucket(buckets []int, limit int) int {
	// find first bucket fitting the limit
	for i := range buckets {
		if buckets[i] <= limit {
			return i
		}
	}
	// every bucket is bigger than limit so take the smallest one
	return len(buckets)-1
}

// looks for encompassing bucket if any found
func existsBetterBucket(buckets []int, value int) (bool, int) {
	for i := len(buckets)-1; i >= 0; i-- {
		// smallest fitting bucket not same as the value
		if buckets[i] > value && buckets[i] % value == 0 {
			return true, i
		}
	}
	return false, -1
}

// merge consequtive packages to minimize packages allocation
func merge(input []int, buckets []int) []int {
	var j int
	merged := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		if i < len(input) - 1 && input[i] == input[i+1] {
			found, index := existsBetterBucket(buckets, input[i])
			if found {
				merged[j] = buckets[index]
				i++
			} else {
				merged[j] = input[i]
			}
			j++
		} else {
			merged[j] = input[i]
			j++
		}
	}

	return merged[:j]
}

// pack packages (sic!) into sorted set of buckets (ordered MAX to MIN)
func pack(items int, buckets []int) []int {
	var results []int

	// use greedy search for best fitting buckets
	var i int
	for items > 0 {
		i = maxFittingBucket(buckets, items)
		results = append(results, buckets[i])
		items = items - buckets[i]
	}

	// try to merge smaller buckets if encompassing candidates found
	results = merge(results, buckets)

	return results
}
