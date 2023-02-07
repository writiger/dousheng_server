package snowfalke

import (
	"fmt"
	"testing"
)

func TestGetLease(t *testing.T) {
	lm := NewLeaseMaker()
	for i := 0; i < 10; i++ {
		id, err := lm.getLease()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(id)
	}
}
