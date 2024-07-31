package main

import (
	"crypto/tls"
	"net/http"

	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func sendMail(name string, email string, message string) {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	messageContent := fmt.Sprintf("Message from:<br><br>%s<br><br>%s<br><br>%s", name, email, message)

	m := gomail.NewMessage()
	m.SetHeader("From", "freshsqueezecleaner@gmail.com")
	m.SetHeader("To", "alexdomzalski@gmail.com") // Can add more recipients with a comma

	// m.SetAddressHeader("Cc", "dan@example.com", "Dan") -- If needing to CC

	m.SetHeader("Subject", "Website Message")
	m.SetBody("text/html", messageContent)

	// m.Attach("") -- To attach a file

	// Configure SMTP
	d := gomail.NewDialer("smtp.gmail.com", 587, "freshsqueezecleaner@gmail.com", os.Getenv("EMAIL_PASSKEY"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func checkForm(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	message := c.PostForm("message")

	sendMail(name, email, message)

	response := "<div><h1 class='text-[rgb(125,112,164)] text-3xl md:text-4xl text-center py-4 pt-7 border-b border-[rgba(166,154,206,0.3)]'>Message Received!</h1><h2 class='text-[rgb(125,112,164)] text-3xl md:text-3xl text-center py-4 pt-7 border-b border-[rgba(166,154,206,0.3)]'>I'll be in touch with you soon!</h2></div>"
	c.Header("Content-Type", "text/html")
	c.String(http.StatusOK, response)
}

func renderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func main() {
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", renderIndex)
	router.POST("/sendMail", checkForm)
	router.Run(":3000")
}
