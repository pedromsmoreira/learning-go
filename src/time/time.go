//Go offers extensive support for times and durations;

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println
	// We’ll start by getting the current time.

	now := time.Now()
	p(now)
	// You can build a time struct by providing the year, month, day, etc. Times are always associated with a Location, i.e. time zone.

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)
	// You can extract the various components of the time value as expected.

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())
	// The Monday-Sunday Weekday is also available.

	p(then.Weekday())
	// These methods compare two times, testing if the first occurs before, after, or at the same time as the second, respectively.

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))
	// The Sub methods returns a Duration representing the interval between two times.

	diff := now.Sub(then)
	p(diff)
	// We can compute the length of the duration in various units.

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())
	// You can use Add to advance a time by a given duration, or with a - to move backwards by a duration.

	p(then.Add(diff))
	p(then.Add(-diff))

	// Use time.Now with Unix or UnixNano to get elapsed time since the Unix epoch
	// in seconds or nanoseconds, respectively.

	now2 := time.Now()
	secs := now2.Unix()
	nanos := now2.UnixNano()
	fmt.Println(now2)
	// Note that there is no UnixMillis, so to get the milliseconds since epoch you’ll
	// need to manually divide from nanoseconds.

	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)
	// You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

	p2 := fmt.Println
	// Here’s a basic example of formatting a time according to RFC3339,
	// using the corresponding layout constant.

	t := time.Now()
	p2(t.Format(time.RFC3339))
	// Time parsing uses the same layout values as Format.

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p2(t1)
	// Format and Parse use example-based layouts.
	// Usually you’ll use a constant from time for these layouts, but you can also supply custom layouts.
	// Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006 to show the pattern
	// with which to format/parse a given time/string.
	// The example time must be exactly as shown: the year 2006, 15 for the hour,
	//  Monday for the day of the week, etc.

	p2(t.Format("3:04PM"))
	p2(t.Format("Mon Jan _2 15:04:05 2006"))
	p2(t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p2(t2)
	// For purely numeric representations you can also use standard string formatting
	// with the extracted components of the time value.

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	// Parse will return an error on malformed input explaining the parsing problem.

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p(e)
}
