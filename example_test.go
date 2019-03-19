package example

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestExample(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// mock behaviour
	var lastGotten uint64
	c := NewMockConsensus(ctrl)
	c.EXPECT().Push(gomock.Any()).
		AnyTimes().
		Do(func(e *Event) {
			if lastGotten < e.Index {
				lastGotten = e.Index
			}
		})
	c.EXPECT().Last().
		AnyTimes().
		DoAndReturn(func() uint64 {
			return lastGotten
		})

	// node test
	n := NewNode(c)
	const Count = 5
	for i := 0; i < 5; i++ {
		n.GenEvent()
	}

	if lastGotten != uint64(Count) {
		t.Fatalf("Only %d of %d events gotten", lastGotten, Count)
	}
}
