package base64encoding

// StandardCodeSet is the default, http safe code set of base64encoding
const StandardCodeSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"

// Base64WebSet is the standard base64 set for encoding data(e.g. images) in html files.
// However, this is not secure for using in URLs due to the '/' character
const Base64WebSet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

// EasilyReadableCodeSet is a set with no characters that look alike (e.g. 0 & O, l & I)
const EasilyReadableCodeSet = "*)23456789abcdefghi_klmnopqrstuvwxyzABCDEFGH+JKLMNOPQRSTUVWXYZ-$"