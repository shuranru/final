package main

import (
	"encoding/json"
	"errors"
	"final/pkg/config"
	"final/pkg/service"
	"final/pkg/structure"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var CountryCode map[string]string
var SMSMMSProviderName map[string]int
var SMSDataSlice []structure.SMSData
var VoiceDataSlice []structure.VoiceData
var EmailDataSlice []structure.EmailData
var BilingData structure.BillingData
var MMSDataSlice []structure.MMSData
var SupportDataSlice []structure.SupportData
var IncidentDataSlice []structure.IncidentData

func main() {

	CountryCode = service.CountryCodeRead(CountryCode)
	SMSMMSProviderName = service.SMSMMSProviderNameRead(SMSMMSProviderName)

	SMSDataSlice = SMSFileRead()
	VoiceDataSlice = VoiceFileRead()
	EmailDataSlice = EmailFileRead()
	BilingData = BilingFileRead()
	MMSDataSlice = MMSWebRead()
	SupportDataSlice = SupportWebRead()
	IncidentDataSlice = IncidentWebRead()

	fmt.Println(SMSDataSlice)
	fmt.Println("")
	//fmt.Println(VoiceDataSlice)
	//fmt.Println("")
	//fmt.Println(EmailDataSlice)
	//fmt.Println("")
	//fmt.Println(BilingData)
	//fmt.Println("")
	//fmt.Println(MMSDataSlice)
	//fmt.Println("")
	//fmt.Println(SupportDataSlice)
	//fmt.Println("")
	//fmt.Println(IncidentDataSlice)

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

func VoiceFileRead() []structure.VoiceData {

	voiceArray := service.FileToSlice(config.VoiceFile)

	var voiceDateTemp []structure.VoiceData
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
		voiceDateTemp = append(voiceDateTemp, structure.VoiceData{voiceTemp[0], bandwidthTemp, responseTimeTemp, voiceTemp[3], stabilityTemp, ttfbTemp, purityTemp, durationTemp})

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
