package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/islekcaganmert/projectile-to-height/internal/ml"
)

func RunRepeatedMeasurements(trials int, epochs int, lr float64) float64 {
	var durations []float64
	for i := 0; i < trials; i++ {
		fmt.Print("Press Enter to start the timer...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		start := time.Now()
		fmt.Print("Press Enter to stop the timer...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		end := time.Now()
		duration := end.Sub(start).Seconds()
		fmt.Printf("Calculated duration: %.6f seconds\n", duration)
		durations = append(durations, duration)
	}
	return ml.GuessTrueValue(durations, epochs, lr)
}
