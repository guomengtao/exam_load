package utils

func CheckMySQLStatus() bool {
	if DB == nil {
		return false
	}
	err := DB.Ping()
	return err == nil
}

func CheckRedisStatus() bool {
	if RedisClient == nil {
		return false
	}
	_, err := RedisClient.Ping(Ctx).Result()
	return err == nil
}