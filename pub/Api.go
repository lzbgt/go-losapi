// Api
package pub

import "time"

type WebApiStub struct {
	SystemStub
	MsgStub
	DataStub
	Sid string
	//	chMsg chan interface{}
}

//nosql数据库接口
type FVPair struct {
	Field string
	Value interface{}
}

type ScorePair struct {
	Score  int64
	Member string
}

type Dataer interface {
	Set(bid, key string, value interface{}) error
	Get(bid, key string) (interface{}, error)
	Expire(bid, key string, duration int64) (int64, error)
	Del(bid string, key ...string) (int64, error)
	Incr(bid, key string) (int64, error)
	Hlen(bid, key string) (int64, error)
	Hdel(bid, key string, field ...string) (int64, error)
	Hset(bid, key, field string, value interface{}) (int64, error)
	Hmclear(bid string, key ...string) (int64, error)
	Hmget(bid, key string, fields ...string) ([]interface{}, error)
	Hmset(bid, key string, values ...FVPair) error
	Hget(bid, key, field string) (interface{}, error)
	Hgetall(bid, key string) ([]FVPair, error)
	Hkeys(bid, key string) ([]string, error)
	Hincrby(bid, key, field string, delta int64) (int64, error)
	Lpush(bid, key string, value ...interface{}) (int64, error)
	Lpop(bid, key string) (interface{}, error)
	Rpush(bid, key string, value ...interface{}) (int64, error)
	Rpop(bid, key string) (interface{}, error)
	Lrange(bid, key string, start, stop int32) ([]interface{}, error)
	Lclear(bid, key string) (int64, error)
	Lmclear(bid string, keys ...string) (int64, error)
	Lindex(bid, k string, index int32) (interface{}, error)
	Lexpire(bid, k string, duration int64) (int64, error)
	Lexpireat(bid, k string, when int64) (int64, error)
	Lttl(bid, k string) (int64, error)
	Lpersist(bid, k string) (int64, error)
	Llen(bid, k string) (int64, error)
	Zadd(bid, key string, args ...ScorePair) (int64, error)
	Zcard(bid, key string) (int64, error)
	Zcount(bid, key string, mins, maxs int64) (int64, error)
	Zrem(bid, key string, members ...string) (int64, error)
	Zscore(bid, key, member string) (int64, error)
	Zrank(bid, key, member string) (int64, error)
	Zrange(bid, key string, mins, maxs int) ([]ScorePair, error)
	Zrangebyscore(bid, key string, mins, maxs int64, offset int, count int) ([]ScorePair, error)
	Sadd(bid, key string, args ...interface{}) (int64, error)
	Scard(bid, key string) (int64, error)
	Sclear(bid, key string) (int64, error)
	Sdiff(bid string, keys ...string) ([]string, error)
	Sinter(bid string, keys ...string) ([]string, error)
	Smclear(bid string, key ...string) (int64, error)
	Smembers(bid, key string) ([]interface{}, error)
	Srem(bid, key string, m interface{}) (int64, error)
	Sunion(bid string, keys ...string) ([]string, error)
}

