package utilities

func reverse_array(old []int) (new []int) {
	start := 0
	end := len(old) - 1
	for start < end {
		new[start] = old[end]
		new[end] = old[start]
		start--
		end--
	}
	return new
}
