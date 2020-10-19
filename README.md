# GoStats
Simple closure function to provide average, min, max of time.Duration over configurable bin size.

```
// We will average over the last 10 samples
var saveSleepTime = gostats.GetDurationStatsFunc(10)

func timerTest(a) {
  start := time.Now()
  defer saveSleepTime(time.Since(start))
  time.Sleep(time.Duration(a)*time.Second)
}


func main() {
  for a:=0;a<10;a++ {
    timerTest()
  }
  // Call with zero to get results without affecting contents
  s := saveSleepTime(0)
  fmt.Printf("%v\n", s)
}
```

