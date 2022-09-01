package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var mode = "1.mx.zip"

func StoreTest(count int, acct string, pass string) {

	client := Client(acct, pass)
	file, _ := ioutil.ReadFile("/opt/ihome/mode/" + mode)

	fmt.Println("上传开始：start time", time.Now().Format("2006-01-02 15:04:05"))
	for i := 0; i < count; i++ {
		key := "large:mode:" + strconv.Itoa(i)
		_ = client.Set(key, file, 10*time.Minute)
		//println(set.String())
	}
	fmt.Println("上传结束：start time", time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println("总耗时：start time", time.Now().Format("2006-01-02 15:04:05"))
	var sampIndex = count - 1
	for i := 0; i < count; i++ {
		if i == sampIndex {
			fmt.Println("单个耗时：start one time", time.Now().Format("2006-01-02 15:04:05"))
		}
		key := "large:mode:" + strconv.Itoa(i)
		_ = client.Get(key)
		//println(get.String())
		if i == sampIndex {
			fmt.Println("单个耗时：end one time", time.Now().Format("2006-01-02 15:04:05"))
		}
	}
	fmt.Println("总耗时：start time", time.Now().Format("2006-01-02 15:04:05"))

	for i := 0; i < count; i++ {
		key := "large:mode:" + strconv.Itoa(i)
		client.Del(key)
	}
}

func Client(acct string, pass string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     acct,
		Password: pass, // no password set
		DB:       2,    // use default DB
	})

	return rdb
}

func main() {
	args := os.Args
	count := args[1]
	acct := args[2]
	pass := args[3]

	num, _ := strconv.Atoi(count)
	println("总执行次数：" + count)
	StoreTest(num, acct, pass)
}
