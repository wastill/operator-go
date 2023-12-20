package commons_operator //nolint:typecheck

// DatabaseSpec defines the desired connection info of Database
type DatabaseSpec struct {
	//+kubebuilder:validation:Required
	Name string `json:"name,omitempty"`
	//+kubebuilder:validation:Required
	Reference string `json:"reference,omitempty"`
	//+kubebuilder:validation:Required
	Credential *DatabaseCredentialSpec `json:"credential,omitempty"`
}

// DatabaseCredentialSpec include  Username and Password or ExistSecret. ExistSecret include Username and Password ,it is encrypted by base64.
type DatabaseCredentialSpec struct {
	ExistSecret string `json:"existingSecret,omitempty"`
	Username    string `json:"username,omitempty"`
	Password    string `json:"password,omitempty"`
}
