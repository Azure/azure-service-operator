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
	if err := fn(); err != nil {
		// allow early exit
		if v, ok := err.(*StopErr); ok {
			return v.Err
		}
		if timeout > 0 {
			time.Sleep(sleep)
			return Retry(timeout-sleep, sleep, fn)
		}
		return err
	}
	return nil
}
