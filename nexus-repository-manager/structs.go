package nexus

type Config struct {
	Url      string
	Username string
	Password string
}

type nexusUser struct {
	UserId        string   `json:"userId"`
	FirstName     string   `json:"firstName"`
	LastName      string   `json:"lastName"`
	EmailAddress  string   `json:"emailAddress"`
	Source        string   `json:"source"`
	Password      string   `json:"password"`
	Status        string   `json:"status"`
	Roles         []string `json:"roles"`
	ExternalRoles []string `json:"externalRoles"`
}

func expandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
}
