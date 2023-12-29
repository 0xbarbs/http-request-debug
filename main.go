package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"net/url"
	"sort"
)

func main() {
	port := flag.Int("port", 8080, "Port to listen on")
	flag.Parse()

	http.HandleFunc("/", handler)

	// Start the HTTP server
	addr := fmt.Sprintf(":%d", *port)
	fmt.Printf("Listening on port %d...\n", *port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	color.Set(color.FgHiGreen, color.Bold)
	fmt.Printf("[%s] %s\n", r.Method, formatURL(r.URL))

	// Print headers in alphabetical order
	keys := make([]string, 0, len(r.Header))
	for key := range r.Header {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		values := r.Header[key]
		sort.Strings(values)
		for _, value := range values {
			color.Set(color.FgHiWhite, color.Bold)
			fmt.Print(key)
			color.Set(color.FgWhite, color.ResetBold)
			fmt.Printf(": %s\n", value)
		}
	}

	// Print body if it exists
	body, err := io.ReadAll(r.Body)
	if err == nil && len(body) > 0 {
		color.Set(color.FgHiWhite, color.ResetBold)
		fmt.Printf("%s\n", body)
	}

	fmt.Println("")
	w.WriteHeader(http.StatusOK)
}

// formatURL formats the URL with query parameters
func formatURL(u *url.URL) string {
	if len(u.RawQuery) > 0 {
		return fmt.Sprintf("%s?%s", u.Path, u.RawQuery)
	}
	return u.Path
}
