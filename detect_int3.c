#include <signal.h>
#include <stdio.h>
#include <stdlib.h>

int debuging;

static void sig_handler(int sig)
{
    debuging = 1;
}

static int detect_int3()
{
#ifdef __amd64__
    debuging = 0;
    debuging++; // for optimize by compile is off
    signal(SIGTRAP, sig_handler);
    __asm__("int3");
    debuging--; // for optimize by compile is off
    if (debuging == 0)
    {
        return 0;
    }
    return 1;
#else
    return 0;
#endif
}
