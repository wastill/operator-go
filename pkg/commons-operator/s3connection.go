package commons_operator

// S3ConnectionSpec defines the desired credential of S3Connection
type S3ConnectionSpec struct {
	// +kubebuilder:validation:Required
	S3Credential *S3Credential `json:"credential,omitempty"`
}

// S3Credential include  AccessKey and SecretKey or ExistingSecret. ExistingSecret include AccessKey and SecretKey ,it is encrypted by base64.
type S3Credential struct {
	// +kubebuilder:validation:Optional
	ExistSecret string `json:"existSecret,omitempty"`
	// +kubebuilder:validation:Optional
	AccessKey string `json:"accessKey,omitempty"`
	// +kubebuilder:validation:Optional
	SecretKey string `json:"secretKey,omitempty"`
}
