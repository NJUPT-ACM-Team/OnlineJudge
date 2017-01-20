package captcha

import (
	locals "OnlineJudge/sessions"

	"github.com/dchest/captcha"

	"io"
)

const (
	CAPTCHA = "_captcha"

	StdWidth  = captcha.StdWidth
	StdHeight = captcha.StdHeight
)

func get(sess locals.Session) (string, bool) {
	t, ok := sess.Get(CAPTCHA)
	if !ok {
		return "", false
	}
	ca, ok := t.(string)
	return ca, ok
}

func set(sess locals.Session, ca string) {
	sess.Set(CAPTCHA, ca)
}

func reset(sess locals.Session) {
	sess.Set(CAPTCHA, nil)
}

func Prepare(sess locals.Session) string {
	ca, ok := get(sess)
	if !ok {
		// key not found, generate new captcha
		ca = captcha.New()
	} else {
		// already have one, reload it
		if !captcha.Reload(ca) {
			// failed to reload, generate new one
			ca = captcha.New()
		}
	}
	set(sess, ca)
	// get(sess)
	return ca
}

func Verify(sess locals.Session, str string) bool {
	ca, ok := get(sess)
	if !ok {
		return false
	}
	// reset(sess)
	return captcha.VerifyString(ca, str)
}

func WriteImage(w io.Writer, ca string, width, height int) {
	captcha.WriteImage(w, ca, width, height)
}
