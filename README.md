# goServe
a simple http server capable of serving resources and pages.
## disclaimer
this software is **NOT** intended for production use, or any use at all at the moment.
i did not consider any security threat while writing this, it's just a side project 
useful to get a simple http server up and running with as little configuration as possible
and with a small resource footprint on the system.
## how to use
after building you just need to execute the generated executable.
if non arguments are given then the server will serve the "web" directory from the current working directory.
the first argument has to be the directory you want to serve and the second is the port on which you want to serve the directory.
