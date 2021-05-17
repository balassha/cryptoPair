package Scheduler

import (
	"fmt"
	"log"
	"time"

	"cryptoCurrencies/Models"
)

var done chan bool
var ticker *time.Ticker
var currencyPairs []Models.CurrencyPair

/*
* This is a Scheduler that periodically sync the data to DB from Remote API
* Sync begins the first request to API and it happens every 2 minutes
 */
func StartDBSync(pair *Models.CurrencyPair) {
	if ticker != nil {
		ticker.Stop()
		done <- true
		log.Println("DB Sync has been stopped succesfully. It will restart soon")
	}

	ticker = time.NewTicker(2 * time.Minute)
	done = make(chan bool)

	pairAlreadyExists := false
	for _, element := range currencyPairs {
		if element.Crypto == pair.Crypto && element.Legacy == pair.Legacy {
			pairAlreadyExists = true
		}
	}

	if !pairAlreadyExists {
		currencyPairs = append(currencyPairs, *pair)
	}

	go func() {
		log.Println("DB Sync Started !")
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				for _, pair := range currencyPairs {
					if err := Models.DBSync(pair); err != nil {
						fmt.Printf("Time : %v, Error : %v", t, err)
					} else {
						fmt.Printf("DB Sync successful at %v for Currency Pair %v\n", t, pair)
					}
				}
			}
		}
	}()
}

/*
* Function to Terminate the Periodic sync when the server is turned off
 */
func StopDBSync() {
	if ticker != nil {
		ticker.Stop()
		done <- true
	}
	log.Println("DB Sync has been stopped succesfully")
}
