package redis

import (
	"context"
	r "github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

// RedisConnection is a connection to a redis database.
type RedisConnection struct {
	client *r.Client
	ctx    context.Context
}

var _ Connection = (*RedisConnection)(nil)

// NewRedisConnection creates a new redis connection.
func NewRedisConnection() Connection {
	return &RedisConnection{
		ctx: context.Background(),
	}
}

// Connect connects to the redis database.
func (c *RedisConnection) Connect(config map[string]interface{}) {
	c.client = r.NewClient(&r.Options{
		Addr:     config["host"].(string) + ":" + strconv.Itoa(config["port"].(int)),
		Password: config["password"].(string),
		DB:       config["database"].(int),
	})
}

// Echo returns the given string.
func (c *RedisConnection) Echo(message string) (string, error) {
	return c.client.Echo(c.ctx, message).Result()
}

// Ping returns the redis ping.
func (c *RedisConnection) Ping() (string, error) {
	return c.client.Ping(c.ctx).Result()
}

// Del deletes the given keys.
//	c.Del("key1", "key2")
//	c.Del("key3")
func (c *RedisConnection) Del(keys ...string) (int64, error) {
	return c.client.Del(c.ctx, keys...).Result()
}

func (c *RedisConnection) Exists(keys ...string) (int64, error) {
	return c.client.Exists(c.ctx, keys...).Result()
}

func (c *RedisConnection) Expire(key string, expiration time.Duration) (bool, error) {
	return c.client.Expire(c.ctx, key, expiration).Result()
}

func (c *RedisConnection) ExpireAt(key string, tm time.Time) (bool, error) {
	return c.client.ExpireAt(c.ctx, key, tm).Result()
}

func (c *RedisConnection) ExpireNX(key string, expiration time.Duration) (bool, error) {
	return c.client.ExpireNX(c.ctx, key, expiration).Result()
}

func (c *RedisConnection) ExpireXX(key string, expiration time.Duration) (bool, error) {
	return c.client.ExpireXX(c.ctx, key, expiration).Result()
}

func (c *RedisConnection) ExpireGT(key string, expiration time.Duration) (bool, error) {
	return c.client.ExpireGT(c.ctx, key, expiration).Result()
}

func (c *RedisConnection) ExpireLT(key string, expiration time.Duration) (bool, error) {
	return c.client.ExpireLT(c.ctx, key, expiration).Result()
}

func (c *RedisConnection) DBSize() (int64, error) {
	return c.client.DBSize(c.ctx).Result()
}

func (c *RedisConnection) Keys(pattern string) ([]string, error) {
	return c.client.Keys(c.ctx, pattern).Result()
}

func (c *RedisConnection) Migrate(host, port, key string, db int, timeout time.Duration) (string, error) {
	return c.client.Migrate(c.ctx, host, port, key, db, timeout).Result()
}

func (c *RedisConnection) Move(key string, db int) (bool, error) {
	return c.client.Move(c.ctx, key, db).Result()
}

func (c *RedisConnection) Persist(key string) (bool, error) {
	return c.client.Persist(c.ctx, key).Result()
}

func (c *RedisConnection) PExpire(key string, expiration time.Duration) (bool, error) {
	return c.client.PExpire(c.ctx, key, expiration).Result()
}

func (c *RedisConnection) PExpireAt(key string, tm time.Time) (bool, error) {
	return c.client.PExpireAt(c.ctx, key, tm).Result()
}

func (c *RedisConnection) PTTL(key string) (time.Duration, error) {
	return c.client.PTTL(c.ctx, key).Result()
}

func (c *RedisConnection) RandomKey() (string, error) {
	return c.client.RandomKey(c.ctx).Result()
}

func (c *RedisConnection) Rename(key, newKey string) (string, error) {
	return c.client.Rename(c.ctx, key, newKey).Result()
}

func (c *RedisConnection) RenameNX(key, newKey string) (bool, error) {
	return c.client.RenameNX(c.ctx, key, newKey).Result()
}

func (c *RedisConnection) Restore(key string, ttl time.Duration, value string) (string, error) {
	return c.client.Restore(c.ctx, key, ttl, value).Result()
}

func (c *RedisConnection) RestoreReplace(key string, ttl time.Duration, value string) (string, error) {
	return c.client.RestoreReplace(c.ctx, key, ttl, value).Result()
}

func (c *RedisConnection) Sort(key string, sort *r.Sort) ([]string, error) {
	return c.client.Sort(c.ctx, key, sort).Result()
}

func (c *RedisConnection) SortStore(key, store string, sort *r.Sort) (int64, error) {
	return c.client.SortStore(c.ctx, key, store, sort).Result()
}

func (c *RedisConnection) SortInterfaces(key string, sort *r.Sort) ([]interface{}, error) {
	return c.client.SortInterfaces(c.ctx, key, sort).Result()
}

func (c *RedisConnection) Touch(keys ...string) (int64, error) {
	return c.client.Touch(c.ctx, keys...).Result()
}

func (c *RedisConnection) TTL(key string) (time.Duration, error) {
	return c.client.TTL(c.ctx, key).Result()
}

func (c *RedisConnection) Type(key string) (string, error) {
	return c.client.Type(c.ctx, key).Result()
}

func (c *RedisConnection) Append(key, value string) (int64, error) {
	return c.client.Append(c.ctx, key, value).Result()
}

func (c *RedisConnection) Decr(key string) (int64, error) {
	return c.client.Decr(c.ctx, key).Result()
}

func (c *RedisConnection) DecrBy(key string, decrement int64) (int64, error) {
	return c.client.DecrBy(c.ctx, key, decrement).Result()
}

func (c *RedisConnection) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}

