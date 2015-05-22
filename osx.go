package macmain

// #include "osx.h"
// #cgo darwin LDFLAGS: -framework Cocoa
import "C"

import (
	"runtime"
)

var tasks chan func()

func init() {
	runtime.LockOSThread()
	tasks = make(chan func(), 10)
}

/* returns true if the currently executing goroutine is running on the
** main thread. probably not reliable unless LockOSThread has been
** called. */
func IsMainThread() bool {
	return C.isMainThread() != 0
}

/* Enters the run loop. Never returns (even after StopLoop is called). */
func RunLoop() {
	if !IsMainThread() {
		panic("tried to enter run-loop from non-main thread :(")
	}
	C.runLoop()
	// [[NSApplication sharedApplication] run]
}

/* Stops the run loop, terminating the application. */
func StopLoop() {
	C.stopLoop()
	// [[NSApplication sharedApplication] terminate:nil]
}

/* Submits a go function to be run on the main thread.
** Doesn't wait for the function to finish (or start!) before returning. */
func Submit(f func()) {
	tasks <- f
	C.postGoFuncEvent()
	// dispatch_async(dispatch_get_main_queue(), ^{ runGoFunc(); });
}

/* Called back from obj-c code. */
//export runGoFunc
func runGoFunc() {
	f := <- tasks
	f()
}
