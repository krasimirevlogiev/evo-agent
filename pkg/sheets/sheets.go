package sheets

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// CreateService creates a new Google Sheets service client.
func CreateService(ctx context.Context, client *http.Client) (*sheets.Service, error) {
	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve Sheets client: %v", err)
	}
	return srv, nil
}

// ReadData reads all data from all sheets of a given spreadsheet.
func ReadData(srv *sheets.Service, spreadsheetId string) (string, error) {
	resp, err := srv.Spreadsheets.Get(spreadsheetId).Do()
	if err != nil {
		return "", fmt.Errorf("unable to retrieve spreadsheet %v: %v", spreadsheetId, err)
	}

	if len(resp.Sheets) == 0 {
		return "", fmt.Errorf("no sheets found in spreadsheet %v", spreadsheetId)
	}

	var allSheetsData string
	for _, sheet := range resp.Sheets {
		sheetName := sheet.Properties.Title
		readRange := sheetName

		valueResp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
		if err != nil {
			log.Printf("unable to retrieve data from sheet '%s' in spreadsheet %v: %v. Skipping sheet.", sheetName, spreadsheetId, err)
			continue
		}

		if len(valueResp.Values) > 0 {
			allSheetsData += fmt.Sprintf("--- Data from Sheet: %s ---\n", sheetName)
			for _, row := range valueResp.Values {
				for _, cell := range row {
					allSheetsData += fmt.Sprintf("%s\t", cell)
				}
				allSheetsData += "\n"
			}
			allSheetsData += "\n"
		}
	}

	if allSheetsData == "" {
		log.Printf("No data found in any sheet for spreadsheet %s.", spreadsheetId)
	}

	return allSheetsData, nil
}

// CreateNewSheetWithSummary creates a new Google Sheet and populates it with the summary.
func CreateNewSheetWithSummary(srv *sheets.Service, title string, summary string) (string, error) {
	spreadsheet := &sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: title,
		},
	}
	spreadsheet, err := srv.Spreadsheets.Create(spreadsheet).Do()
	if err != nil {
		return "", fmt.Errorf("could not create spreadsheet: %v", err)
	}

	writeRange := "Sheet1!A1"
	var vr sheets.ValueRange
	vr.Values = append(vr.Values, []interface{}{summary})

	_, err = srv.Spreadsheets.Values.Update(spreadsheet.SpreadsheetId, writeRange, &vr).ValueInputOption("RAW").Do()
	if err != nil {
		//TODO implement cleanup
		log.Printf("Could not write to spreadsheet %s. Attempting to clean up.", spreadsheet.SpreadsheetId)
		return "", fmt.Errorf("could not write to spreadsheet: %v", err)
	}

	return spreadsheet.SpreadsheetUrl, nil
}

//TODO extend the agent to extract financial data from
//external sources and generate reports based on that with predictions also
