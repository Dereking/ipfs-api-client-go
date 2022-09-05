package IPFSClient

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

// func newFileUploadRequest(uri string, params, formData map[string]string,
// 	fileFormName, filePath string) (*http.Request, error) {
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)
// 	part, err := writer.CreateFormFile(fileFormName, filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	_, err = io.Copy(part, file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for key, val := range formData {
// 		_ = writer.WriteField(key, val)
// 	}
// 	err = writer.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	urlStr := uri + "?1=1"
// 	for key, val := range params {
// 		urlStr = urlStr + "&" + url.QueryEscape(key) + "=" + url.QueryEscape(val)
// 	}
// 	request, err := http.NewRequest("POST", urlStr, body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// log.Println(request)
// 	// log.Println("request.Header")
// 	// log.Println(writer.FormDataContentType())
// 	// log.Println(request.Header)

// 	request.Header.Set("Content-Type", writer.FormDataContentType())
// 	return request, err
// }

// func newFormPostRequest(uri string, params, formData map[string]string) (*http.Request, error) {

// 	body := &bytes.Buffer{}
// 	writer := multipart.NewWriter(body)

// 	for key, val := range formData {
// 		_ = writer.WriteField(key, val)
// 	}
// 	err := writer.Close()
// 	if err != nil {
// 		return nil, err
// 	}

// 	urlStr := uri + "?"
// 	for key, val := range params {
// 		urlStr = urlStr + "&" + url.QueryEscape(key) + "=" + url.QueryEscape(val)
// 	}
// 	request, err := http.NewRequest("POST", urlStr, body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// log.Println(request)
// 	// log.Println("request.Header")
// 	// log.Println(writer.FormDataContentType())
// 	// log.Println(request.Header)

// 	request.Header.Set("Content-Type", writer.FormDataContentType())
// 	return request, err
// }

func PostUrl(urlStr string) ([]byte, error) {

	body := &bytes.Buffer{}

	request, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		return nil, errors.New(urlStr + " err:" + err.Error())
	}

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.New(urlStr + " err:" + err.Error())
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(urlStr + " err:" + err.Error())
	}

	return b, err
}

// the values in params and formData should have  been QueryEscaped
func PostForm(uri string, params, formData map[string][]string) ([]byte, error) {

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	for key, vals := range formData {
		for _, v := range vals {
			_ = writer.WriteField(key, v)
		}
	}
	err := writer.Close()
	if err != nil {
		return nil, errors.New(uri + " err:" + err.Error())
	}

	urlStr := uri + "?"
	for key, vals := range params {
		for _, v := range vals {
			urlStr = urlStr + "&" + key + "=" + v
		}
	}
	request, err := http.NewRequest("POST", urlStr, body)
	if err != nil {
		return nil, errors.New(urlStr + " err:" + err.Error())
	}

	log.Println(urlStr)
	//log.Println(body)
	// log.Println("request.Header")
	// log.Println(writer.FormDataContentType())
	// log.Println(request.Header)

	request.Header.Set("Content-Type", writer.FormDataContentType())

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.New(urlStr + " err:" + err.Error())
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(urlStr + " err:" + err.Error())
	}

	return b, err
}

func PostFormWithFile(uri string, params, formData map[string][]string,
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

	for key, vals := range formData {
		for _, v := range vals {
			_ = writer.WriteField(key, v)
		}
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}

	urlStr := uri + "?"
	for key, vals := range params {
		for _, v := range vals {
			urlStr = urlStr + "&" + key + "=" + v
		}
	}

	//log.Println("urlstr:= ", urlStr)

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
