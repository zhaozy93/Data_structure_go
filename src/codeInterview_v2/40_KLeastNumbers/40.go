package code_40

// 面试题40：最小的k个数
// 题目：输入n个整数，找出其中最小的k个数。例如输入4、5、1、6、2、7、3、8
// 这8个数字，则最小的4个数字是1、2、3、4。

// 使用最大堆来解决问题
type MaxHeap struct {
	vals []int
	len  int
	cap  int
}

func (mh *MaxHeap) Max() (val int, ok bool) {
	if mh.len == 0 {
		return
	}
	ok = true
	val = mh.vals[0]
	return
}

func (mh *MaxHeap) Push(val int) {
	if mh.len < mh.cap {
		mh.vals[mh.len] = val
		mh.Adjust(mh.len, false)
		mh.len++
	}
}

func (mh *MaxHeap) Del() {
	mh.vals[0] = mh.vals[mh.len-1]
	mh.len--
	mh.Adjust(0, true)
}

func (mh *MaxHeap) Adjust(i int, dir bool) {
	if dir {
		lChild := i*2 + 1
		rChild := i*2 + 2
		if lChild <= mh.len && mh.vals[i] < mh.vals[lChild] {
			swap(mh.vals, i, lChild)
			mh.Adjust(lChild, dir)
			return
		}
		if rChild <= mh.len && mh.vals[i] < mh.vals[rChild] {
			swap(mh.vals, i, rChild)
			mh.Adjust(rChild, dir)
			return
		}
	} else {
		father := (i - 1) / 2
		if father >= 0 && mh.vals[i] > mh.vals[father] {
			swap(mh.vals, i, father)
			mh.Adjust(father, dir)
		}
	}
}

func swap(nums []int, i, j int) {
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func FindLeastNums(nums []int, k int) []int {
	mh := &MaxHeap{vals: make([]int, k), cap: k, len: 0}
	for i := 0; i < len(nums); i++ {
		if mh.len < mh.cap {
			mh.Push(nums[i])
		} else {
			val, _ := mh.Max()
			if val > nums[i] {
				mh.Del()
				mh.Push(nums[i])
			}
		}
	}
	return mh.vals
}
