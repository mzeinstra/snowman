package function

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

func getRemoteResource(uri string, config map[interface{}]interface{}) (*string, error) {
	_, err := url.Parse(uri)
	if err != nil {
		return nil, errors.New("Invalid argument given to get_remote template function.")
	}

	req, err := http.NewRequest("GET", uri, bytes.NewBuffer(nil))
	if err != nil {
		return nil, err
	}

	if config != nil {
		if config["headers"] != nil {
			headers := config["headers"].(map[interface{}]interface{})
			for key, value := range headers {
				req.Header.Set(key.(string), value.(string))
			}
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	responseString := string(bodyBytes)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Received bad(HTTP: " + resp.Status + ") response from " + uri + ".")
		fmt.Println(responseString)
		return nil, errors.New("Received bad response from remote resource.")
	}

	return &responseString, nil
}

var remoteFuncs = map[string]interface{}{
	"get_remote": func(uri string) (*string, error) {
		return getRemoteResource(uri, nil)
	},
	"get_remote_with_config": func(uri string, config map[interface{}]interface{}) (*string, error) {
		return getRemoteResource(uri, config)
	},
}

func GetRemoteFuncs() template.FuncMap {
	return template.FuncMap(remoteFuncs)
}
