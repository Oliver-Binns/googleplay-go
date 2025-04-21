package users

type Grant struct {
	Name                string               `json:"name,omitempty"`
	PackageName         string               `json:"packageName,omitempty"`
	AppLevelPermissions []AppLevelPermission `json:"appLevelPermissions"`
}
