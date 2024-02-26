package excrptgrps

import "fmt"

var excerptGrpSums = map[string]float64{
	"Resturant besøg/takeway/mad på skole": -1.0,
	"Helbred (fx tandlœge, medicin m.m.)":  -1.0,
	"ANTON (foder, pleje, godbidder mv.)":  -1.0,
	"Boligudstyr":                          -1.0,
	"Elektronik":                           -1.0,
	"Gaver":                                -1.0,
	"Social arrangementer":                 -1.0,
	"Rejser":                               -1.0,
	"Spotify":                              -1.0,
	"Transport, cykel, parkering m.m.":     -1.0,
	"Golf":                                 -1.0,
	"Dashlane":                             -1.0,
	"Diverse":                              -1.0,
	"Skolebøger m.m.":                      -1.0,
	"Tøj og sko":                           -1.0,
	"Personlig pleje (kosmetik, frisør mv.)": -1.0,
	"Briller abonnement ":                    -1.0,
	"Icloud, onedrive, Prime":                -1.0,
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

var DescToExcerptGrps = map[string]string{
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

func UpdateExcerptSum(excerptGrp string, amount float64) {
	excerptGrpSums[DescToExcerptGrps[excerptGrp]] += float64(amount)
}

func PrintExcerptGrpSum() {
	fmt.Println("###################################################")
	for k, v := range excerptGrpSums {
		fmt.Println(k, ": ", v+1)
	}
	fmt.Println("###################################################")
}

func GetTotal(excrptGrp string) float64 {
	total := excerptGrpSums[excrptGrp]
	if total != 0.0 {
		return excerptGrpSums[excrptGrp] + 1
	}
	return 0.0
}
