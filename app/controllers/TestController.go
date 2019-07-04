package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nsqio/go-nsq"
	"net/http"
	"sync"
)

const NSQ_HOST = "172.16.19.146"

type NSQHandler struct {

}

func (this *NSQHandler) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}


func Test(context *gin.Context) {
	results := gin.H{
		"test": "test-key",
	}

	context.JSON(http.StatusOK, results)
}

func testNSQ() {
	waiter := sync.WaitGroup{}
	waiter.Add(1)

	go func() {
		defer waiter.Done()
		config := nsq.NewConfig()
		config.MaxInFlight = 9

		//建立多个连接
		for i := 0; i < 10; i++ {
			consumer, err := nsq.NewConsumer("test", "struggle", config)
			if nil != err {
				fmt.Println("err", err)
				return
			}
			consumer.AddHandler(&NSQHandler{})
			err = consumer.ConnectToNSQD(NSQ_HOST + ":4151")
			if nil != err {
				fmt.Println("err", err)
				return
			}
		}
		select {}

	}()

	waiter.Wait()
}