package consts

const (
	//DbServer WHERE THE DB IS HOSTED
	dbServer = "localhost"
	//DbPort WHERE THE DB IS LISTENING
	dbPort = "5432"
	//DbName NAME OF DB
	dbName = "restaurant"
	//DbUser User of DB
	dbUser = "postgres"
	//DbPass secret of User
	dbPass      = "postgres"
	redisSecret = dbServer + dbName
	cookieName  = "cnr"
	//CookieUserIdKey value
	CookieUserIdKey = "userId"
)

//GetDNS return the DNS value of the DB.
func GetDNS() string {
	var dns string
	dns = "host=" + dbServer + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " "
	return dns
}

//GetRedisSecret get secret
func GetRedisSecret() string {
	return redisSecret
}

//GetCookieName string
func GetCookieName() string {
	return cookieName
}
