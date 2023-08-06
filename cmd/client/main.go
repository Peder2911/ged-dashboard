package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/peder2911/ged-dashboard/pkg/ged_events"
)

func main(){
   response, err := ged_events.TestCall()
   if err != nil {
      log.Printf("Error when making test call: %s", err)
      return
   }

   data, err := json.Marshal(response)
   if err != nil {
      log.Printf("Error when marshalling data to JSON: %s", err)
      return
   }

   fmt.Println(string(data))
}
