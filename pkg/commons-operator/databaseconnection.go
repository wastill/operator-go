package commons_operator

// DatabaseConnectionSpec defines the desired state of DatabaseConnection
type DatabaseConnectionSpec struct {
	// +kubebuilder:validation:Required
	Provider *DatabaseProvider `json:"provider,omitempty"`
	Default  bool              `json:"default,omitempty"`
}

// DatabaseConnectionCredentialSpec include ExistSecret, it is encrypted by base64.
type DatabaseConnectionCredentialSpec struct {
	ExistSecret string `json:"existingSecret,omitempty"`
}
