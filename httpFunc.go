package IPFSClient

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
)

func newFileUploadRequest(uri string, params, formData map[string]string,
	fileFormName, filePath string) (*http.Request, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileFormName, filePath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range formData {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	urlStr := uri + "?1=1"
	for key, val := range formData {
		urlStr = urlStr + "&" + url.QueryEscape(key) + "=" + url.QueryEscape(val)
	}
	request, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		return nil, err
	}

	// log.Println(request)
	// log.Println("request.Header")
	// log.Println(writer.FormDataContentType())
	// log.Println(request.Header)

	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request, err
}

func newFormPostRequest(uri string, params, formData map[string]string) (*http.Request, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range formData {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	urlStr := uri + "?1=1"
	for key, val := range formData {
		urlStr = urlStr + "&" + url.QueryEscape(key) + "=" + url.QueryEscape(val)
	}
	request, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		return nil, err
	}

	// log.Println(request)
	// log.Println("request.Header")
	// log.Println(writer.FormDataContentType())
	// log.Println(request.Header)

	request.Header.Set("Content-Type", writer.FormDataContentType())
	return request, err
}

func PostForm(uri string, params, formData map[string]string) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, val := range formData {
		_ = writer.WriteField(key, val)
	}
	err := writer.Close()
	if err != nil {
		return nil, err
	}

	urlStr := uri + "?1=1"
	for key, val := range formData {
		urlStr = urlStr + "&" + url.QueryEscape(key) + "=" + url.QueryEscape(val)
	}
	request, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		return nil, err
	}

	// log.Println(request)
	// log.Println("request.Header")
	// log.Println(writer.FormDataContentType())
	// log.Println(request.Header)

	request.Header.Set("Content-Type", writer.FormDataContentType())

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, err
}

func PostFormWithFile(uri string, params, formData map[string]string,
	fileFormName, filePath string) ([]byte, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileFormName, filePath)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	for key, val := range formData {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	urlStr := uri + "?1=1"
	for key, val := range formData {
		urlStr = urlStr + "&" + url.QueryEscape(key) + "=" + url.QueryEscape(val)
	}
	request, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	//log.Println(resp)

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, err
}
