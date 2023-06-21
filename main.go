package captcha_solver_go

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type TypeData struct {
	RecaptchaResponse string `json:"RecaptchaResponse"`
	Result            string `json:"result"`
	RecaptchaId       int    `json:"RecaptchaId"`
}
type TypeResponse struct {
	Message     string   `json:"message"`
	Ok          bool     `json:"ok"`
	Status_code int      `json:"status_code"`
	Data        TypeData `json:"data"`
	AccessToken string   `json:"access_token"`
}
type TypeAccessToken struct {
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	AccessToken string `json:"access_token"`
	StatusCode  int    `json:"status_code"`
	Message     string `json:"message"`
	Ok          bool   `json:"ok"`
}

type AuthClient struct {
	ClientID     string
	ClientSecret string
	Email        string
	Password     string
}


var urlTextCaptcha = "https://app.metabypass.tech/CaptchaSolver/api/v1/services/captchaSolver"
var urlRecaptcha = "https://app.metabypass.tech/CaptchaSolver/api/v1/services/bypassReCaptcha"
var urlGetCaptchaResult = "https://app.metabypass.tech/CaptchaSolver/api/v1/services/getCaptchaResult"


func NewAuthClient(clientID, clientSecret, email, password string) *AuthClient {
	return &AuthClient{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Email:        email,
		Password:     password,
	}
}

func (c *AuthClient) request(payload string, method string, url string, resend401 bool) (TypeData, int, string) {
	accessToken, successful := c.getAccessToken(false)
	if successful {
		newPayload := strings.NewReader(payload)

		client := &http.Client{}
		req, err := http.NewRequest(method, url, newPayload)

		if err != nil {
			fmt.Println(err)
		}
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Authorization", "Bearer "+accessToken)

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
		}

		var data TypeResponse // Change `interface{}` to your expected data structure
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err)
		}

		if data.Status_code == 401 {
			c.getAccessToken(true)
			if resend401 {
				return c.request(payload, method, url, false)
			} else {
				return data.Data, data.Status_code, data.Message
			}
		}

		return data.Data, data.Status_code, data.Message
	} else {
		return TypeData{}, 401, "authentication failed, please check your application info"
	}
}

func (c *AuthClient) getAccessToken(refresh bool) (string, bool) {
	if refresh {
		newAccessToken, successful := c.serviceGetAccessToken()
		saveAccessTokenToFile(newAccessToken)
		return newAccessToken, successful
	}
	accessToken, err := readAccessTokenFromFile("access_token.txt")
	if err == nil {
		if len(accessToken) > 5 {
			return accessToken, true
		} else {
			newAccessToken, successful := c.serviceGetAccessToken()
			saveAccessTokenToFile(newAccessToken)
			return newAccessToken, successful
		}
	} else {
		newAccessToken, successful := c.serviceGetAccessToken()
		saveAccessTokenToFile(newAccessToken)
		return newAccessToken, successful
	}
}


func (c *AuthClient) serviceGetAccessToken() (string, bool) {
	url := "https://app.metabypass.tech/CaptchaSolver/oauth/token"
	method := "POST"
	payload := fmt.Sprintf(`{"grant_type":"password","client_id": "%s" ,"client_secret": "%s","username": "%s","password": "%s"}`, c.ClientID, c.ClientSecret, c.Email, c.Password)
	newPayload := strings.NewReader(payload)

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, newPayload)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, _ := client.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	var data TypeAccessToken // Change `interface{}` to your expected data structure
	err := json.Unmarshal(body, &data)
	fmt.Println(err)

	if len(data.AccessToken) > 5 {
		return data.AccessToken, true
	}
	if data.StatusCode == 401 {
		fmt.Println(data.Message)
	}
	return "", false
}

func saveAccessTokenToFile(accessToken string) error {
	filePath := "./access_token.txt"

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(accessToken)
	if err != nil {
		return err
	}

	return nil
}

func readAccessTokenFromFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	accessToken := string(data)
	return accessToken, nil
}

func imageToBase64(imagePath string) string {
	// Read the image file
	imageData, err := ioutil.ReadFile(imagePath)
	if err != nil {
		log.Fatal("Error reading image file:", err)
	}

	// Convert image data to base64
	base64Data := base64.StdEncoding.EncodeToString(imageData)

	// Print the base64 encoded image data
	return base64Data
}

func (c *AuthClient) textCaptcha(imagePath string) (string, int, string) {
	payloadString := fmt.Sprintf(`{
	"image":"%s"
}`, imageToBase64(imagePath))
	data, code, message := c.request(payloadString, "POST", urlTextCaptcha, true)
	return data.Result, code, message
}

func (c *AuthClient) recaptchaV3(sitekey string, siteUrl string) (string, int, string) {
	payloadString := fmt.Sprintf(`{
	"sitekey":"%s",
	  "version": 3 ,
	  "url": "%s"
}`, sitekey, siteUrl)
	data, code, message := c.request(payloadString, "POST", urlRecaptcha, true)
	return data.RecaptchaResponse, code, message
}

func (c *AuthClient) recaptchaV2(sitekey string, siteUrl string) (string, int, string) {
	payloadString := fmt.Sprintf(`{
	"sitekey":"%s",
	  "version": 2 ,
	  "url": "%s"
}`, sitekey, siteUrl)
	data, code, message := c.request(payloadString, "POST", urlRecaptcha, true)
	if code == 200 {
		fmt.Println("Registration request successfully. Captcha id: " + strconv.Itoa(data.RecaptchaId))
		index := 1
		for index <= 12 {
			time.Sleep(10 * time.Second)
			token, code, message := c.getCaptchaResult(data.RecaptchaId)
			if code == 200 {
				return token, code, message
			} else if code == 201 {
				fmt.Println(strconv.Itoa(index) + "0s, captcha is not ready yet. please wait...")
			} else {
				return token, code, message
			}
			index++
		}
	} else {
		return "", code, message
	}

	return "", 500, "service failed"
}

func (c *AuthClient) getCaptchaResult(recaptchaId int) (string, int, string) {
	strRecaptchaId := strconv.Itoa(recaptchaId)
	payloadString := fmt.Sprintf(`{
	"recaptcha_id":"%s"
}`, strRecaptchaId)
	data, code, message := c.request(payloadString, "GET", urlGetCaptchaResult, true)
	return data.RecaptchaResponse, code, message
}

