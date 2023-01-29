package service

import (
	"errors"
	"final/pkg/structure"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func FileToSlice(filename string) []string {

	contentTemp, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	var content []string = strings.Split(string(contentTemp), "\n")

	return content
}

func WebToByte(url string) ([]byte, int) {

	resp, err := http.Get(url)
	if err != nil {
		errors.New(fmt.Sprint("Ошибка при получении данных. ", err))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errors.New(fmt.Sprint("Ошибка при чтении данных. ", err))
	}
	defer resp.Body.Close()

	return body, resp.StatusCode

}

func CheckBool(check string) bool {
	if check == "1" {
		return true
	} else {
		return false
	}
}

func CountryCodeRead(countryCodeMap map[string]string) map[string]string {

	countryCodeMap = make(map[string]string)

	countryCodeMap["AU"] = "Австралия"
	countryCodeMap["AT"] = "Австрия"
	countryCodeMap["AZ"] = "Азербайджан"
	countryCodeMap["AL"] = "Албания"
	countryCodeMap["DZ"] = "Алжир"
	countryCodeMap["AO"] = "Ангола"
	countryCodeMap["AD"] = "Андорра"
	countryCodeMap["AG"] = "Антигуа и Барбуда"
	countryCodeMap["AR"] = "Аргентина"
	countryCodeMap["AM"] = "Армения"
	countryCodeMap["AF"] = "Афганистан"
	countryCodeMap["BS"] = "Багамские Острова"
	countryCodeMap["BD"] = "Бангладеш"
	countryCodeMap["BB"] = "Барбадос"
	countryCodeMap["BH"] = "Бахрейн"
	countryCodeMap["BZ"] = "Белиз"
	countryCodeMap["BY"] = "Белоруссия"
	countryCodeMap["BE"] = "Бельгия"
	countryCodeMap["BJ"] = "Бенин"
	countryCodeMap["BG"] = "Болгария"
	countryCodeMap["BO"] = "Боливия"
	countryCodeMap["BA"] = "Босния и Герцеговина"
	countryCodeMap["BW"] = "Ботсвана"
	countryCodeMap["BR"] = "Бразилия"
	countryCodeMap["BN"] = "Бруней"
	countryCodeMap["BF"] = "Буркина-Фасо"
	countryCodeMap["BI"] = "Бурунди"
	countryCodeMap["BT"] = "Бутан"
	countryCodeMap["VU"] = "Вануату"
	countryCodeMap["GB"] = "Великобритания"
	countryCodeMap["HU"] = "Венгрия"
	countryCodeMap["VE"] = "Венесуэла"
	countryCodeMap["TL"] = "Восточный Тимор"
	countryCodeMap["VN"] = "Вьетнам"
	countryCodeMap["GA"] = "Габон"
	countryCodeMap["HT"] = "Гаити"
	countryCodeMap["GY"] = "Гайана"
	countryCodeMap["GM"] = "Гамбия"
	countryCodeMap["GH"] = "Гана"
	countryCodeMap["GT"] = "Гватемала"
	countryCodeMap["GN"] = "Гвинея"
	countryCodeMap["GW"] = "Гвинея-Бисау"
	countryCodeMap["DE"] = "Германия"
	countryCodeMap["HN"] = "Гондурас"
	countryCodeMap["GD"] = "Гренада"
	countryCodeMap["GR"] = "Греция"
	countryCodeMap["GE"] = "Грузия"
	countryCodeMap["DK"] = "Дания"
	countryCodeMap["DJ"] = "Джибути"
	countryCodeMap["DM"] = "Доминика"
	countryCodeMap["DO"] = "Доминиканская Республика"
	countryCodeMap["CD"] = "ДР Конго"
	countryCodeMap["EG"] = "Египет"
	countryCodeMap["ZM"] = "Замбия"
	countryCodeMap["ZW"] = "Зимбабве"
	countryCodeMap["IL"] = "Израиль"
	countryCodeMap["IN"] = "Индия"
	countryCodeMap["ID"] = "Индонезия"
	countryCodeMap["JO"] = "Иордания"
	countryCodeMap["IQ"] = "Ирак"
	countryCodeMap["IR"] = "Иран"
	countryCodeMap["IE"] = "Ирландия"
	countryCodeMap["IS"] = "Исландия"
	countryCodeMap["ES"] = "Испания"
	countryCodeMap["IT"] = "Италия"
	countryCodeMap["YE"] = "Йемен"
	countryCodeMap["CV"] = "Кабо-Верде"
	countryCodeMap["KZ"] = "Казахстан"
	countryCodeMap["KH"] = "Камбоджа"
	countryCodeMap["CM"] = "Камерун"
	countryCodeMap["CA"] = "Канада"
	countryCodeMap["QA"] = "Катар"
	countryCodeMap["KE"] = "Кения"
	countryCodeMap["CY"] = "Кипр"
	countryCodeMap["KG"] = "Киргизия"
	countryCodeMap["KI"] = "Кирибати"
	countryCodeMap["KP"] = "КНДР (Корейская Народно-Демократическая Республика)"
	countryCodeMap["CN"] = "Китай (Китайская Народная Республика)"
	countryCodeMap["CO"] = "Колумбия"
	countryCodeMap["KM"] = "Коморы"
	countryCodeMap["CR"] = "Коста-Рика"
	countryCodeMap["CI"] = "Кот-д’Ивуар"
	countryCodeMap["CU"] = "Куба"
	countryCodeMap["KW"] = "Кувейт"
	countryCodeMap["LA"] = "Лаос"
	countryCodeMap["LV"] = "Латвия"
	countryCodeMap["LS"] = "Лесото"
	countryCodeMap["LR"] = "Либерия"
	countryCodeMap["LB"] = "Ливан"
	countryCodeMap["LY"] = "Ливия"
	countryCodeMap["LT"] = "Литва"
	countryCodeMap["LI"] = "Лихтенштейн"
	countryCodeMap["LU"] = "Люксембург"
	countryCodeMap["MU"] = "Маврикий"
	countryCodeMap["MR"] = "Мавритания"
	countryCodeMap["MG"] = "Мадагаскар"
	countryCodeMap["MK"] = "Северная Македония"
	countryCodeMap["BL"] = "Сен-Бартельми"
	countryCodeMap["MW"] = "Малави"
	countryCodeMap["MY"] = "Малайзия"
	countryCodeMap["ML"] = "Мали"
	countryCodeMap["MV"] = "Мальдивы"
	countryCodeMap["MT"] = "Мальта"
	countryCodeMap["MA"] = "Марокко"
	countryCodeMap["MH"] = "Маршалловы Острова"
	countryCodeMap["MX"] = "Мексика"
	countryCodeMap["FM"] = "Микронезия"
	countryCodeMap["MZ"] = "Мозамбик"
	countryCodeMap["MD"] = "Молдавия"
	countryCodeMap["MC"] = "Монако"
	countryCodeMap["MN"] = "Монголия"
	countryCodeMap["MM"] = "Мьянма"
	countryCodeMap["NA"] = "Намибия"
	countryCodeMap["NR"] = "Науру"
	countryCodeMap["NP"] = "Непал"
	countryCodeMap["NE"] = "Нигер"
	countryCodeMap["NG"] = "Нигерия"
	countryCodeMap["NL"] = "Нидерланды"
	countryCodeMap["NI"] = "Никарагуа"
	countryCodeMap["NZ"] = "Новая Зеландия"
	countryCodeMap["NO"] = "Норвегия"
	countryCodeMap["AE"] = "ОАЭ"
	countryCodeMap["OM"] = "Оман"
	countryCodeMap["PK"] = "Пакистан"
	countryCodeMap["PW"] = "Палау"
	countryCodeMap["PA"] = "Панама"
	countryCodeMap["PG"] = "Папуа — Новая Гвинея"
	countryCodeMap["PY"] = "Парагвай"
	countryCodeMap["PE"] = "Перу"
	countryCodeMap["PL"] = "Польша"
	countryCodeMap["PT"] = "Португалия"
	countryCodeMap["CG"] = "Республика Конго"
	countryCodeMap["KR"] = "Республика Корея"
	countryCodeMap["RU"] = "Россия"
	countryCodeMap["RW"] = "Руанда"
	countryCodeMap["RO"] = "Румыния"
	countryCodeMap["SV"] = "Сальвадор"
	countryCodeMap["WS"] = "Самоа"
	countryCodeMap["SM"] = "Сан-Марино"
	countryCodeMap["ST"] = "Сан-Томе и Принсипи"
	countryCodeMap["SA"] = "Саудовская Аравия"
	countryCodeMap["SZ"] = "Эсватини"
	countryCodeMap["SC"] = "Сейшельские Острова"
	countryCodeMap["SN"] = "Сенегал"
	countryCodeMap["VC"] = "Сент-Винсент и Гренадины"
	countryCodeMap["KN"] = "Сент-Китс и Невис"
	countryCodeMap["LC"] = "Сент-Люсия"
	countryCodeMap["RS"] = "Сербия"
	countryCodeMap["SG"] = "Сингапур"
	countryCodeMap["SY"] = "Сирия"
	countryCodeMap["SK"] = "Словакия"
	countryCodeMap["SI"] = "Словения"
	countryCodeMap["SB"] = "Соломоновы Острова"
	countryCodeMap["SO"] = "Сомали"
	countryCodeMap["SD"] = "Судан"
	countryCodeMap["SR"] = "Суринам"
	countryCodeMap["US"] = "США"
	countryCodeMap["SL"] = "Сьерра-Леоне"
	countryCodeMap["TJ"] = "Таджикистан"
	countryCodeMap["TH"] = "Таиланд"
	countryCodeMap["TZ"] = "Танзания"
	countryCodeMap["TG"] = "Того"
	countryCodeMap["TO"] = "Тонга"
	countryCodeMap["TT"] = "Тринидад и Тобаго"
	countryCodeMap["TV"] = "Тувалу"
	countryCodeMap["TN"] = "Тунис"
	countryCodeMap["TM"] = "Туркменистан"
	countryCodeMap["TR"] = "Турция"
	countryCodeMap["UG"] = "Уганда"
	countryCodeMap["UZ"] = "Узбекистан"
	countryCodeMap["UA"] = "Украина"
	countryCodeMap["UY"] = "Уругвай"
	countryCodeMap["FJ"] = "Фиджи"
	countryCodeMap["PH"] = "Филиппины"
	countryCodeMap["FI"] = "Финляндия"
	countryCodeMap["FR"] = "Франция"
	countryCodeMap["HR"] = "Хорватия"
	countryCodeMap["CF"] = "ЦАР"
	countryCodeMap["TD"] = "Чад"
	countryCodeMap["ME"] = "Черногория"
	countryCodeMap["CZ"] = "Чехия"
	countryCodeMap["CL"] = "Чили"
	countryCodeMap["CH"] = "Швейцария"
	countryCodeMap["SE"] = "Швеция"
	countryCodeMap["LK"] = "Шри-Ланка"
	countryCodeMap["EC"] = "Эквадор"
	countryCodeMap["GQ"] = "Экваториальная Гвинея"
	countryCodeMap["ER"] = "Эритрея"
	countryCodeMap["EE"] = "Эстония"
	countryCodeMap["ET"] = "Эфиопия"
	countryCodeMap["ZA"] = "ЮАР"
	countryCodeMap["SS"] = "Южный Судан"
	countryCodeMap["JM"] = "Ямайка"
	countryCodeMap["JP"] = "Япония"

	return countryCodeMap
}

func SMSMMSProviderNameRead(providerNameMap map[string]int) map[string]int {

	providerNameMap = map[string]int{
		"Topolo": 1,
		"Rond":   1,
		"Kildy":  1,
	}

	return providerNameMap
}

func MMSSliceValueDelete(mmsSlice []structure.MMSData, num int) []structure.MMSData {

	lenMmsSlice := len(mmsSlice)
	mmsSlice[num] = mmsSlice[lenMmsSlice-1]
	mmsSlice[lenMmsSlice-1] = structure.MMSData{}
	mmsSlice = mmsSlice[:lenMmsSlice-1]

	return mmsSlice
}

func IncidentSliceValueDelete(incidentSlice []structure.IncidentData, num int) []structure.IncidentData {

	lenIncidentSlice := len(incidentSlice)
	incidentSlice[num] = incidentSlice[lenIncidentSlice-1]
	incidentSlice[lenIncidentSlice-1] = structure.IncidentData{}
	incidentSlice = incidentSlice[:lenIncidentSlice-1]

	return incidentSlice
}
