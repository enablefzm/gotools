package guid

import (
	"crypto/rand"
	"errors"
	"fmt"
)

// Guid is a globally unique 16 byte identifier
type Guid [16]byte

// ErrInvalid is returned when parsing a string that is not formatted
// as a valid guid.
var ErrInvalid = errors.New("guid: invalid format")

// hex character lookup table
var hextable = [...]byte{
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
	0x08, 0x09, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
	0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10, 0x10,
}

// New generates a random RFC 4122-conformant version 4 Guid.
func New() *Guid {
	g := new(Guid)
	if _, err := rand.Read(g[:]); err != nil {
		panic(err)
	}
	g[6] = (g[6] & 0x0f) | 0x40 // version = 4
	g[8] = (g[8] & 0x3f) | 0x80 // variant = RFC 4122
	return g
}

// NewString is a helper function that returns a random RFC 4122-comformant
// version 4 Guid as a string.
func NewString() string {
	return New().String()
}

// IsGuid returns true if the string contains a properly formatted Guid.
func IsGuid(s string) bool {
	if len(s) != 36 {
		return false
	}
	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return false
	}
	for _, sub := range [...]string{
		s[0:8], s[9:13], s[14:18], s[19:23], s[24:36],
	} {
		for i := 0; i < len(sub); i++ {
			if hextable[sub[i]] == 16 {
				return false
			}
		}
	}
	return true
}

// ParseString returns the Guid represented by the string s.
func ParseString(s string) (*Guid, error) {
	if len(s) != 36 {
		return nil, ErrInvalid
	}
	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return nil, ErrInvalid
	}
	g := new(Guid)
	offset := 0
	for _, sub := range [...]string{
		s[0:8], s[9:13], s[14:18], s[19:23], s[24:36],
	} {
		for i := 0; i < len(sub); i, offset = i+2, offset+1 {
			c0 := hextable[sub[i]]
			c1 := hextable[sub[i+1]]
			if c0 == 0x10 || c1 == 0x10 {
				return nil, ErrInvalid
			}
			g[offset] = c0<<4 | c1
		}
	}
	return g, nil
}

// String returns a standard hexadecimal string version of the Guid.
// Lowercase characters are used.
func (g *Guid) String() string {
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		g[0:4], g[4:6], g[6:8], g[8:10], g[10:16])
}

// StringUpper returns a standard hexadecimal string version of the Guid.
// Uppercase characters are used.
func (g *Guid) StringUpper() string {
	return fmt.Sprintf("%X-%X-%X-%X-%X",
		g[0:4], g[4:6], g[6:8], g[8:10], g[10:16])
}

// IsConformant determines if the Guid is RFC 4122-conformant.  If the
// variant is "reserved for future definition" or the version is unknown,
// then it is non-conformant.
func (g *Guid) IsConformant() bool {
	version := (g[6] & 0xf0) >> 4
	if version < 1 || version > 5 {
		return false
	}
	return (g[8] & 0xe0) != 0xe0
}
