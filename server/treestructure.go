package server

import ()

type Node struct {
	Text         string `json:"text"`
	Icon         string `json:"icon,omitempty"`
	SelectedIcon string `json:"selectedIcon,omitempty"`
	Color        string `json:"color,omitempty"`
	BackColor    string `json:"backColor,omitempty"`
	Href         string `json:"href,omitempty"`
	Selectable   bool   `json:"selectable,omitempty"`
	State        struct {
		Checked  bool `json:"checked,omitempty"`
		Disabled bool `json:"disabled,omitempty"`
		Expanded bool `json:"expanded,omitempty"`
		Selected bool `json:"selected,omitempty"`
	} `json:"state,omitempty"`
	Nodes []*Node  `json:"nodes,omitempty"`
	Tags  []string `json:"tags,omitempty"`
}

func (n *Node) addElement(node *Node) {
	toadd := true
	for _, n := range n.Nodes {
		if n.Text == node.Text {
			toadd = false
			// Do the merge
			for _, nn := range node.Nodes {
				n.addElement(nn)
			}
		}
	}
	if toadd {
		n.Nodes = append(n.Nodes, node)
	}
}
