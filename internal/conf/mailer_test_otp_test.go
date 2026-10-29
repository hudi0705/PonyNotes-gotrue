package conf

import (
	"testing"
	"time"
)

func TestMailerTestOTP(t *testing.T) {
	now := time.Now()
	config := MailerConfiguration{}
	if _, ok := config.GetTestOTP("debug@ponynotes.test", now); ok {
		t.Fatal("test OTP must be disabled by default")
	}
	config.TestOTP = map[string]string{"debug@ponynotes.test": "123456"}
	if otp, ok := config.GetTestOTP("DEBUG@ponynotes.test", now); !ok || otp != "123456" {
		t.Fatal("configured test address must use its test OTP")
	}
	if _, ok := config.GetTestOTP("other@ponynotes.test", now); ok {
		t.Fatal("unconfigured addresses must use normal authentication")
	}
	config.TestOTPValidUntil = Time{Time: now.Add(time.Minute)}
	if _, ok := config.GetTestOTP("debug@ponynotes.test", now); !ok {
		t.Fatal("test OTP must work before its configured deadline")
	}
	if _, ok := config.GetTestOTP("debug@ponynotes.test", config.TestOTPValidUntil.Time); ok {
		t.Fatal("test OTP must stop working at its configured deadline")
	}
}