func (c *RedisConnection) GetRange(key string, start, end int64) (string, error) {
	return c.client.GetRange(c.ctx, key, start, end).Result()
}

func (c *RedisConnection) GetSet(key string, value interface{}) (string, error) {
	return c.client.GetSet(c.ctx, key, value).Result()
}

func (c *RedisConnection) GetEx(key string, expiration time.Duration) (string, error) {
	return c.client.GetEx(c.ctx, key, expiration).Result()
}

func (c *RedisConnection) GetDel(key string) (string, error) {
	return c.client.GetDel(c.ctx, key).Result()
}

func (c *RedisConnection) Incr(key string) (int64, error) {
	return c.client.Incr(c.ctx, key).Result()
}

func (c *RedisConnection) IncrBy(key string, value int64) (int64, error) {
	return c.client.IncrBy(c.ctx, key, value).Result()
}

func (c *RedisConnection) IncrByFloat(key string, value float64) (float64, error) {
	return c.client.IncrByFloat(c.ctx, key, value).Result()
}

func (c *RedisConnection) MGet(keys ...string) ([]interface{}, error) {
	return c.client.MGet(c.ctx, keys...).Result()
}

func (c *RedisConnection) MSet(values ...interface{}) (string, error) {
	return c.client.MSet(c.ctx, values...).Result()
}

func (c *RedisConnection) MSetNX(values ...interface{}) (bool, error) {
	return c.client.MSetNX(c.ctx, values...).Result()
}

func (c *RedisConnection) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	return c.client.Set(c.ctx, key, value, expiration).Result()
}

// TODO: more and more methods

// ------

