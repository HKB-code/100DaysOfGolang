package prices

import (
	"fmt"

	"example.com/main/conversion"
	iomanager "example.com/main/ioManager"
)

// type TaxIncludedPriceJob struct {
// 	TaxRate          float64
// 	InputPrices      []float64
// 	TaxIncludedPrice map[string]string
// }

type TaxIncludedPriceJob struct {
	IoManager iomanager.IOManager `json:"-"`
	TaxRate          float64 `json:"tax_rate"`
	InputPrices      []float64 `json:"input_prices"`
	TaxIncludedPrice map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData()error{
	// lines,err:=filemanager.ReadFiles("prices/prices.txt")
	lines,err:=job.IoManager.ReadFiles()
if err != nil {
		fmt.Println(err)
		return err
}
price,err:= conversion.StringToFloat(lines)
if err != nil {
		fmt.Println(err)
		return err
}
	job.InputPrices = price
	return  nil
}


func (job *TaxIncludedPriceJob) Process()error{
	err:=job.LoadData()
	if err != nil {
		fmt.Println(err)
		return  err
}
result := make(map[string]string)
		for _, price := range job.InputPrices {
			taxIncludePrice:= price * (1 + job.TaxRate)
			result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",taxIncludePrice)
		}
		// fmt.Println(result)
		job.TaxIncludedPrice = result

		// filemanager.WriteResult(fmt.Sprintf("result_%.0f.json",job.TaxRate*100),job)
		return job.IoManager.WriteResult(job)
		
	}


// func New(taxRate float64) *TaxIncludedPriceJob {
// 	return &TaxIncludedPriceJob{
// 		InputPrices: []float64{10, 20, 30},
// 		TaxRate:     taxRate,
// 	}
// }
// func New(fm filemanager.FileManager,taxRate float64) *TaxIncludedPriceJob {
// 	return &TaxIncludedPriceJob{
// 		IoManager: fm,
// 		InputPrices: []float64{10, 20, 30},
// 		TaxRate:     taxRate,
// 	}
// }
func New(io iomanager.IOManager,taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IoManager: io,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}