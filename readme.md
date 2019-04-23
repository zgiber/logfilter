# What?

Display logs which are mixed JSON & text outputs in a human readable format. It's written specifically for one use case but displays any JSON.
Log entries must be newline delimited, one JSON log must be a single line.

There are a few fields which it does understand and prints with special formatting:
- time (RFC3339)
- level string
- msg   string

# How?

## Install the tool

`go get -u github.com/zgiber/logfilter`

## Feed it with things

cat my_mixed.log | logfilter
