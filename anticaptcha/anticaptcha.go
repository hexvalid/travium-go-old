package anticaptcha

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"github.com/kataras/golog"
	"net/http"
	"net/url"
	"time"
)

var (
	CfgApiKey             string = "5efd3fe7179bdf7d576bf709d4b698c6"
	CfgCheckInterval      int    = 10
	CfgTimeout            int    = 90
	CfgNoSlotWaitInterval int    = 30
	baseURL                      = &url.URL{Host: "api.anti-captcha.com", Scheme: "http", Path: "/"}
	client                       = &http.Client{
		Timeout: time.Duration(CfgTimeout) * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}}
)

func execRequest(method, url string, bod []byte) (out map[string]interface{}, err error) {
	var res *http.Response
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(bod))
	req.Header.Add("Content-Type", "application/json")
	res, err = client.Do(req)

	out = make(map[string]interface{})
	json.NewDecoder(res.Body).Decode(&out)
	res.Body.Close()
	if out["errorCode"] != nil {
		if out["errorCode"] == "ERROR_NO_SLOT_AVAILABLE" {
			golog.Warnf("AntiCaptcha slotu yok! %d saniye sonra tekrar denenecek...", CfgNoSlotWaitInterval)
			time.Sleep(time.Duration(CfgNoSlotWaitInterval) * time.Second)
		} else {
			return nil, errors.New(out["errorCode"].(string))
		}

	}

	return out, nil
}

// Method to create the task to process the recaptcha, returns the task_id
func createTaskRecaptcha(websiteURL, recaptchaKey string) (float64, error) {
	var body map[string]interface{}

	golog.Debug("ReCaptchaProxyless talebinde bulunuluyor...")
	body = map[string]interface{}{
		"clientKey": CfgApiKey,
		"task": map[string]interface{}{
			"type":          "NoCaptchaTask",
			"websiteURL":    websiteURL,
			"websiteKey":    recaptchaKey,
			"isInvisible":   true,
			"proxyType":     "http",
			"proxyAddress":  "68.183.76.57",
			"proxyPort":     "62342",
			"proxyLogin":    "hexaminer",
			"proxyPassword": "xpc4togcf6244426",
			"cookies":       "group=tr; _ga=GA1.2.625908298.1548968354; _gid=GA1.2.1143182185.1548968354; acptcookiev2=true; _gat=1",
			"userAgent":     "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:64.0) Gecko/20100101 Firefox/64.0",
		},
	}

	b, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	// Make the request
	u := baseURL.ResolveReference(&url.URL{Path: "/createTask"})

	resp, err := execRequest("POST", u.String(), b)
	if err != nil {
		return 0, err
	}

	// TODO treat api errors and handle them properly
	//map[errorId:2 errorCode:ERROR_NO_SLOT_AVAILABLE errorDescription:No idle workers are available at the moment. Please try a bit later or increase your maximum bid in menu Settings - API Setup in Anti-Captcha Customers Area.]
	return resp["taskId"].(float64), nil
}

// Method to check the result of a given task, returns the json returned from the api
func getTaskResult(taskID float64) (map[string]interface{}, error) {
	// Mount the data to be sent
	body := map[string]interface{}{
		"clientKey": CfgApiKey,
		"taskId":    taskID,
	}
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	// Make the request
	u := baseURL.ResolveReference(&url.URL{Path: "/getTaskResult"})
	resp, err := execRequest("POST", u.String(), b)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SendRecaptcha Method to encapsulate the processing of the recaptcha
// Given a url and a key, it sends to the api and waits until
// the processing is complete to return the evaluated key
func SendRecaptcha(websiteURL, recaptchaKey string) (string, error) {
	// Create the task on anti-captcha api and get the task_id
	taskID, err := createTaskRecaptcha(websiteURL, recaptchaKey)
	if err != nil {
		return "", err
	}

	// Check if the result is ready, if not loop until it is
	response, err := getTaskResult(taskID)
	if err != nil {
		return "", err
	}
	for {
		if response["status"] == "processing" {
			time.Sleep(time.Duration(CfgCheckInterval) * time.Second)
			response, err = getTaskResult(taskID)
			if err != nil {
				return "", err
			}
		} else {
			golog.Debug("Captcha hazır")
			break
		}
	}
	return response["solution"].(map[string]interface{})["gRecaptchaResponse"].(string), nil
}

// Method to create the task to process the image captcha, returns the task_id
func createTaskImage(base64img string) (float64, error) {
	// Mount the data to be sent
	body := map[string]interface{}{
		"clientKey": CfgApiKey,
		"task": map[string]interface{}{
			"type": "ImageToTextTask",
			"body": base64img,
		},
	}

	b, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	// Make the request
	u := baseURL.ResolveReference(&url.URL{Path: "/createTask"})
	resp, err := execRequest("POST", u.String(), b)

	if err != nil {
		return 0, err
	}
	// TODO treat api errors and handle them properly
	return resp["taskId"].(float64), nil
}

// SendImage Method to encapsulate the processing of the image captcha
// Given a base64 string from the image, it sends to the api and waits until
// the processing is complete to return the evaluated key
func SendImage(imgString string) (string, error) {
	// Create the task on anti-captcha api and get the task_id
	golog.Debug("Captcha bekleniyor...")
	taskID, err := createTaskImage(imgString)
	if err != nil {
		return "", err
	}

	// Check if the result is ready, if not loop until it is
	response, err := getTaskResult(taskID)
	if err != nil {
		return "", err
	}
	for {
		if response["status"] == "processing" {
			time.Sleep(time.Duration(CfgCheckInterval) * time.Second)
			response, err = getTaskResult(taskID)
			if err != nil {
				return "", err
			}
		} else {
			golog.Debug("Captcha hazır")
			break
		}
	}
	return response["solution"].(map[string]interface{})["text"].(string), nil
}

func GetBalance() (float64, error) {
	golog.Debug("(AntiCaptcha) Bakiye sorgulanıyor...")
	body := map[string]interface{}{
		"clientKey": CfgApiKey,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	// Make the request
	u := baseURL.ResolveReference(&url.URL{Path: "/getBalance"})
	resp, err := execRequest("POST", u.String(), b)
	if err != nil {
		return 0, err
	}
	var balance = resp["balance"].(float64)
	golog.Debugf("(AntiCaptcha) Bakiye: $%f", balance)
	return balance, nil
}
