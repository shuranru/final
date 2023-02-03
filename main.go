package main

import (
	"encoding/json"
	"final/pkg/config"
	"final/pkg/modify"
	"final/pkg/structure"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api", handleConnection).Methods("GET")

	if err := http.ListenAndServe(config.ServerWeb, r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

}

func handleConnection(w http.ResponseWriter, r *http.Request) {

	resultData := getResultData()

	Result := structure.ResultT{}

	if len(resultData.SMS) == 0 ||
		len(resultData.MMS) == 0 ||
		len(resultData.VoiceCall) == 0 ||
		len(resultData.Email) == 0 ||
		len(resultData.Support) == 0 ||
		len(resultData.Incidents) == 0 {

		Result = structure.ResultT{
			false,
			structure.ResultSetT{},
			"Error on collect data",
		}

	} else {
		Result = structure.ResultT{
			true,
			resultData,
			"",
		}
	}

	jsonData, err := json.Marshal(Result)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := "Ошибка конвертации данных в JSON. " + error.Error(err)
		w.Write([]byte(errorMessage))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func getResultData() structure.ResultSetT {

	chanSMS := make(chan [][]structure.SMSData, 1)
	chanMMS := make(chan [][]structure.MMSData, 1)
	chanVoice := make(chan []structure.VoiceCallData, 1)
	chanEmail := make(chan map[string][][]structure.EmailData, 1)
	chanBilling := make(chan structure.BillingData, 1)
	chanSupport := make(chan []int, 1)
	chanIncident := make(chan []structure.IncidentData, 1)

	go modify.SMSModify(chanSMS)
	go modify.MMSModify(chanMMS)
	go modify.VoiceModify(chanVoice)
	go modify.EmailModify(chanEmail)
	go modify.BillingModify(chanBilling)
	go modify.SupportModify(chanSupport)
	go modify.IncidentModify(chanIncident)

	smsData := <-chanSMS
	mmsData := <-chanMMS
	voiceCallData := <-chanVoice
	emailData := <-chanEmail
	billingData := <-chanBilling
	supportData := <-chanSupport
	incidentData := <-chanIncident

	ResultSet := structure.ResultSetT{
		SMS:       smsData,
		MMS:       mmsData,
		VoiceCall: voiceCallData,
		Email:     emailData,
		Billing:   billingData,
		Support:   supportData,
		Incidents: incidentData,
	}

	return ResultSet
}
