package example

//NOTE: mockgen does not work properly out of GOPATH
//go:generate mockgen -package=example -source=example.go -destination=mock_test.go Consensus

type (
	Event struct {
		Index uint64
	}

	Consensus interface {
		Push(*Event)
		Last() uint64
	}

	Node struct {
		consensus Consensus
	}
)

func NewNode(c Consensus) *Node {
	return &Node{
		consensus: c,
	}
}

func (n *Node) GenEvent() {
	next := n.consensus.Last() + 1

	e := &Event{
		Index: next,
	}

	n.consensus.Push(e)
}
