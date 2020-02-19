package tradfricoap

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	coap "github.com/moroen/gocoap/v2"
)

func Observe() error {

	log.Println("Observe called")

	conf, err := GetConfig()
	if err != nil {
		panic(err.Error())
	}

	param := coap.ObserveParams{Host: conf.Gateway, Port: 5684, Id: conf.Identity, Key: conf.Passkey}

	endpoints := `["15001/65554", "15001/65550"]`

	var uris []string

	err = json.Unmarshal([]byte(endpoints), &uris)
	if err != nil {
		panic(err.Error())
	}

	param.Uri = uris

	msg := make(chan []byte)
	sign := make(chan bool)
	errSign := make(chan error)

	// state := 0

	go coap.Observe(param, msg, sign, errSign)
	for {
		select {
		case message, isOpen := <-msg:
			if isOpen == true {
				fmt.Println(string(message))
			} else {
				return nil
			}
		case err = <-errSign:
			return err
		case <-time.After(120 * time.Second):
			sign <- true
			return nil
		}
	}
}
