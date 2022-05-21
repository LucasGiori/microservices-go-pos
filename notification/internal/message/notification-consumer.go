package message

import (
	"encoding/json"
	"fmt"
	"crypto/tls"
	"log"
	"net/smtp"
	"os"

	"microservices/notification/pkg/model"

	"gitlab.com/pos-alfa-microservices-go/core/broker/rabbitmq"
	coreLog "gitlab.com/pos-alfa-microservices-go/core/log"
	"github.com/joho/godotenv"
)

type SendNotificationProcessor interface {
	StartConsume() error
}

type NotificationConsumerProcessor struct {
	queue        string
	client       *rabbitmq.RabbitClient
}

func NewNotificationConsumerProcessor(queue string, client *rabbitmq.RabbitClient) SendNotificationProcessor {
	return &NotificationConsumerProcessor{
		queue:        queue,
		client:       client,
	}
}

func (n NotificationConsumerProcessor) StartConsume() error {
	godotenv.Load(".env")

	consumerNotification := rabbitmq.NewRabbitConsumer(n.client, "notification")
	return consumerNotification.Consume(n.queue, func(body []byte) error {
		bodyNotification := model.Notification{}
	
		if err := json.Unmarshal(body, &bodyNotification); err != nil {
			return fmt.Errorf("fail to unmarshal notification %w", err)
		}

		bodyEmail := "Ticket ID: " + bodyNotification.Id.String() + "\nTicket status: " + bodyNotification.Status +  "\nDescrição: " + bodyNotification.Description

		sendMail(
			os.Getenv("EMAIL"), 
			bodyNotification.Email, 
			bodyEmail,
		)

		coreLog.Logger.Infof("notification data processed")

		return nil
	})
}

func checkErr(err error) {
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

func sendMail(from string, to string, body string) {
	godotenv.Load(".env")

	servername := os.Getenv("SERVER")                
	host := os.Getenv("HOST")                   
	pass := os.Getenv("PASSWORD")

	auth := smtp.PlainAuth("Unialfa", from, pass, host) 
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	msg := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: Ticket UniAlfa\n\n" + body

	//create coneection server TCP
	conn, err := tls.Dial("tcp", servername, tlsConfig)
	checkErr(err)

	//return client SMTP
	c, err := smtp.NewClient(conn, host)
	checkErr(err)

	//authenticate
	err = c.Auth(auth)
	checkErr(err)

	//add sender
	err = c.Mail(from)
	checkErr(err)

	//add recipients
	err = c.Rcpt(to)
	checkErr(err)

	//prepare email
	w, err := c.Data()
	checkErr(err)

	_, err = w.Write([]byte(msg))
	checkErr(err)

	err = w.Close()
	checkErr(err)

	c.Quit()
}