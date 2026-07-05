/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
	length int
}

func Constructor() Codec {
	return Codec{
		length: 5, // include sign (1 char - 0 positive - 1 negative - 2 null)
	}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return string(this.serializeInt(root))
	}
	b := this.serializeInt(root)
	return string(b) + this.serialize(root.Left) + this.serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	b := []byte(data)
	root, _ := this.deserializeBytes(b, 0)

	return root
}

func (this *Codec) deserializeBytes(b []byte, start int) (*TreeNode, int) {
	root := this.deserializeInt(b[start : start+this.length])
	if root == nil {
		return nil, start + this.length
	}
	left, index := this.deserializeBytes(b, start+this.length)
	root.Left = left
	right, index := this.deserializeBytes(b, index)
	root.Right = right
	return root, index
}

func (this *Codec) serializeInt(root *TreeNode) []byte {
	result := make([]byte, this.length)
	for i := 1; i < this.length; i++ {
		result[i] = '0'
	}
	if root == nil {
		result[0] = '2'
		return result
	}
	val := root.Val
	result[0] = '0'
	if root.Val < 0 {
		result[0] = '1'
		val *= -1
	}
	for i := this.length - 1; i > 0; i-- {
		digit := val % 10
		val = val / 10
		result[i] = byte(digit) + '0'
	}
	return result
}

func (this *Codec) deserializeInt(data []byte) *TreeNode {
	if len(data) == 0 || data[0] == '2' {
		return nil
	}
	sign := 1
	if data[0] == 1 {
		sign = -1
	}
	var val int
	for i := 1; i < this.length; i++ {
		val = val*10 + int(data[i]-'0')
	}
	return &TreeNode{
		Val: val * sign,
	}
}
