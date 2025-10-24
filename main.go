package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/islekcaganmert/projectile-to-height/internal/ml"
	projectilephysics "github.com/islekcaganmert/projectile-to-height/internal/projectile_physics"
)

type ScreenKeeper struct {
	Widgets      map[string]fyne.CanvasObject
	LastUpdate   int
	TimerStart   *time.Time
	Measurements []float64
}

func main() {
	a := app.New()
	w := a.NewWindow("Projectile to Height Converter")
	w.Resize(fyne.NewSize(500, 400))

	screenKeeper := &ScreenKeeper{
		Widgets:      make(map[string]fyne.CanvasObject),
		LastUpdate:   0,
		TimerStart:   nil,
		Measurements: make([]float64, 6),
	}

	screenKeeper.Widgets["m1"] = widget.NewLabel("Measurement 1: -.--s")
	screenKeeper.Widgets["m2"] = widget.NewLabel("Measurement 2: -.--s")
	screenKeeper.Widgets["m3"] = widget.NewLabel("Measurement 3: -.--s")
	screenKeeper.Widgets["m4"] = widget.NewLabel("Measurement 4: -.--s")
	screenKeeper.Widgets["m5"] = widget.NewLabel("Measurement 5: -.--s")
	screenKeeper.Widgets["m6"] = widget.NewLabel("Measurement 6: -.--s")

	screenKeeper.Widgets["avg"] = widget.NewLabel("Average: -.--s")
	screenKeeper.Widgets["adj"] = widget.NewLabel("Adjusted Average: -.--s")
	screenKeeper.Widgets["height"] = widget.NewLabel("Calculated Height: -.--m")

	screenKeeper.Widgets["button"] = widget.NewButton("Start Timer", screenKeeper.ButtonPressed)

	w.SetContent(container.NewVBox(
		screenKeeper.Widgets["m1"], screenKeeper.Widgets["m2"], screenKeeper.Widgets["m3"],
		screenKeeper.Widgets["m4"], screenKeeper.Widgets["m5"], screenKeeper.Widgets["m6"],
		widget.NewSeparator(),
		screenKeeper.Widgets["avg"], screenKeeper.Widgets["adj"], screenKeeper.Widgets["height"],
		widget.NewSeparator(),
		screenKeeper.Widgets["button"],
	))
	w.ShowAndRun()
}

func (screenKeeper *ScreenKeeper) UpdateMeasurements(num int) {
	if screenKeeper.TimerStart == nil {
		now := time.Now()
		screenKeeper.TimerStart = &now
		screenKeeper.Widgets[fmt.Sprintf("m%d", num)].(*widget.Label).SetText(
			fmt.Sprintf("Measurement %d: Timing...", num))
		screenKeeper.Widgets["button"].(*widget.Button).SetText("Stop Timer")
	} else {
		duration := time.Since(*screenKeeper.TimerStart).Seconds()
		screenKeeper.Widgets[fmt.Sprintf("m%d", num)].(*widget.Label).SetText(
			fmt.Sprintf("Measurement %d: %.3fs", num, duration))
		screenKeeper.Measurements[num-1] = duration
		screenKeeper.TimerStart = nil
		screenKeeper.Widgets["button"].(*widget.Button).SetText("Start Timer")
		screenKeeper.LastUpdate++
	}
}

func (screenKeeper *ScreenKeeper) ButtonPressed() {
	switch screenKeeper.LastUpdate {
	case 0:
		screenKeeper.UpdateMeasurements(1)
	case 1:
		screenKeeper.UpdateMeasurements(2)
	case 2:
		screenKeeper.UpdateMeasurements(3)
	case 3:
		screenKeeper.UpdateMeasurements(4)
	case 4:
		screenKeeper.UpdateMeasurements(5)
	case 5:
		screenKeeper.UpdateMeasurements(6)
		if screenKeeper.TimerStart == nil {
			avg := ml.GuessTrueValue(screenKeeper.Measurements, 1000, 0.1)
			screenKeeper.Widgets["avg"].(*widget.Label).SetText(fmt.Sprintf("Average: %.3fs", avg))
			adj := ml.GuessTrueValue(screenKeeper.Measurements, 2000, 0.05)
			screenKeeper.Widgets["adj"].(*widget.Label).SetText(fmt.Sprintf("Adjusted Average: %.3fs", adj))
			height := projectilephysics.CalculateHeightFromFallDuration(adj)
			screenKeeper.Widgets["height"].(*widget.Label).SetText(fmt.Sprintf("Calculated Height: %.2fm", height))
			screenKeeper.LastUpdate = 0
		}
	default:
		// Do nothing
	}
}
