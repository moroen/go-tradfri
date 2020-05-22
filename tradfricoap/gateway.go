package tradfricoap

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/buger/jsonparser"
	"github.com/shibukawa/configdir"
	"github.com/tucnak/store"

	coap "github.com/moroen/gocoap/v2"
)

var globalGatewayConfig GatewayConfig

type GatewayConfig struct {
	Gateway  string
	Identity string
	Passkey  string
}

// ErrorNoConfig error
var ErrorNoConfig = errors.New("Tradfri Error: No config")

func (c GatewayConfig) Describe() string {
	out, _ := json.Marshal(c)
	return string(out)
}

func init() {
	// You must init store with some truly unique path first!
	store.Init("tradfri")
}

func SetConfig(c GatewayConfig) {
	globalGatewayConfig = c
}

func GetConfig() (conf GatewayConfig, err error) {
	if globalGatewayConfig == (GatewayConfig{}) {
		err = ErrorNoConfig
	}
	return globalGatewayConfig, err
}

func LoadConfig() (config GatewayConfig, err error) {

	configDir := configdir.New("", "tradfri")
	folder := configDir.QueryFolderContainsFile("gateway.json")

	if folder == nil {
		return config, errors.New("Config not found")
	}

	data, err := folder.ReadFile("gateway.json")
	if err != nil {
		return config, errors.New("Config not found")
	}

	if err := json.Unmarshal(data, &config); err == nil {
		SetConfig(config)
	}

	return config, nil
}

func SaveConfig(conf GatewayConfig) (err error) {
	err = store.Save("gateway.json", &conf)
	if err == nil {
		log.Println("Saved new config: ", conf.Describe())
	}
	return err
}

func CreateIdent(gateway, key, ident string) error {

	payload := fmt.Sprintf("{\"%s\":\"%s\"}", attrIdent, ident)
	URI := uriIdent

	param := coap.RequestParams{Host: gateway, Port: 5684, Uri: URI, Id: "Client_identity", Key: key, Payload: payload}

	res, err := coap.PostRequest(param)
	if err != nil {
		return err
	}

	psk, err := jsonparser.GetString(res, "9091")
	if err != nil {
		return err
	}

	conf := GatewayConfig{Gateway: fmt.Sprintf("%s", gateway), Identity: ident, Passkey: psk}
	SaveConfig(conf)
	SetConfig(conf)
	return nil
}

func GetRequest(URI string) (retmsg []byte, err error) {

	conf, err := GetConfig()
	if err != nil {
		panic(err.Error())
	}

	param := coap.RequestParams{Host: conf.Gateway, Port: 5684, Uri: URI, Id: conf.Identity, Key: conf.Passkey}

	res, err := coap.GetRequest(param)
	if err != nil {
		if err == coap.ErrorHandshake {
			log.Fatalln("Connection timed out")
		} else {
			log.Println(err.Error())
			// panic(err.Error())
		}
	}

	return res, err
}

func PutRequest(URI string, Payload string) (retmsg []byte, err error) {

	conf, err := GetConfig()
	if err != nil {
		panic(err.Error())
	}

	param := coap.RequestParams{Host: conf.Gateway, Port: 5684, Uri: URI, Id: conf.Identity, Key: conf.Passkey, Payload: Payload}

	res, err := coap.PutRequest(param)
	if err != nil {
		panic(err.Error())
	}

	return res, nil
}
