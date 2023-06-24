# Go-based CAPTCHA solver ([Metabypass](https://metabypass.tech))
Free demo (no credit card required) -> https://app.metabypass.tech/application


## Configuration

Get the credentials from the [Application](https://app.metabypass.tech/application) section of the MetaBypass website:

1. Go to [Application Section](https://app.metabypass.tech/application)
2. You can see credentials like the below image



![Uploading 239733451-4420f7ed-1588-412a-b0e8-2876d4ae1854.png…](https://github.com/metabypass/metabypass-python/assets/128980891/4420f7ed-1588-412a-b0e8-2876d4ae1854)


 ## Implementation:

  
To obtain the results for each type of captcha, do the following steps:
   1. Create a new folder in your device and open it by your IDE
       
   2. Write the following command in a terminal to create **go.mod** file
       

      ```go
      
      go mod init <WRITE_NAME_YOU_WANT>
      
      ```
  
   3. Then write the below command to get the package from this repository. It creates **go.sum**  in your directory.
       ``` go

      go get github.com/metabypass/captcha-solver-go
      
       ```

   4. Create **main.go** file, copy the following codes for each type of captcha, and change function inputs with your values.
       
       - **Text_Captcha**
     
         ```go
         package main
         import "fmt"
         import "strconv"
         import captcha_solver_go "github.com/metabypass/captcha-solver-go"
         
         func main() {
         	captcha_solver_go.NewAuthClient("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET", "YOUR_ACCOUNT_EMAIL", "YOUR_ACCOUNT_PASSWORD") // ****CHANGE HERE WITH YOUR VALUE*******
    
         	token, code, message := captcha_solver_go.TextCaptcha("YOUR_CAPTCHA_IMAGE_PATH") // ****CHANGE HERE WITH YOUR VALUE*******
         
         	fmt.Println("code: " + strconv.Itoa(code))
         	fmt.Println("message: " + message)
         	fmt.Println("token: " + token)
         	}
       
         ```
       
       - **Recaptcha V2**
    
         ``` go
         package main
         import "fmt"
         import "strconv"
         import captcha_solver_go "github.com/metabypass/captcha-solver-go"
         
         func main() {
         	captcha_solver_go.NewAuthClient("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET", "YOUR_ACCOUNT_EMAIL", "YOUR_ACCOUNT_PASSWORD") // ****CHANGE HERE WITH YOUR VALUE*******
    
          token, code, message := captcha_solver_go.RecaptchaV2("YOUR_SITE_KEY","YOUR_SITE_URL") // ****CHANGE HERE WITH YOUR VALUE*******
         
         	fmt.Println("code: " + strconv.Itoa(code))
         	fmt.Println("message: " + message)
         	fmt.Println("token: " + token)
         	}

      
         ```
       
        - **Recaptcha V3**
         
           ```go
    
          package main
          import "fmt"
          import "strconv"
          import captcha_solver_go "github.com/metabypass/captcha-solver-go"
          
          func main() {
          	captcha_solver_go.NewAuthClient("YOUR_CLIENT_ID", "YOUR_CLIENT_SECRET", "YOUR_ACCOUNT_EMAIL", "YOUR_ACCOUNT_PASSWORD") // ****CHANGE HERE WITH YOUR VALUE*******
     
     	     token, code, message := captcha_solver_go.RecaptchaV3("YOUR_SITE_KEY","YOUR_SITE_URL") // ****CHANGE HERE WITH YOUR VALUE*******
          
          	fmt.Println("code: " + strconv.Itoa(code))
          	fmt.Println("message: " + message)
          	fmt.Println("token: " + token)
          	}
           
           ```

   4. Write this command in your terminal to get the result:
       ``` go
       go run main.go
       ```


