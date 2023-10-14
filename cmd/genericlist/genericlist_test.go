package genericlist

import (
	"math/rand"
	"testing"
)

func TestGenericList_NewAndAdd(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)

	if list.Count() != 2 {
		t.Error("Expected list count to be 2 after adding two items")
	}
}

func TestGenericList_Remove(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)
	list.Add(6)

	list.Remove(1)

	if list.Get(1) != 6 {
		t.Error("Remove function did not remove the correct element")
	}
}

func TestGenericList_Remove_Panic(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)
	list.Add(6)

	shouldPanic(t, func() { list.Remove(-1) })
}

func TestGenericList_Clone(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	clonedList := list.Clone()

	if clonedList.Count() != list.Count() {
		t.Errorf("Expected cloned list to have %d elements, got %d", list.Count(), clonedList.Count())
	}
	if clonedList.Get(0) != list.Get(0) {
		t.Error("Cloned list data does not match original list data")
	}
}

// And so on for each function...

func TestGenericList_Reverse(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)
	list.Add(6)

	list.Reverse()

	if list.Get(0) != 6 {
		t.Errorf("Expected first element to be 6 after reverse, got %d", list.Get(0))
	}
}

func TestGenericList_Swap(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)

	list.Swap(0, 1)

	if list.Get(0) != 5 || list.Get(1) != 4 {
		t.Error("Elements were not swapped correctly")
	}
}

func TestGenericList_Swap_Panic(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)

	shouldPanic(t, func() { list.Swap(-1, 0) })
}

func TestGenericList_IsEmpty(t *testing.T) {
	list := &GenericList[int]{}
	list.New()

	if !list.IsEmpty() {
		t.Error("Expected list to be empty")
	}

	list.Add(4)

	if list.IsEmpty() {
		t.Error("Expected list not to be empty after adding an element")
	}
}

func TestGenericList_IsNotEmpty(t *testing.T) {
	list := &GenericList[int]{}
	list.New()

	if list.IsNotEmpty() {
		t.Error("Expected list not to be non-empty initially")
	}

	list.Add(4)

	if !list.IsNotEmpty() {
		t.Error("Expected list to be non-empty after adding an element")
	}
}

func TestGenericList_First(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)

	if list.First() != 4 {
		t.Errorf("Expected first element to be 4, got %d", list.First())
	}
}

func TestGenericList_First_Empty(t *testing.T) {
	list := &GenericList[int]{}
	list.New()

	if list.First() != 0 {
		t.Errorf("Expected first element to be 0, got %d", list.First())
	}
}

func TestGenericList_Last_Empty(t *testing.T) {
	list := &GenericList[int]{}
	list.New()

	if list.Last() != 0 {
		t.Errorf("Expected last element to be 0, got %d", list.Last())
	}
}

func TestGenericList_Last(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)

	if list.Last() != 5 {
		t.Errorf("Expected last element to be 5, got %d", list.Last())
	}
}

func TestGenericList_ToSlice(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)

	slice := list.ToSlice()

	if len(slice) != 2 || slice[0] != 4 || slice[1] != 5 {
		t.Error("Failed to convert the list to a slice correctly")
	}
}

func TestGenericList_FromSlice(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	slice := []int{4, 5}
	list.FromSlice(slice)

	if list.Count() != 2 || list.Get(0) != 4 || list.Get(1) != 5 {
		t.Error("Failed to create a list from a slice correctly")
	}
}

func TestGenericList_CopyFromSlice(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	slice := []int{4, 5}
	list.CopyFromSlice(slice)
	slice[0] = 10

	if list.Get(0) != 4 {
		t.Error("List was modified after modifying the original slice in CopyFromSlice")
	}
}

func TestGenericList_CopyToSlice(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)
	slice := make([]int, 2)
	list.CopyToSlice(slice)

	if slice[0] != 4 || slice[1] != 5 {
		t.Error("Failed to copy list data to slice")
	}
}

func TestGenericList_Copy(t *testing.T) {
	list := &GenericList[int]{}
	list.New()
	list.Add(4)
	list.Add(5)

	copiedList := list.Copy()
	copiedList.Add(6)

	if list.Count() != 2 {
		t.Error("Original list was modified after copying")
	}
}

func TestGenericList_CopyFrom(t *testing.T) {
	sourceList := &GenericList[int]{}
	sourceList.New()
	sourceList.Add(4)
	sourceList.Add(5)

	destList := &GenericList[int]{}
	destList.New()
	destList.CopyFrom(sourceList)
	sourceList.Add(6)

	if destList.Count() != 2 {
		t.Error("Destination list didn't copy the data from source list correctly")
	}
}

func TestGenericList_CopyFrom_Panics(t *testing.T) {
	sourceList := &GenericList[int]{}
	sourceList.New()
	sourceList.Add(4)
	sourceList.Add(5)

	shouldPanic(t, func() { sourceList.CopyFrom(nil) })
}

func TestGenericList_CopyTo(t *testing.T) {
	sourceList := &GenericList[int]{}
	sourceList.New()
	sourceList.Add(4)
	sourceList.Add(5)

	destList := &GenericList[int]{}
	destList.New()
	sourceList.CopyTo(destList)
	sourceList.Add(6)

	if destList.Count() != 2 {
		t.Error("Source list didn't copy the data to destination list correctly", sourceList, destList)
	}
}

func TestGenericList_CopyTo_Panics(t *testing.T) {
	sourceList := &GenericList[int]{}
	sourceList.New()
	sourceList.Add(4)
	sourceList.Add(5)

	shouldPanic(t, func() { sourceList.CopyTo(nil) })
}

