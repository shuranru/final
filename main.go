package main

import (
	"encoding/json"
	"final/pkg/config"
	"final/pkg/modify"
	"final/pkg/read"
	"final/pkg/structure"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
)

func main() {

	//fmt.Println(getResultData())

	r := mux.NewRouter()

	r.HandleFunc("/", handleConnection)
	r.HandleFunc("/json", handleConnectionJson)

	if err := http.ListenAndServe(config.ServerWeb, r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))

}

func handleConnectionJson(w http.ResponseWriter, r *http.Request) {
	jsonData, err := json.Marshal(getResultData())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errorMessage := "Ошибка конвертации данных в JSON."
		errorMessage += error.Error(err)
		w.Write([]byte(errorMessage))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func getResultData() structure.ResultSetT {

	var wg sync.WaitGroup

	chanSMS := make(chan [][]structure.SMSData, 1)
	chanMMS := make(chan [][]structure.MMSData, 1)
	chanVoice := make(chan []structure.VoiceCallData, 1)
	chanEmail := make(chan map[string][][]structure.EmailData, 1)
	chanBilling := make(chan structure.BillingData, 1)
	chanSupport := make(chan []int, 1)
	chanIncident := make(chan []structure.IncidentData, 1)

	chanError := make(chan string, 8)

	wg.Add(7)

	go modify.SMSModify(read.SMSFileRead(), &wg, chanSMS, chanError)
	go modify.MMSModify(read.MMSWebRead(), &wg, chanMMS, chanError)
	go modify.VoiceModify(read.VoiceFileRead(), &wg, chanVoice, chanError)
	go modify.EmailModify(read.EmailFileRead(), &wg, chanEmail, chanError)
	readBillingStruct, readBillingError := read.BillingFileRead()
	go modify.BillingModify(readBillingStruct, readBillingError, &wg, chanBilling, chanError)
	go modify.SupportModify(read.SupportWebRead(), &wg, chanSupport, chanError)
	go modify.IncidentModify(read.IncidentWebRead(), &wg, chanIncident, chanError)

	wg.Wait()
	close(chanError)

	smsData := <-chanSMS
	mmsData := <-chanMMS
	voiceCallData := <-chanVoice
	emailData := <-chanEmail
	billingData := <-chanBilling
	supportData := <-chanSupport
	incidentData := <-chanIncident

	var resultErrorMessage string
	lenChanError := len(chanError)
	if lenChanError > 0 {
		for i := 1; i <= lenChanError; i++ {
			resultErrorMessage += <-chanError + "; "
		}
	}

	//fmt.Println("=============================")
	//fmt.Println(resultErrorMessage)

	ResultSet := structure.ResultSetT{
		SMS:       smsData,
		MMS:       mmsData,
		VoiceCall: voiceCallData,
		Email:     emailData,
		Billing:   billingData,
		Support:   supportData,
		Incidents: incidentData,
	}

	//Result := structure.ResultT{}
	//
	//if resultErrorMessage == "" {
	//	Result = structure.ResultT{
	//		true,
	//		ResultSet,
	//		"",
	//	}
	//} else {
	//	Result = structure.ResultT{
	//		false,
	//		structure.ResultSetT{},
	//		resultErrorMessage,
	//	}
	//}
	//
	//return Result

	return ResultSet
}
