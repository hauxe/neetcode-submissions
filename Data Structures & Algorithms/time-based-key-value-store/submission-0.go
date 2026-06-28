var length = 8

type TimeMap struct {
	m map[string][][]byte
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][][]byte),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], encode(value, timestamp))
}

func (this *TimeMap) Get(key string, timestamp int) string {
	list := this.m[key]

	begin, end := 0, len(list)-1
	maxValue, maxTimeStamp := "", -1
	for begin <= end {
		middle := (begin + end) / 2
		middleValue, middleTimestamp := decode(list[middle])
		if middleTimestamp <= timestamp {
			if maxTimeStamp < middleTimestamp {
				maxTimeStamp = middleTimestamp
				maxValue = middleValue
			}
			begin = middle + 1
		} else {
			end = middle - 1
		}
	}

	return maxValue
}

func encode(value string, timestamp int) []byte {
	v := []byte(value)
	result := make([]byte, length+len(v))
	for i := length - 1; i >= 0; i-- {
		result[i] = byte(timestamp%10) + '0'
		timestamp /= 10
	}
	copy(result[length:], v)
	return result
}

func decode(v []byte) (string, int) {
	timestamp := 0
	for i := 0; i < length; i++ {
		timestamp = 10*timestamp + int(v[i]-'0')
	}
	return string(v[length:]), timestamp
}