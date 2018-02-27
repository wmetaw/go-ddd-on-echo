package handlers

// Custom Error Code
const (
	// Client
	CodeClientHoge = 1000

	// server
	CodeServerHoge = 2000

	// infra
	CodeInfraHoge = 3000

	// Third Party
	CodeTpHoge = 4000

	// Other
	CodeOtherHoge = 5000
)

var codeText = map[int]string{
	CodeClientHoge: "CP-1000",
	CodeServerHoge: "SP-2000",
	CodeInfraHoge:  "IP-3000",
	CodeTpHoge:     "TP-4000",
	CodeOtherHoge:  "OP-5000",
}

// codeText returns a text for the Custom ErrorCode. It returns the empty
// string if the code is unknown.
func ErrCodeText(code int) string {
	return codeText[code]
}
