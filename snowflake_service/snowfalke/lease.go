// 使用etcd的lease机制维护一个在[0,1024)循环的int

package snowfalke

import (
	"context"
	"fmt"
	. "go.etcd.io/etcd/client/v3"
	"log"
	"strconv"
	"time"
)

type LeaseMaker struct {
	client *Client
	key    string
	lease  LeaseID
}

func NewLeaseMaker() *LeaseMaker {
	lm := new(LeaseMaker)
	config := Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	var err error
	// 建立连接
	lm.client, err = New(config)
	if err != nil {
		panic(fmt.Sprintf("conect to etcd faild, err:%v\n", err))
	}
	// 使用grant申请lease
	grant, err := lm.client.Grant(context.Background(), 2)
	lm.lease = grant.ID
	lm.key = strconv.FormatInt(int64(lm.lease), 10)
	//服务在线时续租
	ch, _ := lm.client.KeepAlive(context.Background(), lm.lease)
	go func() {
		for {
			if _, ok := <-ch; !ok {
				panic("err : etcd")
			}
		}
	}()
	return lm
}

func (lm LeaseMaker) getLease() (int64, error) {
	// 初始化为0
	resp, err := lm.client.Txn(context.Background()).
		If(Compare(Version(lm.key), "=", 0)).
		Then(OpPut(lm.key, "0", WithLease(lm.lease))).
		Commit()
	if err != nil {
		log.Fatal(err)
	}
	if resp.Succeeded {
		fmt.Println("初始化当前lease")
	}
	return getAndIncr(lm)
}

func getAndIncr(lm LeaseMaker) (int64, error) {
	kv := NewKV(lm.client)
	getResp, err := kv.Get(context.TODO(), lm.key)
	if err != nil {
		return -1, err
	}
	res, _ := strconv.ParseInt(string(getResp.Kvs[0].Value), 10, 64)
	// workerID上限设置为1024 在这里取模
	newValue := (res + 1) % 1024
	_, err = lm.client.Txn(context.Background()).
		If(Compare(Version(lm.key), "!=", 0)).
		Then(OpPut(lm.key, strconv.FormatInt(newValue, 10), WithLease(lm.lease))).
		Commit()
	if err != nil {
		fmt.Println(err.Error())
	}
	return res, nil
}
