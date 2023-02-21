package constants

// GOlang does not have macros.
// This allows us to emulate macros and have a central place to get and modify those values
const (
	ProjectName = "My Project"
	Title       = "Awesome Title"
	version     = 0x1

	SA_port = "4500"

	Timeout = 5

	IKE_HEADER_SIZE = 160

	IKE_SA_INIT = uint8(34)
	IKE_AUTH    = uint8(35)

	QKD_PAYLOAD = 128
	QKD_KEY_ID  = 125
)