type DataStub struct {
	Set           func(sid, bid, key string, value interface{}) error
	Get           func(sid, bid, key string) (interface{}, error)
	Expire        func(sid, bid, key string, dua int64) (int64, error)
	Del           func(sid, bid string, key ...string) (int64, error)
	Incr          func(sid, bid, key string) (int64, error)
	Hmclear       func(sid, bid string, key ...string) (int64, error)
	Hdel          func(sid, bid, key string, field ...string) (int64, error)
	Hlen          func(sid, bid, key string) (int64, error)
	Hset          func(sid, bid, key, field string, value interface{}) (int64, error)
	Hget          func(sid, bid, key, field string) (interface{}, error)
	Hmget         func(sid, bid, key string, fields ...string) ([]interface{}, error)
	Hmset         func(sid, bid, key string, values ...FVPair) error
	Hgetall       func(sid, bid, key string) ([]FVPair, error)
	Hkeys         func(sid, bid, key string) ([]string, error)
	Hincrby       func(sid, bid, key, field string, delta int64) (int64, error)
	Lpush         func(sid, bid, key string, value ...interface{}) (int64, error)
	Lpop          func(sid, bid, key string) (interface{}, error)
	Rpush         func(sid, bid, key string, value ...interface{}) (int64, error)
	Rpop          func(sid, bid, key string) (interface{}, error)
	Lrange        func(sid, bid, key string, start, stop int32) ([]interface{}, error)
	Lclear        func(sid, bid, key string) (int64, error)
	Lmclear       func(sid, bid string, keys ...string) (int64, error)
	Lindex        func(sid, bid, k string, index int32) (interface{}, error)
	Lexpire       func(sid, bid, k string, duration int64) (int64, error)
	Lexpireat     func(sid, bid, k string, when int64) (int64, error)
	Lttl          func(sid, bid, k string) (int64, error)
	Lpersist      func(sid, bid, k string) (int64, error)
	Llen          func(sid, bid, k string) (int64, error)
	Zadd          func(sid, bid, key string, args ...ScorePair) (int64, error)
	Zcard         func(sid, bid, key string) (int64, error)
	Zcount        func(sid, bid, key string, mins, maxs int64) (int64, error)
	Zrem          func(sid, bid, key string, members ...string) (int64, error)
	Zscore        func(sid, bid, key, member string) (int64, error)
	Zrank         func(sid, bid, key, member string) (int64, error)
	Zrange        func(sid, bid, key string, mins, maxs int) ([]ScorePair, error)
	Zrangebyscore func(sid, bid, key string, mins, maxs int64, offset int, count int) ([]ScorePair, error)
	Sadd          func(sid, bid, key string, args ...interface{}) (int64, error)
	Scard         func(sid, bid, key string) (int64, error)
	Sclear        func(sid, bid, key string) (int64, error)
	Sdiff         func(sid, bid string, keys ...string) ([]string, error)
	Sinter        func(sid, bid string, keys ...string) ([]string, error)
	Smclear       func(sid, bid string, key ...string) (int64, error)
	Smembers      func(sid, bid, key string) ([]interface{}, error)
	Srem          func(sid, bid, key string, m interface{}) (int64, error)
	Sunion        func(sid, bid string, keys ...string) ([]string, error)
	Copyblock     func(sid, bid string, from uint64) ([]byte, error) //备份数据块
	//	Copyuser  func(sid string) (*Man, error)                     //待删除
	//Replication func(sid, bid string) ([]byte, error) //备份
	/*
		HEXISTS
		HINCRBY
		HINCRBYFLOAT
		HSETNX
		HVALS
		HSCAN==
		BLPOP
		BRPOP
		BRPOPLPUSH
		LINSERT
		LPUSHX
		LREM
		LSET
		LTRIM
		RPOPLPUSH
		RPUSHX
	*/
}

type Message struct {
	Tm    time.Time //消息发生的时间
	From  string
	To    string
	AppID string      //空表示系统消息
	Msg   string      //表示命令，是由appid约定的，如果appid为空，则是系统消息
	Data  interface{} //应用自定义的数据格式
	Sign  string      //发送者签名
	//如果考虑消息的自净，应当加入一个消息的有效期，到期自动删除
}

type Msgs []*Message

//消息接口
type MsgStub struct {
	//SendMsg 消息发送
	//参数：发送者ID，from；接收者ID：to；appID:消息所属应用ID；msg:消息文本内容；
	//		data:消息对象内容；ppt:消息的校验信息
	//返回值：正常则为空，出错则返回错误信息
	SendMsg func(sid string, msg *Message) error
	//ReadMsg：消息接收
	//返回值：消息数组；返回不为空就需要重新读数据
	ReadMsg func(sid string) ([]*Message, error)
	PullMsg func(sid string, timeout int) (*Message, error)
}
type WebUser struct {
	Id      string     `json:"id"`
	Name    string     `json:"name"`
	Friends []*WebUser `json:"friends"`
	//Url     string     `json:"url"` //目前好象没价值了，暂时去掉
}

