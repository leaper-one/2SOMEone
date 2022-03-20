package util

import (
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

func SendMail(email string) (string, error) {
	// 生成6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	// TODO: 储存验证码
	now := time.Now()
	t := fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	html := fmt.Sprintf(`<div>
		<div>
			尊敬的%s，您好！
		</div>
		<div style="padding: 8px 40px 8px 50px;">
			<p>您于 %s 提交的邮箱验证，本次验证码为 %s，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
		</div>
		<div>
			<p>此邮箱为 LEAPERone 系统邮箱，请勿回复。</p>
		</div>	
	</div>`, email, t, vcode)

	m := gomail.NewMessage()
	m.SetAddressHeader("From", "longguancheng@zohomail.com", "LEAPERone 验证系统")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "[我的验证码]邮箱验证") //设置邮件主题
	m.SetBody("text/html", html)          //设置邮件正文
	// 第一个参数是host 第三个参数是发送邮箱，第四个参数 是邮箱密码
	d := gomail.NewDialer("smtp.zoho.com", 465, "longguancheng@zohomail.com", "sexy0756")
	if err := d.DialAndSend(m); err != nil {
		fmt.Println("错误：", err)
		return "",err
	}
	return vcode,nil
}
