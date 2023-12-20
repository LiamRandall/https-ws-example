package main

import (
    "fmt"
    "net/http"
)

func setSecurityHeaders(w http.ResponseWriter) {

    // Content Security Policy 
	// Note for simplicity of demo adding 'unsafe-inline' directive for script-src to execute in body below;
	// allowing inline script execution potentially enables XSS
    w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; object-src 'none';")

    // X-Content-Type-Options
    w.Header().Set("X-Content-Type-Options", "nosniff")

    // X-Frame-Options
    w.Header().Set("X-Frame-Options", "DENY")

    // X-XSS-Protection
    w.Header().Set("X-XSS-Protection", "1; mode=block")

    // Strict-Transport-Security
    // Uncomment the next line if serving over HTTPS
    // w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

    // Referrer-Policy
    w.Header().Set("Referrer-Policy", "no-referrer")

    // Feature-Policy
    // Example policy, restricts various browser features
    w.Header().Set("Feature-Policy", "geolocation 'none'; microphone 'none'; camera 'none'")
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        setSecurityHeaders(w)
        fmt.Fprintf(w, `<html>
<head><title>WebSocket Test</title></head>
<body>
<script>
    var ws = new WebSocket("ws://localhost:8081/ws");
    ws.onopen = function() {
        console.log("Connected to WebSocket");
    };
    ws.onerror = function(error) {
        console.log("WebSocket Error: " + error);
    };
</script>
</body>
</html>`)
    })

    fmt.Println("HTTP server listening on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
