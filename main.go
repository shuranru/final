package main

import (
	"encoding/json"
	"errors"
	"final/pkg/config"
	"final/pkg/list"
	"final/pkg/service"
	"final/pkg/structure"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

func main() {

	fmt.Println(getResultData())

	r := mux.NewRouter()

	r.HandleFunc("/", handleConnection)

	if err := http.ListenAndServe(config.ServerWeb, r); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))

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

	wg.Add(7)

	go SMSModify(SMSFileRead(), &wg, chanSMS)
	go MMSModify(MMSWebRead(), &wg, chanMMS)
	go VoiceFileRead(&wg, chanVoice)
	go EmailModify(EmailFileRead(), &wg, chanEmail)
	go BillingFileRead(&wg, chanBilling)
	go SupportModify(SupportWebRead(), &wg, chanSupport)
	go IncidentModify(IncidentWebRead(), &wg, chanIncident)

	wg.Wait()

	SMSData := <-chanSMS
	MMSData := <-chanMMS
	VoiceCallData := <-chanVoice
	EmailData := <-chanEmail
	BillingData := <-chanBilling
	SupportData := <-chanSupport
	IncidentData := <-chanIncident

	Result := structure.ResultSetT{
		SMS:       SMSData,
		MMS:       MMSData,
		VoiceCall: VoiceCallData,
		Email:     EmailData,
		Billing:   BillingData,
		Support:   SupportData,
		Incidents: IncidentData,
	}

	return Result
}

func SMSFileRead() []structure.SMSData {
	smsArray := service.FileToSlice(config.SMSFile)

	var smsDateTemp []structure.SMSData
	for _, v := range smsArray {
		smsTemp := strings.Split(v, ";")
		if len(smsTemp) != 4 {
			continue
		}
		smsTemp[0] = strings.ToUpper(smsTemp[0])
		if _, ok := list.CountryCodeMap[smsTemp[0]]; !ok {
			continue
		}
		bandwidthTemp, err := strconv.Atoi(smsTemp[1])
		if err != nil {
			continue
		}
		if bandwidthTemp < 0 || bandwidthTemp > 100 {
			continue
		}
		responseTimeTemp, err := strconv.Atoi(smsTemp[2])
		if err != nil {
			continue
		}
		if responseTimeTemp < 0 {
			continue
		}
		if _, ok := list.ProviderNameMap[smsTemp[3]]; !ok {
			continue
		}
		smsDateTemp = append(smsDateTemp, structure.SMSData{smsTemp[0], smsTemp[1], smsTemp[2], smsTemp[3]})

	}

	return smsDateTemp

}

func MMSWebRead() []structure.MMSData {

	content, status := service.WebToByte(config.MMSWeb)

	var mmsDataTemp []structure.MMSData

	if status == 200 {
		if err := json.Unmarshal(content, &mmsDataTemp); err != nil {
			errors.New(fmt.Sprint("Ошибка в формате JSON. ", err))
		}

		for n, _ := range mmsDataTemp {
			if _, ok := list.CountryCodeMap[mmsDataTemp[n].Country]; !ok {
				mmsDataTemp = service.MMSSliceValueDelete(mmsDataTemp, n)
				continue
			}
			if _, ok := list.ProviderNameMap[mmsDataTemp[n].Provider]; !ok {
				mmsDataTemp = service.MMSSliceValueDelete(mmsDataTemp, n)
				continue
			}
		}
	}

	return mmsDataTemp

}

func VoiceFileRead(wg *sync.WaitGroup, chanVoice chan []structure.VoiceCallData) {

	voiceArray := service.FileToSlice(config.VoiceFile)
	if len(voiceArray) == 0 {
		fmt.Println("Voice: Ошибка во входящих данных . Нет данных")
		chanVoice <- []structure.VoiceCallData{}
		wg.Done()
		return
	}

	var voiceDateTemp []structure.VoiceCallData

	for _, v := range voiceArray {
		voiceTemp := strings.Split(v, ";")
		if len(voiceTemp) != 8 {
			continue
		}
		voiceTemp[0] = strings.ToUpper(voiceTemp[0])
		_, ok := list.CountryCodeMap[voiceTemp[0]]
		if ok != true {
			continue
		}
		bandwidthTemp, err := strconv.Atoi(voiceTemp[1])
		if err != nil {
			continue
		}
		if bandwidthTemp < 0 || bandwidthTemp > 100 {
			continue
		}
		responseTimeTemp, err := strconv.Atoi(voiceTemp[2])
		if err != nil {
			continue
		}
		if responseTimeTemp < 0 {
			continue
		}
		stabilityTemp64, err := strconv.ParseFloat(voiceTemp[4], 32)
		if err != nil {
			continue
		}
		stabilityTemp := float32(stabilityTemp64)
		if stabilityTemp < 0 {
			continue
		}
		ttfbTemp, err := strconv.Atoi(voiceTemp[5])
		if err != nil {
			continue
		}
		if ttfbTemp < 0 {
			continue
		}
		purityTemp, err := strconv.Atoi(voiceTemp[6])
		if err != nil {
			continue
		}
		if purityTemp < 0 {
			continue
		}
		durationTemp, err := strconv.Atoi(voiceTemp[7])
		if err != nil {
			continue
		}
		if durationTemp < 0 {
			continue
		}

		if voiceTemp[3] != "TransparentCalls" && voiceTemp[3] != "E-Voice" && voiceTemp[3] != "JustPhone" {
			continue
		}
		voiceDateTemp = append(voiceDateTemp, structure.VoiceCallData{voiceTemp[0], bandwidthTemp, responseTimeTemp, voiceTemp[3], stabilityTemp, ttfbTemp, purityTemp, durationTemp})

	}

	chanVoice <- voiceDateTemp
	wg.Done()
	return

}

