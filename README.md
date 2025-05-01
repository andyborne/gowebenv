This is a debugging tool for examining the environment that a process is
launched in, e.g. by Docker, `systemd`, `supervisor`, or some other process
manager.

Since a likely use case is to check that secrets are successfully passed in,

*NEVER EXPOSE THIS TO THE PUBLIC INTERNET.*

Environments usually contain sensitive information whether the operator
deliberately added any or not.

It is highly recommended to bind to `localhost` and a firewalled port, access
the web output only on the same local machine, and terminate the process
immediately after gathering the needed `env` data.

While running, the following pages are served:
`/args`          --  The ARGV elements from the process's invocation
`/env`           --  The process's environment variables
`/headers`       --  Request headers received
`/process`       --  The `PID` (process ID) and `PPID` (parent `PID`)
`/` (all others) --  This document
