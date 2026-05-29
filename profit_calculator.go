package main

import(
	"fmt"
	"os"
	"errors"
) 

// Goals
// 1)Validate user inputs
// =>show error message and exit if invalid input is provided
// - No negative numbers
// - No zero
// 2)Store calculated results in to file

func main() {
	revenue, err := getUserInput("Revenue: ")
	if(err != nil){
		fmt.Println(err)
		return
	}
	// modifiedRevenue := fmt.Sprintf("%.2f", revenue)
    // os.WriteFile("reveneFile.txt", []byte(modifiedRevenue), 0644)

	expenses, err := getUserInput("Expenses: ")
	if(err != nil){
		fmt.Println(err)
		return
	}
	// modifiedExpense := fmt.Sprintf("%.2f", expenses)
    // os.WriteFile("expenseFile.txt", []byte(modifiedExpense), 0644)

	taxRate, err := getUserInput("Tax Rate: ")
	if(err != nil){
		fmt.Println(err)
		return
	}
	// modifiedTaxRate := fmt.Sprintf("%.2f", taxRate)
    // os.WriteFile("taxFile.txt", []byte(modifiedTaxRate), 0644)

	// if(err1 != null || err2 != null || err3 != null){
	// 	fmt.Println(err1)
	// 	return
	// }

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.3f\n", ratio)
	storeValuesInFile(ebt, profit, ratio)
}

func storeValuesInFile(ebt, profit, ratio float64){
	resultData := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(resultData), 0644)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if(userInput <= 0){
		return 0, errors.New("Invalid input Provided. Exiting the program.")
	}
	return userInput, nil
}
