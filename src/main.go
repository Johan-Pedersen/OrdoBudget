package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func main() {
	spreadsheetId := "1Dg3qfLZd3S2ISqYLA7Av-D3njmiWPlcq-tQAodhgeAc"
	ctx := context.Background()
	b, err := os.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/drive")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	readRange := "Udtrœk!B2:C"
	valRange, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}

	// requests := make([]*sheets.Request, 0)
	for i, elm := range valRange.Values {
		fmt.Println(i, " Udtrœk ", elm[1], " beløb ", elm[0])
		// i := int64(i)
		// req := cutPasteSingleReq(i+1, 1, i+1,
		// 	int64(textToCol(elm[0].(string))))
		//
		// requests = append(requests, req)
	}

	// cutPasteReq := cutPasteSingleReq(1, 1, 5, 1)
	// Create the BatchUpdateRequest
	// batchUpdateReq := &sheets.BatchUpdateSpreadsheetRequest{
	// 	Requests: requests,
	// }

	// Execute the BatchUpdate request
	// _, err = srv.Spreadsheets.BatchUpdate(spreadsheetId, batchUpdateReq).Context(ctx).Do()

	// if err != nil {
	// 	log.Fatalf("Unable to perform CutPaste operation: %v", err)
	// }
	log.Println("Data moved successfully!")
}

func cutPasteSingleReq(fromRow, fromCol, toRow, toCol int64) *sheets.Request {
	cutPasteReq := &sheets.Request{
		CutPaste: &sheets.CutPasteRequest{
			Source: &sheets.GridRange{
				EndColumnIndex:   fromCol + 1,
				EndRowIndex:      fromRow + 1,
				SheetId:          1472288449,
				StartColumnIndex: fromCol,
				StartRowIndex:    fromRow,
			},
			Destination: &sheets.GridCoordinate{
				ColumnIndex: toCol,
				RowIndex:    toRow,
				SheetId:     1472288449,
			},
			PasteType: "PASTE_NORMAL", // Adjust paste type as needed
		},
	}
	return cutPasteReq
}
