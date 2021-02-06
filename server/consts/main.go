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
	dbPass = "postgres"
)

//GetDNS return the DNS value of the DB.
func GetDNS() string {
	var dns string
	dns = "host=" + dbServer + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort + " "
	return dns
}
