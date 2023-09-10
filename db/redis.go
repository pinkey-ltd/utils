package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
)

var _rdb *redis.Client

func RedisDB() *redis.Client {
	return _rdb
}

func RedisConn() (bool, error) {
	rd := redis.NewClient(&redis.Options{
		Addr:     "10.2x.1xx.1xx:6379", // url
		Password: "XixxxxU",
		DB:       0, // 0号数据库
	})
	res, err := rd.Ping().Result()
	if err != nil {
		log.Println("ping err :", err)
		return false, err
	}
	_rdb = rd
	log.Println(res)
	return true, nil
}

// string操作
func SetAndGet() {
	// set操作：第三个参数是过期时间，如果是0表示不会过期。
	defer func(_rdb *redis.Client) {
		err := _rdb.Close()
		if err != nil {
			log.Println("[Redis IO error: ]", err)
		}
	}(_rdb) // 记得关闭连接
	err := _rdb.Set("prüfung", "hallo kugou", 0).Err()

	if err != nil {
		log.Println("set err :", err)
		return
	}
	// get操作
	val, err := _rdb.Get("prüfung").Result()
	if err != nil {
		log.Println("get err :", err)
		return
	}
	log.Println("k1 ==", val)
}

// 声明一个全局的rdb变量
var rdb *redis.Client

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "10.23.1x.1x:6379",
		Password: "XixxxxxU", // no password set
		DB:       7,          // use default DB
	})

	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// initClientS 哨兵模式
func initClientS() (err error) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

// cluster
func initClientC() (err error) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func redisExample() {
	err := rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	val2, err := rdb.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name does not exist")
	} else if err != nil {
		fmt.Printf("get name failed, err:%v\n", err)
		return
	} else {
		fmt.Println("name", val2)
	}
}
