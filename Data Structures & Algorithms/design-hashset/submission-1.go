type MyHashSet struct {
	arr [1_000_001]bool
}

func Constructor() MyHashSet {
	return MyHashSet{}
}

func (this *MyHashSet) Add(key int) {
	this.arr[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.arr[key] = false
}

func (this *MyHashSet) Contains(key int) bool {
	return this.arr[key]
}