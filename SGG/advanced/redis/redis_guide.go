package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// 创建上下文
var ctx1 = context.Background()

/*
每次执行一次 Redis 命令（如 HGetAll），都会经历以下过程：

TCP 三次握手（消耗网络往返时间 RTT）。

身份验证/握手（Redis 密码验证）。

发送命令并等待结果。

TCP 四次挥手（断开连接）。
*/
// 创建单机 redis 客户端
func createClient1() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// 测试连接
	pong, err := client.Ping(ctx1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis连接成功：", pong)

	return client
}

/*
初始化 (Initialization)：
程序启动时，连接池会根据 MinIdleConns 预先创建几个连接，静静地躺在池子里，处于 Idle（空闲） 状态。

申请 (Borrow/Get)：
当你的代码调用 client.Get() 时，连接池会检查内部队列。如果有空闲连接，直接拨给你使用，将其标记为 Active（活跃）。

使用 (Execute)：
你在该连接上执行 Redis 指令。

归还 (Release/Put)：
这是最关键的一步。 执行完后，连接并不关闭，而是清除当前状态，重新回到池子的空闲队列中，等待下一个人。
*/
// 创建连接池的连接
func createClientWithPool() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		PoolSize:     10,              // 连接池大小
		MinIdleConns: 5,               // 最小空闲连接数
		MaxIdleConns: 10,              // 最大空闲连接数
		DialTimeout:  5 * time.Second, // 连接超时
		ReadTimeout:  3 * time.Second, // 读超时
		WriteTimeout: 3 * time.Second, // 写超时
		PoolTimeout:  4 * time.Second, // 获取连接超时
	})

	return client
}

