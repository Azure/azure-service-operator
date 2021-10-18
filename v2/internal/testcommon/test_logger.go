/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/go-logr/logr"
)

// TODO: I'm not sure the best way to configure this so for now I'm just going to set it to 5
var LogLevel = 5

var _ logr.Logger = &TestLogger{}

func NewTestLogger(t *testing.T) logr.Logger {
	return &TestLogger{
		t:       t,
		values:  nil,
		enabled: true,
		name:    "",
	}
}

// TestLogger is a logr.Logger wrapper around t.Log, so that we can use it in the controller
type TestLogger struct {
	t      *testing.T
	values []interface{}

	enabled bool
	name    string
}

// kvListFormat was adapted from the klog method of the same name and formats a keysAndValues list into
// a sequence of "key"="value" pairs.
func kvListFormat(b *bytes.Buffer, keysAndValues ...interface{}) {
	for i := 0; i < len(keysAndValues); i += 2 {
		var v interface{}
		k := keysAndValues[i]
		if i+1 < len(keysAndValues) {
			v = keysAndValues[i+1]
		} else {
			v = "(MISSING)"
		}
		b.WriteByte(' ')
		if _, ok := v.(fmt.Stringer); ok {
			b.WriteString(fmt.Sprintf("%s=%q", k, v))
		} else {
			b.WriteString(fmt.Sprintf("%s=%#v", k, v))
		}
	}
}

func (t *TestLogger) clone() *TestLogger {
	var clonedValues []interface{}
	clonedValues = append(clonedValues, t.values...)

	result := &TestLogger{
		t:       t.t,
		enabled: t.enabled,
		values:  clonedValues,
		name:    t.name,
	}
	return result
}

// makeHeader makes a header that looks a lot (but not quite exactly) like the klog header
// The idea is that looking at test logs and looking at actual container logs should be basically
// the same.
func (t *TestLogger) makeHeader() string {
	now := time.Now()
	timeStr := now.Format(time.RFC3339)

	nameString := ""
	if t.name != "" {
		nameString = fmt.Sprintf(" %s", t.name)
	}

	severity := "I" // TODO: Technically there are multiple severities?

	return fmt.Sprintf("%s%s]%s", severity, timeStr, nameString)
}

func (t *TestLogger) Info(msg string, keysAndValues ...interface{}) {
	if t.Enabled() {
		t.t.Helper()

		b := &bytes.Buffer{}

		header := t.makeHeader()
		kvListFormat(b, t.values...)
		kvListFormat(b, keysAndValues...)

		t.t.Log(fmt.Sprintf("%s \"msg\"=\"%s\"%s", header, msg, b))
	}
}

func (t *TestLogger) Enabled() bool {
	return t.enabled
}

func (t *TestLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	if t.Enabled() {
		b := &bytes.Buffer{}
		header := t.makeHeader()
		kvListFormat(b, t.values...)
		kvListFormat(b, keysAndValues...)

		t.t.Log(fmt.Sprintf("%s \"msg\"=\"%s\" \"err\"=\"%s\"%s", header, msg, err, b))
	}
}

func (t *TestLogger) V(level int) logr.Logger {
	result := t.clone()
	if level <= LogLevel {
		return result
	}

	result.enabled = false
	return result
}

func (t *TestLogger) WithValues(keysAndValues ...interface{}) logr.Logger {
	result := t.clone()
	result.values = append(result.values, keysAndValues...)

	return result
}

func (t *TestLogger) WithName(name string) logr.Logger {
	result := t.clone()
	result.name = name
	return result
}
