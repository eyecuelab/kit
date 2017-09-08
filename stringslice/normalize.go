package stringslice

import "github.com/eyecuelab/kit/str"

//NFKD normalized each string to NFKD
func NFKD(a []string) []string {
	return str.Map(str.NFKD, a)
}

//NFC normalizes each string to NFC
func NFC(a []string) []string {
	return str.Map(str.NFC, a)
}

//NFD normalizes each string to NKD
func NFD(a []string) []string {
	return str.Map(str.NFD, a)
}

//NKFC normalizes each string to NKFC
func NKFC(a []string) []string {
	return str.Map(str.NKFC, a)
}
