package template

import (
	"fmt"
	"testing"
)

func TestOpt_GenAndSendOTP(t *testing.T) {
	smsOTP := &sms{}
	o := otp{
		iOtp: smsOTP,
	}
	o.genAndSendOTP(4)
	fmt.Println("")
	emailOTP := &email{}
	o = otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}
