# Go-based CAPTCHA solver ([Metabypass](https://metabypass.tech))
Free demo (no credit card required) -> https://app.metabypass.tech/application


## Configuration

Get the following credentials from the [Application](https://app.metabypass.tech/application) section of the MetaBypass website:

``` go
var CLIENT_ID = "YOUR_CLIENT_ID"                 //****CHANGE HERE WITH YOUR VALUE*******
var CLIENT_SECRET = "YOUR_CLIENT_SECRET"        //****CHANGE HERE WITH YOUR VALUE*******
var EMAIL = "YOUR_ACCOUNT_EMAIL"               //****CHANGE HERE WITH YOUR VALUE*******
var PASSWORD = "YOUR_ACCOUNT_PASSWORD"        //****CHANGE HERE WITH YOUR VALUE*******
```

1. Go to [Application Section](https://app.metabypass.tech/application)
2. You can see credentials like the below image



![Uploading 239733451-4420f7ed-1588-412a-b0e8-2876d4ae1854.png…](https://github.com/metabypass/metabypass-python/assets/128980891/4420f7ed-1588-412a-b0e8-2876d4ae1854)


 ## Implementation:

  - **Text_Captcha**
    
    The below function gets a bytes object for transferring images as text in API requests

    ``` go
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
    ```

    The following function returns the result of your captcha image

    ``` go
    func textCaptcha(imagePath string) (string, int, string) {
    	payloadString := fmt.Sprintf(`{"image":"%s"}`, imageToBase64(imagePath))
    	data, code, message := request(payloadString, "POST", urlTextCaptcha, true)
    	return data.Result, code, message
    }
    ```






