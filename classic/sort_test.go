package classic

import "testing"

var (
	nums = []int{1, 3, 22, 12, 43, 6, 2, 12, 345, 2, -10}
)

func TestBubbleSort(t *testing.T) {
	BubbleSort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Fail()
		}
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(nums)
	}
}

func TestSelectSort(t *testing.T) {
	SelectSort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Fail()
		}
	}
}

func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort(nums)
	}
}

func TestInsertSort(t *testing.T) {
	InsertSort(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Fail()
		}
	}
}

func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort(nums)
	}
}

func TestMergeSort(t *testing.T) {
	MergeSort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Fail()
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(nums, 0, len(nums)-1)
	}
}

func TestQuickSort(t *testing.T) {
	QuickSort(nums, 0, len(nums)-1)
	t.Logf("%v", nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			t.Fail()
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(nums, 0, len(nums)-1)
	}
}