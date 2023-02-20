package covermaker

import (
	"fmt"
	"testing"
)

func TestGetSnapshot(t *testing.T) {
	t.Run("测试", func(t *testing.T) {
		fmt.Println(GetSnapshot("../static/videos/630185627949727744.mp4", "../static/covers/630185627949727744", 1))
	})
}
