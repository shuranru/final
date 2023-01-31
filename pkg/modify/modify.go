package modify

import (
	"final/pkg/config"
	"final/pkg/list"
	"final/pkg/structure"
	"fmt"
	"sync"
)

func SMSModify(smsTemp []structure.SMSData, wg *sync.WaitGroup, chanSMS chan [][]structure.SMSData, chanError chan string) {

	lenSMSData := len(smsTemp)
	if lenSMSData == 0 {
		errorMessage := "SMS: Ошибка во входящих данных. Нет данных"
		fmt.Println(errorMessage)
		chanError <- errorMessage
		chanSMS <- [][]structure.SMSData{}
		wg.Done()
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
	wg.Done()
	return

}

func MMSModify(mmsTemp []structure.MMSData, wg *sync.WaitGroup, chanMMS chan [][]structure.MMSData, chanError chan string) {

	lenMMSData := len(mmsTemp)
	if lenMMSData == 0 {
		errorMessage := "MMS: Ошибка во входящих данных. Нет данных"
		fmt.Println(errorMessage)
		chanError <- errorMessage
		chanMMS <- [][]structure.MMSData{}
		wg.Done()
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
	wg.Done()
	return

}

func VoiceModify(voiceTemp []structure.VoiceCallData, wg *sync.WaitGroup, chanVoice chan []structure.VoiceCallData, chanError chan string) {

	if len(voiceTemp) == 0 {
		errorMessage := "Voice: Ошибка во входящих данных. Нет данных"
		fmt.Println(errorMessage)
		chanError <- errorMessage
		chanVoice <- []structure.VoiceCallData{}
		wg.Done()
		return
	}

	chanVoice <- voiceTemp
	wg.Done()
	return

}

func EmailModify(emailData []structure.EmailData, wg *sync.WaitGroup, chanEmail chan map[string][][]structure.EmailData, chanError chan string) {

	lenEmailData := len(emailData)
	if lenEmailData == 0 {
		errorMessage := "Email: Ошибка во входящих данных. Нет данных"
		fmt.Println(errorMessage)
		chanError <- errorMessage
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

	for n := range emailDataCountry {
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

func BillingModify(billingTemp structure.BillingData, err string, wg *sync.WaitGroup, chanBilling chan structure.BillingData, chanError chan string) {

	if err != "" {
		fmt.Println(err)
		chanError <- err
		chanBilling <- billingTemp
		wg.Done()
		return
	}

	chanBilling <- billingTemp
	wg.Done()
	return

}

func SupportModify(supportData []structure.SupportData, wg *sync.WaitGroup, chanSupport chan []int, chanError chan string) {

	if len(supportData) == 0 {
		errorMessage := "Support: Ошибка во входящих данных. Нет данных"
		fmt.Println(errorMessage)
		chanError <- errorMessage
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

func IncidentModify(incident []structure.IncidentData, wg *sync.WaitGroup, chanIncident chan []structure.IncidentData, chanError chan string) {

	lenIncident := len(incident)
	if lenIncident == 0 {
		errorMessage := "Incident: Ошибка во входящих данных. Нет данных"
		fmt.Println(errorMessage)
		chanError <- errorMessage
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
