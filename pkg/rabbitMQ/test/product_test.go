package test

import (
	"GameAdmin/pkg/rabbitMQ"
	"fmt"
	"sync"
	"testing"
)

func TestProduct(t *testing.T) {
	initrabbitmq()
	rund()
}

var oncePool sync.Once
var instanceRPool *rabbitMQ.RabbitMQPool

func initrabbitmq() *rabbitMQ.RabbitMQPool {
	oncePool.Do(func() {
		instanceRPool = rabbitMQ.NewProductPool()
		//err := instanceRPool.Connect("192.168.1.169", 5672, "admin", "admin")
		err := instanceRPool.ConnectVirtualHost("192.168.1.169", 5672, "temptest", "test123456", "/temptest1")
		if err != nil {
			fmt.Println(err)
		}
	})
	return instanceRPool
}

func rund() {

	var wg sync.WaitGroup

	//wg.Add(1)
	//go func() {
	//	fmt.Println("aaaaaaaaaaaaaaaaaaaaaa")
	//	defer wg.Done()
	//	runtime.SetMutexProfileFraction(1)  // 开启对锁调用的跟踪
	//	runtime.SetBlockProfileRate(1)      // 开启对阻塞操作的跟踪
	//	err:= http.ListenAndServe("0.0.0.0:8080", nil)
	//	fmt.Println(err)
	//}()

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			data := rabbitMQ.GetRabbitMqDataFormat("testChange31", rabbitMQ.EXCHANGE_TYPE_DIRECT, "testQueue31", "", []byte(fmt.Sprintf("这里是数据%d", num)))
			err := instanceRPool.Push(data)
			if err != nil {
				fmt.Println(err)
			}
		}(i)
	}

	wg.Wait()
}