type Token struct {
	Id   uint64 //swarm last id
	Mac  uint64 //当前数据的校验值
	Time uint   //数据不变的情况下，每登录一次就刷新一次加1
	Sign []byte //签名，mac和time一起
}

type JSHostInfo struct {
	Id   string
	Url  string
	Bids []string
}

type JSwarmInfo struct {
	Token  *Token
	Nodes  []*JSHostInfo
	Father string
	Son    []string
}

type LoginReply struct {
	Sid   string
	User  *WebUser
	Swarm *JSwarmInfo
}

type CreditInfo struct {
	IDA   string
	IDB   string
	MaxA  int64
	MaxB  int64
	Cur   int64     //当前消耗掉的信用，以用户A为参考
	Last  time.Time //递增编号，可以设置为时间，确保有唯一ID
	Valid time.Time //有效期，需要的时候加上,相关的时间值可以和用户最后上限时间匹配的
}

type Credit struct {
	CreditInfo
	SignID   string //签名A，或B，或中间人，结算时用
	Signinfo []byte //自己的签名没必要保留，可以现签
	//Logs    string //互动日志 后期结算用,不一定用的上，用时再加
}
type BaseNodeInfo struct {
	OwnerId  string //owner的ID
	ShortUrl []byte //压缩过的URL
}

type InviteInfo struct {
	Credit
	HostId   string
	ShortUrl []byte //以上变量为邀请信息
}

type JSwarms map[string]*JSwarmInfo
type InvCodeInfo struct {
	Validity    uint
	FriendCount uint
	Money       uint
}
type SwapInfo struct {
	Id     string
	Url    []byte
	Orphan map[string]*Token
}

type SwapInfoReply struct {
	Son       map[string]*Token        //对方接受你作上线,所以无视你的通行证
	Father    map[string]*Token        //对方接收你作下线
	Introduce map[string]*BaseNodeInfo //对方推荐了一个新的节点
}
type MapRelation map[string]uint8
type SystemStub struct {
	Register           func(par ...string) (string, error)
	Login              func(p1, p2 string, tp ...string) (*LoginReply, error)
	Logout             func(sid, info string) error
	GetVar             func(sid, name string, args ...string) (interface{}, error)
	Act                func(sid, name string, args ...string) error
	GetGobVar          func(sid, name string, args ...string) ([]byte, error)
	Exit               func(string) //停止本服务
	Restart            func(string)
	Invite             func(sid, uid string) (string, error)
	Accept             func(sid, inv string) error
	AcceptReply        func(sid, invUId string, invCredit *Credit, replyInv *InviteInfo) error
	Removeuser         func(sid, bid string) error
	Test               func(sid, bid string) error
	UpdateData         func(sid string, data []byte) error
	KeepFriendHostInfo func(sid string, hostsinfo []*JSHostInfo) (*JSwarms, error)
	Setnodeip          func(sid, ip string) error
	Proxyget           func(sid, url string) (string, error)
	CreateInvCode      func(sid, bid string, validity, friendcount, money uint) (string, error)
	GetInvCodeInfo     func(sid, bid string, invCode string) (*InvCodeInfo, error)
	UpdateInvCode      func(sid, bid string, invCode string, validity, friendcount, money uint) error
	DeleteInvCode      func(sid, bid, invCode string) error
	SetInvTemplate     func(sid, bid, invcode, template string) error
	GetInvTemplate     func(sid, bid, invCode string) (string, error)
	GetAppDownloadKey  func(sid, bid, invcode, appID, appName string) (string, error)
	ListenToMe         func(sid, bid string, token *Token) (bool, error)
	SwapNodeInfo       func(sid string, info *SwapInfo) (*SwapInfoReply, error)
	ConfirmRelation    func(sid string, ret MapRelation) (MapRelation, error)
	SendMail           func(sid, to, subject, body, mailtype string) error
}

