package config

type Config struct {
	Databases Database `json:"databases"`
	Server    Server   `json:"server"`
	Security  Security `json:"security"`
}

type Server struct {
	Port *string `json:"port"`
}

// Database contains data needed for database connections
type Database struct {
	Nick               *string `json:"nick"`
	Name               *string `json:"name"`
	Username           *string `json:"username"`
	Password           *string `json:"password"`
	Host               *string `json:"hostname"`
	Port               *string `json:"port"`
	MaxConn            *int    `json:"max_conn"`
	MaxIdle            *int    `json:"max_idle"`
	TransactionTimeout *int    `json:"transaction_timeout"`
}

type Security struct {
	AppKey         *string `json:"app_key"`
	AppSecutiry    *string `json:"app_security"`
	Key            *string `json:"key"`
	MaxSessionTime *int    `json:"max_session_time"`
}
