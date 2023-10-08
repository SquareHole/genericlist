// Package genericlist provides a versatile and type-safe generic list for any type
// that implements the 'comparable' contract, which allows for comparisons using
// the standard comparison operators (<, >, ==, etc.).
package genericlist

// GenericList represents a type-safe, dynamic list of comparable items.
type GenericList[T comparable] struct {
	data []T
}

// New initializes an empty list.
func (g *GenericList[T]) New() {
	g.data = []T{}
}

// Add appends the provided value to the end of the list.
func (g *GenericList[T]) Add(value T) {
	g.data = append(g.data, value)
}

// Remove deletes the element at the specified index from the list.
// Panics if the index is out of bounds.
func (g *GenericList[T]) Remove(i int) {
	if i < 0 || i > len(g.data)-1 {
		panic("index is out of range")
	}
	g.data = append(g.data[:i], g.data[i+1:]...)
}

// RemoveByValue deletes the first occurrence of the specified value from the list.
// Panics if the value is uninitialized.
func (g *GenericList[T]) RemoveByValue(value T) {
	if value == *new(T) {
		panic("value is nil")
	}
	for i := 0; i < len(g.data); i++ {
		if g.data[i] == value {
			g.Remove(i)
			return
		}
	}
}

// Clear erases all elements from the list, resetting it to empty.
func (g *GenericList[T]) Clear() {
	g.data = []T{}
}

// Count returns the number of elements currently in the list.
func (g *GenericList[T]) Count() int {
	return len(g.data)
}

// Clone creates and returns a new list with elements identical to the current list.
func (g *GenericList[T]) Clone() *GenericList[T] {
	return &GenericList[T]{data: g.data}
}

// Merge appends all elements from the provided list to the current list.
func (g *GenericList[T]) Merge(list *GenericList[T]) {
	if list == nil {
		panic("list is nil")
	}
	g.data = append(g.data, list.data...)
}

// Get retrieves the element at the specified index from the list.
func (g *GenericList[T]) Get(i int) T {
	if i < 0 || i > len(g.data)-1 {
		panic("index is out of range")
	}
	return g.data[i]
}

// IndexOf finds and returns the index of the first occurrence of the provided value.
// Returns -1 if the value is not found in the list.
func (g *GenericList[T]) IndexOf(value T) int {
	if value == *new(T) {
		panic("value is nil")
	}
	for i := 0; i < len(g.data); i++ {
		if g.data[i] == value {
			return i
		}
	}
	return -1
}

// Contains checks if the list contains the provided value.
func (g *GenericList[T]) Contains(value T) bool {
	if value == *new(T) {
		panic("value is nil")
	}
	for i := 0; i < len(g.data); i++ {
		if g.data[i] == value {
			return true
		}
	}
	return false
}

// ForEach executes the provided function for each element in the list.
func (g *GenericList[T]) ForEach(f func(*T)) {
	for i := range g.data {
		f(&g.data[i])
	}
}

// Map produces a new list by applying the provided function to each element of the current list.
func (g *GenericList[T]) Map(f func(T) T) *GenericList[T] {
	result := GenericList[T]{}
	for _, item := range g.data {
		result.Add(f(item))
	}
	return &result
}

// Filter creates a new list containing only elements for which the provided function returns true.
func (g *GenericList[T]) Filter(f func(T) bool) *GenericList[T] {
	result := GenericList[T]{}
	for _, item := range g.data {
		if f(item) {
			result.Add(item)
		}
	}
	return &result
}

// Reduce reduces the list to a single value using the provided function.
// Starts with the first element as the initial accumulator and applies the function for each subsequent element.
func (g *GenericList[T]) Reduce(f func(T, T) T) T {
	if len(g.data) == 0 {
		return *new(T)
	}
	acc := g.data[0]
	for _, item := range g.data[1:] {
		acc = f(acc, item)
	}
	return acc
}

// Any checks if any element in the list satisfies the condition in the provided function.
func (g *GenericList[T]) Any(f func(T) bool) bool {
	for _, item := range g.data {
		if f(item) {
			return true
		}
	}
	return false
}

// All verifies if all elements in the list satisfy the condition in the provided function.
func (g *GenericList[T]) All(f func(T) bool) bool {
	for _, item := range g.data {
		if !f(item) {
			return false
		}
	}
	return true
}

// Reverse reverses the order of elements in the list.
func (g *GenericList[T]) Reverse() {
	for i := 0; i < len(g.data)/2; i++ {
		g.data[i], g.data[len(g.data)-i-1] = g.data[len(g.data)-i-1], g.data[i]
	}
}

// Swap exchanges positions of elements at the specified indices in the list.
func (g *GenericList[T]) Swap(i, j int) {
	if i < 0 || i > len(g.data)-1 || j < 0 || j > len(g.data)-1 {
		panic("index is out of range")
	}
	g.data[i], g.data[j] = g.data[j], g.data[i]
}

// IsEmpty checks if the list is currently empty.
func (g *GenericList[T]) IsEmpty() bool {
	return len(g.data) == 0
}

// IsNotEmpty checks if the list contains any elements.
func (g *GenericList[T]) IsNotEmpty() bool {
	return len(g.data) > 0
}

// First retrieves the first element of the list.
func (g *GenericList[T]) First() T {
	if g.IsEmpty() {
		return *new(T)
	}
	return g.data[0]
}

// Last fetches the last element of the list.
func (g *GenericList[T]) Last() T {
	if g.IsEmpty() {
		return *new(T)
	}
	return g.data[len(g.data)-1]
}

// ToSlice converts the list to a standard Go slice.
func (g *GenericList[T]) ToSlice() []T {
	return g.data
}

// FromSlice initializes the list with the elements from the provided slice.
func (g *GenericList[T]) FromSlice(data []T) {
	g.data = data
}

// CopyFromSlice replaces the current list with a deep copy of the provided slice.
func (g *GenericList[T]) CopyFromSlice(data []T) {
	g.data = append([]T{}, data...)
}

// CopyTo copies the elements of the current list (g) into the provided list.
// The provided list's data will be replaced by the elements from the current list.
// Panics if the provided list is nil.
func (g *GenericList[T]) CopyTo(list *GenericList[T]) {
	if list == nil {
		panic("list is nil")
	}

	// Create a slice with the same length and capacity as g.data
	list.data = make([]T, len(g.data))

	copy(list.data, g.data)
}

// CopyToSlice copies the list to the given slice
func (g *GenericList[T]) CopyToSlice(data []T) {
	copy(data, g.data)
}

// Copy creates a new list with elements identical to the current list.
func (g *GenericList[T]) Copy() *GenericList[T] {
	return &GenericList[T]{data: g.data}
}

// CopyFrom replaces the current list's (g) data with the elements from the provided list.
// The current list's data will be overwritten with the data from the provided list.
// Panics if the provided list is nil.
func (g *GenericList[T]) CopyFrom(list *GenericList[T]) {
	if list == nil {
		panic("list is nil")
	}
	g.data = make([]T, len(list.data))
	copy(g.data, list.data)
}
