package googleplay

type AppLevelPermission string

const (
	UnspecifiedAppLevelPermission AppLevelPermission = "APP_LEVEL_PERMISSION_UNSPECIFIED"
	CanAccessApp                                     = "CAN_ACCESS_APP"
	CanViewFinancialData                             = "CAN_VIEW_FINANCIAL_DATA"
	CanManagePermissions                             = "CAN_MANAGE_PERMISSIONS"
	CanReplyToReviews                                = "CAN_REPLY_TO_REVIEWS"
	CanManagePublicAPKs                              = "CAN_MANAGE_PUBLIC_APKS"
	CanManageTrackAPKs                               = "CAN_MANAGE_TRACK_APKS"
	CanManageTrackUsers                              = "CAN_MANAGE_TRACK_USERS"
	CanManagePublicListing                           = "CAN_MANAGE_PUBLIC_LISTING"
	CanManageDraftApps                               = "CAN_MANAGE_DRAFT_APPS"
	CanManageOrders                                  = "CAN_MANAGE_ORDERS"
	CanManageAppContent                              = "CAN_MANAGE_APP_CONTENT"
	CanViewNonFinancialData                          = "CAN_VIEW_NON_FINANCIAL_DATA"
	CanViewAppQuality                                = "CAN_VIEW_APP_QUALITY"
	CanManageDeeplinks                               = "CAN_MANAGE_DEEPLINKS"
)
