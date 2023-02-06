package util

import (
	"fmt"
	"testing"
)

func TestPassword(t *testing.T) {
	passwdSaved, err := HashAndSalt("testing")
	if err != nil {
		fmt.Println(err.Error())
	}
	t.Run("测试1", func(t *testing.T) {
		res := ComparePasswords(passwdSaved, "testing")
		if !res {
			t.Fatal("理应返回正确")
		}
	})
	t.Run("测试2", func(t *testing.T) {
		res := ComparePasswords(passwdSaved, "hello")
		if res {
			t.Fatal("理应返回错误")
		}
	})
}
