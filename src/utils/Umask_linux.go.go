
package utils

import "syscall"

func Umask()  {
	syscall.Umask(0)
}