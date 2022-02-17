# sentry-test
A simple program to log a test event to your Sentry instance.

Pre-built binaries can be found on the [releases page](https://github.com/rxdn/sentry-test/releases).

# Running
Simply run the program with the `-dsn` flag specifying your Sentry DSN.

Example: `./sentry-test -dsn='https://abcdefghij123456789@sentry.example.com/1'`

# Building
The program can be build with `go build main.go` 
