package main

import (
	"fmt"
	"time"
)

func main() {
	// 1: Current time
	now := time.Now()
	fmt.Println("Now:", now)

	// 2: Extract parts
	fmt.Println("Year:", now.Year())
	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Hour:", now.Hour())
	fmt.Println("Minute:", now.Minute())
	fmt.Println("Second:", now.Second())

	// 3: Format time (Go uses special layout - exact reference time)
	fmt.Println("Formatted:", now.Format("2006-01-02 15:04:05"))

	// 4: Custom format (DD/MM/YYYY)
	fmt.Println("Custom Format:", now.Format("02/01/2006"))

	// 5: Parse string to time
	str := "2025-06-27 21:30:00"
	t, err := time.Parse("2006-01-02 15:04:05", str)
	if err == nil {
		fmt.Println("Parsed time:", t)
	}

	// 6: Add & Subtract duration
	after1Hour := now.Add(time.Hour)
	before30Min := now.Add(-30 * time.Minute)
	fmt.Println("After 1 hour:", after1Hour)
	fmt.Println("30 mins ago:", before30Min)

	// 7: Time difference (Duration)
	diff := after1Hour.Sub(now)
	fmt.Println("Duration between:", diff)

	// 8: Sleep for 2 seconds
	fmt.Print("Sleeping...")
	time.Sleep(2 * time.Second)
	fmt.Println(" Done!")

	// 9: Unix timestamp
	fmt.Println("Unix timestamp:", now.Unix())
	fmt.Println("Unix Milli:", now.UnixMilli())
}
