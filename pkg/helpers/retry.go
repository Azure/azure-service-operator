// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package helpers

import "time"

func NewStop(e error) *StopErr {
	return &StopErr{e}
}

type StopErr struct {
	Err error
}

func (s *StopErr) Error() string {
	return s.Err.Error()
}

func Retry(timeout time.Duration, sleep time.Duration, fn func() error) error {
	started := time.Now()
	for {
		err := fn()
		if err == nil {
			return nil
		}
		// allow early exit
		if v, ok := err.(*StopErr); ok {
			return v.Err
		}
		if time.Since(started) >= timeout {
			return err
		}
		time.Sleep(sleep)
	}
}