func TestGenericList_RemoveByValue(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		value    int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 6, []int{1, 2, 3, 4, 5}}, // Value not present
		{[]int{}, 1, []int{}},                           // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		g.RemoveByValue(tt.value)

		if !slicesEqual(g.data, tt.expected) {
			t.Errorf("Expected %v after removing value %v, but got %v", tt.expected, tt.value, g.data)
		}
	}
}

func TestGenericList_Merge(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		value    []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}, []int{}}, // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		g.Merge(&GenericList[int]{data: tt.value})

		if !slicesEqual(g.data, tt.expected) {
			t.Errorf("Expected %v after removing value %v, but got %v", tt.expected, tt.value, g.data)
		}
	}
}

func TestGenericList_Merge_Panics(t *testing.T) {
	// Define some test cases
	gl := &GenericList[int]{data: []int{1, 2, 3}}
	shouldPanic(t, func() { gl.Merge(nil) })

}

func TestGenericList_RemoveByValue_Panics(t *testing.T) {
	list := &GenericList[string]{}
	list.Add("test")

	shouldPanic(t, func() { list.RemoveByValue("") })
}

func TestGenericList_Clear(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{}}, // Normal list
		{[]int{}, []int{}},              // Already empty list
		{[]int{1}, []int{}},             // Single item list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		g.Clear()

		if len(g.data) != len(tt.expected) {
			t.Errorf("Expected list of length %d after clearing, but got %d", len(tt.expected), len(g.data))
		}
	}
}

func TestGenericList_Get_Panics(t *testing.T) {
	gl := &GenericList[int]{data: []int{1, 2, 3}}
	shouldPanic(t, func() { gl.Get(-1) })
}

func TestGenericList_IndexOf(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		value    int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 6, -1}, // Value not present
		{[]int{}, 1, -1},              // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		actual := g.IndexOf(tt.value)

		if actual != tt.expected {
			t.Errorf("Expected %d after searching for value %d, but got %d", tt.expected, tt.value, actual)
		}
	}
}

func TestGenericList_IndexOf_Panics(t *testing.T) {
	list := &GenericList[string]{}
	list.Add("test")

	shouldPanic(t, func() { list.IndexOf("") })
}

func TestGenericList_Contains(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		value    int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, 3, true},
		{[]int{1, 2, 3, 4, 5}, 1, true},
		{[]int{1, 2, 3, 4, 5}, 5, true},
		{[]int{1, 2, 3, 4, 5}, 6, false}, // Value not present
		{[]int{}, 1, false},              // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		actual := g.Contains(tt.value)

		if actual != tt.expected {
			t.Errorf("Expected %t after searching for value %d, but got %t", tt.expected, tt.value, actual)
		}
	}
}

func TestGenericList_Contains_Panics(t *testing.T) {
	list := &GenericList[string]{}
	list.Add("test")

	shouldPanic(t, func() { list.Contains("") })
}

func TestGenericList_ForEach(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 3, 4, 5, 6}},
		{[]int{}, []int{}}, // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		g.ForEach(func(i *int) { *i++ })

		if !slicesEqual(g.data, tt.expected) {
			t.Errorf("Expected %v after applying ForEach, but got %v", tt.expected, g.data)
		}
	}
}

func TestGenericList_Map(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4, 6, 8, 10}},
		{[]int{}, []int{}}, // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		actual := g.Map(func(i int) int { return i * 2 })

		if !slicesEqual(actual.data, tt.expected) {
			t.Errorf("Expected %v after applying Map, but got %v", tt.expected, actual.data)
		}
	}
}

func TestGenericList_Filter(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{2, 4}},
		{[]int{}, []int{}}, // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		actual := g.Filter(func(i int) bool { return i%2 == 0 })

		if !slicesEqual(actual.data, tt.expected) {
			t.Errorf("Expected %v after applying Filter, but got %v", tt.expected, actual.data)
		}
	}
}

func TestGenericList_Reduce(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 15},
		{[]int{}, 0}, // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		actual := g.Reduce(func(a, b int) int { return a + b })

		if actual != tt.expected {
			t.Errorf("Expected %d after applying Reduce, but got %d", tt.expected, actual)
		}
	}
}

func TestGenericList_Any(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{}, false}, // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		actual := g.Any(func(i int) bool { return i == 3 })

		if actual != tt.expected {
			t.Errorf("Expected %t after applying Any, but got %t", tt.expected, actual)
		}
	}
}

func TestGenericList_All(t *testing.T) {
	// Define some test cases
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, false},
		{[]int{3, 3, 3, 3, 3}, true},
		{[]int{}, true}, // Empty list
	}

	for _, tt := range tests {
		g := &GenericList[int]{data: tt.input}
		actual := g.All(func(i int) bool { return i == 3 })

		if actual != tt.expected {
			t.Errorf("Expected %t after applying All, but got %t", tt.expected, actual)
		}
	}
}

// Helper function to compare two slices for equality
func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// shouldPanic is a helper function for testing panics.
func shouldPanic(t *testing.T, f func()) {
	t.Helper()
	defer func() { _ = recover() }()
	f()
	t.Errorf("should have panicked")
}

func BenchmarkAdd(b *testing.B) {
	list := &GenericList[int]{}
	list.New()

	for i := 0; i < b.N; i++ {
		list.Add(rand.Intn(100))
	}
}

func BenchmarkRemove(b *testing.B) {
	list := &GenericList[int]{}
	list.New()

	for i := 0; i < 1000; i++ {
		list.Add(rand.Int())
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.RemoveByValue(rand.Int())
	}
}

func BenchmarkMap(b *testing.B) {
	list := &GenericList[int]{}
	list.New()

	for i := 0; i < 1000; i++ {
		list.Add(rand.Intn(100))
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		list.Map(func(item int) int {
			return item * 2
		})
	}
}
