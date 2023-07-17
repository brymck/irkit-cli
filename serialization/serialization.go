package serialization

import "strings"

const (
	hextable  = "0123456789ABCDEF"
	separator = '/'
)

type Serializer struct {
	sb strings.Builder
}

func NewSerializer() *Serializer {
	return &Serializer{}
}

func (s *Serializer) WriteByte(b byte) {
	highOrderNibble := b >> 4
	if highOrderNibble != 0 {
		s.sb.WriteByte(hextable[highOrderNibble])
	}
	s.sb.WriteByte(hextable[b&0x0f])
}

func (s *Serializer) WriteBytes(src []byte) {
	for _, b := range src {
		s.WriteByte(b)
	}
}

func (s *Serializer) WriteString(src string) {
	s.WriteBytes([]byte(src))
}

func (s *Serializer) WriteSeparator() {
	s.sb.WriteRune(separator)
}

func (s *Serializer) String() string {
	return s.sb.String()
}
