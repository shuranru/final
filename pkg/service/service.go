package service

import (
	"errors"
	"final/pkg/structure"
	"fmt"
	"io"
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
	body, err := io.ReadAll(resp.Body)
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
