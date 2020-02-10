package main

func Index(items [] int, predicate func(int) bool) int {
	for idx, item := range items {
		if predicate(item) {
			return idx
		}
	}
	return -1
}

func All(items []int, predicate func(int) bool) bool {
	return Index(items, func(i int) bool {
		return !predicate(i)
	}) == -1
}

func Any(items []int, predicate func(int) bool) bool {
	return Index(items, predicate) != -1
}

func None(items []int, predicate func(int) bool) bool {
	return Index(items, predicate) == -1
}

func Find(items []int, predicate func(int) bool) int {
	return items[Index(items, predicate)]
}

func main() {

}