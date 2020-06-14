package http

import (
	"time"

	"gitlab.66ifuel.com/golang-tools/golib/common"
)

type HttpRequest struct {
	httpMethod  string
	requestId   uint32
	startTime   time.Time
	remoteAddr  string
	requestTime string
	userAgent   string
}

type HttpResponse struct {
	status  int
	errCode common.ErrorCodeType
}
