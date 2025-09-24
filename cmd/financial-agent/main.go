package main

import (
	"context"
	"fmt"
	"log"

	"financial-agent/pkg/auth"
	"financial-agent/pkg/gemini"
	"financial-agent/pkg/gmail"
	"financial-agent/pkg/sheets"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Financial Agent Initialized")

	client := auth.SetupGoogleClient()

	sheetsSrv, err := sheets.CreateService(context.Background(), client)
	if err != nil {
		log.Fatalf("Unable to create Sheets service: %v", err)
	}

	//Spreadsheets ids you want the agent to consider
	spreadsheetIDs := []string{
		"example_spreadsheet_id",
		"example_spreadsheet_id",
	}

	var allData string
	for _, id := range spreadsheetIDs {
		data, err := sheets.ReadData(sheetsSrv, id)
		if err != nil {
			log.Printf("Could not read data from sheet %s: %v", id, err)
			continue
		}
		allData += data + "\n\n"
	}

	fmt.Println("--- Combined Data from Sheets ---")
	fmt.Println(allData)
	fmt.Println("---------------------------------")

	summary, err := gemini.Summarize(context.Background(), allData)
	if err != nil {
		log.Fatalf("Failed to summarize data: %v", err)
	}

	fmt.Println("--- Financial Summary ---")
	fmt.Println(summary)
	fmt.Println("-------------------------")

	reportTitle := "Financial Summary Report"
	newSheetURL, err := sheets.CreateNewSheetWithSummary(sheetsSrv, reportTitle, summary)
	if err != nil {
		log.Fatalf("Failed to create new sheet with summary: %v", err)
	}

	fmt.Printf("Summary report created: %s\n", newSheetURL)

	gmailSrv, err := gmail.CreateService(context.Background(), client)
	if err != nil {
		log.Fatalf("Unable to create Gmail service: %v", err)
	}
	//the email of the recipient you want the agent to send the report to
	recipientEmail := "evlogievkrasi@gmail.com"
	emailSubject := "Your Automated Financial Summary Report is Ready"
	emailBody := fmt.Sprintf(`
		<html>
			<body>
				<p>Hello,</p>
				<p>Your automated financial analysis is complete. The summary report has been generated and is available in a new Google Sheet.</p>
				<p>You can access the report here: <a href="%s">Financial Summary Report</a></p>
				<p>Best regards,<br>Your Financial AI Agent</p>
			</body>
		</html>
	`, newSheetURL)

	err = gmail.SendEmail(gmailSrv, recipientEmail, emailSubject, emailBody)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	}

	fmt.Printf("Email sent successfully to %s\n", recipientEmail)
}
