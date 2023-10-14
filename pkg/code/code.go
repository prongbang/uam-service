package code

const (
	StatusOK                       = "SUC200"
	StatusCreated                  = "SUC201"
	StatusNoContent                = "SUC204"
	StatusBadRequest               = "CLI400"
	StatusUnAuthorization          = "CLI401"
	StatusForbidden                = "CLI403"
	StatusNotFound                 = "CLI404"
	StatusDataNotFound             = "DTA001"
	StatusDataDuplicated           = "DUP001"
	StatusDatabaseExecuteFailed    = "DTB001"
	StatusDatabaseNoDataUpdate     = "DTB002"
	StatusDatabaseQueryFailed      = "DTB003"
	StatusDatabaseReferred         = "DTB004"
	StatusInformationRequired      = "INF001"
	StatusInformationExecuteFailed = "INF002"
	StatusInformationNotMatch      = "INF003"
	StatusInformationExists        = "INF004"
	StatusTokenExpired             = "TKN001"
	StatusTokenInvalid             = "TKN002"
	StatusFileTypeNotSupported     = "FTY001"
	StatusFormatNotSupported       = "FMT001"
	StatusFtpConnectionFail        = "FTP001"
	StatusNotifyFTPFail            = "NTF001"
	StatusNotifyEmailFail          = "NTF002"
	StatusNotifyLineFail           = "NTF003"
	StatusDataInvalid              = "IVL001"
)
