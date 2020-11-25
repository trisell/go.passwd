package gopasswd

import (
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
	errorValue := "Unable to Find UID"
	if test2.Error() != errorValue {
		t.Errorf("error == %v; Expected %s", test2, errorValue)
	}
}

func TestPutPWEnt(t *testing.T) {

	testUser := &Passwd{
		pwName:   "testUser",
		pwPasswd: "et22rdre",
		pwShell:  "/bin/bash",
	}

	_, errorValue := PutPWEnt(testUser)

	if errorValue.Error() != "Unable to create user" {
		t.Errorf("error == %v; Expected Unable to create user", errorValue)
	}

}
