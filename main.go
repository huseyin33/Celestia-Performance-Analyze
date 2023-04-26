package main

//"github.com/shirou/gopsutil/cpu"
import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"time"
)

func main() {

	// run for a day and take the cpu usage to make it graph
	usage, _ := cpu.Percent(time.Hour*6, true)
	fmt.Println(usage)
	// Define the data as a plotter.Values
	p := plot.New()

	// Create bar chart data
	data := plotter.Values(usage)

	// Create a new bar chart plotter
	bar, err := plotter.NewBarChart(data, vg.Points(10))
	if err != nil {
		panic(err)
	}

	// Add the bar chart plotter to the plot
	p.Add(bar)

	// Set the plot title and labels
	p.Title.Text = "CPU Core Usage For 6 Hours"
	p.X.Label.Text = "CPU Core"
	p.Y.Label.Text = "Usage Percentage"

	// Save the plot to a PNG file
	err = p.Save(4*vg.Inch, 4*vg.Inch, "cpuPercentageChart.png")
	if err != nil {
		panic(err)
	}
}
