package shooting

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	Method      string
	Url         string
	HttpVersion string
	Host        string

	Headers map[string][]string

	Data []byte

	Name   string
	Raw    string
	Input  map[string]string
	Output map[string]string
}

func NewFromRaw(rawHttp string) (*Request, error) {
	var request Request
	var err error
	headers := make(map[string][]string)
	request.Raw = rawHttp

	lines := strings.Split(rawHttp, "\r\n")

	// request line: GET / HTTP/1.1
	requestLine := strings.Split(lines[0], " ")
	request.Method = requestLine[0] // TODO: Validate method error with switch case http.Method...
	request.Url = requestLine[1]
	request.HttpVersion = requestLine[2]

	// Headers and Body
	isHeader := true
	body := ""
	for i := 1; i < len(lines); i++ {

		if lines[i] == "" && isHeader {
			isHeader = false
		} else if isHeader {
			header := strings.SplitN(lines[i], ":", 2)
			if strings.EqualFold(header[0], "Host") {
				// Extract Host
				request.Host = strings.TrimSpace(header[1])
			} else {
				// Other headers
				h := make([]string, 0)
				h = append(h, strings.TrimSpace(header[1]))
				headers[header[0]] = h
			}
		} else {
			body += lines[i] + "\r\n"
		}
	}

	request.Headers = headers
	request.Data = []byte(body)

	return &request, err
}

func (r *Request) LoadRequest() error {
	var err error
	headers := make(map[string][]string)

	lines := strings.Split(r.Raw, "\r\n")

	// request line: GET / HTTP/1.1
	requestLine := strings.Split(lines[0], " ")
	r.Method = requestLine[0] // TODO: Validate method error with switch case http.Method...
	r.Url = requestLine[1]
	r.HttpVersion = requestLine[2]

	// Headers and Body
	isHeader := true
	body := ""
	for i := 1; i < len(lines); i++ {

		if lines[i] == "" && isHeader {
			isHeader = false
		} else if isHeader {
			header := strings.SplitN(lines[i], ":", 2)
			if strings.EqualFold(header[0], "Host") {
				// Extract Host
				r.Host = strings.TrimSpace(header[1])
			} else {
				// Other headers
				h := make([]string, 0)
				h = append(h, strings.TrimSpace(header[1]))
				headers[header[0]] = h
			}
		} else {
			body += lines[i] + "\r\n"
		}
	}

	r.Headers = headers
	r.Data = []byte(body)

	return err
}

func (r Request) SendRequest(client *http.Client) (*http.Response, error) {
	bodyReader := bytes.NewReader(r.Data)

	target := fmt.Sprintf("https://%s%s", r.Host, r.Url)

	req, err := http.NewRequest(r.Method, target, bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header = r.Headers

	return client.Do(req)
}
