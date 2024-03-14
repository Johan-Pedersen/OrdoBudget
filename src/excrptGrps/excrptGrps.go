package excrptgrps

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var excrptGrpTotals = map[string]float64{}

var excrptGrps = []ExcrptGrpParent{}

var resume = []string{}

// Marshal and unmarshal json
type Data struct {
	ExcrptMappings map[string]map[string][]string
}

type ExcrptGrp struct {
	// Used to make lookup in excerptMappings array
	ind int

	// Name of the ExcrptGrp
	name string

	// Matches for this excrptGrp
	mappings []string

	// Defines the type of this excerpt
	parent string
}

type ExcrptGrpParent struct {
	name       string
	excrptGrps []ExcrptGrp
}

func UpdateExcrptTotal(date, excrpt string, amount float64) {
	excrptGrpName := ""

	// ignore case
	excrpt = strings.ToLower(strings.Trim(excrpt, " "))

	ind := -1
	// Find correct excrpt grp
	for _, parent := range excrptGrps {
		for i := range parent.excrptGrps {
			for _, match := range parent.excrptGrps[i].mappings {
				if strings.Contains(excrpt, strings.ToLower(match)) {
					excrptGrpName = parent.excrptGrps[i].name
					break
				}
			}
			if excrptGrpName != "" {
				break
			}
		}

		if excrptGrpName != "" {
			break
		}
	}

	// No matches for group
	if excrptGrpName == "" {
		fmt.Println("Can't match to group:", date, excrpt, ":", amount)
		fmt.Println("Select group")
		fmt.Scan(&ind)

		for _, parent := range excrptGrps {
			for _, excrptGrp := range parent.excrptGrps {
				if excrptGrp.ind == ind {
					excrptGrpName = excrptGrp.name
					break
				}
			}

			if excrptGrpName != "" {
				break
			}

		}
	}

	if excrptGrpName != "Ignored" {
		// Update correct excrpt grp
		excrptGrpTotals[excrptGrpName] += float64(amount)
	}

	// Update resume

	UpdateResume(date, excrpt, excrptGrpName, amount)
}

func UpdateResume(date, excrpt, excrptGrpName string, amount float64) {
	resume = append(resume, date+" "+excrpt+" "+strconv.FormatFloat(amount, 'f', -1, 64)+": "+excrptGrpName)
}

func PrintExcrptGrpTotals() {
	fmt.Println("###################################################")
	for k, v := range excrptGrpTotals {
		fmt.Println(k, ": ", v+1)
	}
	fmt.Println("###################################################")
}

func PrintExcrptGrps() {
	fmt.Println("Excerpt groups")
	fmt.Println("###################################################")
	for _, parent := range excrptGrps {
		fmt.Println("\n************", parent.name, "************")
		for _, excrptGrp := range parent.excrptGrps {
			fmt.Println(excrptGrp.ind, ":", excrptGrp.name)
		}
	}
	fmt.Println("###################################################")
}

func PrintResume() {
	fmt.Println("Resume")
	fmt.Println("###################################################")
	for _, str := range resume {
		fmt.Println(str)
	}
	fmt.Println("###################################################")
}

func GetTotal(excrptGrp ExcrptGrp) float64 {
	total := excrptGrpTotals[excrptGrp.name]
	if total != 0.0 {
		if excrptGrp.parent == "Indkomst efter skat" {
			return excrptGrpTotals[excrptGrp.name] + 1
		} else {
			return -1 * (excrptGrpTotals[excrptGrp.name] + 1)
		}
	}
	return 0.0
}

func InitExcrptGrps() {
	// Open the JSON file
	file, err := os.Open("excrptGrpData.json")
	if err != nil {
		log.Fatalln("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the JSON data from the file
	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("Error reading JSON data:", err)
		return
	}

	// Create an instance of the struct to hold the unmarshaled data
	var data Data

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalln("Error unmarshaling JSON data:", err)
		return
	}

	i := 0
	// Init excrptGrp Totals
	for parentName, excrpts := range data.ExcrptMappings {

		grps := []ExcrptGrp{}
		for excrptGrp, mappings := range excrpts {
			excrptGrpTotals[excrptGrp] = -1.0
			grps = append(grps,
				createExcrptGrp(i, excrptGrp, parentName, mappings))
			i += 1
		}
		parent := ExcrptGrpParent{parentName, grps}
		excrptGrps = append(excrptGrps, parent)
	}
}

func createExcrptGrp(ind int, name, parent string, mappings []string) ExcrptGrp {
	return ExcrptGrp{
		ind:      ind,
		name:     name,
		mappings: mappings,
		parent:   parent,
	}
}

/*
Get excerpt group based on name OR index(ind).
Both can be specified, but it's not necessary.
If you don't want to use ind, make ind < 0.
if you don't want to use name, make name = "".
*/
func GetExcrptGrp(name string, ind int) (ExcrptGrp, error) {
	for _, parent := range excrptGrps {
		for _, excrptGrp := range parent.excrptGrps {
			if excrptGrp.name == name || excrptGrp.ind == ind {
				return excrptGrp, nil
			}
		}
	}
	return ExcrptGrp{}, errors.New("Excrpt grp: " + name + ", not found")
}
