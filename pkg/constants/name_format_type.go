package constants

type NameFormatType string

const (
	NameFormat_None          NameFormatType = "none"
	NameFormat_FirstLast                    = "first_last"
	NameFormat_FirstL                       = "first_last_initial"
	NameFormat_CertificateID                = "cid"
)
