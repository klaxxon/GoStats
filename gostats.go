package gostats

import (
  time
)

// avgTime returns the average and max duration in buffer
func avgTime(binsize int) func(time.Duration) DurationStats {
	bins := make([]time.Duration, binsize)
	var sum time.Duration
	var count int
	pos := 0
	return func(t time.Duration) DurationStats {
		if t != 0 {
			sum -= bins[pos]
			sum += t
			bins[pos] = t
			pos++
			if count < binsize {
				count++
			}
			if pos >= binsize {
				pos = 0
			}
		}
		if count == 0 {
			return DurationStats{}
		}

		min := time.Duration(1<<63 - 1)
		var max time.Duration
		for _, a := range bins {
			if a > max {
				max = a
			}
			if a < min {
				min = a
			}
		}
		return DurationStats{Avg: time.Duration(sum / time.Duration(count)), Max: max, Min: min}
	}
}
