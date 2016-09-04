# ping-fastcgi

Send an http get request to a fcgi instance and verify the response contains a certain string.

## Why
I wanted a simple check of the health of a running php-fpm process in a kubernetes pod.
Since php-fpm has a [ping.path](http://php.net/manual/en/install.fpm.configuration.php) configuration option this:
- bypases any other dependencies (no http server needed, direct fcgi communication)
- no need to write any php script (although you can !)


## Usage

```
NAME:
   ping-fastcgi - Commandline tool to sent a http ping request to fastcgi

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --host value      Ip or hostname of the fastcgi process (default: "127.0.0.1")
   --port value      Port of the fastcgi process (default: 9000)
   --uri value       Uri of the ping endpoint, default is /ping (default: "/ping")
   --response value  Expected response from the endpoint (default: "pong")
   --help, -h        show help
   --version, -v     print the version

```

# Todo
- timeout on connect
- timeout on receive answer

I didn't bother implementing both since fo my usecase this is handled by kubernetes.


# Credits

This small thing would not be possible without fcgiclient written by Junqing Tan

