package main

import (
	"log"
	"net/http"
)

func main() {
	/*mes := byoaproto.ResponseEntity{CampaignID: 10, StrategyID: 15}
	marshaller := &jsonpb.Marshaler{}
	buf := bytes.Buffer{}
	marshaller.Marshal(&buf, &mes)
	fmt.Printf("%s\n", buf.Bytes())*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	})
	log.Fatal(http.ListenAndServe(":9998", nil))

}
