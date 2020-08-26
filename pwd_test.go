package gopasswd

import (
	"fmt"
	"testing"
)

func TestGetPWUid(t *testing.T) {
	test, _ := GetPWUid(0)
	if test.pwName != "root" {
		t.Errorf("pwName == %s; want root", test.pwName)
	}
	if test.pwDir != "/root" {
		t.Errorf("pwDir == %s; want /root", test.pwDir)
	}

	_, test2 := GetPWUid(93948)
	if test2 !=  {
		t.Errorf("pwName == %e; wanted error", test2)
	}
}
