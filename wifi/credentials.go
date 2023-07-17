package wifi

import (
	"github.com/brymck/irkit-cli/crc8"
	"github.com/brymck/irkit-cli/serialization"
)

type Credentials struct {
	Security     byte
	SSID         string
	Password     string
	WifiIsSet    bool
	WifiWasValid bool
	DeviceKey    string
}

func (c *Credentials) Serialize() string {
	s := serialization.NewSerializer()
	s.WriteByte(c.Security)
	s.WriteSeparator()
	s.WriteString(c.SSID)
	s.WriteSeparator()
	if c.Security == SECURITY_WEP {
		// This might be a bug in the IRKit implementation?
		s2 := serialization.NewSerializer()
		s2.WriteString(c.Password)
		s.WriteString(s2.String())
	} else {
		s.WriteString(c.Password)
	}
	s.WriteSeparator()
	s.WriteString(c.DeviceKey)
	s.WriteSeparator()
	s.WriteByte(2) // 2: TELEC, 1: FCC, 0: ETSI
	s.WriteSeparator()
	s.WriteSeparator()
	s.WriteSeparator()
	s.WriteSeparator()
	s.WriteSeparator()
	s.WriteSeparator()
	s.WriteByte(c.CheckByte())
	return s.String()
}

func (c *Credentials) CheckByte() byte {
	crc := crc8.Crc8FromByte(c.Security, 0x00)
	crc = crc8.Crc8FromString(c.SSID, 33, crc)
	crc = crc8.Crc8FromString(c.Password, 64, crc)
	crc = crc8.Crc8FromBoolean(c.WifiIsSet, crc)
	crc = crc8.Crc8FromBoolean(c.WifiWasValid, crc)
	return crc8.Crc8FromString(c.DeviceKey, 33, crc)
}
