package img

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func MiddlewareImg(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		r.Body = http.MaxBytesReader(w, r.Body, 10<<20) // 10 MB

		buf, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Cannot read file", http.StatusBadRequest)
			return
		}

		contentType := http.DetectContentType(buf)
		if contentType != "image/jpeg" && contentType != "image/png" {
			log.Println("Uploaded content type:", contentType)
			log.Printf("Content-Type (header): %s", r.Header.Get("Content-Type"))
			log.Printf("Detected (magic): %s", contentType)

			http.Error(w, "Unsupported image type", http.StatusUnsupportedMediaType)
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(buf))

		next.ServeHTTP(w, r)
	})
}
