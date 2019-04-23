package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// read from stdin line by line
	scn := bufio.NewScanner(os.Stdin)

	for scn.Scan() {
		m := map[string]interface{}{}
		b := scn.Bytes()
		err := json.Unmarshal(b, &m)
		if err != nil {
			fmt.Println(scn.Text())
			continue
		}

		se := &standardEntry{}
		err = json.Unmarshal(b, se)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(se)
	}

	if err := scn.Err(); err != nil {
		fmt.Println(err)
	}
}

func clean(m map[string]interface{}) {
	knownFields := []string{"time", "level", "file", "func", "msg", "trace_id", "traceid"}
	for _, field := range knownFields {
		delete(m, field)
	}
}

type standardEntry struct {
	Time         time.Time
	Level        string
	File         string
	Func         string
	Msg          string
	TraceID      string
	customFields map[string]interface{}
}

func (se *standardEntry) String() string {
	customFields := []string{}
	for k, v := range se.customFields {
		customFields = append(customFields, fmt.Sprintf("%s=%s", strings.ToTitle(k), v))
	}

	out := ""
	if !se.Time.IsZero() {
		out = fmt.Sprintf("%v - %s", se.Time, out)
	}

	if se.Level != "" {
		out = fmt.Sprintf("%s[%s]", out, strings.ToUpper(se.Level))
	}

	if se.File != "" {
		out = fmt.Sprintf("%s File=%s", out, se.File)
	}

	if se.Func != "" {
		out = fmt.Sprintf("%s Func=%s", out, se.Func)
	}

	return fmt.Sprintf("%s '%s' %s",
		out, se.Msg, strings.Join(customFields, " "))
}
