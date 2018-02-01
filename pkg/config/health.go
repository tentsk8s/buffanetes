package config

// Health is config directive that tells Buffanetes how to check for readiness and liveness
type Health struct {
	Alive string `toml:"alive"`
	Ready string `toml:"ready"`
}
