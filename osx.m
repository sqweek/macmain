#import <Cocoa/Cocoa.h>
#include "_cgo_export.h"

int
isMainThread()
{
	return [NSThread isMainThread];
}

void
runLoop()
{
	/* cocoa initialises its locks & other concurrency protections only if it
	** notices an NSThread being spawned. So, start one that does nothing.
	** ref: https://developer.apple.com/library/mac/documentation/Cocoa/Conceptual/Multithreading/CreatingThreads/CreatingThreads.html
	** (section "Protecting the Cocoa Frameworks") */
	NSThread* nop = [NSThread alloc];
	[[nop init] start];

	[[NSApplication sharedApplication] run];
}

void
stopLoop()
{
	[[NSApplication sharedApplication] terminate:nil];
}

void
postGoFuncEvent()
{
//	[[NSObject alloc] performSelectorOnMainThread:@selector(runGoFunc)];
	/* it sounded like dispatch_async might be a newer API? I found it
	** before I got the performSelector approach to work... */
	dispatch_async(dispatch_get_main_queue(), ^{ runGoFunc(); });
}
