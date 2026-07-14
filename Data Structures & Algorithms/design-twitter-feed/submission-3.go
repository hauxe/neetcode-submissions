type KHeap struct {
	heap [][2]int
}

func (this *KHeap) Push(val [2]int) {
	this.heap = append(this.heap, val)
	child := len(this.heap) - 1
	parent := (child - 1) / 2

	for child != parent && this.heap[parent][1] < this.heap[child][1] {
		this.heap[child], this.heap[parent] = this.heap[parent], this.heap[child]
		child, parent = parent, (parent-1)/2
	}
}

func (this *KHeap) Pop() [2]int {
	if len(this.heap) == 0 {
		return [2]int{}
	}
	top := this.heap[0]
	this.heap[0] = this.heap[len(this.heap)-1]
	this.heap = this.heap[:len(this.heap)-1]

	i := 0
	for {
		largest := i
		left, right := i*2+1, i*2+2
		if left < len(this.heap) && this.heap[left][1] > this.heap[largest][1] {
			largest = left
		}
		if right < len(this.heap) && this.heap[right][1] > this.heap[largest][1] {
			largest = right
		}

		if largest == i {
			break
		}
		this.heap[i], this.heap[largest] = this.heap[largest], this.heap[i]
		i = largest
	}

	return top
}

type Twitter struct {
	followerMap map[int]map[int]struct{}
	tweetMap    map[int][][2]int
	maxFeed     int
	counter     int
}

func Constructor() Twitter {
	return Twitter{
		followerMap: make(map[int]map[int]struct{}),
		tweetMap:    make(map[int][][2]int),
		maxFeed:     10,
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.Follow(userId, userId)
	this.counter++
	this.tweetMap[userId] = append(this.tweetMap[userId], [2]int{tweetId, this.counter})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	kHeap := &KHeap{}
	var result []int
	for i := 0; i < this.maxFeed; i++ {
		for followeeId := range this.followerMap[userId] {
			tweetMap := this.tweetMap[followeeId]
			j := len(tweetMap) - 1 - i
			if j >= 0 {
				kHeap.Push(tweetMap[j])
			}
		}
		feed := kHeap.Pop()
		if feed[1] == 0 {
			return result
		}
		result = append(result, feed[0])
	}
	return result
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.followerMap[followerId] == nil {
		this.followerMap[followerId] = make(map[int]struct{})
	}
	this.followerMap[followerId][followeeId] = struct{}{}
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.followerMap[followerId] == nil || followerId == followeeId {
		return
	}
	delete(this.followerMap[followerId], followeeId)
}