/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package testcommon

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/go-logr/logr"
)

// TODO: I'm not sure the best way to configure this so for now I'm just going to set it to 5
var LogLevel = 5

var _ logr.LogSink = &TestLogger{}

func NewTestLogger(t *testing.T) logr.Logger {
	panicCount := 0
	return logr.New(&TestLogger{
		t:       t,
		values:  nil,
		enabled: true,
		name:    "",

		// There are known logging races that can cause the envtest controller to log after a test has exited. This most commonly
		// happens when the test context cleanup code is running. It uses DeleteResourceAndWait, and it is technically possible that
		// the "Wait" part notices that the resource is deleted and allows the test to complete before the reconciler has
		// actually finished its final delete reconcile pass. In this case a small number of logs (1-2) can be emitted after the test
		// has technically finished. This can easily be reproduced by forcing a wait in the controller before it emits its final logs.
		// Note that multiple controllers in test all share the same logger, so technically (if we get really unlucky) this can still
		// be hit. You can increase the ignoredPanicMax to reduce the probability of that. With that said, the max is low now because
		// we really don't (in general) want to tolerate this situation as it means we lose test logs and have stuff happening after
		// a test should have ended.
		ignoredPanicMax:   2,
		ignoredPanicCount: &panicCount,
	})
}

// TestLogger is a logr.Logger wrapper around t.Log, so that we can use it in the controller
type TestLogger struct {
	t      *testing.T
	values []interface{}

	enabled bool
	name    string

	ignoredPanicCount *int // ptr to share state across copies
	ignoredPanicMax   int
}

func (_ TestLogger) Init(info logr.RuntimeInfo) {
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
		t:                 t.t,
		enabled:           t.enabled,
		values:            clonedValues,
		name:              t.name,
		ignoredPanicMax:   t.ignoredPanicMax,
		ignoredPanicCount: t.ignoredPanicCount,
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

func (t *TestLogger) handlePanic() {
	if err := recover(); err != nil {
		errStr, ok := err.(string)

		if !ok || !strings.Contains(errStr, "Log in goroutine after") {
			panic(err)
		}

		// increment count of panics
		*t.ignoredPanicCount += 1

		if *t.ignoredPanicCount > t.ignoredPanicMax {
			panic(err)
		}

		// If we're under the count we swallow the panic
	}
}

func (t *TestLogger) Info(level int, msg string, keysAndValues ...interface{}) {
	defer t.handlePanic()

	if t.Enabled(level) {
		t.t.Helper()

		b := &bytes.Buffer{}

		header := t.makeHeader()
		kvListFormat(b, t.values...)
		kvListFormat(b, keysAndValues...)

		t.t.Log(fmt.Sprintf("%s \"msg\"=\"%s\"%s", header, msg, b))
	}
}

func (t *TestLogger) Enabled(_level int) bool {
	return t.enabled
}

func (t *TestLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	defer t.handlePanic()

	if t.Enabled(0) {
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
		return logr.New(result)
	}

	result.enabled = false
	return logr.New(result)
}

func (t *TestLogger) WithValues(keysAndValues ...interface{}) logr.LogSink {
	result := t.clone()
	result.values = append(result.values, keysAndValues...)

	return result
}

func (t *TestLogger) WithName(name string) logr.LogSink {
	result := t.clone()
	result.name = name
	return result
}
