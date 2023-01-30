package main

import (
	"encoding/json"
	"errors"
	"final/pkg/config"
	"final/pkg/service"
	"final/pkg/structure"
	"fmt"
	"github.com/gorilla/mux"
	"math"
	"net/http"
	"strconv"
	"strings"
)

var CountryCode map[string]string
var SMSMMSProviderName map[string]int

var Result structure.ResultSetT

func main() {

	CountryCode = service.CountryCodeRead(CountryCode)
	SMSMMSProviderName = service.SMSMMSProviderNameRead(SMSMMSProviderName)

	SMSData := SMSModify(SMSFileRead())
	MMSData := MMSModify(MMSWebRead())
	VoiceCallData := VoiceFileRead()
	EmailData := EmailModify(EmailFileRead())
	BilingData := BilingFileRead()
	SupportData := SupportModify(SupportWebRead())
	IncidentData := IncidentModify(IncidentWebRead())

	Result := structure.ResultSetT{
		SMS:       SMSData,
		MMS:       MMSData,
		VoiceCall: VoiceCallData,
		Email:     EmailData,
		Billing:   BilingData,
		Support:   SupportData,
		Incidents: IncidentData,
	}

	fmt.Println(Result)
	//fmt.Println("")
	//fmt.Println(SMSModify(SMSDataSlice))
	//fmt.Println(VoiceCallDataSlice)
	//fmt.Println("\n\n============================")
	//fmt.Println(EmailDataSlice)
	//fmt.Println("\n\n============================")
	//fmt.Println(BilingData)
	//fmt.Println("")
	//fmt.Println(MMSDataSlice)
	//fmt.Println("\n\n============================")
	//fmt.Println(SupportData)

	r := mux.NewRouter()

	//done := make(chan os.Signal, 1)
	//signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	r.HandleFunc("/", handleConnection)

	//if err := http.ListenAndServe(config.ServerWeb, r); err != nil && err != http.ErrServerClosed {
	//	log.Fatalf("Ошибка запуска сервера: %v", err)
	//}

}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	//userId := chi.URLParam(r, "user_id")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Ok"))
}

func getResultData() structure.ResultSetT {

	return structure.ResultSetT{}
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
		_, ok := CountryCode[smsTemp[0]]
		if ok != true {
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
		//if smsTemp[3] != "Topolo" && smsTemp[3] != "Rond" && smsTemp[3] != "Kildy" {
		//	continue
		//}
		if _, ok := SMSMMSProviderName[smsTemp[3]]; !ok {
			continue
		}
		smsDateTemp = append(smsDateTemp, structure.SMSData{smsTemp[0], smsTemp[1], smsTemp[2], smsTemp[3]})

	}

	return smsDateTemp

}

