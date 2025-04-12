package users

type AccessState string

const (
	UnspecifiedAccessState AccessState = "ACCESS_STATE_UNSPECIFIED"
	Invited                AccessState = "INVITED"
	InvitationExpired      AccessState = "INVITATION_EXPIRED"
	AccessGranted          AccessState = "ACCESS_GRANTED"
	AccessExpired          AccessState = "ACCESS_EXPIRED"
)
