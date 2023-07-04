// 在 Go 语言中，数组和切片的 append 操作的时间复杂度是不同的。

// 数组（Array）的 append 操作：

// 当向数组追加元素时，由于数组的长度是固定的，无法直接追加元素。需要创建一个新的更大的数组，并将原始数组的内容复制到新数组中。
// 这意味着在数组上进行 append 操作的时间复杂度为 O(n)，其中 n 是数组的长度。
// 切片（Slice）的 append 操作：

// 切片的 append 操作更为灵活，可以直接在原始切片的末尾追加元素，并自动扩容切片的容量。
// 当切片的容量不足以容纳新的元素时，Go 会创建一个新的更大的底层数组，并将原始切片的元素复制到新数组中。然后，在新数组上进行追加操作。
// 平均情况下，切片的 append 操作的时间复杂度为 O(1)。但在扩容时，复杂度可能为 O(n)，其中 n 是当前切片的长度。
// 需要注意的是，切片的 append 操作可能会触发底层数组的重新分配和复制，因此在进行大量追加操作时，建议事先设置切片的容量，以避免频繁的扩容操作，提高性能。

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
}
