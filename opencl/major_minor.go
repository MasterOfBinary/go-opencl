package opencl

import (
	"strconv"
	"strings"
)

// MajorMinor contains a major and minor version number.
type MajorMinor struct {
	Major uint8
	Minor uint8
}

// String converts a MajorMinor to a string in the format <major>.<minor>.
func (p MajorMinor) String() string {
	return strconv.FormatUint(uint64(p.Major), 10) + "." + strconv.FormatUint(uint64(p.Minor), 10)
}

// ParseMajorMinor parses a string in the <major>.<minor> format and returns a MajorMinor.
func ParseMajorMinor(input string) (MajorMinor, error) {
	elems := strings.Split(input, ".")
	if len(elems) != 2 {
		return MajorMinor{}, ErrorParsingVersion
	}

	maj, errMaj := strconv.ParseUint(elems[0], 10, 64)
	if errMaj != nil {
		return MajorMinor{}, ErrorParsingVersion
	}

	min, errMin := strconv.ParseUint(elems[1], 10, 64)
	if errMin != nil {
		return MajorMinor{}, ErrorParsingVersion
	}

	return MajorMinor{
		Major: uint8(maj),
		Minor: uint8(min),
	}, nil
}
