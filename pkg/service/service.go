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

	countryCodeMap["AB"] = "Abkhazia"
	countryCodeMap["AU"] = "Australia"
	countryCodeMap["AT"] = "Austria"
	countryCodeMap["AZ"] = "Azerbaijan"
	countryCodeMap["AL"] = "Albania"
	countryCodeMap["DZ"] = "Algeria"
	countryCodeMap["AS"] = "American Samoa"
	countryCodeMap["AI"] = "Anguilla"
	countryCodeMap["AO"] = "Angola"
	countryCodeMap["AD"] = "Andorra"
	countryCodeMap["AQ"] = "Antarctica"
	countryCodeMap["AG"] = "Antigua and Barbuda"
	countryCodeMap["AR"] = "Argentina"
	countryCodeMap["AM"] = "Armenia"
	countryCodeMap["AW"] = "Aruba"
	countryCodeMap["AF"] = "Afghanistan"
	countryCodeMap["BS"] = "Bahamas"
	countryCodeMap["BD"] = "Bangladesh"
	countryCodeMap["BB"] = "Barbados"
	countryCodeMap["BH"] = "Bahrain"
	countryCodeMap["BY"] = "Belarus"
	countryCodeMap["BZ"] = "Belize"
	countryCodeMap["BE"] = "Belgium"
	countryCodeMap["BJ"] = "Benin"
	countryCodeMap["BM"] = "Bermuda"
	countryCodeMap["BG"] = "Bulgaria"
	countryCodeMap["BO"] = "Bolivia, plurinational state of"
	countryCodeMap["BQ"] = "Bonaire, Sint Eustatius and Saba"
	countryCodeMap["BA"] = "Bosnia and Herzegovina"
	countryCodeMap["BW"] = "Botswana"
	countryCodeMap["BR"] = "Brazil"
	countryCodeMap["IO"] = "British Indian Ocean Territory"
	countryCodeMap["BN"] = "Brunei Darussalam"
	countryCodeMap["BF"] = "Burkina Faso"
	countryCodeMap["BI"] = "Burundi"
	countryCodeMap["BT"] = "Bhutan"
	countryCodeMap["VU"] = "Vanuatu"
	countryCodeMap["HU"] = "Hungary"
	countryCodeMap["VE"] = "Venezuela"
	countryCodeMap["VG"] = "Virgin Islands, British"
	countryCodeMap["VI"] = "Virgin Islands, U.S."
	countryCodeMap["VN"] = "Vietnam"
	countryCodeMap["GA"] = "Gabon"
	countryCodeMap["HT"] = "Haiti"
	countryCodeMap["GY"] = "Guyana"
	countryCodeMap["GM"] = "Gambia"
	countryCodeMap["GH"] = "Ghana"
	countryCodeMap["GP"] = "Guadeloupe"
	countryCodeMap["GT"] = "Guatemala"
	countryCodeMap["GN"] = "Guinea"
	countryCodeMap["GW"] = "Guinea-Bissau"
	countryCodeMap["DE"] = "Germany"
	countryCodeMap["GG"] = "Guernsey"
	countryCodeMap["GI"] = "Gibraltar"
	countryCodeMap["HN"] = "Honduras"
	countryCodeMap["HK"] = "Hong Kong"
	countryCodeMap["GD"] = "Grenada"
	countryCodeMap["GL"] = "Greenland"
	countryCodeMap["GR"] = "Greece"
	countryCodeMap["GE"] = "Georgia"
	countryCodeMap["GU"] = "Guam"
	countryCodeMap["DK"] = "Denmark"
	countryCodeMap["JE"] = "Jersey"
	countryCodeMap["DJ"] = "Djibouti"
	countryCodeMap["DM"] = "Dominica"
	countryCodeMap["DO"] = "Dominican Republic"
	countryCodeMap["EG"] = "Egypt"
	countryCodeMap["ZM"] = "Zambia"
	countryCodeMap["EH"] = "Western Sahara"
	countryCodeMap["ZW"] = "Zimbabwe"
	countryCodeMap["IL"] = "Israel"
	countryCodeMap["IN"] = "India"
	countryCodeMap["ID"] = "Indonesia"
	countryCodeMap["JO"] = "Jordan"
	countryCodeMap["IQ"] = "Iraq"
	countryCodeMap["IR"] = "Iran, Islamic Republic of"
	countryCodeMap["IE"] = "Ireland"
	countryCodeMap["IS"] = "Iceland"
	countryCodeMap["ES"] = "Spain"
	countryCodeMap["IT"] = "Italy"
	countryCodeMap["YE"] = "Yemen"
	countryCodeMap["CV"] = "Cape Verde"
	countryCodeMap["KZ"] = "Kazakhstan"
	countryCodeMap["KH"] = "Cambodia"
	countryCodeMap["CM"] = "Cameroon"
	countryCodeMap["CA"] = "Canada"
	countryCodeMap["QA"] = "Qatar"
	countryCodeMap["KE"] = "Kenya"
	countryCodeMap["CY"] = "Cyprus"
	countryCodeMap["KG"] = "Kyrgyzstan"
	countryCodeMap["KI"] = "Kiribati"
	countryCodeMap["CN"] = "China"
	countryCodeMap["CC"] = "Cocos (Keeling) Islands"
	countryCodeMap["CO"] = "Colombia"
	countryCodeMap["KM"] = "Comoros"
	countryCodeMap["CG"] = "Congo"
	countryCodeMap["CD"] = "Congo, Democratic Republic of the"
	countryCodeMap["KP"] = "Korea, Democratic People's republic of"
	countryCodeMap["KR"] = "Korea, Republic of"
	countryCodeMap["CR"] = "Costa Rica"
	countryCodeMap["CI"] = "Cote d'Ivoire"
	countryCodeMap["CU"] = "Cuba"
	countryCodeMap["KW"] = "Kuwait"
	countryCodeMap["CW"] = "Curaçao"
	countryCodeMap["LA"] = "Lao People's Democratic Republic"
	countryCodeMap["LV"] = "Latvia"
	countryCodeMap["LS"] = "Lesotho"
	countryCodeMap["LB"] = "Lebanon"
	countryCodeMap["LY"] = "Libyan Arab Jamahiriya"
	countryCodeMap["LR"] = "Liberia"
	countryCodeMap["LI"] = "Liechtenstein"
	countryCodeMap["LT"] = "Lithuania"
	countryCodeMap["LU"] = "Luxembourg"
	countryCodeMap["MU"] = "Mauritius"
	countryCodeMap["MR"] = "Mauritania"
	countryCodeMap["MG"] = "Madagascar"
	countryCodeMap["YT"] = "Mayotte"
	countryCodeMap["MO"] = "Macao"
	countryCodeMap["MW"] = "Malawi"
	countryCodeMap["MY"] = "Malaysia"
	countryCodeMap["ML"] = "Mali"
	countryCodeMap["UM"] = "United States Minor Outlying Islands"
	countryCodeMap["MV"] = "Maldives"
	countryCodeMap["MT"] = "Malta"
	countryCodeMap["MA"] = "Morocco"
	countryCodeMap["MQ"] = "Martinique"
	countryCodeMap["MH"] = "Marshall Islands"
	countryCodeMap["MX"] = "Mexico"
	countryCodeMap["FM"] = "Micronesia, Federated States of"
	countryCodeMap["MZ"] = "Mozambique"
	countryCodeMap["MD"] = "Moldova"
	countryCodeMap["MC"] = "Monaco"
	countryCodeMap["MN"] = "Mongolia"
	countryCodeMap["MS"] = "Montserrat"
	countryCodeMap["MM"] = "Burma"
	countryCodeMap["NA"] = "Namibia"
	countryCodeMap["NR"] = "Nauru"
	countryCodeMap["NP"] = "Nepal"
	countryCodeMap["NE"] = "Niger"
	countryCodeMap["NG"] = "Nigeria"
	countryCodeMap["NL"] = "Netherlands"
	countryCodeMap["NI"] = "Nicaragua"
	countryCodeMap["NU"] = "Niue"
	countryCodeMap["NZ"] = "New Zealand"
	countryCodeMap["NC"] = "New Caledonia"
	countryCodeMap["NO"] = "Norway"
	countryCodeMap["AE"] = "United Arab Emirates"
	countryCodeMap["OM"] = "Oman"
	countryCodeMap["BV"] = "Bouvet Island"
	countryCodeMap["IM"] = "Isle of Man"
	countryCodeMap["NF"] = "Norfolk Island"
	countryCodeMap["CX"] = "Christmas Island"
	countryCodeMap["HM"] = "Heard Island and McDonald Islands"
	countryCodeMap["KY"] = "Cayman Islands"
	countryCodeMap["CK"] = "Cook Islands"
	countryCodeMap["TC"] = "Turks and Caicos Islands"
	countryCodeMap["PK"] = "Pakistan"
	countryCodeMap["PW"] = "Palau"
	countryCodeMap["PS"] = "Palestinian Territory, Occupied"
	countryCodeMap["PA"] = "Panama"
	countryCodeMap["VA"] = "Holy See (Vatican City State)"
	countryCodeMap["PG"] = "Papua New Guinea"
	countryCodeMap["PY"] = "Paraguay"
	countryCodeMap["PE"] = "Peru"
	countryCodeMap["PN"] = "Pitcairn"
	countryCodeMap["PL"] = "Poland"
	countryCodeMap["PT"] = "Portugal"
	countryCodeMap["PR"] = "Puerto Rico"
	countryCodeMap["MK"] = "Macedonia, The Former Yugoslav Republic Of"
	countryCodeMap["RE"] = "Reunion"
	countryCodeMap["RU"] = "Russian Federation"
	countryCodeMap["RW"] = "Rwanda"
	countryCodeMap["RO"] = "Romania"
	countryCodeMap["WS"] = "Samoa"
	countryCodeMap["SM"] = "San Marino"
	countryCodeMap["ST"] = "Sao Tome and Principe"
	countryCodeMap["SA"] = "Saudi Arabia"
	countryCodeMap["SH"] = "Saint Helena, Ascension And Tristan Da Cunha"
	countryCodeMap["MP"] = "Northern Mariana Islands"
	countryCodeMap["BL"] = "Saint Barthélemy"
	countryCodeMap["MF"] = "Saint Martin (French Part)"
	countryCodeMap["SN"] = "Senegal"
	countryCodeMap["VC"] = "Saint Vincent and the Grenadines"
	countryCodeMap["KN"] = "Saint Kitts and Nevis"
	countryCodeMap["LC"] = "Saint Lucia"
	countryCodeMap["PM"] = "Saint Pierre and Miquelon"
	countryCodeMap["RS"] = "Serbia"
	countryCodeMap["SC"] = "Seychelles"
	countryCodeMap["SG"] = "Singapore"
	countryCodeMap["SX"] = "Sint Maarten"
	countryCodeMap["SY"] = "Syrian Arab Republic"
	countryCodeMap["SK"] = "Slovakia"
	countryCodeMap["SI"] = "Slovenia"
	countryCodeMap["GB"] = "United Kingdom"
	countryCodeMap["US"] = "United States"
	countryCodeMap["SB"] = "Solomon Islands"
	countryCodeMap["SO"] = "Somalia"
	countryCodeMap["SD"] = "Sudan"
	countryCodeMap["SR"] = "Suriname"
	countryCodeMap["SL"] = "Sierra Leone"
	countryCodeMap["TJ"] = "Tajikistan"
	countryCodeMap["TH"] = "Thailand"
	countryCodeMap["TW"] = "Taiwan, Province of China"
	countryCodeMap["TZ"] = "Tanzania, United Republic Of"
	countryCodeMap["TL"] = "Timor-Leste"
	countryCodeMap["TG"] = "Togo"
	countryCodeMap["TK"] = "Tokelau"
	countryCodeMap["TO"] = "Tonga"
	countryCodeMap["TT"] = "Trinidad and Tobago"
	countryCodeMap["TV"] = "Tuvalu"
	countryCodeMap["TN"] = "Tunisia"
	countryCodeMap["TM"] = "Turkmenistan"
	countryCodeMap["TR"] = "Turkey"
	countryCodeMap["UG"] = "Uganda"
	countryCodeMap["UZ"] = "Uzbekistan"
	countryCodeMap["UA"] = "Ukraine"
	countryCodeMap["WF"] = "Wallis and Futuna"
	countryCodeMap["UY"] = "Uruguay"
	countryCodeMap["FO"] = "Faroe Islands"
	countryCodeMap["FJ"] = "Fiji"
	countryCodeMap["PH"] = "Philippines"
	countryCodeMap["FI"] = "Finland"
	countryCodeMap["FK"] = "Falkland Islands (Malvinas)"
	countryCodeMap["FR"] = "France"
	countryCodeMap["GF"] = "French Guiana"
	countryCodeMap["PF"] = "French Polynesia"
	countryCodeMap["TF"] = "French Southern Territories"
	countryCodeMap["HR"] = "Croatia"
	countryCodeMap["CF"] = "Central African Republic"
	countryCodeMap["TD"] = "Chad"
	countryCodeMap["ME"] = "Montenegro"
	countryCodeMap["CZ"] = "Czech Republic"
	countryCodeMap["CL"] = "Chile"
	countryCodeMap["CH"] = "Switzerland"
	countryCodeMap["SE"] = "Sweden"
	countryCodeMap["SJ"] = "Svalbard and Jan Mayen"
	countryCodeMap["LK"] = "Sri Lanka"
	countryCodeMap["EC"] = "Ecuador"
	countryCodeMap["GQ"] = "Equatorial Guinea"
	countryCodeMap["AX"] = "Åland Islands"
	countryCodeMap["SV"] = "El Salvador"
	countryCodeMap["ER"] = "Eritrea"
	countryCodeMap["SZ"] = "Eswatini"
	countryCodeMap["EE"] = "Estonia"
	countryCodeMap["ET"] = "Ethiopia"
	countryCodeMap["ZA"] = "South Africa"
	countryCodeMap["GS"] = "South Georgia and the South Sandwich Islands"
	countryCodeMap["OS"] = "South Ossetia"
	countryCodeMap["SS"] = "South Sudan"
	countryCodeMap["JM"] = "Jamaica"
	countryCodeMap["JP"] = "Japan"

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
