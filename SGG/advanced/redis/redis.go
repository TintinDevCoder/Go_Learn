package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// go语言操作redis数据库，使用go-redis库
// go-redis是一个go语言的redis客户端库，支持redis的各种数据类型和命令，使用简单，性能高效
// 使用：在需要使用redis的go项目中，导入go-redis库，创建redis客户端，使用客户端的方法进行数据操作
// go-redis库的安装：在终端中运行以下命令安装go-redis库
// go get github.com/redis/go-redis/v9

// 创建一个上下文对象，用于在后续的redis操作中传递上下文信息
var ctx = context.Background()

// 创建redis客户端，连接redis服务器，返回redis客户端对象
func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // redis服务器地址和端口
		Password: "",               // redis服务器密码，如果没有设置密码则留空
		DB:       0,                // redis数据库索引，默认为0
	})

	// 通过client.Ping()检查是否成功连接到redis服务器
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err) // 如果连接失败，抛出异常
	}
	fmt.Println("Redis连接成功:", pong) // 输出连接成功的消息
	// 返回redis客户端对象，供后续的redis操作使用
	return client
}

// String的操作
func stringOperation(client *redis.Client) {
	// 第三个参数是过期时间，过期时间为0表示永不过期
	err1 := client.Set(ctx, "name", "dd", 0).Err()
	if err1 != nil {
		panic(err1)
	}
	val1, err2 := client.Get(ctx, "name").Result()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("name:", val1)

	// 设置过期时间
	err3 := client.Set(ctx, "age", "20", 1*time.Second).Err()
	if err3 != nil {
		panic(err3)
	}
	// 自增操作
	client.Incr(ctx, "age")
	client.Incr(ctx, "age")
	// 自减操作
	client.Decr(ctx, "age")

	val2, err4 := client.Get(ctx, "age").Result()
	if err4 != nil {
		panic(err4)
	}
	fmt.Println("age:", val2)

	// 因为 key "age" 的过期时间是一秒钟, 因此当一秒后, 此 key 会自动被删除了.
	time.Sleep(1 * time.Second)
	val3, err5 := client.Get(ctx, "age").Result()
	if err5 == redis.Nil {
		fmt.Println("age key 不存在了")
	} else if err5 != nil {
		fmt.Println("获取 age key 失败:", err5)
	} else {
		fmt.Println("age:", val3)
	}
}

// list的操作
// list可以在头部或尾部添加元素，也可以在头部或尾部删除元素，还可以获取指定范围内的元素等
func listOperation(client *redis.Client) {
	//在名称为 fruit 的list尾添加一个值为value的元素
	client.RPush(ctx, "fruit", "apple")
	//在名称为 fruit 的list头添加一个值为value的 元素
	client.LPush(ctx, "fruit", "banana")
	length, err1 := client.LLen(ctx, "fruit").Result()
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("fruit list的长度:", length)

	// 获取名称为 fruit 的list中索引为0的元素
	val1, err2 := client.LIndex(ctx, "fruit", 0).Result()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("fruit list中索引为0的元素:", val1)

	// 删除名称为 fruit 的list中的首元素
	val2, err3 := client.LPop(ctx, "fruit").Result()
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("fruit: ", val2)

	// 删除名称为 fruit 的list中的尾元素
	val3, err4 := client.RPop(ctx, "fruit").Result()
	if err4 != nil {
		panic(err4)
	}
	fmt.Println("fruit: ", val3)
}

// set的操作
// set是一个无序集合，集合中的元素是唯一的，可以添加元素、删除元素、判断元素是否存在等
func setOperation(client *redis.Client) {
	// 向 mset 中添加元素
	client.SAdd(ctx, "mset", "a", "b", "c")
	// 获取 mset 中的所有元素
	val1, err1 := client.SMembers(ctx, "mset").Result()
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("mset中的元素:", val1)

	// 判断 mset 中是否存在元素 a
	exists, err2 := client.SIsMember(ctx, "mset", "a").Result()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("mset中是否存在元素a:", exists)

	// 求交集, 即既在黑名单中, 又在白名单中的元素
	client.SAdd(ctx, "wset", "b", "c")
	names, err := client.SInter(ctx, "mset", "wset").Result()
	if err != nil {
		panic(err)
	}
	// 获取到的元素是 "the Elder"
	fmt.Println("mset and wset inter result: ", names)
}

// hash的操作
// hash是一个键值对集合，类似于map，可以添加字段、删除字段、获取字段值等
func hashOperation(client *redis.Client) {
	// 向 user:1 中添加字段 name 和 age
	val, err := client.HSet(ctx, "user:1", map[string]string{"name": "test", "age": "20"}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("向 user:1 中添加字段 name 和 age 的结果:", val)

	// 获取 user:1 中字段 name 的值
	val1, err1 := client.HGet(ctx, "user:1", "name").Result()
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("user:1 中字段 name 的值:", val1)

	// 获取 user:1 中所有字段和值
	val2, err2 := client.HGetAll(ctx, "user:1").Result()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("user:1 中所有字段和值:", val2)

	// 删除 user:1 中字段 age
	client.HDel(ctx, "user:1", "age")

	// 获取 user:1 中字段 age 的值
	val3, err3 := client.HGet(ctx, "user:1", "age").Result()
	if err3 == redis.Nil {
		fmt.Println("user:1 中字段 age 不存在了")
	} else if err3 != nil {
		fmt.Println("获取 user:1 中字段 age 失败:", err3)
	} else {
		fmt.Println("user:1 中字段 age 的值:", val3)
	}

	// 获取名为 user:1 的hash中所有字段的数量
	length, err4 := client.HLen(ctx, "user:1").Result()
	if err4 != nil {
		panic(err4)
	}
	fmt.Println("user:1 中所有字段的数量:", length)
}

// 连接池的使用
// go-redis库默认使用连接池来管理与redis服务器的连接，连接池可以提高性能和资源利用率，避免频繁创建和关闭连接带来的开销
// 连接池的配置可以通过redis.Options结构体中的相关字段进行设置，例如MaxIdle、MaxActive、IdleTimeout等
func redisPool() {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10, // 连接池的最大连接数
		MinIdleConns: 5,  // 连接池的最小空闲连接数
	})
	defer client.Close() // 在函数结束时关闭redis客户端，释放资源
	wg := sync.WaitGroup{}
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			for j := 0; j < 100; j++ {
				client.Set(ctx, fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
				client.Get(ctx, fmt.Sprintf("name%d", j)).Result()
			}

			fmt.Printf("PoolStats, TotalConns: %d\n", client.PoolStats().TotalConns)
		}()
	}

	wg.Wait()
}
func main() {
	// 创建redis客户端，连接redis服务器，返回redis客户端对象
	client := createClient()
	// String的操作
	fmt.Println("String的操作：")
	stringOperation(client)

	// list的操作
	fmt.Println("list的操作：")
	listOperation(client)

	// set的操作
	fmt.Println("set的操作：")
	setOperation(client)

	// hash的操作
	fmt.Println("hash的操作：")
	hashOperation(client)

}
