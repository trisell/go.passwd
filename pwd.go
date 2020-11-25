package gopasswd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Passwd is of same type as pwd from the C type
type Passwd struct {
	pwName   string
	pwPasswd string
	pwUID    int
	pwGid    int
	pwGecos  string
	pwDir    string
	pwShell  string
}

// GetPWUid returns user from /etc/passwd via uid
func GetPWUid(uid int) (*Passwd, error) {

	file, err := os.Open("/etc/passwd")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		user := strings.Split(scanner.Text(), ":")
		userUID, _ := strconv.Atoi(user[2])
		if userUID == uid {

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
	return &Passwd{}, fmt.Errorf("Unable to Find UID")
}

func PutPWEnt(newUser *Passwd) (bool, error) {

	userNameValidation, err := regexp.MatchString("[a-z_][a-z0-9_-]*[$]?", newUser.pwName)
	if err != nil {
		panic(err)
	}

	if userNameValidation {

		uuid, err := getLastUserUid()
		if err != nil {
			return false, err
		}

		if uuid < 1000 {
			file, err := getPWFile() 
			if err != nil {
				return false, err
			}

			passwdEntry := newUser.pwName + ":" + "x" + ":" + "1000" + ":" + "1000" + ":" + ":" + "/home/" + newUser.pwName + ":" + newUser.pwShell

			file.WriteString(passwdEntry)
			return true, nil
		} else {
			file, err := getPWDFile()
			if err != nil {
				panic(err)
			}

			passwdEntry := newUser.pwName + ":" + "x" + ":" + (getLastUserUid + 1) + ":" + (getLastUserUid + 1) ":" "/home/" + newUser.pwName + ":" + newUser.PwShell

			file.WriteString(passwdEntry)

			return true, nil
		}
	}

	return false, fmt.Errorf("Unable to create user")
}

func getPWDFile() (*os.File, error) {
	file, err := os.OpenFile("/home/dom/passwd", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return file, err
	}
	defer file.Close()

	return file, nil
}

func getLastUserUid() (int, error) {

	file, err := os.Open("/etc/passwd")
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var last []byte
	for scanner.Scan() {
		last = scanner.Bytes()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	lastUser := strings.Split(string(last), ":")

	return strconv.Atoi(lastUser[2])

}
