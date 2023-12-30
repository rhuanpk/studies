package consts

import "strconv"

type serverHost string

// ServerHost is the server host.
const ServerHost serverHost = "localhost"

// Protocol return the host protocol: `http`.
func (sh serverHost) Protocol() string {
	return "http"
}

// Parse return the server host parsed like "http://<host>".
func (sh serverHost) Parse() string {
	return sh.Protocol() + "://" + string(sh)
}

type serverPort int

// ServerPort is the api port.
const ServerPort serverPort = 9999

// String return the api port like string.
func (sp serverPort) String() string {
	return strconv.Itoa(int(sp))
}

// Parse return the api port parsed like ":<port>".
func (sp serverPort) Parse() string {
	return ":" + sp.String()
}
