package removeDuplicates

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	numUnique := removeDuplicates([]int{00, 10, 20, 30, 40, 40, 40, 50, 60, 60, 70, 80, 80})
	if numUnique != 9 {
		t.Error("mismatching number of uniques, expected 9, got ", numUnique)
	}

	//TODO: add assert

}
