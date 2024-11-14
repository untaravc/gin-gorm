package request

type VersionCheckRequest struct {
	AppVersion string `json:"app_version" binding:"required"`
}
