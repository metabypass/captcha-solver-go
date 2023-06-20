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

  - Download 'main.go'
  - To obtain the results for each type of captcha, **uncomment** the following codes in the 'main' function:
    - **Text_Captcha**
      
      ```go
      //token, code, message := textCaptcha("YOUR_CAPTCHA_IMAGE_PATH")
      ```
  
       The 'textCaptcha' function gets the captcha imagePath, then use the 'imageToBase64' function to get a bytes object for transferring images as text in API requests, and finally sends a request to get the result of your captcha image
   
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
    - **ReCaptcha V2**
      
      ```go
      //token, code, message := recaptchaV2("YOUR_SITE_KEY", "YOUR_SITE_URL")
      ```
      
    - **ReCaptcha V3**
      
      ```go
      //token, code, message := recaptchaV3("YOUR_SITE_KEY", "YOUR_SITE_URL")
      ```
    