// SetEX(ctx context.Context, key string, value interface{}, expiration time.Duration) *StatusCmd
// SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
// SetXX(ctx context.Context, key string, value interface{}, expiration time.Duration) *BoolCmd
// SetRange(ctx context.Context, key string, offset int64, value string) *IntCmd
// StrLen(ctx context.Context, key string) *IntCmd
// Copy(ctx context.Context, sourceKey string, destKey string, db int, replace bool) *IntCmd
//
// GetBit(ctx context.Context, key string, offset int64) *IntCmd
// SetBit(ctx context.Context, key string, offset int64, value int) *IntCmd
// BitCount(ctx context.Context, key string, bitCount *BitCount) *IntCmd
// BitOpAnd(ctx context.Context, destKey string, keys ...string) *IntCmd
// BitOpOr(ctx context.Context, destKey string, keys ...string) *IntCmd
// BitOpXor(ctx context.Context, destKey string, keys ...string) *IntCmd
// BitOpNot(ctx context.Context, destKey string, key string) *IntCmd
// BitPos(ctx context.Context, key string, bit int64, pos ...int64) *IntCmd
// BitField(ctx context.Context, key string, args ...interface{}) *IntSliceCmd
//
// Scan(ctx context.Context, cursor uint64, match string, count int64) *ScanCmd
// ScanType(ctx context.Context, cursor uint64, match string, count int64, keyType string) *ScanCmd
// SScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
// HScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
// ZScan(ctx context.Context, key string, cursor uint64, match string, count int64) *ScanCmd
//
// HDel(ctx context.Context, key string, fields ...string) *IntCmd
// HExists(ctx context.Context, key, field string) *BoolCmd
// HGet(ctx context.Context, key, field string) *StringCmd
// HGetAll(ctx context.Context, key string) *StringStringMapCmd
// HIncrBy(ctx context.Context, key, field string, incr int64) *IntCmd
// HIncrByFloat(ctx context.Context, key, field string, incr float64) *FloatCmd
// HKeys(ctx context.Context, key string) *StringSliceCmd
// HLen(ctx context.Context, key string) *IntCmd
// HMGet(ctx context.Context, key string, fields ...string) *SliceCmd
// HSet(ctx context.Context, key string, values ...interface{}) *IntCmd
// HMSet(ctx context.Context, key string, values ...interface{}) *BoolCmd
// HSetNX(ctx context.Context, key, field string, value interface{}) *BoolCmd
// HVals(ctx context.Context, key string) *StringSliceCmd
// HRandField(ctx context.Context, key string, count int, withValues bool) *StringSliceCmd
//
// BLPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
// BRPop(ctx context.Context, timeout time.Duration, keys ...string) *StringSliceCmd
// BRPopLPush(ctx context.Context, source, destination string, timeout time.Duration) *StringCmd
// LIndex(ctx context.Context, key string, index int64) *StringCmd
// LInsert(ctx context.Context, key, op string, pivot, value interface{}) *IntCmd
// LInsertBefore(ctx context.Context, key string, pivot, value interface{}) *IntCmd
// LInsertAfter(ctx context.Context, key string, pivot, value interface{}) *IntCmd
// LLen(ctx context.Context, key string) *IntCmd
// LPop(ctx context.Context, key string) *StringCmd
// LPopCount(ctx context.Context, key string, count int) *StringSliceCmd
// LPos(ctx context.Context, key string, value string, args LPosArgs) *IntCmd
// LPosCount(ctx context.Context, key string, value string, count int64, args LPosArgs) *IntSliceCmd
// LPush(ctx context.Context, key string, values ...interface{}) *IntCmd
// LPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
// LRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
// LRem(ctx context.Context, key string, count int64, value interface{}) *IntCmd
// LSet(ctx context.Context, key string, index int64, value interface{}) *StatusCmd
// LTrim(ctx context.Context, key string, start, stop int64) *StatusCmd
// RPop(ctx context.Context, key string) *StringCmd
// RPopCount(ctx context.Context, key string, count int) *StringSliceCmd
// RPopLPush(ctx context.Context, source, destination string) *StringCmd
// RPush(ctx context.Context, key string, values ...interface{}) *IntCmd
// RPushX(ctx context.Context, key string, values ...interface{}) *IntCmd
// LMove(ctx context.Context, source, destination, srcpos, destpos string) *StringCmd
// BLMove(ctx context.Context, source, destination, srcpos, destpos string, timeout time.Duration) *StringCmd
//
// SAdd(ctx context.Context, key string, members ...interface{}) *IntCmd
// SCard(ctx context.Context, key string) *IntCmd
// SDiff(ctx context.Context, keys ...string) *StringSliceCmd
// SDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
// SInter(ctx context.Context, keys ...string) *StringSliceCmd
// SInterStore(ctx context.Context, destination string, keys ...string) *IntCmd
// SIsMember(ctx context.Context, key string, member interface{}) *BoolCmd
// SMIsMember(ctx context.Context, key string, members ...interface{}) *BoolSliceCmd
// SMembers(ctx context.Context, key string) *StringSliceCmd
// SMembersMap(ctx context.Context, key string) *StringStructMapCmd
// SMove(ctx context.Context, source, destination string, member interface{}) *BoolCmd
// SPop(ctx context.Context, key string) *StringCmd
// SPopN(ctx context.Context, key string, count int64) *StringSliceCmd
// SRandMember(ctx context.Context, key string) *StringCmd
// SRandMemberN(ctx context.Context, key string, count int64) *StringSliceCmd
// SRem(ctx context.Context, key string, members ...interface{}) *IntCmd
// SUnion(ctx context.Context, keys ...string) *StringSliceCmd
// SUnionStore(ctx context.Context, destination string, keys ...string) *IntCmd
//
// XAdd(ctx context.Context, a *XAddArgs) *StringCmd
// XDel(ctx context.Context, stream string, ids ...string) *IntCmd
// XLen(ctx context.Context, stream string) *IntCmd
// XRange(ctx context.Context, stream, start, stop string) *XMessageSliceCmd
// XRangeN(ctx context.Context, stream, start, stop string, count int64) *XMessageSliceCmd
// XRevRange(ctx context.Context, stream string, start, stop string) *XMessageSliceCmd
// XRevRangeN(ctx context.Context, stream string, start, stop string, count int64) *XMessageSliceCmd
// XRead(ctx context.Context, a *XReadArgs) *XStreamSliceCmd
// XReadStreams(ctx context.Context, streams ...string) *XStreamSliceCmd
// XGroupCreate(ctx context.Context, stream, group, start string) *StatusCmd
// XGroupCreateMkStream(ctx context.Context, stream, group, start string) *StatusCmd
// XGroupSetID(ctx context.Context, stream, group, start string) *StatusCmd
// XGroupDestroy(ctx context.Context, stream, group string) *IntCmd
// XGroupCreateConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
// XGroupDelConsumer(ctx context.Context, stream, group, consumer string) *IntCmd
// XReadGroup(ctx context.Context, a *XReadGroupArgs) *XStreamSliceCmd
// XAck(ctx context.Context, stream, group string, ids ...string) *IntCmd
// XPending(ctx context.Context, stream, group string) *XPendingCmd
// XPendingExt(ctx context.Context, a *XPendingExtArgs) *XPendingExtCmd
// XClaim(ctx context.Context, a *XClaimArgs) *XMessageSliceCmd
// XClaimJustID(ctx context.Context, a *XClaimArgs) *StringSliceCmd
// XAutoClaim(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimCmd
// XAutoClaimJustID(ctx context.Context, a *XAutoClaimArgs) *XAutoClaimJustIDCmd
//
// // TODO: XTrim and XTrimApprox remove in v9.
// XTrim(ctx context.Context, key string, maxLen int64) *IntCmd
// XTrimApprox(ctx context.Context, key string, maxLen int64) *IntCmd
// XTrimMaxLen(ctx context.Context, key string, maxLen int64) *IntCmd
// XTrimMaxLenApprox(ctx context.Context, key string, maxLen, limit int64) *IntCmd
// XTrimMinID(ctx context.Context, key string, minID string) *IntCmd
// XTrimMinIDApprox(ctx context.Context, key string, minID string, limit int64) *IntCmd
// XInfoGroups(ctx context.Context, key string) *XInfoGroupsCmd
// XInfoStream(ctx context.Context, key string) *XInfoStreamCmd
// XInfoStreamFull(ctx context.Context, key string, count int) *XInfoStreamFullCmd
// XInfoConsumers(ctx context.Context, key string, group string) *XInfoConsumersCmd
//
// BZPopMax(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
// BZPopMin(ctx context.Context, timeout time.Duration, keys ...string) *ZWithKeyCmd
//
//
// ZAdd(ctx context.Context, key string, members ...*Z) *IntCmd
// ZAddNX(ctx context.Context, key string, members ...*Z) *IntCmd
// ZAddXX(ctx context.Context, key string, members ...*Z) *IntCmd
// ZAddCh(ctx context.Context, key string, members ...*Z) *IntCmd
// ZAddNXCh(ctx context.Context, key string, members ...*Z) *IntCmd
// ZAddXXCh(ctx context.Context, key string, members ...*Z) *IntCmd
// ZAddArgs(ctx context.Context, key string, args ZAddArgs) *IntCmd
// ZAddArgsIncr(ctx context.Context, key string, args ZAddArgs) *FloatCmd
// ZIncr(ctx context.Context, key string, member *Z) *FloatCmd
// ZIncrNX(ctx context.Context, key string, member *Z) *FloatCmd
// ZIncrXX(ctx context.Context, key string, member *Z) *FloatCmd
// ZCard(ctx context.Context, key string) *IntCmd
// ZCount(ctx context.Context, key, min, max string) *IntCmd
// ZLexCount(ctx context.Context, key, min, max string) *IntCmd
// ZIncrBy(ctx context.Context, key string, increment float64, member string) *FloatCmd
// ZInter(ctx context.Context, store *ZStore) *StringSliceCmd
// ZInterWithScores(ctx context.Context, store *ZStore) *ZSliceCmd
// ZInterStore(ctx context.Context, destination string, store *ZStore) *IntCmd
// ZMScore(ctx context.Context, key string, members ...string) *FloatSliceCmd
// ZPopMax(ctx context.Context, key string, count ...int64) *ZSliceCmd
// ZPopMin(ctx context.Context, key string, count ...int64) *ZSliceCmd
// ZRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
// ZRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
// ZRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
// ZRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
// ZRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
// ZRangeArgs(ctx context.Context, z ZRangeArgs) *StringSliceCmd
// ZRangeArgsWithScores(ctx context.Context, z ZRangeArgs) *ZSliceCmd
// ZRangeStore(ctx context.Context, dst string, z ZRangeArgs) *IntCmd
// ZRank(ctx context.Context, key, member string) *IntCmd
// ZRem(ctx context.Context, key string, members ...interface{}) *IntCmd
// ZRemRangeByRank(ctx context.Context, key string, start, stop int64) *IntCmd
// ZRemRangeByScore(ctx context.Context, key, min, max string) *IntCmd
// ZRemRangeByLex(ctx context.Context, key, min, max string) *IntCmd
// ZRevRange(ctx context.Context, key string, start, stop int64) *StringSliceCmd
// ZRevRangeWithScores(ctx context.Context, key string, start, stop int64) *ZSliceCmd
// ZRevRangeByScore(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
// ZRevRangeByLex(ctx context.Context, key string, opt *ZRangeBy) *StringSliceCmd
// ZRevRangeByScoreWithScores(ctx context.Context, key string, opt *ZRangeBy) *ZSliceCmd
// ZRevRank(ctx context.Context, key, member string) *IntCmd
// ZScore(ctx context.Context, key, member string) *FloatCmd
// ZUnionStore(ctx context.Context, dest string, store *ZStore) *IntCmd
// ZUnion(ctx context.Context, store ZStore) *StringSliceCmd
// ZUnionWithScores(ctx context.Context, store ZStore) *ZSliceCmd
// ZRandMember(ctx context.Context, key string, count int, withScores bool) *StringSliceCmd
// ZDiff(ctx context.Context, keys ...string) *StringSliceCmd
// ZDiffWithScores(ctx context.Context, keys ...string) *ZSliceCmd
// ZDiffStore(ctx context.Context, destination string, keys ...string) *IntCmd
//
// PFAdd(ctx context.Context, key string, els ...interface{}) *IntCmd
// PFCount(ctx context.Context, keys ...string) *IntCmd
// PFMerge(ctx context.Context, dest string, keys ...string) *StatusCmd
//
// BgRewriteAOF(ctx context.Context) *StatusCmd
// BgSave(ctx context.Context) *StatusCmd
// ClientKill(ctx context.Context, ipPort string) *StatusCmd
// ClientKillByFilter(ctx context.Context, keys ...string) *IntCmd
// ClientList(ctx context.Context) *StringCmd
// ClientPause(ctx context.Context, dur time.Duration) *BoolCmd
// ClientID(ctx context.Context) *IntCmd
// ConfigGet(ctx context.Context, parameter string) *SliceCmd
// ConfigResetStat(ctx context.Context) *StatusCmd
// ConfigSet(ctx context.Context, parameter, value string) *StatusCmd
// ConfigRewrite(ctx context.Context) *StatusCmd
// DBSize(ctx context.Context) *IntCmd
// FlushAll(ctx context.Context) *StatusCmd
// FlushAllAsync(ctx context.Context) *StatusCmd
// FlushDB(ctx context.Context) *StatusCmd
// FlushDBAsync(ctx context.Context) *StatusCmd
// Info(ctx context.Context, section ...string) *StringCmd
// LastSave(ctx context.Context) *IntCmd
// Save(ctx context.Context) *StatusCmd
// Shutdown(ctx context.Context) *StatusCmd
// ShutdownSave(ctx context.Context) *StatusCmd
// ShutdownNoSave(ctx context.Context) *StatusCmd
// SlaveOf(ctx context.Context, host, port string) *StatusCmd
// Time(ctx context.Context) *TimeCmd
// DebugObject(ctx context.Context, key string) *StringCmd
// ReadOnly(ctx context.Context) *StatusCmd
// ReadWrite(ctx context.Context) *StatusCmd
// MemoryUsage(ctx context.Context, key string, samples ...int) *IntCmd
//
// Eval(ctx context.Context, script string, keys []string, args ...interface{}) *Cmd
// EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *Cmd
// ScriptExists(ctx context.Context, hashes ...string) *BoolSliceCmd
// ScriptFlush(ctx context.Context) *StatusCmd
// ScriptKill(ctx context.Context) *StatusCmd
// ScriptLoad(ctx context.Context, script string) *StringCmd
//
// Publish(ctx context.Context, channel string, message interface{}) *IntCmd
// PubSubChannels(ctx context.Context, pattern string) *StringSliceCmd
// PubSubNumSub(ctx context.Context, channels ...string) *StringIntMapCmd
// PubSubNumPat(ctx context.Context) *IntCmd
//
// ClusterSlots(ctx context.Context) *ClusterSlotsCmd
// ClusterNodes(ctx context.Context) *StringCmd
// ClusterMeet(ctx context.Context, host, port string) *StatusCmd
// ClusterForget(ctx context.Context, nodeID string) *StatusCmd
// ClusterReplicate(ctx context.Context, nodeID string) *StatusCmd
// ClusterResetSoft(ctx context.Context) *StatusCmd
// ClusterResetHard(ctx context.Context) *StatusCmd
// ClusterInfo(ctx context.Context) *StringCmd
// ClusterKeySlot(ctx context.Context, key string) *IntCmd
// ClusterGetKeysInSlot(ctx context.Context, slot int, count int) *StringSliceCmd
// ClusterCountFailureReports(ctx context.Context, nodeID string) *IntCmd
// ClusterCountKeysInSlot(ctx context.Context, slot int) *IntCmd
// ClusterDelSlots(ctx context.Context, slots ...int) *StatusCmd
// ClusterDelSlotsRange(ctx context.Context, min, max int) *StatusCmd
// ClusterSaveConfig(ctx context.Context) *StatusCmd
// ClusterSlaves(ctx context.Context, nodeID string) *StringSliceCmd
// ClusterFailover(ctx context.Context) *StatusCmd
// ClusterAddSlots(ctx context.Context, slots ...int) *StatusCmd
// ClusterAddSlotsRange(ctx context.Context, min, max int) *StatusCmd
//
// GeoAdd(ctx context.Context, key string, geoLocation ...*GeoLocation) *IntCmd
// GeoPos(ctx context.Context, key string, members ...string) *GeoPosCmd
// GeoRadius(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *GeoLocationCmd
// GeoRadiusStore(ctx context.Context, key string, longitude, latitude float64, query *GeoRadiusQuery) *IntCmd
// GeoRadiusByMember(ctx context.Context, key, member string, query *GeoRadiusQuery) *GeoLocationCmd
// GeoRadiusByMemberStore(ctx context.Context, key, member string, query *GeoRadiusQuery) *IntCmd
// GeoSearch(ctx context.Context, key string, q *GeoSearchQuery) *StringSliceCmd
// GeoSearchLocation(ctx context.Context, key string, q *GeoSearchLocationQuery) *GeoSearchLocationCmd
// GeoSearchStore(ctx context.Context, key, store string, q *GeoSearchStoreQuery) *IntCmd
// GeoDist(ctx context.Context, key string, member1, member2, unit string) *FloatCmd
// GeoHash(ctx context.Context, key string, members ...string) *StringSliceCmd
