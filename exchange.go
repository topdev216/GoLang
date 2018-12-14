package main

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"log"
	"net/http"
	"strings"
)

type CreateItem struct {
	XMLName            struct{}          `xml:"m:CreateItem"`
	MessageDisposition string            `xml:"MessageDisposition,attr"`
	SavedItemFolderId  SavedItemFolderId `xml:"m:SavedItemFolderId"`
	Items              Messages          `xml:"m:Items"`
}

type Messages struct {
	Message []Message `xml:"t:Message"`
}

type SavedItemFolderId struct {
	DistinguishedFolderId DistinguishedFolderId `xml:"t:DistinguishedFolderId"`
}

type DistinguishedFolderId struct {
	Id string `xml:"Id,attr"`
}

type Message struct {
	ItemClass    string     `xml:"t:ItemClass"`
	Subject      string     `xml:"t:Subject"`
	Body         Body       `xml:"t:Body"`
	Sender       OneMailbox `xml:"t:Sender"`
	ToRecipients XMailbox   `xml:"t:ToRecipients"`
}

type Body struct {
	BodyType string `xml:"BodyType,attr"`
	Body     []byte `xml:",chardata"`
}

type OneMailbox struct {
	Mailbox Mailbox `xml:"t:Mailbox"`
}

type XMailbox struct {
	Mailbox []Mailbox `xml:"t:Mailbox"`
}

type Mailbox struct {
	EmailAddress string `xml:"t:EmailAddress"`
}

func BuildEmail(from string, to []string, subject string, body []byte) ([]byte, error) {
	log.Println("Building Email data packet format...")
	c := new(CreateItem)
	c.MessageDisposition = "SendOnly" // SendAndSaveCopy/SendOnly/SaveOnly
	c.SavedItemFolderId.DistinguishedFolderId.Id = "sentitems"
	m := new(Message)
	m.ItemClass = "IPM.Note"
	m.Subject = subject
	m.Body.BodyType = "HTML" // HTML/TEXT
	m.Body.Body = body
	m.Sender.Mailbox.EmailAddress = from
	mb := make([]Mailbox, len(to))
	for i, addr := range to {
		mb[i].EmailAddress = addr
	}
	m.ToRecipients.Mailbox = append(m.ToRecipients.Mailbox, mb...)
	c.Items.Message = append(c.Items.Message, *m)
	return xml.MarshalIndent(c, "", "  ")
}

func Issue(ewsAddr string, username string, password string, body []byte) (*http.Response, error) {
	log.Println("Issuing a request...")
	b2 := []byte(`<?xml version="1.0" encoding="utf-8" ?><soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages" xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Header><t:RequestServerVersion Version="Exchange2007_SP1" /></soap:Header><soap:Body>`)
	b2 = append(b2, body...)
	b2 = append(b2, "\n  </soap:Body>\n</soap:Envelope>"...)
	req, err := http.NewRequest("POST", ewsAddr, bytes.NewReader(b2))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "text/xml")
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}} // 不验证SSL证书
	// client := new(http.Client) // 验证SSL证书
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error { return http.ErrUseLastResponse }
	return client.Do(req)
}

func SendMail(sendto []string, title string, body []byte) {
	log.Println("Sending an Email...")
	b, err := BuildEmail("sender@mail", sendto, title, body)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(b))

	resp, err := Issue("https://xxxxxxxxxxxxx.com.cn/EWS/Exchange.asmx",
		"domain\\username", "password", b)
	if err != nil {
		log.Println(err.Error())
	}

	if resp.StatusCode == 200 {
		log.Println(strings.Join(sendto, ", ") + " has/have been sent successfully")
	}
}

func main() {
	SendMail([]string{"receiver@mail"}, "Mail Title", []byte("<h1>Mail Content</h1>"))
}
