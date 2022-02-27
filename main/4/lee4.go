package main

func main() {
	x := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	println(x)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	var sum = make([]int, len1+len2)
	index := 0
	j := 0
	k := 0
	for {
		var num1, num2 = -100001, -100001
		if j < len1 {
			num1 = nums1[j]

		}
		if k < len2 {
			num2 = nums2[k]
		}
		if num1 == -100001 && num2 == -100001 {
			break
		}
		if num1 == -100001 {
			sum[index] = num2
			index++
			k++
			continue
		}
		if num2 == -100001 {
			sum[index] = num1
			index++
			j++
			continue
		}

		if num1 < num2 {
			sum[index] = num1
			index++
			j++
		} else {
			sum[index] = num2
			index++
			k++
		}
	}
	sumCount := len1 + len2
	if sumCount%2 == 0 {
		return float64(sum[sumCount/2]+sum[sumCount/2-1]) / 2
	} else {
		return float64(sum[sumCount/2])

	}

}
