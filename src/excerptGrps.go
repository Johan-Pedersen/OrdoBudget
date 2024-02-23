package main

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
	"Resturant besøg/takeway/mad på skole": 1,
	"Helbred (fx tandlœge, medicin m.m.)":  2,
	"ANTON (foder, pleje, godbidder mv.)":  3,
	"Boligudstyr":                          4,
	"Elektronik":                           5,
	"Gaver":                                6,
	"Social arrangementer":                 7,
	"Rejser":                               8,
	"Spotify":                              9,
	"Transport, cykel, parkering m.m.":     10,
	"Golf":                                 11,
	"Dashlane":                             12,
	"Diverse":                              13,
	"Skolebøger m.m.":                      14,
	"Tøj og sko":                           15,
	"Personlig pleje (kosmetik, frisør mv.)": 16,
	"Briller abonnement ":                    17,
	"Icloud, onedrive, Prime":                18,
}

var excerptGrps = map[string]string{
	"skatteguiden.dk GF2023":        "Diverse",
	"Silkeborg Ry Golfklub -":       "Golf",
	"MobilePay Rejsekort":           "Transport, cykel, parkering m.m.",
	"SILKEBORG RY GOLFKLUDen 03.02": "Golf",
	"Akademiker a-kas se":           "Fagforening og A-kasse",
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
