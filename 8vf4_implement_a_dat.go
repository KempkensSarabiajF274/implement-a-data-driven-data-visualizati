// 8vf4_implement_a_dat.go
// A Data-Driven Data Visualization Controller

package main

import (
	"encoding/csv"
	"fmt"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"io"
	"log"
)

// DataPoint represents a single data point
type DataPoint struct {
	X float64
	Y float64
}

// VisualizationController controls the data visualization
type VisualizationController struct {
	DataPoints []DataPoint
}

// NewVisualizationController creates a new instance of VisualizationController
func NewVisualizationController(dataPoints []DataPoint) *VisualizationController {
	return &VisualizationController{DataPoints: dataPoints}
}

// LoadData loads data from a CSV file
func LoadData(filename string) ([]DataPoint, error) {
	f, err := fopen(filename, "r")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	dataPoints := make([]DataPoint, 0)
	for _, record := range records {
		x, err := strconv.ParseFloat(record[0], 64)
		if err != nil {
			return nil, err
		}
		y, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, err
		}
		dataPoints = append(dataPoints, DataPoint{X: x, Y: y})
	}

	return dataPoints, nil
}

// VisualizeData visualizes the data
func (vc *VisualizationController) VisualizeData() {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	p.Title.Text = "Data Visualization"
	p.X.Label.Text = "X Axis"
	p.Y.Label.Text = "Y Axis"

	xVals := make([]float64, len(vc.DataPoints))
	yVals := make([]float64, len(vc.DataPoints))
	for i, dp := range vc.DataPoints {
		xVals[i] = dp.X
		yVals[i] = dp.Y
	}

	line, err := plotter.NewLine(xVals, yVals)
	if err != nil {
		log.Panic(err)
	}

	line.Color = vg.ColorNRGBA{R: 255, A: 255}
	line.Width = vg.Length(2)

	p.Add(line)

	err = p.Save(8*vg.Inch, 8*vg.Inch, "png", "data_visualization.png")
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	dataPoints, err := LoadData("data.csv")
	if err != nil {
		log.Fatal(err)
	}

	controller := NewVisualizationController(dataPoints)
	controller.VisualizeData()
}