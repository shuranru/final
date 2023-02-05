package main

import (
	"encoding/json"
	"final/pkg/config"
	"final/pkg/modify"
	"final/pkg/service"
	"final/pkg/structure"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
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

	if len(resultData.SMS) == 0 {

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

	resultSet := structure.ResultSetT{}

	chanSMS := make(chan structure.SMSDataWError, 1)
	chanMMS := make(chan structure.MMSDataWError, 1)
	chanVoice := make(chan structure.VoiceCallDataWError, 1)
	chanEmail := make(chan structure.EmailDataWError, 1)
	chanBilling := make(chan structure.BillingDataWError, 1)
	chanSupport := make(chan structure.SupportDataWError, 1)
	chanIncident := make(chan structure.IncidentDataWError, 1)

	chanError := make(chan string, 8)

	go modify.SMSModify(chanSMS, chanError)
	go modify.MMSModify(chanMMS, chanError)
	go modify.VoiceModify(chanVoice, chanError)
	go modify.EmailModify(chanEmail, chanError)
	go modify.BillingModify(chanBilling, chanError)
	go modify.SupportModify(chanSupport, chanError)
	go modify.IncidentModify(chanIncident, chanError)

	timeRequest := time.Now().Unix()

	for {
		if len(chanError) > 0 {
			resultSet = structure.ResultSetT{}
			break
		} else if len(chanSMS) > 0 &&
			len(chanMMS) > 0 &&
			len(chanVoice) > 0 &&
			len(chanEmail) > 0 &&
			len(chanBilling) != 0 &&
			len(chanSupport) > 0 &&
			len(chanIncident) > 0 {

			smsData := <-chanSMS
			mmsData := <-chanMMS
			voiceCallData := <-chanVoice
			emailData := <-chanEmail
			billingData := <-chanBilling
			supportData := <-chanSupport
			incidentData := <-chanIncident

			resultSet = structure.ResultSetT{
				SMS:       smsData.SMSDataStruct,
				MMS:       mmsData.MMSDataStruct,
				VoiceCall: voiceCallData.VoiceCallDataStruct,
				Email:     emailData.EmailData,
				Billing:   billingData.BillingDataStruct,
				Support:   supportData.SupportData,
				Incidents: incidentData.IncidentDataStruct,
			}
			break
		} else if time.Now().Unix() > timeRequest+config.TimeOut {
			resultSet = structure.ResultSetT{}
			service.LogWrite("Ошибка получения данных, сработал таймаут в "+strconv.Itoa(int(config.TimeOut))+" сек.", "warn")
			break
		}
	}

	return resultSet
}
