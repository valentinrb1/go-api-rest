package models

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at"`
}

type Info struct {
	CPUUsage   string `json:"cpu_usage"`
	SystemLoad string `json:"system_load"`
	Memfree    string `json:"memfree"`
	Memswap    string `json:"memswap"`
}