func (this *WebApiStub) Get(bid, key string) (interface{}, error) {
	return this.DataStub.Get(this.Sid, bid, key)
}

func (this *WebApiStub) Set(bid, key string, value interface{}) error {
	return this.DataStub.Set(this.Sid, bid, key, value)
}

func (this *WebApiStub) Expire(bid, key string, dua int64) (int64, error) {
	return this.DataStub.Expire(this.Sid, bid, key, dua)
}

func (this *WebApiStub) Del(bid string, key ...string) (int64, error) {
	return this.DataStub.Del(this.Sid, bid, key...)
}

func (this *WebApiStub) Incr(bid string, key string) (int64, error) {
	return this.DataStub.Incr(this.Sid, bid, key)
}

func (this *WebApiStub) Hmclear(bid string, key ...string) (int64, error) {
	return this.DataStub.Hmclear(this.Sid, bid, key...)
}

func (this *WebApiStub) Hlen(bid, key string) (int64, error) {
	return this.DataStub.Hlen(this.Sid, bid, key)
}

func (this *WebApiStub) Hdel(bid, key string, field ...string) (int64, error) {
	return this.DataStub.Hdel(this.Sid, bid, key, field...)
}

func (this *WebApiStub) Hget(bid, key string, field string) (interface{}, error) {
	return this.DataStub.Hget(this.Sid, bid, key, field)
}

func (this *WebApiStub) Hgetall(bid, key string) ([]FVPair, error) {
	return this.DataStub.Hgetall(this.Sid, bid, key)
}

func (this *WebApiStub) Hkeys(bid, key string) ([]string, error) {
	return this.DataStub.Hkeys(this.Sid, bid, key)
}

func (this *WebApiStub) Hset(bid, key, field string, value interface{}) (int64, error) {
	return this.DataStub.Hset(this.Sid, bid, key, field, value)
}

func (this *WebApiStub) Hmget(bid, key string, fields ...string) ([]interface{}, error) {
	return this.DataStub.Hmget(this.Sid, bid, key, fields...)
}

func (this *WebApiStub) Hmset(bid, key string, values ...FVPair) error {
	return this.DataStub.Hmset(this.Sid, bid, key, values...)
}

func (this *WebApiStub) Hincrby(bid, key, field string, delta int64) (int64, error) {
	return this.DataStub.Hincrby(this.Sid, bid, key, field, delta)
}

func (this *WebApiStub) Lpush(bid, key string, value ...interface{}) (int64, error) {
	return this.DataStub.Lpush(this.Sid, bid, key, value)
}

func (this *WebApiStub) Lpop(bid, key string) (interface{}, error) {
	return this.DataStub.Lpop(this.Sid, bid, key)
}

func (this *WebApiStub) Rpush(bid, key string, value ...interface{}) (int64, error) {
	return this.DataStub.Rpush(this.Sid, bid, key, value)
}

func (this *WebApiStub) Rpop(bid, key string) (interface{}, error) {
	return this.DataStub.Rpop(this.Sid, bid, key)
}

func (this *WebApiStub) Lrange(bid, key string, start, stop int32) ([]interface{}, error) {
	return this.DataStub.Lrange(this.Sid, bid, key, start, stop)
}

func (this *WebApiStub) Lclear(bid, key string) (int64, error) {
	return this.DataStub.Lclear(this.Sid, bid, key)
}

func (this *WebApiStub) Lmclear(bid string, keys ...string) (int64, error) {
	return this.DataStub.Lmclear(this.Sid, bid, keys...)
}

func (this *WebApiStub) Lindex(bid, key string, index int32) (interface{}, error) {
	return this.DataStub.Lindex(this.Sid, bid, key, index)
}

