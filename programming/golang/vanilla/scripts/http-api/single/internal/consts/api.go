package consts

type apiEndpoint string

// APIEndpoint is the endpoint path.
const APIEndpoint apiEndpoint = "xpto"

// Parse return the endpoint path parsed like "/<endpoint>".
func (ae apiEndpoint) Parse(isMult bool) string {
	endpoint := "/" + string(ae)
	return map[bool]string{
		true:  endpoint + "/",
		false: endpoint,
	}[isMult]
}

// URL build the full URL of endpoint.
func (ae apiEndpoint) URL() string {
	return ServerHost.Parse() + ServerPort.Parse() + ae.Parse(false)
}
