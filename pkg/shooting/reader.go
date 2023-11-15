package shooting

import (
	"os"
	"strings"
)

type Reader interface {
	init()
	readCustomFile(string) ([]*Request, error)
}

func ReadFromFile(filePath string, r Reader) ([]*Request, error) {
	r.init()
	return r.readCustomFile(filePath)
}

type RawReader struct {
	keyDelimiter     string
	requestDelimiter string
}

func (r *RawReader) init() {
	if r.requestDelimiter == "" {
		// default requestDelimiter
		r.requestDelimiter = "====k4it0z11====\r\n"
	}

	if r.keyDelimiter == "" {
		// default keyDelimiter
		r.keyDelimiter = "\r\n====minaris====\r\n"
	}
}

func (r *RawReader) readCustomFile(filePath string) ([]*Request, error) {

	result := make([]*Request, 0)

	raw, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	rawHttpString := string(raw)

	blocks := strings.Split(rawHttpString, r.requestDelimiter)
	for i := range blocks {
		inputMap := make(map[string]string)
		outputMap := make(map[string]string)
		// fmt.Println(requests[i])
		KeyAndRawRequest := strings.Split(blocks[i], r.keyDelimiter)

		// TODO: Implement validation
		name := KeyAndRawRequest[0]
		inputs := strings.Split(KeyAndRawRequest[1], "\r\n")
		outputs := strings.Split(KeyAndRawRequest[2], "\r\n")
		rawRequest := KeyAndRawRequest[3]

		// req, err := NewFromRaw(rawRequest)
		// if err != nil {
		// 	return nil, err
		// }

		req := new(Request)
		req.Raw = rawRequest

		for _, inp := range inputs {
			inputPair := strings.SplitN(inp, "--->", 2)
			inputMap[inputPair[0]] = inputPair[1]
		}
		req.Input = inputMap

		for _, outp := range outputs {
			outputPair := strings.SplitN(outp, "=", 2)
			outputMap[outputPair[0]] = outputPair[1]
		}
		req.Output = outputMap

		req.Name = name

		result = append(result, req)
	}

	return result, nil
}
