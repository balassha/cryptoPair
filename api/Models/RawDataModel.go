package Models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"cryptoCurrencies/Database"
	"cryptoCurrencies/Utils"
)

const (
	RemoteAPI    = "https://min-api.cryptocompare.com/data/pricemultifull?fsyms="
	RAW_Type     = "RAW"
	DISPLAY_Type = "DISPLAY"
)

//Gets data from Remote API and converts it to Response Type
func GetRawData(pair *CurrencyPair) (interface{}, RAW, DISPLAY, error) {

	data, err := http.Get(RemoteAPI + pair.Crypto + "&tsyms=" + pair.Legacy)
	if err != nil {
		return nil, RAW{}, DISPLAY{}, fmt.Errorf("error while retreiving data from remote host : %v", err)
	}

	rawData := make(map[string]map[string]map[string]map[string]interface{})

	err = json.NewDecoder(data.Body).Decode(&rawData)
	if err != nil {
		return nil, RAW{}, DISPLAY{}, fmt.Errorf("error while decoding data : %v", err)
	}

	raw := make(map[string]interface{})
	display := make(map[string]interface{})

	for _, v := range Utils.ResponseKeys {
		raw[v] = rawData[RAW_Type][pair.Crypto][pair.Legacy][v]
		display[v] = rawData[DISPLAY_Type][pair.Crypto][pair.Legacy][v]
	}

	rawData[RAW_Type][pair.Crypto][pair.Legacy] = raw
	rawData[DISPLAY_Type][pair.Crypto][pair.Legacy] = display

	rawStruct := RAW{
		CHANGE24HOUR:    raw["CHANGE24HOUR"].(float64),
		CHANGEPCT24HOUR: raw["CHANGEPCT24HOUR"].(float64),
		OPEN24HOUR:      raw["OPEN24HOUR"].(float64),
		VOLUME24HOUR:    raw["VOLUME24HOUR"].(float64),
		VOLUME24HOURTO:  raw["VOLUME24HOURTO"].(float64),
		LOW24HOUR:       raw["LOW24HOUR"].(float64),
		HIGH24HOUR:      raw["HIGH24HOUR"].(float64),
		PRICE:           raw["PRICE"].(float64),
		LASTUPDATE:      raw["LASTUPDATE"].(float64),
		SUPPLY:          raw["SUPPLY"].(float64),
		MKTCAP:          raw["MKTCAP"].(float64),
	}

	displayStruct := DISPLAY{
		CHANGE24HOUR:    display["CHANGE24HOUR"].(string),
		CHANGEPCT24HOUR: display["CHANGEPCT24HOUR"].(string),
		OPEN24HOUR:      display["OPEN24HOUR"].(string),
		VOLUME24HOUR:    display["VOLUME24HOUR"].(string),
		VOLUME24HOURTO:  display["VOLUME24HOURTO"].(string),
		LOW24HOUR:       display["LOW24HOUR"].(string),
		HIGH24HOUR:      display["HIGH24HOUR"].(string),
		PRICE:           display["PRICE"].(string),
		LASTUPDATE:      display["LASTUPDATE"].(string),
		SUPPLY:          display["SUPPLY"].(string),
		MKTCAP:          display["MKTCAP"].(string),
	}

	return rawData, rawStruct, displayStruct, nil
}

func GetAllDataFromDB(pairData *Pairs, pair *CurrencyPair) (err error) {
	if err = Database.DB.Where(map[string]interface{}{"fsyms": pair.Crypto, "tsyms": pair.Legacy}).First(pairData).Error; err != nil {
		return err
	}
	return nil
}

func DBSync(currencyPair CurrencyPair) error {

	_, raw, display, err := GetRawData(&currencyPair)
	if err != nil {
		return fmt.Errorf("error occured while DBSync : %v", err)
	}

	pair := &Pairs{
		Fsyms:   currencyPair.Crypto,
		Tsyms:   currencyPair.Legacy,
		Raw:     raw,
		Display: display,
	}

	//Update data if it already exists
	if err = GetAllDataFromDB(pair, &currencyPair); err != nil {
		if err := Database.DB.Create(pair).Error; err != nil {
			return fmt.Errorf("error occured while DBSync : %v", err)
		}
	} else {
		Database.DB.Save(pair)
	}

	return nil
}
