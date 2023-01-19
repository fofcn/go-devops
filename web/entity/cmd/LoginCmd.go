package cmd

type LoginCmd struct {
	User string `json:"username"`
	Pass string `json:"password"`
}
