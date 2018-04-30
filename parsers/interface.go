package parsers

import "io"

// BasicBookInfo book info
type BasicBookInfo struct {
	Author    string
	Title     string
	Coverpage string
	Extension string
	Words     map[string]string
}

// GetEPUBInfo return basic book info from epub
func GetEPUBInfo(rdr io.Reader) (result BasicBookInfo, err error) {
	return GetEPUBInfo(rdr)
}