func (this *WebApiStub) Lexpire(bid, key string, duration int64) (int64, error) {
	return this.DataStub.Lexpire(this.Sid, bid, key, duration)
}

func (this *WebApiStub) Lexpireat(bid, key string, when int64) (int64, error) {
	return this.DataStub.Lexpireat(this.Sid, bid, key, when)
}

func (this *WebApiStub) Lttl(bid, key string) (int64, error) {
	return this.DataStub.Lttl(this.Sid, bid, key)
}

func (this *WebApiStub) Lpersist(bid, key string) (int64, error) {
	return this.DataStub.Lpersist(this.Sid, bid, key)
}

func (this *WebApiStub) Llen(bid, key string) (int64, error) {
	return this.DataStub.Llen(this.Sid, bid, key)
}

func (this *WebApiStub) Zadd(bid, key string, args ...ScorePair) (int64, error) {
	return this.DataStub.Zadd(this.Sid, bid, key, args...)
}

func (this *WebApiStub) Zcard(bid, key string) (int64, error) {
	return this.DataStub.Zcard(this.Sid, bid, key)
}

func (this *WebApiStub) Zcount(bid, key string, mins, maxs int64) (int64, error) {
	return this.DataStub.Zcount(this.Sid, bid, key, mins, maxs)
}

func (this *WebApiStub) Zrem(bid, key string, members ...string) (int64, error) {
	return this.DataStub.Zrem(this.Sid, bid, key, members...)
}

func (this *WebApiStub) Zscore(bid, key, member string) (int64, error) {
	return this.DataStub.Zscore(this.Sid, bid, key, member)
}

func (this *WebApiStub) Zrank(bid, key, member string) (int64, error) {
	return this.DataStub.Zrank(this.Sid, bid, key, member)
}

func (this *WebApiStub) Zrange(bid, key string, mins, maxs int) ([]ScorePair, error) {
	return this.DataStub.Zrange(this.Sid, bid, key, mins, maxs)
}

func (this *WebApiStub) Zrangebyscore(bid, key string, mins, maxs int64, offset int, count int) ([]ScorePair, error) {
	return this.DataStub.Zrangebyscore(this.Sid, bid, key, mins, maxs, offset, count)
}

func (this *WebApiStub) Sadd(bid, key string, args ...interface{}) (int64, error) {
	return this.DataStub.Sadd(this.Sid, bid, key, args...)
}

func (this *WebApiStub) Scard(bid, key string) (int64, error) {
	return this.DataStub.Sadd(this.Sid, bid, key)
}

func (this *WebApiStub) Sclear(bid, key string) (int64, error) {
	return this.DataStub.Sclear(this.Sid, bid, key)
}

func (stub *WebApiStub) Sdiff(bid string, keys ...string) ([]string, error) {
	return stub.DataStub.Sdiff(stub.Sid, bid, keys...)
}

func (stub *WebApiStub) Sinter(bid string, keys ...string) ([]string, error) {
	return stub.DataStub.Sinter(stub.Sid, bid, keys...)
}

func (stub *WebApiStub) Sunion(bid string, keys ...string) ([]string, error) {
	return stub.DataStub.Sunion(stub.Sid, bid, keys...)
}

func (this *WebApiStub) Smembers(bid, key string) ([]interface{}, error) {
	return this.DataStub.Smembers(this.Sid, bid, key)
}

func (this *WebApiStub) Srem(bid, key string, m interface{}) (int64, error) {
	return this.DataStub.Srem(this.Sid, bid, key, m)
}

func (this *WebApiStub) Smclear(bid string, key ...string) (int64, error) {
	return this.DataStub.Smclear(this.Sid, bid, key...)
}

func (this *WebApiStub) Setdata(bid string, value interface{}) (key string, err error) {
	if key, err = Obj2Key(value); err == nil {
		return key, this.DataStub.Set(this.Sid, bid, key, value)
	}
	return
}
