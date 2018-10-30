package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type SearchResult struct {
	Infos []Info `json:"Search"`
}

type Info struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	ImdbID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

const (
	API_URL    = "http://www.omdbapi.com"
	POSTER_URL = "http://img.omdbapi.com"
)

func searchQuery(term, token string) (*SearchResult, error) {
	v := url.Values{}
	v.Add("s", term)
	v.Add("apikey", token)

	url := fmt.Sprintf("%s/?%s", API_URL, v.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http.Get(%s) failed: %s", url, err)
	}
	defer resp.Body.Close()

	var ret SearchResult
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll() failed: %s", err)
	}

	err = json.Unmarshal(b, &ret)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal() failed: %s", err)
	}

	return &ret, nil
}

func downloadPoster(ret *SearchResult, id, token string) (string, error) {
	v := url.Values{}
	v.Add("apikey", token)
	v.Add("i", id)

	url := fmt.Sprintf("%s/?%s", POSTER_URL, v.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("http.Get(%s) failed: %s", url, err)
	}
	defer resp.Body.Close()

	path := id + ".jpg"
	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("os.Crete(%s) failed: %s", path, err)
	}

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", fmt.Errorf("io.Copy() failed (path: %s): %s", path, err)
	}

	return path, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: poster SEARCH_TERM\n")
		os.Exit(1)
	}

	token := os.Getenv("OMDB_API_TOKEN")

	ret, err := searchQuery(os.Args[1], token)
	if err != nil {
		fmt.Fprintf(os.Stderr, "searchQuery(%s) failed: %s\n", os.Args[1], err)
		os.Exit(2)
	}

	for _, i := range ret.Infos {
		path, err := downloadPoster(ret, i.ImdbID, token)
		if err != nil {
			fmt.Fprintf(os.Stderr, "downloadPoster() failed: %s\n", err)
		}
		fmt.Println("title: " + i.Title)
		fmt.Println("output: " + path)
	}
}
