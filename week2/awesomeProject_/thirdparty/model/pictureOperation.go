package model

import "net/http"

func SendDelete(path string, accessToken string) {
	req, _ := http.NewRequest(http.MethodDelete, "http://localhost:9090/picutre/manage", nil)
	req.Header.Add("path", path)
	req.Header.Add("access_token", accessToken)
	req.Header.Add("Content-Type", "text/plain")
	_, _ = http.DefaultClient.Do(req)
}

func SendShow(accessToken string) {
	req, _ := http.NewRequest(http.MethodGet, "http://localhost:9090/picutre/manage", nil)
	req.Header.Add("access_token", accessToken)
	req.Header.Add("Content-Type", "text/plain")
	_, _ = http.DefaultClient.Do(req)
}

func SendAdd(path string, accessToken string) {
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:9090/picutre/manage", nil)
	req.Header.Add("path", path)
	req.Header.Add("access_token", accessToken)
	req.Header.Add("Content-Type", "text/plain")
	_, _ = http.DefaultClient.Do(req)

}
