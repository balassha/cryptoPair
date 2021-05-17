package Models

import "github.com/jinzhu/gorm"

type RAW struct {
	CHANGE24HOUR    float64 `json:"CHANGE24HOUR"`
	CHANGEPCT24HOUR float64 `json:"CHANGEPCT24HOUR"`
	OPEN24HOUR      float64 `json:"OPEN24HOUR"`
	VOLUME24HOUR    float64 `json:"VOLUME24HOUR"`
	VOLUME24HOURTO  float64 `json:"VOLUME24HOURTO"`
	LOW24HOUR       float64 `json:"LOW24HOUR"`
	HIGH24HOUR      float64 `json:"HIGH24HOUR"`
	PRICE           float64 `json:"PRICE"`
	LASTUPDATE      float64 `json:"LASTUPDATE"`
	SUPPLY          float64 `json:"SUPPLY"`
	MKTCAP          float64 `json:"MKTCAP"`
}

type DISPLAY struct {
	CHANGE24HOUR    string `json:"CHANGE24HOUR" gorm:"column:CHANGE24HOUR_DISPLAY"`
	CHANGEPCT24HOUR string `json:"CHANGEPCT24HOUR" gorm:"column:CHANGEPCT24HOUR_DISPLAY"`
	OPEN24HOUR      string `json:"OPEN24HOUR" gorm:"column:OPEN24HOUR_DISPLAY"`
	VOLUME24HOUR    string `json:"VOLUME24HOUR" gorm:"column:VOLUME24HOURDISPLAY"`
	VOLUME24HOURTO  string `json:"VOLUME24HOURTO" gorm:"column:VOLUME24HOURTO_DISPLAY"`
	LOW24HOUR       string `json:"LOW24HOUR" gorm:"column:LOW24HOUR_DISPLAY"`
	HIGH24HOUR      string `json:"HIGH24HOUR" gorm:"column:HIGH24HOUR_DISPLAY"`
	PRICE           string `json:"PRICE" gorm:"column:PRICE_DISPLAY"`
	LASTUPDATE      string `json:"LASTUPDATE" gorm:"column:LASTUPDATE_DISPLAY"`
	SUPPLY          string `json:"SUPPLY" gorm:"column:SUPPLY_DISPLAY"`
	MKTCAP          string `json:"MKTCAP" gorm:"column:MKTCAP_DISPLAY"`
}

type Pairs struct {
	gorm.Model
	Fsyms   string
	Tsyms   string
	Raw     RAW     `gorm:"embedded"`
	Display DISPLAY `gorm:"embedded"`
}

func (p *Pairs) TableName() string {
	return "Pairs"
}
