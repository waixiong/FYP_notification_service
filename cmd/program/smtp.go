package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

// StartTLS Email Example

func main() {

	// from := mail.Address{"Wai Xiong", "wx.chee@getitqec.com"}
	from := mail.Address{"Wai Xiong", "waixiong@siswa.um.edu.my"}
	subj := "mjml Testing"
	body := `<!doctype html><html xmlns="http://www.w3.org/1999/xhtml" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office"><head><title></title><!--[if !mso]><!-- --><meta http-equiv="X-UA-Compatible" content="IE=edge"><!--<![endif]--><meta http-equiv="Content-Type" content="text/html; charset=UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><style type="text/css">#outlook a { padding:0; }
	body { margin:0;padding:0;-webkit-text-size-adjust:100%;-ms-text-size-adjust:100%; }
	table, td { border-collapse:collapse;mso-table-lspace:0pt;mso-table-rspace:0pt; }
	img { border:0;height:auto;line-height:100%; outline:none;text-decoration:none;-ms-interpolation-mode:bicubic; }
	p { display:block;margin:13px 0; }</style><!--[if mso]>
  <xml>
  <o:OfficeDocumentSettings>
	<o:AllowPNG/>
	<o:PixelsPerInch>96</o:PixelsPerInch>
  </o:OfficeDocumentSettings>
  </xml>
  <![endif]--><!--[if lte mso 11]>
  <style type="text/css">
	.mj-outlook-group-fix { width:100% !important; }
  </style>
  <![endif]--><style type="text/css">@media only screen and (min-width:480px) {
  .mj-column-per-100 { width:100% !important; max-width: 100%; }
}</style><style type="text/css">@media only screen and (max-width:480px) {
table.mj-full-width-mobile { width: 100% !important; }
td.mj-full-width-mobile { width: auto !important; }
}</style></head><body><div><!--[if mso | IE]><table align="center" border="0" cellpadding="0" cellspacing="0" class="" style="width:600px;" width="600" ><tr><td style="line-height:0px;font-size:0px;mso-line-height-rule:exactly;"><![endif]--><div style="margin:0px auto;max-width:600px;"><table align="center" border="0" cellpadding="0" cellspacing="0" role="presentation" style="width:100%;"><tbody><tr><td style="direction:ltr;font-size:0px;padding:20px 0;text-align:center;"><!--[if mso | IE]><table role="presentation" border="0" cellpadding="0" cellspacing="0"><tr><td class="" style="vertical-align:top;width:600px;" ><![endif]--><div class="mj-column-per-100 mj-outlook-group-fix" style="font-size:0px;text-align:left;direction:ltr;display:inline-block;vertical-align:top;width:100%;"><table border="0" cellpadding="0" cellspacing="0" role="presentation" style="vertical-align:top;" width="100%"><!-- mj-image width="100px" src="/assets/img/logo-small.png"></mj-image --><tr><td align="center" style="font-size:0px;padding:10px 25px;word-break:break-word;"><table border="0" cellpadding="0" cellspacing="0" role="presentation" style="border-collapse:collapse;border-spacing:0px;"><tbody><tr><td style="width:360px;"><img height="auto" src="https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcT7I4jnjHhdCQZoRJZGKj6xKBOcAKAmx6XrJHADO0DcVgMD7w99&usqp=CAU" style="border:0;display:block;outline:none;text-decoration:none;height:auto;width:100%;font-size:13px;" width="360"></td></tr></tbody></table></td></tr><tr><td style="font-size:0px;padding:10px 25px;word-break:break-word;"><p style="border-top:solid 4px #F45E43;font-size:1;margin:0px auto;width:100%;"></p><!--[if mso | IE]><table align="center" border="0" cellpadding="0" cellspacing="0" style="border-top:solid 4px #F45E43;font-size:1;margin:0px auto;width:550px;" role="presentation" width="550px" ><tr><td style="height:0;line-height:0;"> &nbsp;
</td></tr></table><![endif]--></td></tr><tr><td align="left" style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:helvetica;font-size:24px;line-height:1;text-align:left;color:#F45E43;">Hello World</div></td></tr><tr><td style="font-size:0px;padding:10px 25px;word-break:break-word;"><p style="border-top:solid 4px #F45E43;font-size:1;margin:0px auto;width:100%;"></p><!--[if mso | IE]><table align="center" border="0" cellpadding="0" cellspacing="0" style="border-top:solid 4px #F45E43;font-size:1;margin:0px auto;width:550px;" role="presentation" width="550px" ><tr><td style="height:0;line-height:0;"> &nbsp;
</td></tr></table><![endif]--></td></tr><tr><td align="left" style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Arial;font-size:18px;line-height:1;text-align:left;color:#000000;"><center>The quick brown fox jumps over the lazy dog, the quick brown fox jumps over the lazy dog, and the quick brown fox jumps over the lazy dog</center></div></td></tr><tr><td style="font-size:0px;padding:10px 25px;word-break:break-word;"><p style="border-top:solid 2px #F45E43;font-size:1;margin:0px auto;width:100%;"></p><!--[if mso | IE]><table align="center" border="0" cellpadding="0" cellspacing="0" style="border-top:solid 2px #F45E43;font-size:1;margin:0px auto;width:550px;" role="presentation" width="550px" ><tr><td style="height:0;line-height:0;"> &nbsp;
</td></tr></table><![endif]--></td></tr><tr><td align="left" style="font-size:0px;padding:10px 25px;word-break:break-word;"><div style="font-family:Arial;font-size:12px;line-height:1;text-align:left;color:#000000;"><center>This is computer generated email, please do not reply.</center></div></td></tr></table></div><!--[if mso | IE]></td></tr></table><![endif]--></td></tr></tbody></table></div><!--[if mso | IE]></td></tr></table><![endif]--></div></body></html>`
	// h, _ := template.ParseGlob(raw)
	// buf := new(bytes.Buffer)
	// if err := h.Execute(buf, h); err != nil {
	// 	log.Panic(err)
	// }
	// body := buf.String()
	// fmt.Printf("BODY : \n%v\n", body)

	// Connect to the SMTP Server
	// servername := "smtp.zoho.com:587"
	servername := "smtp.gmail.com:587"

	host, _, _ := net.SplitHostPort(servername)

	// auth := smtp.PlainAuth("", from.Address, "98WyE736a2n9", host)
	auth := smtp.PlainAuth("", from.Address, "425cwx000", host)

	// TLS config
	// certPath, _ := filepath.Abs(commons.ENVVariable("CRT_PATH"))
	// b, _ := ioutil.ReadFile(certPath)
	// cp := x509.NewCertPool()
	// if !cp.AppendCertsFromPEM(b) {
	// 	log.Fatalf("fail to dial: %v", errors.New("credentials: failed to append certificates"))
	// }
	tlsconfig := &tls.Config{
		// InsecureSkipVerify: true,
		ServerName: host,
		// RootCAs:            cp,
	}

	c, err := smtp.Dial(servername)
	if err != nil {
		log.Panic(err)
	}

	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	to := []mail.Address{
		// mail.Address{"", "soh@getitqec.com"},
		mail.Address{"", "wx.chee@getitqec.com"},
		// mail.Address{"", "david.cch@getitqec.com"},
	}

	for _, rep := range to {
		fmt.Printf("Sending to %v...\n", rep.Address)
		// Setup headers
		headers := make(map[string]string)
		headers["From"] = from.String()
		headers["To"] = rep.String()
		headers["Subject"] = subj
		headers["MIME-version"] = "1.0"
		headers["Content-Type"] = "text/html"
		headers["charset"] = "UTF-8"

		// Setup message
		message := ""
		for k, v := range headers {
			message += fmt.Sprintf("%s: %s\r\n", k, v)
		}
		message += "\r\n" + body

		if err = c.Mail(from.Address); err != nil {
			log.Panic(err)
		}

		// To && From
		if err = c.Mail(from.Address); err != nil {
			log.Panic(err)
		}
		if err = c.Rcpt(rep.Address); err != nil {
			log.Panic(err)
		}

		// Data
		w, err := c.Data()
		if err != nil {
			log.Panic(err)
		}

		_, err = w.Write([]byte(message))
		if err != nil {
			log.Panic(err)
		}

		err = w.Close()
		if err != nil {
			log.Panic(err)
		}
	}

	c.Quit()

}
