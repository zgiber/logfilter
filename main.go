package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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
	// maintain the ordering using the slice
	// for known fields
	expectedFields := []string{
		"time", "level", "msg", "file", "func", "trace_id"}

	out := ""
	for _, fieldKey := range expectedFields {
		v, exists := e[fieldKey]
		if !exists {
			continue
		}

		// some basic formatting for known fields
		switch fieldKey {
		case "time":
			out = fmt.Sprint(v)
		case "level":
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

	return out
}
