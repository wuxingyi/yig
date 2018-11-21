package api

import (
	"github.com/journeymidnight/yig/crypto"
	"net/http"
)

// hasServerSideEncryptionHeader returns true if the given HTTP header
// contains server-side-encryption.
func hasServerSideEncryptionHeader(header http.Header) bool {
	return crypto.S3.IsRequested(header) || crypto.SSEC.IsRequested(header)
}
