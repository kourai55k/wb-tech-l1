package main

import "fmt"

// У нас есть интерфейс Notifier, который отправляет уведомления
// структура EmailNotifier реализует этот интерфейс,
// а структура SMSNotifier - нет.
// Поэтому мы создадим адаптер для структуры SMSNotifier,
// который будет реализовывать интерфейс Notifier

// Notifier - интерфейс для отправки уведомлений
type Notifier interface {
	SendNotification(message string)
}

// EmailNotifier - структура для отправки уведомлений через email
type EmailNotifier struct {
	emailAddress string
}

// SendNotification - метод для отправки уведомления по электронной почте
func (e EmailNotifier) SendNotification(message string) {
	// Отправляем уведомление через email
	fmt.Printf("Sent email to %s: %s\n", e.emailAddress, message)
}

// SMSNotifier - структура для отправки уведомлений через SMS
type SMSNotifier struct {
	phoneNumber string
}

// SendSMS - метод для отправки уведомления через SMS
func (s SMSNotifier) SendSMS(message string) {
	fmt.Printf("Sent SMS to %s: %s\n", s.phoneNumber, message)
}

// SMSAdapter - адаптер для smsNotifier
type SMSAdapter struct {
	// передаем структуру для отправки SMS
	smsNotifier SMSNotifier
}

// SendNotification - реализация метода отправки сообщения для адаптера
func (sa SMSAdapter) SendNotification(message string) {
	// вызываем метод для отправки уведомления через SMS
	sa.smsNotifier.SendSMS(message)
}

func main() {
	// слайс, который хранит элементы, которые реализуют интерфейс Notifier
	var notifiers []Notifier

	// добавляем в слайс EmailNotifier
	emailNotifier := EmailNotifier{emailAddress: "user@example.com"}
	notifiers = append(notifiers, emailNotifier)

	// добавляем в слайс SMSNotifier через адаптер
	smsNotifier := SMSNotifier{phoneNumber: "+1234567890"}
	smsAdapter := SMSAdapter{smsNotifier: smsNotifier}
	notifiers = append(notifiers, smsAdapter)

	// отправляем уведомления
	for _, notifier := range notifiers {
		notifier.SendNotification("Hello! This is a notification.")
	}
}
