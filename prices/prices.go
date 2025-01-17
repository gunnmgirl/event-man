package prices

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadPrices() error {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("Could not open file")
		fmt.Println(err)
		return err
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Could not read file")
		fmt.Println(err)
		file.Close()
		return err
	}

	prices := make([]float64, len(lines))

	for index, line := range lines {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			file.Close()
			return errors.New("could not convert price to float64 value")
		}
		prices[index] = price
	}

	job.InputPrices = prices
	return nil
}

func (job *TaxIncludedPriceJob) Process() {
	result := make(map[string]string)
	job.LoadPrices()

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}
