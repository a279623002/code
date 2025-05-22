package platform

type Platform interface {
	Start()
	Run(string, string)
}

type PlatformManager struct {
	Platform
}

func NewPlatformManager(platform Platform) *PlatformManager {
	return &PlatformManager{platform}
}

