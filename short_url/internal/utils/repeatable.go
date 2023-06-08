package utils

import "time"

// DoWhitTries receives a function and call it attempt times with a delay between calls
func DoWhitTries(f func() error, attempt int, delay time.Duration) (err error) {
	for attempt > 0 {
		if err = f(); err != nil {
			time.Sleep(delay)
			attempt--
			continue
		}
		return nil
	}
	return
}