func SMSModify(smsTemp []structure.SMSData) [][]structure.SMSData {

	for n, _ := range smsTemp {
		smsTemp[n].Country = CountryCode[smsTemp[n].Country]
	}

	sizeSMSData := len(smsTemp)

	returnSMSTemp := make([][]structure.SMSData, 2)

	smsProviderTemp := make([]structure.SMSData, sizeSMSData)
	copy(smsProviderTemp, smsTemp)

	for i := 0; i <= sizeSMSData-1; i++ {
		for j := sizeSMSData - 1; j >= i+1; j-- {
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

	//for i := 0; i <= sizeSMSData-1; i++ {
	//	for j := sizeSMSData - 1; j >= i+1; j-- {
	//
	//	}
	//}

	return returnSMSTemp

}

func VoiceFileRead() []structure.VoiceCallData {

	voiceArray := service.FileToSlice(config.VoiceFile)

	var voiceDateTemp []structure.VoiceCallData
	for _, v := range voiceArray {
		voiceTemp := strings.Split(v, ";")
		if len(voiceTemp) != 8 {
			continue
		}
		voiceTemp[0] = strings.ToUpper(voiceTemp[0])
		_, ok := CountryCode[voiceTemp[0]]
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

	return voiceDateTemp

}

func EmailFileRead() []structure.EmailData {
	emailArray := service.FileToSlice(config.EmailFile)

	emailProvidersMap := map[string]int{
		"Gmail":       1,
		"Yahoo":       1,
		"Hotmail":     1,
		"MSN":         1,
		"Orange":      1,
		"Comcast":     1,
		"AOL":         1,
		"Live":        1,
		"RediffMail":  1,
		"GMX":         1,
		"Proton Mail": 1,
		"Yandex":      1,
		"Mail.ru":     1,
	}

	var emailDateTemp []structure.EmailData
	for _, v := range emailArray {
		emailTemp := strings.Split(v, ";")
		if len(emailTemp) != 3 {
			continue
		}
		emailTemp[0] = strings.ToUpper(emailTemp[0])
		_, ok := CountryCode[emailTemp[0]]
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
		_, ok = emailProvidersMap[emailTemp[1]]
		if ok != true {
			continue
		}
		emailDateTemp = append(emailDateTemp, structure.EmailData{emailTemp[0], emailTemp[1], deliveryTimeTemp})

	}

	return emailDateTemp

}

func EmailModify(emailData []structure.EmailData) map[string][][]structure.EmailData {

	emailDataCountry := make(map[string][]structure.EmailData)

	lenEmailData := len(emailData)

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

	return returnEmailData

}

func BilingFileRead() structure.BillingData {

	billingBool := strings.Split(service.FileToSlice(config.BillingFile)[0], "")
	var billingDec uint8 = 0

	lenBillingBool := len(billingBool)

	for n, v := range billingBool {
		if v != "1" && v != "0" {
			panic(errors.New("Ошибка во входящих данных билинга"))
		}
		if v == "1" {
			billingDec += uint8(math.Pow(2, float64(lenBillingBool)-1-float64(n)))
		}
	}

	billingDateTemp := structure.BillingData{
		service.CheckBool(billingBool[0]),
		service.CheckBool(billingBool[1]),
		service.CheckBool(billingBool[2]),
		service.CheckBool(billingBool[3]),
		service.CheckBool(billingBool[4]),
		service.CheckBool(billingBool[5]),
	}

	return billingDateTemp

}

func MMSWebRead() []structure.MMSData {

	content, status := service.WebToByte(config.MMSWeb)

	var mmsDataTemp []structure.MMSData

	if status == 200 {
		if err := json.Unmarshal(content, &mmsDataTemp); err != nil {
			errors.New(fmt.Sprint("Ошибка в формате JSON. ", err))
		}

		for n, _ := range mmsDataTemp {
			if _, ok := CountryCode[mmsDataTemp[n].Country]; !ok {
				mmsDataTemp = service.MMSSliceValueDelete(mmsDataTemp, n)
				continue
			}
			if _, ok := SMSMMSProviderName[mmsDataTemp[n].Provider]; !ok {
				mmsDataTemp = service.MMSSliceValueDelete(mmsDataTemp, n)
				continue
			}
		}
	}

	return mmsDataTemp

}

func MMSModify(mmsTemp []structure.MMSData) [][]structure.MMSData {

	for n, _ := range mmsTemp {
		mmsTemp[n].Country = CountryCode[mmsTemp[n].Country]
	}

	sizeMMSData := len(mmsTemp)

	returnMMSTemp := make([][]structure.MMSData, 2)

	mmsProviderTemp := make([]structure.MMSData, sizeMMSData)
	copy(mmsProviderTemp, mmsTemp)

	for i := 0; i <= sizeMMSData-1; i++ {
		for j := sizeMMSData - 1; j >= i+1; j-- {
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

	//for i := 0; i <= sizeMMSData-1; i++ {
	//	for j := sizeMMSData - 1; j >= i+1; j-- {
	//		if mmsProviderTemp[j].Provider < mmsProviderTemp[j-1].Provider {
	//			mmsProviderTemp[j], mmsProviderTemp[j-1] = mmsProviderTemp[j-1], mmsProviderTemp[j]
	//		}
	//	}
	//}

	return returnMMSTemp

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

func SupportModify(supportData []structure.SupportData) []int {

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

	return supportModifyTemp

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

func IncidentModify(incident []structure.IncidentData) []structure.IncidentData {

	lenIncident := len(incident)
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

	return incidentTemp

}
