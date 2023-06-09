package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Set taxID and filename
	taxID := "1234567890123"
	filename := taxID + "_pn_" + time.Now().Format("20060102") + "_001.txt"

	// Create file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Write column headers
	_, err = file.WriteString("PNRef|TAXID|PaymentDate|Status\n")
	if err != nil {
		fmt.Println("Unable to write to file:", err)
	}

	// Generate and write rows
	for i := 1; i <= 25; i++ {
		pn := fmt.Sprintf("PN%04d", i)
		date := time.Now().AddDate(0, 0, i-1).Format("2006-01-02")
		status := "COMPLETED"

		_, err = file.WriteString(fmt.Sprintf("%s|%s|%s|%s\n", pn, taxID, date, status))
		if err != nil {
			fmt.Println("Unable to write to file:", err)
		}
	}

	fmt.Println("File successfully written.")
}
