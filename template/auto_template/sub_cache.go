package auto_template

import _ "embed"

//go:embed api_enter.go.tpl
var Api []byte

//go:embed router_enter.go.tpl
var Router []byte

//go:embed service_enter.go.tpl
var Service []byte