// 字符串（String）操作
func stringOperations(client *redis.Client) {
	// 设置值
	err := client.Set(ctx1, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	// 设置带过期时间
	err = client.Set(ctx1, "key_with_ttl", "value", 10*time.Second).Err()

	// 获取值
	val, err := client.Get(ctx1, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key:", val)

	// 获取不存在的键
	val2, err := client.Get(ctx1, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2:", val2)
	}

	// 自增/自减
	client.Incr(ctx1, "counter")
	client.IncrBy(ctx1, "counter", 10)
	client.Decr(ctx1, "counter")

	// 批量操作
	client.MSet(ctx1, "key1", "value1", "key2", "value2")
	values, err := client.MGet(ctx1, "key1", "key2").Result()
	fmt.Println("values：", values)

	// 追加
	client.Append(ctx1, "key", " appended")
}

// 哈希（Hash）操作
func hashOperations(client *redis.Client) {
	// 设置哈希字段
	client.HSet(ctx1, "user:1", "name", "John")
	client.HSet(ctx1, "user:1", "age", "30")

	// 批量设置
	client.HSet(ctx1, "user:2", map[string]string{
		"name":  "Alice",
		"age":   "25",
		"email": "alice@example.com",
	})

	// 获取字段
	name, err := client.HGet(ctx1, "user:1", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name：", name)

	// 获取所有字段
	allFields, err := client.HGetAll(ctx1, "user:2").Result()
	for field, value := range allFields {
		fmt.Printf("%s: %s\n", field, value)
	}

	// 获取多个字段
	values, err := client.HMGet(ctx1, "user:2", "name", "email").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("values：", values)

	// 删除字段
	client.HDel(ctx1, "user:2", "email")

	// 字段自增
	client.HIncrBy(ctx1, "user:1", "age", 1)

	// 检查字段是否存在
	exists, err := client.HExists(ctx1, "user:1", "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("exists：", exists)

}

// 列表（List）操作
func listOperations(client *redis.Client) {
	// 从左侧插入
	client.LPush(ctx1, "mylist", "world")
	client.LPush(ctx1, "mylist", "hello")

	// 从右侧插入
	client.RPush(ctx1, "mylist", "!")

	// 获取列表长度
	length, err := client.LLen(ctx1, "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("length：", length)

	// 获取元素
	val, err := client.LIndex(ctx1, "mylist", 0).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("val：", val)

	// 获取范围
	items, err := client.LRange(ctx1, "mylist", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("items：", items)

	// 弹出元素
	poppl, err := client.LPop(ctx1, "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("poppl：", poppl)

	poppr, err := client.RPop(ctx1, "mylist").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("poppr：", poppr)

	// 修剪列表
	client.LTrim(ctx1, "mylist", 0, 2)
}

// 集合（Set）操作
func setOperations(client *redis.Client) {
	// 添加元素
	client.SAdd(ctx1, "myset", "a", "b", "c")

	// 获取所有元素
	members, err := client.SMembers(ctx1, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("members：", members)

	// 检查元素是否存在
	isMember, err := client.SIsMember(ctx1, "myset", "a").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("isMember：", isMember)

	// 获取集合大小
	size, err := client.SCard(ctx1, "myset").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("size：", size)

	// 移除元素
	client.SRem(ctx1, "myset", "a")

	// 集合运算
	client.SAdd(ctx1, "set1", "a", "b", "c")
	client.SAdd(ctx1, "set2", "c", "d", "e")

	// 交集
	inter, err := client.SInter(ctx1, "set1", "set2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("inter：", inter)

	// 并集
	union, err := client.SUnion(ctx1, "set1", "set2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("union：", union)

	// 差集
	diff, err := client.SDiff(ctx1, "set1", "set2").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("diff：", diff)
}

// 有序集合（Sorted Set）操作
func sortedSetOperations(client *redis.Client) {
	// 添加元素（带分数）
	client.ZAdd(ctx1, "leaderboard", redis.Z{
		Score:  100,
		Member: "player1",
	})

	// 批量添加
	client.ZAdd(ctx1, "leaderboard", []redis.Z{
		{Score: 200, Member: "player2"},
		{Score: 150, Member: "player3"},
	}...)

	// 获取排名（升序）
	rank, err := client.ZRank(ctx1, "leaderboard", "player1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("rank：", rank)

	// 获取排名（降序）
	revRank, err := client.ZRevRank(ctx1, "leaderboard", "player1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("revRank：", revRank)

	// 获取分数
	score, err := client.ZScore(ctx1, "leaderboard", "player1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("score：", score)

	// 获取范围（按分数）
	items, err := client.ZRangeByScore(ctx1, "leaderboard", &redis.ZRangeBy{
		Min: "100",
		Max: "200",
	}).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("items：", items)

	// 获取范围（按排名）
	items, err = client.ZRange(ctx1, "leaderboard", 0, 2).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("items：", items)

	// 增加分数
	client.ZIncrBy(ctx1, "leaderboard", 50, "player1")
}

// 管道（Pipeline）
func pipelineExample(client *redis.Client) {
	// 创建管道
	pipe := client.Pipeline()

	// 添加多个命令
	pipe.Set(ctx1, "key1", "value1", 0)
	pipe.Set(ctx1, "key2", "value2", 0)
	pipe.Get(ctx1, "key1")
	pipe.Get(ctx1, "key2")

	// 执行管道
	cmds, err := pipe.Exec(ctx1)
	if err != nil {
		panic(err)
	}

	// 处理结果
	for _, cmd := range cmds {
		fmt.Println(cmd.String())
	}
}

// 简化的管道使用
func pipelineSimple(client *redis.Client) {
	// 1. 声明两个指针变量，类型为 *redis.StringCmd。
	// 它们目前是 nil，相当于两个“空凭证”，用来准备接收后续管道返回的结果句柄。
	var get1, get2 *redis.StringCmd

	// 2. 调用 Pipelined 方法。它接收一个匿名函数（闭包），并自动处理管道的创建和 Exec() 的调用。
	// ctx1 用于控制整个网络请求的生命周期（如超时控制）。
	cmds, err := client.Pipelined(ctx1, func(pipe redis.Pipeliner) error {

		// 3. 在管道中排队一个 GET 命令。
		// 注意：此时并没有发生网络请求，pipe.Get 立即返回一个关联了 "key1" 的命令对象赋给 get1。
		get1 = pipe.Get(ctx1, "key1")

		// 4. 同理，在管道中排队第二个 GET 命令并赋值给 get2。
		get2 = pipe.Get(ctx1, "key2")

		// 5. 闭包返回 nil。此时 Pipelined 内部会自动触发管道的正式提交（即执行 Exec）。
		return nil
	})

	// 6. 检查整个管道在网络传输或执行过程中是否发生了错误（如 Redis 挂了或网络断了）。
	if err != nil {
		panic(err) // 如果管道执行失败，直接抛出异常
	}

	// 7. 从 get1 这个凭证中提取真正的执行结果。
	// 由于前面 Pipelined 已经执行完毕，此时 get1 内部已经填充了来自 Redis 的数据。
	val1, err := get1.Result()
	if err != nil {
		// 8. 这里的错误通常是业务级的，比如 "key1" 在 Redis 里根本不存在（redis.Nil）。
		panic(err)
	}
	// 9. 打印获取到的第一个值。
	fmt.Println(val1)

	// 10. 同理，提取并处理第二个命令的结果。
	val2, err := get2.Result()
	if err != nil {
		panic(err)
	}
	// 11. 打印获取到的第二个值。
	fmt.Println(val2)

	// 也可以通过cmds进行调用
	// cmds 就是那 2 个 Get 命令的结果集合
	for _, cmd := range cmds {
		// 将接口断言为具体的 StringCmd 来取值
		fmt.Println(cmd.(*redis.StringCmd).Val())
	}
}

// 事务（Transaction）
func transactionExample(client *redis.Client) {
	// 1. 调用 Watch 方法监控 "counter" 这个键。
	// 如果在执行过程中，有其他客户端修改了 "counter"，整个事务逻辑会报错并退出。
	err := client.Watch(ctx1, func(tx *redis.Tx) error {

		// 2. 在事务开始前，获取当前的最新的值。
		// 注意：这里的 tx 是一个特殊的事务连接，它会保持对键的监控。
		n, err := tx.Get(ctx1, "counter").Int()

		// 3. 错误处理：如果是网络错误则返回；如果键不存在（redis.Nil），我们可以视为初始值为 0。
		if err != nil && err != redis.Nil {
			return err
		}

		// 4. 进入真正的事务提交阶段（相当于 Redis 的 MULTI/EXEC）。
		// TxPipelined 会将内部的命令打包，并在最后检查 Watch 的键是否被改动。
		_, err = tx.TxPipelined(ctx1, func(pipe redis.Pipeliner) error {

			// 5. 定义修改逻辑：将刚才读到的值 n 加 1，并写回 Redis。
			// 这里的逻辑是：基于我刚才看到的 n，我计划把它变成 n+1。
			pipe.Set(ctx1, "counter", n+1, 0)

			// 6. 返回 nil 告知 TxPipelined 可以提交命令了。
			return nil
		})

		// 7. 返回 TxPipelined 的执行结果。
		// 如果在这个 func 执行期间，"counter" 被人改了，这里会返回 redis.TxFailedErr。
		return err

		// 8. 最后的参数 "counter" 是告诉 Watch 到底要“盯死”哪个键。
	}, "counter") // Watch中counter是要监控的键，如果在事务执行过程中这个键被修改了，整个事务会失败

	// 9. 检查整个 Watch 流程是否报错。
	if err != nil {
		// 10. 在高并发 IM 系统中，如果是事务失败（redis.TxFailedErr），
		// 正常的做法通常是启动一个 for 循环进行“重试”，而不是直接 panic。
		panic(err)
	}
}

// 发布/订阅（Pub/Sub）
// PUBLISH channel1 "test_message"用于测试
func pubSubExample(client *redis.Client) {

	// 订阅
	pubsub := client.Subscribe(ctx1, "channel1")
	defer pubsub.Close()

	// 接收消息
	ch := pubsub.Channel()
	for msg := range ch {
		fmt.Printf("收到消息: %s from %s\n", msg.Payload, msg.Channel)
	}
}

// 发布消息
func publishMessage(client *redis.Client) {
	err := client.Publish(ctx1, "channel1", "hello world").Err()
	if err != nil {
		panic(err)
	}
}

// 键操作
func keyOperations(client *redis.Client) {
	// 检查键是否存在
	exists, err := client.Exists(ctx, "key1", "key2").Result()
	fmt.Println("exists：", exists)

	// 删除键
	deleted, err := client.Del(ctx, "key1", "key2").Result()
	fmt.Println("deleted：", deleted)

	// 设置过期时间
	client.Expire(ctx, "key", 10*time.Second)
	client.ExpireAt(ctx, "key", time.Now().Add(10*time.Second))

	// 获取剩余生存时间
	ttl, err := client.TTL(ctx, "key").Result()
	fmt.Println("ttl：", ttl)

	// 移除过期时间
	client.Persist(ctx, "key")

	// 重命名键
	client.Rename(ctx, "oldkey", "newkey")

	// 扫描键
	var cursor uint64
	var keys []string
	for {
		var scannedKeys []string
		scannedKeys, cursor, err = client.Scan(ctx, cursor, "pattern:*", 10).Result()
		if err != nil {
			panic(err)
		}
		keys = append(keys, scannedKeys...)
		if cursor == 0 {
			break
		}
	}
}
func main() {
	pool := createClientWithPool()
	/*	stringOperations(pool)
		hashOperations(pool)
		listOperations(pool)
		setOperations(pool)*/
	// 开启协程去订阅，不阻塞主流程
	go pubSubExample(pool)

	// 给一点点时间让订阅连接建立成功
	time.Sleep(100 * time.Millisecond)

	// 现在发布消息，订阅者就能收到了
	publishMessage(pool)
}
