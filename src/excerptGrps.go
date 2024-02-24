package main

import "fmt"

var excerptGrpSums = map[string]float64{
	"Resturant besøg/takeway/mad på skole": 0.0,
	"Helbred (fx tandlœge, medicin m.m.)":  0.0,
	"ANTON (foder, pleje, godbidder mv.)":  0.0,
	"Boligudstyr":                          0.0,
	"Elektronik":                           0.0,
	"Gaver":                                0.0,
	"Social arrangementer":                 0.0,
	"Rejser":                               0.0,
	"Spotify":                              0.0,
	"Transport, cykel, parkering m.m.":     0.0,
	"Golf":                                 0.0,
	"Dashlane":                             0.0,
	"Diverse":                              0.0,
	"Skolebøger m.m.":                      0.0,
	"Tøj og sko":                           0.0,
	"Personlig pleje (kosmetik, frisør mv.)": 0.0,
	"Briller abonnement ":                    0.0,
	"Icloud, onedrive, Prime":                0.0,
}

var excerptGrpRows = map[string]int{
	"Resturant besøg/takeway/mad på skole": 47,
	"Helbred (fx tandlœge, medicin m.m.)":  48,
	"ANTON (foder, pleje, godbidder mv.)":  49,
	"Boligudstyr":                          50,
	"Elektronik":                           51,
	"Gaver":                                52,
	"Social arrangementer":                 53,
	"Rejser":                               54,
	"Spotify":                              55,
	"Transport, cykel, parkering m.m.":     56,
	"Golf":                                 57,
	"Dashlane":                             58,
	"Bryllup":                              59,
	"Diverse":                              60,
	"Skolebøger m.m.":                      61,
	"Tøj og sko":                           62,
	"Personlig pleje (kosmetik, frisør mv.)": 63,
	"Briller abonnement ":                    64,
	"Icloud, onedrive, Prime":                65,
}

var excerptGrps = map[string]string{
	"skatteguiden.dk GF2023":        "Diverse",
	"Silkeborg Ry Golfklub -":       "Golf",
	"MobilePay Rejsekort":           "Transport, cykel, parkering m.m.",
	"SILKEBORG RY GOLFKLUDen 03.02": "Golf",
	"Akademiker a-kasse":            "Fagforening og A-kasse",
	"Til madkonto":                  "Mad",
	"Johan til fælles":              "Fælles Udgifter",
	"Johan forsikringer":            "Forsikringer",
	"LØNOVERFØRSEL":                 "Løn",
	"FOETEX":                        "Mad",
	"SPOTIFY":                       "Spotify",
	"GOLF":                          "Golf",
	"Golf":                          "Golf",
	"REMA 1000":                     "Mad",
}

func updateExcerptSum(excerptGrp string, amount float64) {
	excerptGrpSums[excerptGrps[excerptGrp]] += float64(amount)
}

func printExcerptGrpSum() {
	for k, v := range excerptGrpSums {
		fmt.Println(k, ": ", v)
	}
}
