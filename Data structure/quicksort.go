package qsort
 
func quickSort(values []int, left int, right int) {
	if left < right {
 
		 
		temp := values[left]
 
		 
		i, j := left, right
		for {
				for values[j] >= temp && i < j {
				j--
			}
 
	for values[i] <= temp && i < j {
				i++
			}
 
			 
			if i >= j {
				break
			}
 			values[i], values[j] = values[j], values[i]
		}
 
		values[left] = values[i]
		values[i] = temp
 

		quickSort(values, left, i-1)
		quickSort(values, i+1, right)
	}
}
 
func QuickSort(values []int) {
	quickSort(values, 0, len(values)-1)
}