func EmailFileRead() []structure.EmailData {
	emailArray := service.FileToSlice(config.EmailFile)

	var emailDateTemp []structure.EmailData
	for _, v := range emailArray {
		emailTemp := strings.Split(v, ";")
		if len(emailTemp) != 3 {
			continue
		}
		emailTemp[0] = strings.ToUpper(emailTemp[0])
		_, ok := list.CountryCodeMap[emailTemp[0]]
		if ok != true {
			continue
		}
		deliveryTimeTemp, err := strconv.Atoi(emailTemp[2])
		if err != nil {
			continue
		}
		if deliveryTimeTemp < 0 {
			continue
		}
		_, ok = list.EmailProvidersMap[emailTemp[1]]
		if ok != true {
			continue
		}
		emailDateTemp = append(emailDateTemp, structure.EmailData{emailTemp[0], emailTemp[1], deliveryTimeTemp})

	}

	return emailDateTemp

}

func BillingFileRead(wg *sync.WaitGroup, chanBilling chan structure.BillingData) {

	billingBool := strings.Split(service.FileToSlice(config.BillingFile)[0], "")
	if len(billingBool) == 0 {
		fmt.Println("Billing: Ошибка во входящих данных. Нет данных")
		chanBilling <- structure.BillingData{}
		wg.Done()
		return
	}

	var billingDateTemp structure.BillingData

	var billingDec uint8 = 0

	lenBillingBool := len(billingBool)

	for n, v := range billingBool {
		if v != "1" || v != "0" || lenBillingBool != 6 {
			fmt.Println("Billing: Ошибка во входящих данных билинга")
			chanBilling <- billingDateTemp
			wg.Done()
			return
		}
		if v == "1" {
			billingDec += uint8(math.Pow(2, float64(lenBillingBool)-1-float64(n)))
		}
	}

	billingDateTemp = structure.BillingData{
		service.CheckBool(billingBool[0]),
		service.CheckBool(billingBool[1]),
		service.CheckBool(billingBool[2]),
		service.CheckBool(billingBool[3]),
		service.CheckBool(billingBool[4]),
		service.CheckBool(billingBool[5]),
	}
	chanBilling <- billingDateTemp
	wg.Done()
	return

}

func SupportWebRead() []structure.SupportData {

	content, status := service.WebToByte(config.SupportWeb)

	var supportDataTemp []structure.SupportData

	if status == 200 {
		if err := json.Unmarshal(content, &supportDataTemp); err != nil {
			errors.New(fmt.Sprint("Ошибка в формате JSON. ", err))
		}
	} else {
		errors.New(fmt.Sprint("Ошибка при получении данных."))
	}

	return supportDataTemp

}

func IncidentWebRead() []structure.IncidentData {

	content, status := service.WebToByte(config.IncidentWeb)

	var incidentDataTemp []structure.IncidentData

	if status == 200 {
		if err := json.Unmarshal(content, &incidentDataTemp); err != nil {
			errors.New(fmt.Sprint("Ошибка в формате JSON. ", err))
		}
		for n, _ := range incidentDataTemp {
			if incidentDataTemp[n].Status != "active" && incidentDataTemp[n].Status != "closed" {
				incidentDataTemp = service.IncidentSliceValueDelete(incidentDataTemp, n)
				continue
			}
		}
	} else {
		errors.New(fmt.Sprint("Ошибка при получении данных."))
	}

	return incidentDataTemp

}

func SMSModify(smsTemp []structure.SMSData, wg *sync.WaitGroup, chanSMS chan [][]structure.SMSData) {

	lenSMSData := len(smsTemp)
	if lenSMSData == 0 {
		fmt.Println("SMS: Ошибка во входящих данных. Нет данных")
		chanSMS <- [][]structure.SMSData{}
		wg.Done()
		return
	}

	returnSMSTemp := make([][]structure.SMSData, 2)

	for n, _ := range smsTemp {
		smsTemp[n].Country = list.CountryCodeMap[smsTemp[n].Country]
	}

	smsProviderTemp := make([]structure.SMSData, lenSMSData)
	copy(smsProviderTemp, smsTemp)

	for i := 0; i <= lenSMSData-1; i++ {
		for j := lenSMSData - 1; j >= i+1; j-- {
			if smsTemp[j].Country < smsTemp[j-1].Country {
				smsTemp[j], smsTemp[j-1] = smsTemp[j-1], smsTemp[j]
			}
			if smsProviderTemp[j].Provider < smsProviderTemp[j-1].Provider {
				smsProviderTemp[j], smsProviderTemp[j-1] = smsProviderTemp[j-1], smsProviderTemp[j]
			}
		}
	}
	returnSMSTemp[0] = smsTemp
	returnSMSTemp[1] = smsProviderTemp

	chanSMS <- returnSMSTemp
	wg.Done()
	return

}

