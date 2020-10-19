# GoStats
Simple closure function to provide average, min, max of time.Duration over configurable bin size.

```
package main

import (
	"fmt"
	"github.com/klaxxon/gostats"
	"time"
)

// We will average over the last 10 samples
var saveSleepTime = gostats.GetDurationStatsFunc(10)

func timerTest(a int) {
	start := time.Now()
	defer saveSleepTime(time.Since(start))
	fmt.Printf("Sleeping %d mS\n", a)
	time.Sleep(time.Duration(a) * time.Millisecond)
}

func main() {
	// Measure how long each call takes also
	calls := gostats.GetDurationStatsFunc(10)
	for a := 0; a < 10; a++ {
		start := time.Now()
		timerTest(a * 100)
		calls(time.Since(start))
	}
	// Call with zero to get results without affecting contents
	fmt.Printf("%v\n", saveSleepTime(0))
	fmt.Printf("%v\n", calls(0))
}

```

