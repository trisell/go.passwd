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

	_, test2 := GetPWUid(93941)
	if test2.Error() != "Unable to find UID" {
		t.Errorf("pwName == %s; wanted error", test2.Error())
	}
}

func TestGetpwnam(t *testing.T) {
	testResult, _ := Getpwnam("root")

	if testResult.pwName != "root" {
		t.Errorf("pwName == %s; want root", testResult.pwName)
	}

	if testResult.pwDir != "/root" {
		t.Errorf("Expected pwDir to be /root; Received: %s", testResult.pwDir)
	}

	_, testError := Getpwnam("Test_User_That_I_Hope_NEVER_Exists")

	if testError.Error() != "Unable to locate user with username" {
		t.Errorf("Expected error Message 'Unable to locate user with username; Received: %e", testError)
	}
}

func TestPutpwname(t *testing.T) {
	testUser := &Passwd{
		pwName:   "TestUser123456",
		pwPasswd: "x",
		pwUID:    2001,
		pwGid:    2001,
		pwGecos:  "Test User May Delete",
		pwDir:    "/home/TestUser123456",
		pwShell:  "/bin/bash",
	}
	testSuccess, _ := Putpwnam(testUser)

	if testSuccess != true {
		t.Errorf("Expected True; Received: %t", testSuccess)
	}

	testFailure, _ := Putpwnam(testUser)

	if testFailure != false {
		t.Errorf("Expected false; Recieved: %t", testFailure)
	}

	testResult, _ := Getpwnam("TestUser123456")

	if testResult.pwName != "TestUser123456" {
		t.Errorf("Expected userName %s; Expected: 'TestUser123456'", testResult.pwName)
	}

	testCleanup, _ := Rempwnam("TestUser123456")

	if testCleanup != true {
		t.Errorf("Test Cleanup Failed")
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
