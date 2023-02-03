package read

import (
	"encoding/json"
	"final/pkg/config"
	"final/pkg/list"
	"final/pkg/service"
	"final/pkg/structure"
	"math"
	"strconv"
	"strings"
)

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
		smsDateTemp = append(smsDateTemp, structure.SMSData{Country: smsTemp[0], Bandwidth: smsTemp[1], ResponseTime: smsTemp[2], Provider: smsTemp[3]})

	}

	return smsDateTemp

}

func MMSWebRead() []structure.MMSData {

	content, status := service.WebToByte(config.MMSWeb)

	var mmsDataTemp []structure.MMSData

	if status == 200 {
		if err := json.Unmarshal(content, &mmsDataTemp); err != nil {
			service.LogWrite("MMS: Ошибка в формате JSON: "+error.Error(err), "warn")
			return []structure.MMSData{}
		}

		for n := range mmsDataTemp {
			if _, ok := list.CountryCodeMap[mmsDataTemp[n].Country]; !ok {
				mmsDataTemp = service.MMSSliceValueDelete(mmsDataTemp, n)
				continue
			}
			if _, ok := list.ProviderNameMap[mmsDataTemp[n].Provider]; !ok {
				mmsDataTemp = service.MMSSliceValueDelete(mmsDataTemp, n)
				continue
			}
		}
	} else {
		service.LogWrite("MMS: Ошибка при получении данных.", "warn")
		return []structure.MMSData{}
	}

	return mmsDataTemp

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
		voiceDateTemp = append(voiceDateTemp, structure.VoiceCallData{Country: voiceTemp[0], Bandwidth: bandwidthTemp, ResponseTime: responseTimeTemp, Provider: voiceTemp[3], ConnectionStability: stabilityTemp, TTFB: ttfbTemp, VoicePurity: purityTemp, MedianOfCallsTime: durationTemp})

	}

	return voiceDateTemp

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
		emailDateTemp = append(emailDateTemp, structure.EmailData{Country: emailTemp[0], Provider: emailTemp[1], DeliveryTime: deliveryTimeTemp})

	}

	return emailDateTemp

}

func BillingFileRead() (structure.BillingData, string) {

	billingBool := strings.Split(service.FileToSlice(config.BillingFile)[0], "")
	if len(billingBool) == 0 {
		return structure.BillingData{}, "Billing: Ошибка во входящих данных. Нет данных"
	}

	var billingDateTemp structure.BillingData

	var billingDec uint8 = 0

	lenBillingBool := len(billingBool)

	for n, v := range billingBool {
		if (v != "1" && v != "0") || lenBillingBool != 6 {
			return structure.BillingData{}, "Billing: Ошибка во входящих данных. Нет данных"
		}
		if v == "1" {
			billingDec += uint8(math.Pow(2, float64(lenBillingBool)-1-float64(n)))
		}
	}

	billingDateTemp = structure.BillingData{
		CreateCustomer: service.CheckBool(billingBool[0]),
		Purchase:       service.CheckBool(billingBool[1]),
		Payout:         service.CheckBool(billingBool[2]),
		Recurring:      service.CheckBool(billingBool[3]),
		FraudControl:   service.CheckBool(billingBool[4]),
		CheckoutPage:   service.CheckBool(billingBool[5]),
	}

	return billingDateTemp, ""

}

func SupportWebRead() []structure.SupportData {

	content, status := service.WebToByte(config.SupportWeb)

	var supportDataTemp []structure.SupportData

	if status == 200 {
		if err := json.Unmarshal(content, &supportDataTemp); err != nil {
			service.LogWrite("Support: Ошибка в формате JSON: "+error.Error(err), "warn")
			return []structure.SupportData{}
		}
	} else {
		service.LogWrite("Support: Ошибка при получении данных.", "warn")
		return []structure.SupportData{}
	}

	return supportDataTemp

}

func IncidentWebRead() []structure.IncidentData {

	content, status := service.WebToByte(config.IncidentWeb)

	var incidentDataTemp []structure.IncidentData

	if status == 200 {
		if err := json.Unmarshal(content, &incidentDataTemp); err != nil {
			service.LogWrite("Incident: Ошибка в формате JSON: "+error.Error(err), "warn")
			return []structure.IncidentData{}
		}
		for n := range incidentDataTemp {
			if incidentDataTemp[n].Status != "active" && incidentDataTemp[n].Status != "closed" {
				incidentDataTemp = service.IncidentSliceValueDelete(incidentDataTemp, n)
				continue
			}
		}
	} else {
		service.LogWrite("Incident:  Ошибка при получении данных.", "warn")
		return []structure.IncidentData{}
	}

	return incidentDataTemp

}
