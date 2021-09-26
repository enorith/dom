package dom

import "bytes"

type Attribute struct {
	name  []byte
	value []byte
}

type Node struct {
	name       []byte
	attributes []Attribute
	content    []byte
	nodes      []*Node
}

func (n *Node) SetContent(content []byte) *Node {
	n.content = content
	return n
}

func (n *Node) SetAttribute(name []byte, value []byte) *Node {
	n.attributes = append(n.attributes, Attribute{name: name, value: value})
	return n
}

func (n *Node) RemoveAttribute(name []byte) *Node {
	index := -1
	for i, a := range n.attributes {
		if bytes.Equal(a.name, name) {
			index = i
		}
	}
	if index > -1 {
		n.attributes = removeAttribute(n.attributes, index)
	}

	return n
}

func (n *Node) SetAttributes(attributes []Attribute) *Node {
	n.attributes = attributes
	return n
}

func (n *Node) AddNode(node *Node) *Node {
	n.nodes = append(n.nodes, node)
	return n
}

func (n *Node) SetNodes(nodes []*Node) *Node {
	n.nodes = nodes
	return n
}

func NewNode(name []byte) *Node {
	return &Node{name: name, attributes: make([]Attribute, 0), nodes: make([]*Node, 0)}
}

func removeAttribute(attrs []Attribute, index int) []Attribute {
	return append(attrs[:index], attrs[index+1:]...)
}
