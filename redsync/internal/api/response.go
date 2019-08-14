package api

import "net/http"

// GeneralSuccess defines general api success response
var GeneralSuccess = Response{
	HTTPStatusCode: http.StatusOK,
	Code:           1000,
	Message:        "Ok",
}
