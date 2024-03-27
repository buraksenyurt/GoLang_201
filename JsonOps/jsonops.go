package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	sampleData := `
	{
		"data": {
			"id":"11230495409546",
			"title":"Atari Pac Man",
			"year":1984,
			"point": 7.6,
			"object":"Arcade"
		}
	}
	`

	var m map[string]map[string]interface {
	}

	err := json.Unmarshal([]byte(sampleData), &m)
	checkError(err)
	fmt.Println(m)
	fmt.Println("************")
	content, err := json.Marshal(m)
	checkError(err)
	fmt.Println(string(content))

	startTime := time.Now()
	currencyDay := new(CurrencyDay)
	currencyDate := time.Now()
	currencyDay.GetData(currencyDate)
	elapsedTime := time.Since(startTime)
	fmt.Printf("Currencies fetched in %s\n", elapsedTime)
}

func (curr *CurrencyDay) GetData(currencyDate time.Time) {
	xDate := currencyDate
	t := new(tarih_Date)
	currDay := t.getData(currencyDate, xDate)
	for {
		if currDay == nil {
			currencyDate = currencyDate.AddDate(0, 0, -1)
			currDay := t.getData(currencyDate, xDate)
			if currDay != nil {
				break
			} else {
				break
			}
		}
	}

}

func (curr *tarih_Date) getData(currencyDate time.Time, xDate time.Time) *CurrencyDay {
	currDay := new(CurrencyDay)
	var res *http.Response
	var err error
	var url string

	url = "https://www.tcmb.gov.tr/kurlar/" + currencyDate.Format("200601") + "/" + currencyDate.Format("02012006") + ".xml"
	fmt.Println("Url info is ", url)

	res, err = http.Get(url)
	checkError(err)
	defer res.Body.Close()

	if res.StatusCode != http.StatusNotFound {
		tarih := new(tarih_Date)
		dcdr := xml.NewDecoder(res.Body)
		err = dcdr.Decode(tarih)
		checkError(err)
		curr = &tarih_Date{}
		currDay.Id = xDate.Format("02012006")
		currDay.Date = xDate
		currDay.DayNo = tarih.Bulten_No
		currDay.Currencies = make([]Currency, len(tarih.Currency))
		for i, curr := range tarih.Currency {
			currDay.Currencies[i].Code = curr.CurrencyCode
			currDay.Currencies[i].CurrencyName = curr.CurrencyName
			currDay.Currencies[i].CurrencyNameTR = curr.Isim
			currDay.Currencies[i].BanknoteBuying, _ = strconv.ParseFloat(curr.BanknoteBuying, 64)
			currDay.Currencies[i].BanknoteSelling, _ = strconv.ParseFloat(curr.BanknoteSelling, 64)
			currDay.Currencies[i].ForexBuying, _ = strconv.ParseFloat(curr.ForexBuying, 64)
			currDay.Currencies[i].ForexSelling, _ = strconv.ParseFloat(curr.ForexSelling, 64)
			currDay.Currencies[i].CrossOrder, _ = strconv.Atoi(curr.CrossRateOther)
			currDay.Currencies[i].CrossRateUSD, _ = strconv.ParseFloat(curr.CrossRateUSD, 64)
			currDay.Currencies[i].CrossRateOther, _ = strconv.ParseFloat(curr.CrossRateOther, 64)
			currDay.Currencies[i].Unit, _ = strconv.Atoi(curr.Unit) //ASCII to Int
		}

		fmt.Println(currDay)
		SaveJson("Currencies.json", currDay)
	} else {
		currDay = nil
	}

	return currDay
}

func SaveJson(fileName string, payload interface{}) {
	file, err := os.Create(fileName)
	checkError(err)
	encoder := json.NewEncoder(file)
	err = encoder.Encode(payload)
	checkError(err)
	file.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

type CurrencyDay struct {
	Id         string
	Date       time.Time
	Currencies []Currency
	DayNo      string
}

type Currency struct {
	Code            string
	CrossOrder      int
	Unit            int
	CurrencyNameTR  string
	CurrencyName    string
	BanknoteBuying  float64
	BanknoteSelling float64
	ForexBuying     float64
	ForexSelling    float64
	CrossRateUSD    float64
	CrossRateOther  float64
}

type tarih_Date struct {
	XMLName   xml.Name `xml:"Tarih_Date"`
	Tarih     string   `xml:"Tarih,attr"`
	Date      string   `xml:"Date,attr"`
	Bulten_No string   `xml:"Bulten_No,attr"`
	Currency  []xmlCurrency
}

type xmlCurrency struct {
	Kod             string `xml:"Kod,attr"`
	CrossOrder      string `xml:"CrossOrder,attr"`
	CurrencyCode    string `xml:"CurrencyCode,attr"`
	Unit            string `xml:"Unit"`
	Isim            string `xml:"Isim"`
	CurrencyName    string `xml:"CurrencyName"`
	ForexBuying     string `xml:"ForexBuying"`
	ForexSelling    string `xml:"ForexSelling"`
	BanknoteBuying  string `xml:"BanknoteBuying"`
	BanknoteSelling string `xml:"BanknoteSelling"`
	CrossRateUSD    string `xml:"CrossRateUSD"`
	CrossRateOther  string `xml:"CrossRateOther"`
}
