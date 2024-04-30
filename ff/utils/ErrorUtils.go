package ff_utils

import "log"

type ErrorUtils struct{}

func (this *ErrorUtils) FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
