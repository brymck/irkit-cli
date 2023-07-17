package wifi

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	c := Credentials{
		Security:  SECURITY_WPA_WPA2,
		SSID:      "foo",
		Password:  "bar",
		WifiIsSet: true,
	}

	expected := "8/666F6F/626172//2//////48"
	c.Serialize()

	if c.Serialize() != expected {
		t.Errorf("Expected %q, but got %q", expected, c.Serialize())
	}
}
