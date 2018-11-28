package main

import (
	"fmt"
	"math/rand"
	"time"
)

func RandInt(min, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Int63n(max-min)
}

func get_random() []int64 {

	result := []int64{}
	for true {
		if len(result) == 10 {
			break
		}
		result = append(result, RandInt(0, 100))
	}
	return result
}

// 步骤
// 冒泡排序算法的运作如下：

// 1.比较相邻的元素。如果第一个比第二个大，就交换他们两个。
// 2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
// 3.针对所有的元素重复以上的步骤，除了最后一个。
// 4.持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

func bubble_sort(arr []int64) []int64 {
	size := len(arr)
	for i := 0; i < size; i++ {
		flag := true
		for j := 1; j < size-i; j++ {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
				flag = false
			}
		}
		if flag {
			return arr
		}
	}
	return arr
}

// 选择排序
// 原理
// 选择排序（Selection sort）是一种简单直观的排序算法。它的工作原理大致是将后面的元素最小元素一个个取出然后按顺序放置。

// 步骤
// 1.在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
// 2.再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
// 3.重复第二步，直到所有元素均排序完毕。

func select_sort(arr []int64) []int64 {
	size := len(arr)

	for i := 0; i < size; i++ {
		min := i
		tmp := min
		for j := i + 1; j < size; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		if min != tmp {
			arr[i], arr[min] = arr[min], arr[i]
			min = tmp
		}
	}
	return arr
}

// 插入排序
// 原理
// 插入排序（Insertion Sort）是一种简单直观的排序算法。它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。

// 步骤
// 1.从第一个元素开始，该元素可以认为已经被排序
// 2.取出下一个元素，在已经排序的元素序列中从后向前扫描
// 3.如果该元素（已排序）大于新元素，将该元素移到下一位置
// 4.重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
// 5.将新元素插入到该位置后
// 6.重复步骤2~5

func insert_sort(arr []int64) []int64 {
	size := len(arr)

	for i := 1; i < size; i++ {
		if arr[i] < arr[i-1] {
			tmp := arr[i]
			index := i
			for j := i - 1; j >= 0; j-- {
				if arr[j] > tmp {
					arr[j+1] = arr[j]
					index = j
				} else {
					break
				}
			}
			arr[index] = tmp
		}
	}
	return arr
}

// 归并排序
// 原理
// 归并操作(归并算法)，指的是将两个已经排序的序列合并成一个序列的操作。归并排序算法依赖归并操作。

// 步骤
// 迭代法
// 1.申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列
// 2.设定两个指针，最初位置分别为两个已经排序序列的起始位置
// 3.比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置
// 4.重复步骤3直到某一指针到达序列尾
// 5.将另一序列剩下的所有元素直接复制到合并序列尾

func merge_sort(arr []int64) []int64 {
	size := len(arr)
	f := func(l, r []int64) []int64 {
		result := []int64{}
		l_size := len(l)
		r_size := len(r)
		l_index, r_index := 0, 0
		for l_index < l_size && r_index < r_size {
			if l[l_index] < r[r_index] {
				result = append(result, l[l_index])
				l_index += 1
			} else {
				result = append(result, r[r_index])
				r_index += 1
			}
		}
		result = append(result, l[l_index:]...)
		result = append(result, r[r_index:]...)
		return result
	}
	if size == 1 {
		return arr
	}
	mid := size / 2
	l := merge_sort(arr[:mid])
	r := merge_sort(arr[mid:])
	return f(l, r)
}

// 堆排序
// 原理
// 堆排序（Heapsort）是指利用堆这种数据结构所设计的一种排序算法。堆积是一个近似完全二叉树的结构，并同时满足堆积的性质：即子结点的键值或索引总是小于（或者大于）它的父节点。

// 步骤
// 1.创建最大堆:将堆所有数据重新排序，使其成为最大堆
// 2.最大堆调整:作用是保持最大堆的性质，是创建最大堆的核心子程序
// 3.堆排序:移除位在第一个数据的根节点，并做最大堆调整的递归运算

func heap_sort(arr []int64) []int64 {

	maxHeap := func(lst []int64, start, end int) {
		root := start
		for true {
			child := 2*root + 1
			if child > end {
				break
			}
			if (child+1) <= end && lst[child] < lst[child+1] {
				child += 1
			}
			if lst[root] < lst[child] {
				lst[root], lst[child] = lst[child], lst[root]
				root = child
			} else {
				break
			}
		}
	}

	max_heap_sort := func(arr []int64) {
		for i := len(arr) / 2; i >= 0; i-- {
			maxHeap(arr, i, len(arr)-1)
		}

		for end := len(arr) - 1; end > 0; end-- {
			arr[0], arr[end] = arr[end], arr[0]
			maxHeap(arr, 0, end-1)
		}
	}

	min_heap := func(lst []int64, start, end int) {
		leaf := start
		for true {
			parant := (leaf - 1) / 2
			if leaf < 0 {
				break
			}
			if (leaf+1) <= end && lst[leaf] > lst[leaf+1] {
				leaf += 1
			}

			if lst[leaf] < lst[parant] {
				lst[leaf], lst[parant] = lst[parant], lst[leaf]
				leaf = parant
			} else {
				break
			}
		}
	}

	min_heap_sort := func(arr []int64) {
		for i := len(arr) - 1; i >= len(arr)/2; i-- {
			min_heap(arr, i, len(arr)-1)
		}
		//fmt.Println(arr)
		for end := len(arr) - 1; end > 0; end-- {
			arr[0], arr[end] = arr[end], arr[0]
			min_heap(arr, end-2, end-1)
		}
	}
	max_heap_sort(arr)
	min_heap_sort([]int64{1})
	return arr
}

// 快速排序
// 原理
// 快速排序使用分治法（Divide and conquer）策略来把一个序列（list）分为两个子序列（sub-lists）。

// 步骤
// 1.从数列中挑出一个元素，称为”基准”（pivot），
// 2.重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区结束之后，该基准就处于数列的中间位置。这个称为分区（partition）操作。
// 3.递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。

func quick_sort(arr []int64) []int64 {
	less := []int64{}
	more := []int64{}
	pivotList := []int64{}
	if len(arr) <= 1 {
		return arr
	} else {
		pivot := arr[0]
		for _, v := range arr {
			if v > pivot {
				more = append(more, v)
			} else if v < pivot {
				less = append(less, v)
			} else {
				pivotList = append(pivotList, v)
			}
		}
		//fmt.Println(more, less)
		less = quick_sort(less)
		more = quick_sort(more)

	}
	return append(append(less, pivotList...), more...)
}

func main() {
	arr := get_random()
	fmt.Println("origin:", arr)
	fmt.Println("bubble_sort:", bubble_sort(arr))
	arr = get_random()
	fmt.Println("origin:", arr)
	fmt.Println("select_sort:", select_sort(arr))
	arr = get_random()
	fmt.Println("origin:", arr)
	fmt.Println("insert_sort:", insert_sort(arr))
	arr = get_random()
	fmt.Println("origin:", arr)
	fmt.Println("merge_sort:", merge_sort(arr))
	arr = get_random()
	arr = []int64{7, 97, 98, 12, 80, 16, 85, 10, 10, 31}
	fmt.Println("origin:", arr)
	fmt.Println("heap_sort:", heap_sort(arr))
	arr = get_random()
	fmt.Println("origin:", arr)
	fmt.Println("quick_sort:", quick_sort(arr))
}
