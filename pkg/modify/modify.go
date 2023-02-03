package modify

import (
	"final/pkg/config"
	"final/pkg/list"
	"final/pkg/read"
	"final/pkg/service"
	"final/pkg/structure"
)

func SMSModify(chanSMS chan [][]structure.SMSData) {

	smsTemp := read.SMSFileRead()

	lenSMSData := len(smsTemp)
	if lenSMSData == 0 {
		service.LogWrite("SMS: Ошибка во входящих данных. Нет данных", "warn")
		chanSMS <- [][]structure.SMSData{}
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

	chanSMS <- returnSMSTemp
	return

}

func MMSModify(chanMMS chan [][]structure.MMSData) {

	mmsTemp := read.MMSWebRead()

	lenMMSData := len(mmsTemp)
	if lenMMSData == 0 {
		service.LogWrite("MMS: Ошибка во входящих данных. Нет данных", "warn")
		chanMMS <- [][]structure.MMSData{}
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

	chanMMS <- returnMMSTemp
	return

}

func VoiceModify(chanVoice chan []structure.VoiceCallData) {

	voiceTemp := read.VoiceFileRead()

	if len(voiceTemp) == 0 {
		service.LogWrite("Voice: Ошибка во входящих данных. Нет данных", "warn")
		chanVoice <- []structure.VoiceCallData{}
		return
	}

	chanVoice <- voiceTemp
	return

}

func EmailModify(chanEmail chan map[string][][]structure.EmailData) {

	emailData := read.EmailFileRead()

	lenEmailData := len(emailData)
	if lenEmailData == 0 {
		service.LogWrite("Email: Ошибка во входящих данных. Нет данных", "warn")
		chanEmail <- map[string][][]structure.EmailData{}
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

	chanEmail <- returnEmailData
	return

}

func BillingModify(chanBilling chan structure.BillingData) {

	billingTempStruct, err := read.BillingFileRead()

	if err != "" {
		service.LogWrite(err, "warn")
		chanBilling <- billingTempStruct
		return
	}

	chanBilling <- billingTempStruct
	return

}

func SupportModify(chanSupport chan []int) {

	supportData := read.SupportWebRead()

	if len(supportData) == 0 {
		service.LogWrite("Support: Ошибка во входящих данных. Нет данных", "warn")
		chanSupport <- []int{}
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
	return

}

func IncidentModify(chanIncident chan []structure.IncidentData) {

	incident := read.IncidentWebRead()

	lenIncident := len(incident)
	if lenIncident == 0 {
		service.LogWrite("Incident: Ошибка во входящих данных. Нет данных", "warn")
		chanIncident <- []structure.IncidentData{}
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
	return
}
