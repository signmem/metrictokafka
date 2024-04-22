package g

type GlobalConfig struct {
	Debug			bool		`json:"debug"`
	LogFile			string		`json:"logfile"`
	LogMaxAge		int			`json:"logmaxage"`
	LogRotateAge	int			`json:"logrotateage"`
	Kafka 			*KConfig	`json:"kafka"`
	Http 			*HTTP		`json:"http"`
}

type HTTP struct {
	Address			string		`json:"address"`
	Port			string		`json:"port"`
}

type KConfig struct {
	Enable 		bool 			`json:"enable"`
	Topic 		string			`json:"topic"`
	Servers 	[]string 		`json:"servers"`
}