func MMSModify(mmsTemp []structure.MMSData, wg *sync.WaitGroup, chanMMS chan [][]structure.MMSData) {

	lenMMSData := len(mmsTemp)
	if lenMMSData == 0 {
		fmt.Println("MMS: Ошибка во входящих данных. Нет данных")
		chanMMS <- [][]structure.MMSData{}
		wg.Done()
		return
	}

	returnMMSTemp := make([][]structure.MMSData, 2)

	for n, _ := range mmsTemp {
		mmsTemp[n].Country = list.CountryCodeMap[mmsTemp[n].Country]
	}

	mmsProviderTemp := make([]structure.MMSData, lenMMSData)
	copy(mmsProviderTemp, mmsTemp)

	for i := 0; i <= lenMMSData-1; i++ {
		for j := lenMMSData - 1; j >= i+1; j-- {
			if mmsTemp[j].Country < mmsTemp[j-1].Country {
				mmsTemp[j], mmsTemp[j-1] = mmsTemp[j-1], mmsTemp[j]
			}
			if mmsProviderTemp[j].Provider < mmsProviderTemp[j-1].Provider {
				mmsProviderTemp[j], mmsProviderTemp[j-1] = mmsProviderTemp[j-1], mmsProviderTemp[j]
			}
		}
	}
	returnMMSTemp[0] = mmsTemp
	returnMMSTemp[1] = mmsProviderTemp

	chanMMS <- returnMMSTemp
	wg.Done()
	return

}

func EmailModify(emailData []structure.EmailData, wg *sync.WaitGroup, chanEmail chan map[string][][]structure.EmailData) {

	lenEmailData := len(emailData)
	if lenEmailData == 0 {
		fmt.Println("Email: Ошибка во входящих данных. Нет данных")
		chanEmail <- map[string][][]structure.EmailData{}
		wg.Done()
		return
	}

	emailDataCountry := make(map[string][]structure.EmailData)

	for i := 0; i <= lenEmailData-1; i++ {
		for j := lenEmailData - 1; j >= i+1; j-- {
			if emailData[j].DeliveryTime < emailData[j-1].DeliveryTime {
				emailData[j], emailData[j-1] = emailData[j-1], emailData[j]
			}
		}
	}

	for _, v := range emailData {
		emailDataCountry[v.Country] = append(emailDataCountry[v.Country], v)
	}

	returnEmailData := make(map[string][][]structure.EmailData, len(emailDataCountry))

	for n, _ := range emailDataCountry {
		returnEmailData[n] = make([][]structure.EmailData, 2)
		lenEmailDataCountry := len(emailDataCountry[n])
		for i := 0; i <= 2; i++ {
			returnEmailData[n][0] = append(returnEmailData[n][0], emailDataCountry[n][i])
			returnEmailData[n][1] = append(returnEmailData[n][1], emailDataCountry[n][lenEmailDataCountry-1-i])
		}
	}

	chanEmail <- returnEmailData
	wg.Done()
	return

}

func SupportModify(supportData []structure.SupportData, wg *sync.WaitGroup, chanSupport chan []int) {

	if len(supportData) == 0 {
		fmt.Println("Support: Ошибка во входящих данных. Нет данных")
		chanSupport <- []int{}
		wg.Done()
		return
	}

	supportModifyTemp := make([]int, 2)

	ticketCount := 0

	for _, v := range supportData {
		ticketCount += v.ActiveTickets
	}

	if ticketCount >= 0 && ticketCount < 9 {
		supportModifyTemp[0] = 1
	} else if ticketCount >= 9 && ticketCount <= 16 {
		supportModifyTemp[0] = 2
	} else {
		supportModifyTemp[0] = 3
	}

	supportModifyTemp[1] = int(float32(ticketCount) * config.TimeTicket)

	chanSupport <- supportModifyTemp
	wg.Done()
	return

}

func IncidentModify(incident []structure.IncidentData, wg *sync.WaitGroup, chanIncident chan []structure.IncidentData) {

	lenIncident := len(incident)
	if lenIncident == 0 {
		fmt.Println("Incident: Ошибка во входящих данных. Нет данных")
		chanIncident <- []structure.IncidentData{}
		wg.Done()
		return
	}

	incidentTemp := make([]structure.IncidentData, lenIncident)

	activeCount := 0
	closeCount := lenIncident - 1

	for i := 0; i <= lenIncident-1; i++ {
		if incident[i].Status == "active" {
			incidentTemp[activeCount] = incident[i]
			activeCount++
		} else {
			incidentTemp[closeCount] = incident[i]
			closeCount--
		}
	}

	chanIncident <- incidentTemp
	wg.Done()
	return
}
