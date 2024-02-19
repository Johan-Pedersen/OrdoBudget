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

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	// Own
	// https://docs.google.com/spreadsheets/d/1Dg3qfLZd3S2ISqYLA7Av-D3njmiWPlcq-tQAodhgeAc/edit#gid=0
	spreadsheetId := "1Dg3qfLZd3S2ISqYLA7Av-D3njmiWPlcq-tQAodhgeAc"
	// readRange := "Udtrœk!C2:C"
	// resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve data from sheet: %v", err)
	// }
	//
	// if len(resp.Values) == 0 {
	// 	fmt.Println("No data found.")
	// } else {
	// 	fmt.Println("Name, Major:")
	// for _, row := range resp.Values {
	// Print columns A and E, which correspond to indices 0 and 4.
	// fmt.Printf("%s\n", row[0])
	//
	// // Tekst : kolonne
	// udtrœksKlasser := map[string]string{
	// 	"Resturant besøg/takeway/mad på skole": "E",
	// 	"Helbred (fx tandlœge, medicin m.m.)":  "F",
	// 	"ANTON (foder, pleje, godbidder mv.)":  "G",
	// 	"Boligudstyr":                          "H",
	// 	"Elektronik":                           "I",
	// 	"Gaver":                                "J",
	// 	"Social arrangementer":                 "K",
	// 	"Rejser":                               "L",
	// 	"Spotify":                              "M",
	// 	"Transport, cykel, parkering m.m.":     "N",
	// 	"Golf":                                 "O",
	// 	"Dashlane":                             "P",
	// 	"Diverse":                              "Q",
	// 	"Skolebøger m.m.":                      "R",
	// 	"Tøj og sko":                           "S",
	// 	"Personlig pleje (kosmetik, frisør mv.)": "T",
	// 	"Briller abonnement ":                    "U",
	// 	"Icloud, onedrive, Prime":                "V",
	// }
	//
	// maping := map[string]string{
	// 	"skatteguiden.dk GF2023":        "Diverse",
	// 	"Silkeborg Ry Golfklub -":       "Golf",
	// 	"MobilePay Rejsekort":           "Transport, cykel, parkering m.m.",
	// 	"SILKEBORG RY GOLFKLUDen 03.02": "Golf",
	// }
	//
	// println(udtrœksKlasser[maping[row[0].(string)]])

	// cutpaste := sheets.CutPasteRequest{
	// 	Destination: &sheets.GridCoordinate{
	// 		ColumnIndex: 4,
	// 		RowIndex:    2,
	// 		SheetId:     1,
	// 	},
	// 	PasteType: "PASTE_NORMAL",
	// 	Source: &sheets.GridRange{
	// 		EndColumnIndex:   1,
	// 		EndRowIndex:      33,
	// 		SheetId:          1,
	// 		StartColumnIndex: 1,
	// 		StartRowIndex:    2,
	// 	},
	// }
	// requests := []sheets.Request{
	// 	CutPaste: cutpaste,
	// }
	//
	// request := sheets.BatchUpdateSpreadsheetRequest{
	// 	IncludeSpreadsheetInResponse: false,
	// 	Requests: requests,
	// }
	// }
	// Create the CutPasteRequest
	// }
	cutPasteReq := &sheets.Request{
		CutPaste: &sheets.CutPasteRequest{
			Source: &sheets.GridRange{
				EndColumnIndex:   0,
				EndRowIndex:      31,
				SheetId:          1472288449,
				StartColumnIndex: 0,
				StartRowIndex:    1,
			},
			Destination: &sheets.GridCoordinate{
				ColumnIndex: 5,
				RowIndex:    1,
				SheetId:     1472288449,
			},
			PasteType: "PASTE_NORMAL", // Adjust paste type as needed
		},
	}

	// Create the BatchUpdateRequest
	batchUpdateReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{cutPasteReq},
	}

	// Execute the BatchUpdate request
	_, err = srv.Spreadsheets.BatchUpdate(spreadsheetId, batchUpdateReq).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to perform CutPaste operation: %v", err)
	}

	log.Println("Data moved successfully!")
}
