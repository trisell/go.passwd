package gopasswd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// Passwd is of same type as pwd from the C type
type Passwd struct {
	pwName   string
	pwPasswd string
	pwUID    uint32
	pwGid    uint32
	pwGecos  string
	pwDir    string
	pwShell  string
}

// GetPWUid returns user from /etc/passwd via uid
func GetPWUid(uid int) (*Passwd, error) {

	file, err := os.Open("/etc/passwd")
	if err != nil {
		return nil, err
	}
}

	for scanner.Scan() {
		user := strings.Split(scanner.Text(), ":")
		userUID, _ := strconv.Atoi(user[2])
		if userUID == uid {
			fmt.Println(user)
			userGid, _ := strconv.Atoi(user[3])
			return &Passwd{
				pwName:   user[0],
				pwPasswd: user[1],
				pwUID:    userUID,
				pwGid:    userGid,
				pwGecos:  user[4],
				pwDir:    user[5],
				pwShell:  user[6],
			}, nil
		}
	}
	return &Passwd{}, fmt.Errorf("Unable to find UID")
}

//Getpwnam returns a user from /etc/passwd
func Getpwnam(name string) (*Passwd, error) {

	file, err := os.Open("/etc/passwd")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		user := strings.Split(scanner.Text(), ":")
		userName := user[0]
		if userName == name {
			userUID, _ := strconv.Atoi(user[2])
			userGid, _ := strconv.Atoi(user[3])
			return &Passwd{
				pwName:   user[0],
				pwPasswd: user[1],
				pwUID:    userUID,
				pwGid:    userGid,
				pwGecos:  user[4],
				pwDir:    user[5],
				pwShell:  user[6],
			}, nil
		}
	}

	return nil, fmt.Errorf("Unable to locate user with username")
}

// Putpwnam inserts a user struct into the /etc/passwrd file, if struct contains a password other then x
// that password is placed in the /etc/shadow file
func Putpwnam(user *Passwd) (bool, error) {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		return false, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		userEntry := strings.Split(scanner.Text(), ":")
		userName := userEntry[0]
		if userName == user.pwName {
			return false, fmt.Errorf("User with username %s already exists", user.pwName)
		}

	}

	fileWrite, err := os.OpenFile("/etc/passwd", os.O_APPEND|os.O_WRONLY, 0744)

	fileWrite.WriteString(user.pwName + ":x:" + fmt.Sprint(user.pwUID) + ":" + fmt.Sprint(user.pwGid) + ":" + user.pwGecos + ":" + user.pwDir + ":" + user.pwShell)

	return true, nil
}

//Rempwnam removes a user from /etc/passwd
func Rempwnam(name string) (bool, error) {

	file, err := os.Open("/etc/passwd")
	if err != nil {
		return false, err
	}
	defer file.Close()

	tempFile, err := ioutil.TempFile("", "")
	if err != nil {
		return false, err
	}
	defer os.Remove(tempFile.Name())

	scanner := bufio.NewScanner(file)

	userExists := false
	for scanner.Scan() {
		userEntry := strings.Split(scanner.Text(), ":")
		userName := userEntry[0]
		if userName == name {
			userExists = true
			continue
		}

		if _, err := tempFile.WriteString(scanner.Text() + "\n"); err != nil {
			return false, err
		}
	}

	if err := file.Close(); err != nil {
		return false, err
	}

	if err := tempFile.Close(); err != nil {
		return false, err
	}

	if userExists == true {
		if err := os.Rename(tempFile.Name(), "/etc/passwd"); err != nil {
			return false, err

		}
	}

	return true, nil
}
