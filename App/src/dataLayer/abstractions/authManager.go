package dataLayerAbstractions

type AuthManager interface {
	OpenSession()
	GetExistingSession() any
	Close()
}
