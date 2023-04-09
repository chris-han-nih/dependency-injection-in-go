package solid

import (
	"fmt"
	"io"
)

type CalculatorGood struct {
	data map[string]float64
}

func (c *CalculatorGood) Calculate(path string) error {
	return nil
}

func (c *CalculatorGood) getData() map[string]float64 {
	return nil
}

type Printer interface {
	Output(data map[string]float64)
}

type DefaultPrinter struct {
	Writer io.Writer
}

func (d *DefaultPrinter) Output(data map[string]float64) {
	for path, result := range data {
		fmt.Fprintf(d.Writer, "%s -> %.1f\n", path, result)
	}
}

type CSVPrinter struct {
	Writer io.Writer
}

func (c *CSVPrinter) Output(data map[string]float64) {
	for path, result := range data {
		fmt.Fprintf(c.Writer, "%s,%.1f\n", path, result)
	}
}
