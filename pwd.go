package gopasswd

/* #cgo CFLAGS: -g -Wall
 #include <stdio.h>
 #include <stdlib.h>
 #include <pwd.h>
 FILE * getpwdFile(){
	FILE *fp;
	fp = fopen("/etc/passwd", "w+");
	return fp;
 }

*/
import "C"
import "unsafe"

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

// Getpwnam Returns User name using glibc user functions
func Getpwnam(name string) *Passwd {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cpw := C.getpwnam(cname)
	return &Passwd{
		pwUID:   uint32(cpw.pw_uid),
		pwGid:   uint32(cpw.pw_uid),
		pwDir:   C.GoString(cpw.pw_dir),
		pwShell: C.GoString(cpw.pw_shell),
	}
}

// Getpwuid returns user from /etc/passwd via uid
func Getpwuid(uid uint32) (*Passwd, error) {
	cuid := C.uint(uid)
	cpw := C.getpwuid(cuid)
	return &Passwd{
		pwUID:   uint32(cpw.pw_uid),
		pwName:  C.GoString(cpw.pw_name),
		pwDir:   C.GoString(cpw.pw_dir),
		pwShell: C.GoString(cpw.pw_shell),
		pwGecos: C.GoString(cpw.pw_gecos),
	}, nil
}

// Putpwent inserts user into /etc/passwd
func Putpwent(user Passwd) int {
	cPasswd := C.struct_passwd{
		pw_uid: C.uint(user.pwUID),
	}
	cpw := C.putpwent(&cPasswd, C.getpwdFile())

	return int(cpw)
}
