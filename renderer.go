package dom

const MLOpenTag byte = '<'
const MLCloseTag byte = '>'

type Html []byte
type Attrs map[string]string

func (h Html) HTML() []byte {
	return []byte(h)
}

func RenderML(node *Node) (result Html) {
	result = append(result, MLOpenTag)
	result = append(result, node.name...)
	for _, attr := range node.attributes {
		result = append(result, ' ')
		result = append(result, attr.name...)
		result = append(result, '=')
		result = append(result, '"')
		result = append(result, attr.value...)
		result = append(result, '"')
	}

	if len(node.content) > 0 || len(node.nodes) > 0 {
		result = append(result, MLCloseTag)
		for _, n := range node.nodes {
			result = append(result, RenderML(n)...)
		}

		result = append(result, node.content...)

		result = append(result, []byte{MLOpenTag, '/'}...)
		result = append(result, node.name...)
		return append(result, MLCloseTag)
	}

	return append(result, []byte{'/', MLCloseTag}...)
}

func H(tag string, attributes Attrs, content string, nodes ...*Node) Html {
	n := NewNode([]byte(tag))

	for k, v := range attributes {
		n.SetAttribute([]byte(k), []byte(v))
	}
	n.SetContent([]byte(content))

	n.SetNodes(nodes)

	return RenderML(n)
}
