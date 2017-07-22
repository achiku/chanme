package chanme

import (
	"testing"
	"time"
)

func TestBlockHash(t *testing.T) {
	exp := "5ddc7436553cfd6564884cd8fda70e4b9645cbb94ba4f52aae2857d6f67cecfc"
	tm := time.Date(2017, 7, 15, 18, 49, 0, 0, time.UTC)
	tm.UnixNano()
	b := &Block{
		Index:        1,
		PreviousHash: " 9eb930e48257fd6be7ed2d7bb63b06242cd46b4b5e5ad7fa5e3b21e7fe1d2a3b ",
		Timestamp:    tm.UnixNano(),
		Data:         "xxxxxxxxs",
	}
	if h := Hash(b); h != exp {
		t.Errorf("want %s got %s", exp, h)
	}
}
