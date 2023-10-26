package utils

import (
	"regexp"
	"time"
)

var InvitationTimeout = time.Second * 432000

var EmailRegex = regexp.MustCompile(`^[\w.]+(\+(\d|\w)+)?@([\w-]+\.)+[\w-]{2,10}$`)
