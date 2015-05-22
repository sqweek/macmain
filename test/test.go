package main

import (
	"fmt"
	"github.com/sqweek/macmain"
	"time"
)

func main() {
	go stuff()
	fmt.Println("main thread?", macmain.IsMainThread())
	macmain.RunLoop()

	// http://stackoverflow.com/questions/15853230/is-nsapp-terminateid-deprecated
	/* Do not bother to put final cleanup code in your application’s main() function—it
	** will never be executed. If cleanup is necessary, perform that cleanup in the
	** delegate’s applicationWillTerminate: method. */
	fmt.Println("main run-loop terminated, exiting...")
	time.Sleep(2 * time.Second)
}

func stuff() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("hello from background thread", macmain.IsMainThread())
	time.Sleep(500 * time.Millisecond)
	macmain.Submit(func() {
		fmt.Println("hello from main thread", macmain.IsMainThread())
	})
	time.Sleep(500 * time.Millisecond)
	macmain.StopLoop()
}
