package main

import "fmt"

var (
	output int
	done   int
)

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func BinarySearch(a []int, x int) int {
	r := -1 // not found
	start := 0
	end := len(a) - 1
	for start <= end {
		mid := (start + end) / 2
		if a[mid] == x {
			r = mid // found
			break
		} else if a[mid] < x {
			start = mid + 1
		} else if a[mid] > x {
			end = mid - 1
		}
	}
	return r
}

func firstMissingPositive(nums []int) int {

	inputsorted := quickSortStart(nums)
    fmt.Println(inputsorted)
    
	if inputsorted[0] > 1 && len(inputsorted)==1{
		return 1
	}

    if inputsorted[0] == 1 && len(inputsorted)==1 {
		return 2
	}

	cache := BinarySearch(inputsorted, 1)

	if cache == -1 {
		return 1
	}
    
    num:=cache
	for i := cache; i < len(inputsorted)-1; i++ {

        
        
        if inputsorted[i]!=num-cache+1{
            return i-cache+1
        }
		fmt.Println(inputsorted[i], i,num - cache)
        
        if inputsorted[i+1]!=inputsorted[i]{
            num++
        }
	}
    
    if inputsorted[len(inputsorted)-1]-1!=inputsorted[len(inputsorted)-2]{
        return 
    }

	if done == 0 {
		return inputsorted[len(inputsorted)-1] + 1
	}
	return 0
}
