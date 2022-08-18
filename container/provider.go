package container

type Provider interface {
	// Register registers the services into the application.
	Register()

	// Boot boots the application.
	Boot()
}
