const BUCKETS = 700

type MyHashMap struct {
	m [BUCKETS][][2]int
}

func Constructor() MyHashMap {
	return MyHashMap{}
}

func hash(key int) int {
	return key % BUCKETS
}

func (this *MyHashMap) Put(key int, value int) {
	hashed := hash(key)
	for i, item := range this.m[hashed] {
		if item[0] == key {
			item[1] = value
			this.m[hashed][i] = item
			return
		}
	}
	this.m[hashed] = append(this.m[hashed], [2]int{key, value})
}

func (this *MyHashMap) Get(key int) int {
	hashed := hash(key)
	for _, item := range this.m[hashed] {
		if item[0] == key {
			return item[1]
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	hashed := hash(key)
	for i, item := range this.m[hashed] {
		if item[0] == key {
			this.m[hashed] = append(this.m[hashed][:i], this.m[hashed][i+1:]...)
			return
		}
	}
}