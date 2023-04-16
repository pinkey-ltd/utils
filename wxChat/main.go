package wxChat

import (
	"log"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

const token = "se4ie1gjd39sj1rjd"

type query struct {
	Signature    string `json:"signature" xml:"signature" binding:"signature"`
	MsgSignature string `json:"msg_signature" xml:"msg_signature" binding:"msg_signature"`
	Echostr      string `json:"echostr" xml:"echostr" binding:"echostr"`
	EncryptType  string `json:"encrypt_type" xml:"encrypt_type" binding:"encrypt_type"`
	Timestamp    string `json:"timestamp" xml:"timestamp" binding:"timestamp"`
	Nonce        string `json:"nonce" xml:"nonce" binding:"nonce"`
}

type tokenLine []string

func (a tokenLine) Len() int           { return len(a) }
func (a tokenLine) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a tokenLine) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	r := gin.Default()

	r.GET("/auth", func(c *gin.Context) {
		r.LoadHTMLGlob("templates/*")
		q := query{}
		// if errA := c.ShouldBind(&q); errA == nil {
		q.Echostr = c.Query("echostr")
		q.Timestamp = c.Query("timestamp")
		q.Signature = c.Query("signature")
		q.Nonce = c.Query("nonce")
		// 排序
		line := tokenLine{q.Timestamp, token, q.Nonce}
		sort.Strings(line)
		sortLine := strings.Join(line, "")
		// debug
		log.Printf("q.Signature:")
		log.Println(q.Signature)
		log.Printf("sortLine:")
		log.Println(sortLine)
		log.Printf("SHA1:")
		log.Println(SHA1(sortLine))
		if SHA1(sortLine) == q.Signature {
			c.HTML(200, "index.tmpl", gin.H{
				"title": q.Echostr,
			})
		} else {
			c.JSON(403, gin.H{
				"error": "auth",
			})
		}
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"echo": "OK",
		})
	})
	// https
	err := r.RunTLS(":443", "cert/ca.pem", "cert/ca.key")
	if err != nil {
		log.Println(err)
	}
}
