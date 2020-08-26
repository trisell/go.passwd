package gopasswd

import (
	"bufio"
	"fmt"
	"os"
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
