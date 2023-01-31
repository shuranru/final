package service

import (
	"final/pkg/structure"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func FileToSlice(filename string) []string {

	var content []string

	contentTemp, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка чтения из файла", filename)
	}
	content = strings.Split(string(contentTemp), "\n")

	return content
}

func WebToByte(url string) ([]byte, int) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка при получении данных. ", err)
		return nil, 502
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при получении данных. ", err)
		return nil, 500
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
