package main

import (
	"net/http"
	"os"
	"syscall"
)

func main() {
	config := getConfig()

	// Before we start serving pages, we change to the store dir and
	// then chroot to pwd. If we fail to do either of these things we
	// exit. This is to prevent inadvertent access to any
	// other parts of the system. This also has the benefit that the
	// rest of the program doesn't need to know the real location of the
	// store; as far as it is concerned it is just using the pwd.
	die(syscall.Chdir(config.PathToStore))
	die(syscall.Chroot("."))

	// We now shed root privileges (which we only needed in the first
	// place to put ourselves in a chroot jail). We set our effective
	// gid and uid to our *real* gid and uid.
	die(syscall.Setgid(syscall.Getgid()))
	die(syscall.Setuid(syscall.Getuid()))

	// Now that we are running as the unprivileged user, check
	// that we can read and write to pwd.
	die(testPwdWriteable())

	// initLogging establishes INFO and ERROR as two global writers
	// writing to stdout and stderr respectively.
	initLogging(os.Stdout, os.Stderr)

	http.HandleFunc("/", putHandler)
	http.HandleFunc(pasteRoot, getHandler)

	INFO.Printf(
		"Starting server on port %d serving from %q.",
		config.Port, config.PathToStore)

	die(http.ListenAndServe(config.getPortString(), nil))
}
