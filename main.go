package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

// The verb represents the structure of a verb in the JSON file.
type Verb struct {
	Verb string   `json:"verb"`
	De   []string `json:"de"`
	It   []string `json:"it"`
}

// parseFlags handles the command-line parsing of flags and returns the verb to search for.
func parseFlags() string {
	var flagVarVerb string
	verbToFindPtr := flag.String("verb", "", "Verbo italiano da cercare (es. Essere)")
	flag.StringVar(&flagVarVerb, "v", "", "Verbo italiano da cercare (abbreviazione di -verb)")
	flag.Parse()
	return *verbToFindPtr
}

// readVerbsFromFile reads the contents of the specified JSON file and returns it as a byte slice.
// Returns an error if reading the file fails.
func readVerbsFromFile(filePath string) ([]byte, error) {
	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("errore nella lettura del file %s: %w", filePath, err)
	}
	return byteValue, nil
}

// unmarshalVerbs parses the JSON from a byte slice and deserializes it into a Verb structure slice.
// Returns an error if the JSON parsing fails.
func unmarshalVerbs(byteValue []byte) ([]Verb, error) {
	var verbi []Verb
	err := json.Unmarshal(byteValue, &verbi)
	if err != nil {
		return nil, fmt.Errorf("errore nell'analisi JSON: %w", err)
	}
	return verbi, nil
}

// displayVerbInfo prints information for a single verb.
func displayVerbInfo(verbo Verb) {
	fmt.Println("\nVerbo (Italiano):", verbo.Verb)

	fmt.Println("\nConiugazioni Tedesche:")
	for _, coniugazione := range verbo.De {
		fmt.Println(coniugazione)
	}

	fmt.Println("\nConiugazioni Italiane:")
	for _, coniugazione := range verbo.It {
		fmt.Println(coniugazione)
	}
	fmt.Println("--------------------")
}

// findAndDisplayVerbs iterates through the verb slice, filters by verb if specified, and prints the information.
func findAndDisplayVerbs(verbi []Verb, verbToFind string) bool {
	verbFound := false
	for _, verbo := range verbi {
		if verbToFind == "" || strings.EqualFold(strings.ToLower(verbo.Verb), strings.ToLower(verbToFind)) {
			verbFound = true
			displayVerbInfo(verbo)
		}
	}
	return verbFound
}

// handle Output Messages handles printing of output messages based on the search results.
func handleOutputMessages(verbToFind string, verbFound bool, verbi []Verb) {
	if verbToFind != "" && !verbFound {
		fmt.Printf("\nVerbo '%s' non trovato nel file verbs.json.\n", verbToFind)
	} else if verbToFind == "" && !verbFound && len(verbi) == 0 {
		fmt.Println("\nNessun verbo trovato nel file verbs.json.")
	} else if verbToFind == "" && verbFound && len(verbi) > 0 {
		fmt.Println("\nElenco di tutti i verbi nel file verbs.json mostrato.")
		fmt.Println("Per visualizzare un verbo specifico, usa il flag: -verb <nome_verbo> oppure -v <nome_verbo>")
	} else if verbToFind != "" && verbFound {
		fmt.Printf("\nInformazioni per il verbo '%s' mostrate.\n", verbToFind)
	}
}

func main() {
	// Parsing flags.
	verbToFind := parseFlags()

	// Reading the JSON file.
	byteValue, err := readVerbsFromFile("./assets/verbs.json")
	if err != nil {
		fmt.Println("Errore:", err)
		return
	}

	// Unmarshalling JSON.
	verbi, err := unmarshalVerbs(byteValue)
	if err != nil {
		fmt.Println("Errore:", err)
		return
	}

	// Verb lookup and printing.
	verbFound := findAndDisplayVerbs(verbi, verbToFind)

	// Output message handling.
	handleOutputMessages(verbToFind, verbFound, verbi)
}
