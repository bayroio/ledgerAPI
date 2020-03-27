package main

import (
	"fmt"
	"log"
    "net/http"
    "io/ioutil"
    "encoding/json"

	"github.com/gorilla/mux"
)

type ledger struct {
	Name    string `json:"Name"`
	Symbol  string `json:"Symbol"`
	Address string `json:"Address"`
}

type allLedgers []ledger

var ledgers = allLedgers{
	{
		Name:       "Frutas y Hortalizas JB",
		Symbol:     "FHJB",
		Address:    "0x0",
    },
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Ahoj Network APIs!")
}

func createLedger(w http.ResponseWriter, r *http.Request) {
	var newLedger ledger
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the ledger name symbol an address only in order to update")
	}
	
	json.Unmarshal(reqBody, &newLedger)
	ledgers = append(ledgers, newLedger)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newLedger)
}

func getOneLedger(w http.ResponseWriter, r *http.Request) {
	ledgerSymbol := mux.Vars(r)["symbol"]

	for _, singleLedger := range ledgers {
		if singleLedger.Symbol == ledgerSymbol {
			json.NewEncoder(w).Encode(singleLedger)
		}
	}
}

func getAllLedgers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(ledgers)
}

func updateLedger(w http.ResponseWriter, r *http.Request) {
	ledgerSymbol := mux.Vars(r)["symbol"]
	var updatedLedger ledger

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the ledger name symbol an address only in order to update")
	}
	json.Unmarshal(reqBody, &updatedLedger)

	for i, singleLedger := range ledgers {
		if singleLedger.Symbol == ledgerSymbol {
            singleLedger.Name = updatedLedger.Name
			singleLedger.Address = updatedLedger.Address
			ledgers = append(ledgers[:i], singleLedger)
			json.NewEncoder(w).Encode(singleLedger)
		}
	}
}

func deleteLedger(w http.ResponseWriter, r *http.Request) {
	ledgerSymbol := mux.Vars(r)["symbol"]

	for i, singleLedger := range ledgers {
		if singleLedger.Symbol == ledgerSymbol {
			ledgers = append(ledgers[:i], ledgers[i+1:]...)
			fmt.Fprintf(w, "The ledger with symbol %v has been deleted successfully", ledgerSymbol)
		}
	}
}
func main() {
	//initEvents()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/ledger", createLedger).Methods("POST")
	router.HandleFunc("/ledgers", getAllLedgers).Methods("GET")
	router.HandleFunc("/ledgers/{symbol}", getOneLedger).Methods("GET")
	router.HandleFunc("/ledgers/{symbol}", updateLedger).Methods("PATCH")
	router.HandleFunc("/ledgers/{symbol}", deleteLedger).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}