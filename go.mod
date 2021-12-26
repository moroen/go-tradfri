module github.com/moroen/go-tradfri

go 1.16

replace github.com/moroen/go-tradfricoap => ../go-tradfricoap/sync

replace github.com/moroen/gocoap/v5 => ../gocoap/v5

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/moroen/go-tradfricoap v0.1.1
	github.com/satori/go.uuid v1.2.0
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
	golang.org/x/text v0.3.6 // indirect
)
