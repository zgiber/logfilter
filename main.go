package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	red    = "\u001b[31m"
	green  = "\u001b[32m"
	yellow = "\u001b[33m"
	white  = "\u001b[37m"
	reset  = "\u001b[0m"
)

func main() {
	// read from stdin line by line
	scn := bufio.NewScanner(os.Stdin)

	for scn.Scan() {
		m := logEntry{}
		b := scn.Bytes()
		err := json.Unmarshal(b, &m)
		if err != nil {
			fmt.Println(scn.Text())
			continue
		}

		fmt.Println(m)
	}

	if err := scn.Err(); err != nil {
		fmt.Println(err)
	}
}

type logEntry map[string]interface{}

func (e logEntry) String() string {
	// maintain the ordering
	// for some known fields
	expectedFields := []string{
		"time", "level", "msg", "file", "func", "trace_id"}

	out := ""
	colored := wrap(white)
	for _, fieldKey := range expectedFields {
		v, exists := e[fieldKey]
		if !exists {
			continue
		}

		// basic formatting for known fields
		switch fieldKey {
		case "time":
			out = fmt.Sprint(v)
		case "level":
			switch strings.ToLower(fmt.Sprint(v)) {
			case "debug":
				colored = wrap(green)
			case "info":
			case "warning":
				colored = wrap(yellow)
			case "error":
				colored = wrap(red)
			}
			out = fmt.Sprintf("%s [%s]", out, v)
		case "msg":
			out = fmt.Sprintf("%s '%s'", out, v)
		case "file":
			out = fmt.Sprintf("%s file=%s", out, v)
		case "func":
			out = fmt.Sprintf("%s func=%s", out, v)
		case "trace_id":
			out = fmt.Sprintf("%s trace_id=%s", out, v)
		}
		delete(e, fieldKey)
	}

	// custom fields in whatever order
	for k, v := range e {
		out = fmt.Sprintf("%s %s=%v", out, k, v)
	}

	return colored(out)
}

func wrap(color string) func(string) string {
	return func(text string) string {
		return fmt.Sprintf("%s%s%s", color, text, reset)
	}
}
