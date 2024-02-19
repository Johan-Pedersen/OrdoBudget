package main

import "fmt"

var udtrœksKlasser = map[string]string{
	"Resturant besøg/takeway/mad på skole": "E",
	"Helbred (fx tandlœge, medicin m.m.)":  "F",
	"ANTON (foder, pleje, godbidder mv.)":  "G",
	"Boligudstyr":                          "H",
	"Elektronik":                           "I",
	"Gaver":                                "J",
	"Social arrangementer":                 "K",
	"Rejser":                               "L",
	"Spotify":                              "M",
	"Transport, cykel, parkering m.m.":     "N",
	"Golf":                                 "O",
	"Dashlane":                             "P",
	"Diverse":                              "Q",
	"Skolebøger m.m.":                      "R",
	"Tøj og sko":                           "S",
	"Personlig pleje (kosmetik, frisør mv.)": "T",
	"Briller abonnement ":                    "U",
	"Icloud, onedrive, Prime":                "V",
}

var mapping = map[string]string{
	"skatteguiden.dk GF2023":        "Diverse",
	"Silkeborg Ry Golfklub -":       "Golf",
	"MobilePay Rejsekort":           "Transport, cykel, parkering m.m.",
	"SILKEBORG RY GOLFKLUDen 03.02": "Golf",
}

func textToCol(text string) int64 {
	maping := mapping[text]
	var res int64
	fmt.Println("mapping ", maping)
	if maping != "" {
		res = colToColInd(udtrœksKlasser[maping])
	}
	if res < 0 {
		return 1
	}
	return res
}
