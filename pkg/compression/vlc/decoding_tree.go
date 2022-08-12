package vlc

import "strings"

type DecodingTree struct {
	Value string
	Zero  *DecodingTree
	One   *DecodingTree
}

func (dt *DecodingTree) Decode(str string) string {
	var buf strings.Builder

	currentNode := dt

	for _, x := range str {
		if currentNode.Value != "" {
			buf.WriteString(currentNode.Value)
			currentNode = dt
		}

		switch string(x) {
		case "0":
			currentNode = currentNode.Zero
		case "1":
			currentNode = currentNode.One
		}
	}

	if currentNode.Value != "" {
		buf.WriteString(currentNode.Value)
	}

	return buf.String()
}

func (dt *DecodingTree) Add(code string, value rune) {
	currentNode := dt

	for _, ch := range code {
		switch string(ch) {
		case "0":
			if currentNode.Zero == nil {
				currentNode.Zero = &DecodingTree{}
			}
			currentNode = currentNode.Zero
		case "1":
			if currentNode.One == nil {
				currentNode.One = &DecodingTree{}
			}
			currentNode = currentNode.One
		}
	}

	currentNode.Value = string(value)
}

func (et encodingTable) DecodingTree() DecodingTree {
	res := DecodingTree{}

	for char, code := range et {
		res.Add(code, char)
	}

	return res
}
