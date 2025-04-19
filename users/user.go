package users

type User struct {
	Name                        string                     `json:"name"`
	Email                       string                     `json:"email"`
	DeveloperAccountPermissions []DeveloperLevelPermission `json:"developerAccountPermissions"`
	Partial                     bool                       `json:"partial,omitempty"`
	AccessState                 AccessState                `json:"accessState,omitempty"`
	Grants                      []Grant                    `json:"grants,omitempty"`
}
