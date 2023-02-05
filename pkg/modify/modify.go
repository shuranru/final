package modify

import (
	"final/pkg/config"
	"final/pkg/list"
	"final/pkg/read"
	"final/pkg/service"
	"final/pkg/structure"
)

func SMSModify(chanSMS chan structure.SMSDataWError, chanError chan string) {

	smsTemp := read.SMSFileRead()

	lenSMSData := len(smsTemp)
	if lenSMSData == 0 {
		errorMessage := "SMS: Ошибка во входящих данных. Нет данных"
		service.LogWrite(errorMessage, "warn")
		//chanSMS <- structure.SMSDataWError{[][]structure.SMSData{}, errors.New(errorMessage)}
		chanError <- errorMessage
		return
	}

	returnSMSTemp := make([][]structure.SMSData, 2)

	for n := range smsTemp {
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

	chanSMS <- structure.SMSDataWError{returnSMSTemp, nil}
	return

}

func MMSModify(chanMMS chan structure.MMSDataWError, chanError chan string) {

	mmsTemp := read.MMSWebRead()

	lenMMSData := len(mmsTemp)
	if lenMMSData == 0 {
		errorMessage := "MMS: Ошибка во входящих данных. Нет данных"
		service.LogWrite(errorMessage, "warn")
		//chanMMS <- structure.MMSDataWError{[][]structure.MMSData{}, errors.New(errorMessage)}
		chanError <- errorMessage
		return
	}

	returnMMSTemp := make([][]structure.MMSData, 2)

	for n := range mmsTemp {
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

	chanMMS <- structure.MMSDataWError{returnMMSTemp, nil}
	return

}

func VoiceModify(chanVoice chan structure.VoiceCallDataWError, chanError chan string) {

	voiceTemp := read.VoiceFileRead()

	if len(voiceTemp) == 0 {
		errorMessage := "Voice: Ошибка во входящих данных. Нет данных"
		service.LogWrite(errorMessage, "warn")
		//chanVoice <- structure.VoiceCallDataWError{[]structure.VoiceCallData{}, errors.New(errorMessage)}
		chanError <- errorMessage
		return
	}

	chanVoice <- structure.VoiceCallDataWError{voiceTemp, nil}
	return

}

func EmailModify(chanEmail chan structure.EmailDataWError, chanError chan string) {

	emailData := read.EmailFileRead()

	lenEmailData := len(emailData)
	if lenEmailData == 0 {
		errorMessage := "Email: Ошибка во входящих данных. Нет данных"
		service.LogWrite(errorMessage, "warn")
		//chanEmail <- structure.EmailDataWError{map[string][][]structure.EmailData{}, errors.New(errorMessage)}
		chanError <- errorMessage
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

	for n := range emailDataCountry {
		returnEmailData[n] = make([][]structure.EmailData, 2)
		lenEmailDataCountry := len(emailDataCountry[n])
		for i := 0; i <= 2; i++ {
			returnEmailData[n][0] = append(returnEmailData[n][0], emailDataCountry[n][i])
			returnEmailData[n][1] = append(returnEmailData[n][1], emailDataCountry[n][lenEmailDataCountry-1-i])
		}
	}

	chanEmail <- structure.EmailDataWError{returnEmailData, nil}
	return

}

func BillingModify(chanBilling chan structure.BillingDataWError, chanError chan string) {

	billingTempStruct, errorMessage := read.BillingFileRead()

	if errorMessage != "" {
		service.LogWrite(errorMessage, "warn")
		//chanBilling <- structure.BillingDataWError{structure.BillingData{}, errors.New(errorMessage)}
		chanError <- errorMessage
		return
	}

	chanBilling <- structure.BillingDataWError{billingTempStruct, nil}
	return

}

func SupportModify(chanSupport chan structure.SupportDataWError, chanError chan string) {

	supportData := read.SupportWebRead()

	if len(supportData) == 0 {
		errorMessage := "Support: Ошибка во входящих данных. Нет данных"
		service.LogWrite(errorMessage, "warn")
		//chanSupport <- structure.SupportDataWError{[]int{}, errors.New(errorMessage)}
		chanError <- errorMessage
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

	chanSupport <- structure.SupportDataWError{supportModifyTemp, nil}
	return

}

func IncidentModify(chanIncident chan structure.IncidentDataWError, chanError chan string) {

	incident := read.IncidentWebRead()

	lenIncident := len(incident)
	if lenIncident == 0 {
		errorMessage := "Incident: Ошибка во входящих данных. Нет данных"
		service.LogWrite(errorMessage, "warn")
		//chanIncident <- structure.IncidentDataWError{IncidentDataStruct: []structure.IncidentData{}, Error: errors.New(errorMessage)}
		chanError <- errorMessage
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

	chanIncident <- structure.IncidentDataWError{incidentTemp, nil}
	return
}
