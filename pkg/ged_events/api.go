package ged_events

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GedResponse struct {
   TotalCount int
   TotalPages int
   PreviousPageUrl string
   NextPageUrl string
   Result []GedEvent
}

type GedEvent struct {
   Id int `json:"id"`
   Year int `json:"year"`
}

func TestCall() (GedResponse, error) {
   response := GedResponse{}

   req, err := http.NewRequest("GET",fmt.Sprint("./ged/",API_VERSION, "?pagesize=5&page=1"), nil)
   if err != nil {
      return response, fmt.Errorf("Error when creating request: %s",err)
   }

   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return response, fmt.Errorf("Error when making request: %s",err)
   }

   body,err := ioutil.ReadAll(res.Body)
   err = json.Unmarshal(body, &response)
   return response, nil
}
