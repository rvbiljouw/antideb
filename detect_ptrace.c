#include <stdio.h>
#include <sys/ptrace.h>

#include "antideb.h"

#ifndef PTRACE_TRACEME
#define PTRACE_TRACEME 0
#define PTRACE_DETACH 11
#endif

static int detect_ptrace(void)
{
    if (ptrace(PTRACE_TRACEME, 0, NULL, NULL) == -1)
    {
        return RESULT_YES;
    }

    ptrace(PTRACE_DETACH, 0, NULL, NULL);
    return RESULT_NO;
}
