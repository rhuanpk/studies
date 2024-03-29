package consts

import "strconv"

// PROTO Type
type serverProto string

const (
	http  serverProto = "http"
	https serverProto = "https"
)

// HOST Type
type serverHost struct {
	Host  string
	Proto serverProto
}

// Parse return the server host parsed like "<proto>://<host>".
func (sh serverHost) Parse() string {
	return string(sh.Proto) + "://" + sh.Host
}

// PORT Type
type serverPort int

// String return the api port like string.
func (sp serverPort) String() string {
	return strconv.Itoa(int(sp))
}

// Parse return the api port parsed like ":<port>".
func (sp serverPort) Parse() string {
	return ":" + sp.String()
}
