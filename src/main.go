package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	excrptgrps "budgetAutomation/src/excrptGrps"
	"budgetAutomation/src/util"

	req "budgetAutomation/src/requests"

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

	excrptgrps.InitExcrptGrps()

	excrptgrps.PrintExcrptGrps()

	// Get Date, Amount and description
	readRangeExrpt := "Udtrœk!A2:D"
	valRange, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRangeExrpt).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}

	// Which month from 1-12 should be handled
	var month int64 = -1
	fmt.Println("Specify month:")
	fmt.Scan(&month)
	// requests := make([]*sheets.Request, 0)

	isRightMonth := false

	// account balance
	accBalance := -1.0

	// find Excerpt Total for current month.
	for _, elm := range valRange.Values {

		date, description := elm[0].(string), elm[2].(string)

		// s := strings.ReplaceAll(elm[1].(string), ",", ".")

		amount, err := strconv.ParseFloat(elm[1].(string), 64)

		if err != nil {
			log.Println("Could not read amount for", date, ":", description)
		} else {
			// Get excerpt month
			if date != "Reserveret" {

				exrptMonth, err := strconv.ParseInt(strings.Split(date, "/")[1], 0, 64)
				if err != nil {
					log.Fatal("Could not read excerpt date", err)
				}

				if month == exrptMonth {
					isRightMonth = true
					if accBalance == -1.0 {
						s := strings.ReplaceAll(elm[3].(string), ",", ".")

						accBalance, err = strconv.ParseFloat(s, 64)
						if err != nil {
							log.Fatalln("Could not read account balance")
						}
					}

				} else if exrptMonth < month {
					break
				}
				if isRightMonth {
					excrptgrps.UpdateExcrptTotal(date, description, amount)
				} // else {
				// excrptgrps.UpdateResume(date, description, "Not handled", amount)
				//	}
			}
		}
	}
	excrptgrps.PrintExcrptGrpTotals()

	// Find excerpt grps to insert at
	readRangeInserRow := "A1:A"
	rows, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRangeInserRow).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}

	var updateReqs []*sheets.Request

	println("****")
	for i, elm := range rows.Values {
		if len(elm) != 0 {

			total, notFoundErr := excrptgrps.GetTotal(elm[0].(string))

			fmt.Println("notFoundErr", notFoundErr)

			if notFoundErr == nil {
				fmt.Println("total: ", total)
				if total != 0.0 {
					updateReqs = append(updateReqs, req.SingleUpdateReq(total, int64(i), util.MonthToColInd(month), 1685114351))
				} else {
					updateReqs = append(updateReqs, req.SingleUpdateReqBlank(int64(i), util.MonthToColInd(month), 1685114351))
				}
			} else if strings.EqualFold(strings.Trim(elm[0].(string), " "), "Faktisk balance") {
				updateReqs = append(updateReqs, req.SingleUpdateReq(accBalance, int64(i), util.MonthToColInd(month), 1685114351))
			}
		}
	}

	// i := int64(i)
	// req := cutPasteSingleReq(i+1, 1, i+1,
	// 	int64(textToCol(elm[0].(string))))
	//
	// requests = append(requests, req)
	// }

	// updatereq := req.MultiUpdateReq([]float64{5.0, 6.0, 7.0, 9.0}, 0, 5, 1472288449)
	// cutPasteReq := cutPasteSingleReq(1, 1, 5, 1)
	// Create the BatchUpdateRequest
	batchUpdateReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: updateReqs,
	}

	// Execute the BatchUpdate request
	_, err = srv.Spreadsheets.BatchUpdate(spreadsheetId, batchUpdateReq).Context(ctx).Do()

	if err != nil {
		log.Fatalf("Unable to perform update operation: %v", err)
	}
	log.Println("Data moved successfully!")

	excrptgrps.PrintResume()
}
