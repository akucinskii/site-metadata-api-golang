package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /", handleRequest)
    if err := http.ListenAndServe(":8080", mux); err != nil {
        fmt.Println(err.Error())
    }
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")
    if url == "" {
        http.Error(w, "Missing 'url' query parameter", http.StatusBadRequest)
        return
    }

    dataAsString, err := fetchURLContent(url)
    if err != nil {
        http.Error(w, "Failed to fetch site metadata", http.StatusInternalServerError)
        return
    }

    metadata := extractMetadata(dataAsString)
    jsonData, err := json.Marshal(metadata)
    if err != nil {
        http.Error(w, "Failed to convert metadata to JSON", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonData)
}

func fetchURLContent(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    bytes, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(bytes), nil
}

func extractMetadata(htmlContent string) map[string]interface{} {
    metadata := make(map[string]interface{})
    properties := []string{"og:title", "og:description", "og:image"}

    for _, property := range properties {
        metadata[property] = extractProperty(htmlContent, property)
    }

    return metadata
}

func extractProperty(htmlContent string, property string) string {
    doc, err := html.Parse(strings.NewReader(htmlContent))
    if err != nil {
        return ""
    }

    var foundProperty string
    var f func(*html.Node)
    f = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "meta" {
            for _, attr := range n.Attr {
                if attr.Key == "property" && attr.Val == property {
                    foundProperty = getMetaContent(n)
                    return
                }
            }
        }
        for c := n.FirstChild; c != nil; c = c.NextSibling {
            f(c)
        }
    }
    f(doc)
    return foundProperty
}

func getMetaContent(n *html.Node) string {
    for _, attr := range n.Attr {
        if attr.Key == "content" {
            return attr.Val
        }
    }
    return ""
}