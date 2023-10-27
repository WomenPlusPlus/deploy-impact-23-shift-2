package utils

import (
	"regexp"
	"time"
)

var InvitationTimeout = time.Hour * 24 * 30

var EmailRegex = regexp.MustCompile(`^[\w.]+(\+(\d|\w)+)?@([\w-]+\.)+[\w-]{2,10}$`)
