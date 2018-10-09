package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

const (
	// APIURL github api endpoint
	APIURL       = "https://api.github.com"
	templatePath = "./template.json"
	acceptHeader = "application/vnd.github.v3.text-match+json"
)

// issueReq post data of Create() and Edit()
type issueReq struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"` // Markdown format
	State     string   `json:"state,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Labels    []string `json:"lables,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

// Issue result of Create(), Get(), Edit() and Close()
type Issue struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	State   string `json:"state"`
	Number  int    `json:"number"`
	HTMLURL string `json:"html_url"`
}

// Issues result of Search()
type Issues struct {
	TotalCount int      `json:"total_count"`
	Items      []*Issue `json:"items"`
}

// HTTPReq store data for HTTP Request (get, post, patch)
type HTTPReq struct {
	method       string
	url          string
	body         io.Reader
	auth         bool
	expectedCode int
}

func (r *HTTPReq) do() (*bytes.Buffer, error) {
	fmt.Fprintf(os.Stderr, "method:%s url:%s body:%t nil:%t\n", r.method, r.url, r.body, nil)

	req, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return nil, err
	}

	if r.auth {
		token := fmt.Sprintf("token %s", os.Getenv("GO_ISSUE_TOKEN"))
		req.Header.Set("Authorization", token)
	}

	req.Header.Set("Accept", acceptHeader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != r.expectedCode {
		return nil, fmt.Errorf("invalid http status! got:%d(%s) want:%d", res.StatusCode, res.Status, r.expectedCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(b), nil
}

// helper functions
func createTempFileFrom(src io.Reader) (string, error) {
	dst, err := ioutil.TempFile("", "go-issue.*")
	if err != nil {
		return "", fmt.Errorf("ioutil.TempFile failed: %s", err)
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return "", fmt.Errorf("io.Copy failed: %s", err)
	}

	return dst.Name(), nil
}

func edit(path string) (*bytes.Buffer, error) {
	editor, err := exec.LookPath(os.Getenv("EDITOR"))
	if err != nil {
		return nil, fmt.Errorf("os.Getenv failed: %s", err)
	}

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("os.Command.Run failed: %s", err)
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadFile failed: %s", err)
	}

	return bytes.NewBuffer(b), nil
}

// IssueSearch queries the GitHub issue tracker
func IssueSearch(terms []string) (*Issues, error) {
	req := &HTTPReq{
		method:       "GET",
		url:          APIURL + "/search/issues?q=" + url.QueryEscape(strings.Join(terms, " ")),
		expectedCode: http.StatusOK,
	}

	res, err := req.do()
	if err != nil {
		return nil, err
	}

	var ret Issues
	err = json.NewDecoder(res).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// IssueGet get a specified issue
func IssueGet(owner, repo, num string) (*Issue, error) {
	req := &HTTPReq{
		method:       "GET",
		url:          fmt.Sprintf("%s/repos/%s/%s/issues/%s", APIURL, owner, repo, num),
		expectedCode: http.StatusOK,
	}

	res, err := req.do()
	if err != nil {
		return nil, err
	}

	var ret Issue
	err = json.NewDecoder(res).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// IssueCreate create an issue
func IssueCreate(owner, repo string) (*Issue, error) {
	fp, err := os.Open(templatePath)
	if err != nil {
		return nil, err
	}

	path, err := createTempFileFrom(fp)
	if err != nil {
		return nil, fmt.Errorf("createTempFileFromTemplate() failed: %s", err)
	}
	defer os.Remove(path)

	body, err := edit(path)
	if err != nil {
		return nil, fmt.Errorf("edit failed: %s", err)
	}

	req := &HTTPReq{
		method:       "POST",
		url:          fmt.Sprintf("%s/repos/%s/%s/issues", APIURL, owner, repo),
		body:         body,
		auth:         true,
		expectedCode: http.StatusCreated,
	}

	res, err := req.do()
	if err != nil {
		return nil, fmt.Errorf("post() failed: %s", err)
	}

	var ret Issue
	err = json.NewDecoder(res).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// IssueClose close the specified issue
func IssueClose(owner, repo, num string) (*Issue, error) {
	body, err := json.Marshal(&issueReq{State: "closed"})
	if err != nil {
		return nil, err
	}

	req := &HTTPReq{
		method:       "PATCH",
		url:          fmt.Sprintf("%s/repos/%s/%s/issues/%s", APIURL, owner, repo, num),
		body:         bytes.NewBuffer(body),
		auth:         true,
		expectedCode: http.StatusOK,
	}

	res, err := req.do()
	if err != nil {
		return nil, err
	}

	var ret Issue
	err = json.NewDecoder(res).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}

// IssueEdit get specified issue, edit the issue and update
func IssueEdit(owner, repo, num string) (*Issue, error) {
	// get current issue info
	issue, err := IssueGet(owner, repo, num)
	if err != nil {
		return nil, err
	}

	// create temporary file
	src, err := json.MarshalIndent(&issue, "", "  ")
	if err != nil {
		return nil, err
	}

	path, err := createTempFileFrom(bytes.NewBuffer(src))
	if err != nil {
		return nil, err
	}

	// edit issue and post
	body, err := edit(path)
	if err != nil {
		return nil, fmt.Errorf("edit() failed: %s", err)
	}

	req := &HTTPReq{
		method:       "POST",
		url:          fmt.Sprintf("%s/repos/%s/%s/issues/%s", APIURL, owner, repo, num),
		body:         body,
		auth:         true,
		expectedCode: http.StatusOK,
	}

	res, err := req.do()
	if err != nil {
		return nil, fmt.Errorf("post() failed: %s", err)
	}

	var ret Issue
	err = json.NewDecoder(res).Decode(&ret)
	if err != nil {
		return nil, err
	}

	return &ret, nil
}
