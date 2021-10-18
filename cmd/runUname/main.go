package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	var unameRes unix.Utsname
	if err := unix.Uname(&unameRes); err != nil {
		// Actually, this might be useful, since this means it's not unix.
		fmt.Println("Unable to get uname: ", err)
	}

	// We may need a regex at some point... when we have more than one installer.
	fmt.Println("Version: ", string(unameRes.Version[:]))
	fmt.Println("Sysname: ", string(unameRes.Sysname[:]))
	fmt.Println("Release: ", string(unameRes.Release[:]))
	fmt.Println("Machine: ", string(unameRes.Machine[:]))
}
