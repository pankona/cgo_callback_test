

#include <stdio.h>
#include <unistd.h>

#include "libhello.h"

typedef void (*cbFunc)();

static void
callback(int in_arg) {
	printf("this is callback function of C. in_arg = %d\n", in_arg);
}

int
main() {
	Hello(callback);
	sleep(2);
	InvokeCallback();
	sleep(2);
	InvokeCallback();
	sleep(2);
	InvokeCallback();
	return 0;
}

