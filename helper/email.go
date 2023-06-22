package helper

import (
	"log"
	"net/smtp"
)

func Contactmail(mg, sub, mail string) {
	auth := smtp.PlainAuth(
		"",
		"pukarlamichhane567@gmail.com",
		"fypslllzmikwosok",
		"smtp.gmail.com",
	)
	body := "\nEmail:" + mail + "\nProblem:" + mg
	msg := "Subject:" + sub + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"pukarlamichhane567@gmail.com",
		[]string{"pukarlamichhane767@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		log.Println(err)
	}

}

func Ordermail(name, address, items, quntity, phone string) {
	auth := smtp.PlainAuth(
		"",
		"pukarlamichhane567@gmail.com",
		"fypslllzmikwosok",
		"smtp.gmail.com",
	)

	body := "\nPhone:" + phone + "\n:Quantity" + quntity + "\n:Itemname" + items + "\n:name:" + name + "\n:address" + address
	msg := "Subject:" + "About order" + "\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"pukarlamichhane567@gmail.com",
		[]string{"pukarlamichhane767@gmail.com"},
		[]byte(msg),
	)
	if err != nil {
		log.Println(err)
	}
}
