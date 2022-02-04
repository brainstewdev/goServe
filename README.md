# goServe
goServe is a simple web server intended to be simple and fast to configure and use.
no configuration files are required and no tinkering involved- just build and run.
## how to build
you will need to build the executable yourself, if there are no releases available for your platform or OS. 
to build the Executable you just need to have go installed on your machine and execute
the go build command inside the directory in which you have cloned the repository.
## Usage
the program takes two flags: the directory to serve and the port on which to serve.
the directory can be set using the -d flag and it defaults to the web directory
in the working directory if not set. 
the port can be set using the -p flag and it will defaults to the port 80 if not set.
## disclaimer
this software is **NOT** intended for production use, or any use at all at the moment.
i did not consider any security threat while writing this, it's just a side project 
useful to get a simple http server up and running with as little configuration as possible
and with a small resource footprint on the system.

