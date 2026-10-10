type CountSquares struct {
	count  map[[2]int]int
	points [][2]int
}

func Constructor() CountSquares {
	return CountSquares{
		count: make(map[[2]int]int),
	}
}

func (this *CountSquares) Add(point []int) {
	p := [2]int{point[0], point[1]}
	if this.count[p] == 0 {
		this.points = append(this.points, p)
	}
	this.count[p]++
}

func (this *CountSquares) Count(point []int) int {
	total := 0
	for _, p := range this.points {
		dx, dy := abs(point[0]-p[0]), abs(point[1]-p[1])
		if dx == 0 || dx != dy {
			continue
		}
		total += this.count[[2]int{p[0], p[1]}] * this.count[[2]int{p[0], point[1]}] * this.count[[2]int{point[0], p[1]}]
	}
	return total
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}