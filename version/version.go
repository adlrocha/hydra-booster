package version

const (
	// Version number of the Hydra Booster node, it should be kept in sync with the current release tag.
	Version = "0.4.2"
	// UserAgent is the string passed by the identify protocol to other nodes in the network.
	UserAgent = "hydra-booster/" + Version
)
