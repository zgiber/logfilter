# What?

Display logs which are mixed JSON & text outputs in a human readable format. It's written specifically for one use case but displays any JSON.
Log entries must be newline delimited, one JSON log entry must be a single line.

There are a few fields which it does understand and prints with special formatting:
- time (RFC3339)
- level (string)
- msg   (string)

# How?

## Install the tool

`go get -u github.com/zgiber/logfilter`

## Feed it with things

```
▸ logfilter git:(master) cat ~/jsonlog.json
{"level":"debug","file":"stuff.go:34","time":"2018-01-01T23:12:34.456Z","msg":"something happened"}
{"level":"info","file":"stuff.go:34","time":"2018-01-01T23:12:34.456Z","msg":"something happened"}
{"level":"warning","file":"stuff.go:34","time":"2018-01-01T23:12:34.456Z","msg":"something happened"}
{"level":"error","file":"stuff.go:34","time":"2018-01-01T23:12:35.456Z","msg":"something else happened","trace_id":"000001","my mood is":"great"}
01-01-2018T11:11:11.111 INFO Some process has started
01-01-2018T11:11:12.111 ERROR
  stacktrace
  comes in multiple lines
  which are not linked by anything
{"level":"INFO","file":"stuff.go:34","time":"2018-01-01T23:12:36.456Z","msg":"something else happened again"}
{"level":"ERROR","file":"stuff.go:34","time":"2018-01-01T23:12:37.456Z","msg":"something else happened yet again\n\tin multiple lines\n\tbecause we can"}

▸ logfilter git:(master) cat ~/jsonlog.json | logfilter
2018-01-01T23:12:34.456Z [debug] 'something happened' file=stuff.go:34
2018-01-01T23:12:34.456Z [info] 'something happened' file=stuff.go:34
2018-01-01T23:12:34.456Z [warning] 'something happened' file=stuff.go:34
2018-01-01T23:12:35.456Z [error] 'something else happened' file=stuff.go:34 trace_id=000001 my mood is=great
01-01-2018T11:11:11.111 INFO Some process has started
01-01-2018T11:11:12.111 ERROR
  stacktrace
  comes in multiple lines
  which are not linked by anything
2018-01-01T23:12:36.456Z [INFO] 'something else happened again' file=stuff.go:34
2018-01-01T23:12:37.456Z [ERROR] 'something else happened yet again
	in multiple lines
	because we can' file=stuff.go:34
```

# Notes

At the moment this is a single purpose tool, without much reason to have it changed.
One nice feature would be to support multiline JSON without too much overhead.
