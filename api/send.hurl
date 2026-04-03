POST http://localhost:8080/api/v1/messages/send
Content-Type: application/json

{
  	"FromUserID": 1,
	"ToUserID": 2,
	"Content": "hello dude."
}

HTTP 201