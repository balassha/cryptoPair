package Controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"cryptoCurrencies/Models"
	"cryptoCurrencies/Scheduler"
)

const (
	BTC     = "BTC"
	USD     = "USD"
	FSYMS   = "fsyms"
	TSYMS   = "tsyms"
	RAW     = "RAW"
	DISPLAY = "DISPLAY"
)

func GetRawDataController(c *gin.Context) {
	var pairs Models.CurrencyPair
	var pairData Models.Pairs
	pairs.Crypto, pairs.Legacy = c.DefaultQuery(FSYMS, BTC), c.DefaultQuery(TSYMS, USD)

	if response, _, _, err := Models.GetRawData(&pairs); err != nil {
		log.Println("Unable to communicate to Remote API, Trying to fetch data from DB.")
		if err := Models.GetAllDataFromDB(&pairData, &pairs); err != nil {
			log.Printf("Data not found in DB. Error : %v", err)
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			log.Println("Found Data in DB.")
			c.JSON(http.StatusOK, BuildHTTPResponse(&pairData))
		}
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		//Schedule DB Update
		Scheduler.StartDBSync(&pairs)
		c.JSON(http.StatusOK, response)
	}
}

func DBDataController(c *gin.Context) {
	var pairs Models.CurrencyPair
	var pairData Models.Pairs
	pairs.Crypto, pairs.Legacy = c.DefaultQuery(FSYMS, BTC), c.DefaultQuery(TSYMS, USD)

	if err := Models.GetAllDataFromDB(&pairData, &pairs); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, BuildHTTPResponse(&pairData))
	}
}

// A helper to build a Datastructure that will house the incoming response
// remote host
func BuildHTTPResponse(pairData *Models.Pairs) map[string]interface{} {
	m := make(map[string]interface{})
	fr := make(map[string]interface{})
	fd := make(map[string]interface{})
	tr := make(map[string]Models.RAW)
	td := make(map[string]Models.DISPLAY)
	tr[pairData.Tsyms] = pairData.Raw
	td[pairData.Tsyms] = pairData.Display
	fr[pairData.Fsyms] = tr
	fd[pairData.Fsyms] = td
	m[RAW] = fr
	m[DISPLAY] = fd

	return m
}
