package snowfalke

import (
	"fmt"
	"testing"
	"time"
)

func TestNewUUID(t *testing.T) {

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(NewUUID())
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			fmt.Println(NewUUID())
		}
	}()

	time.Sleep(2 * time.Second)
}
