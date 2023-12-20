package commons_operator

const S3BucketFinalizer = "s3bucket.finalizers.stack.zncdata.net"

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// S3BucketSpec defines the desired fields of S3Bucket
type S3BucketSpec struct {

	// +kubebuilder:validation:Required
	Reference string `json:"reference,omitempty"`

	// +kubebuilder:validation:Optional
	Credential *S3BucketCredential `json:"credential,omitempty"`

	// +kubebuilder:validation:Required
	Name string `json:"name,omitempty"`
}

// S3BucketCredential defines the desired secret of S3Bucket
type S3BucketCredential struct {
	// +kubebuilder:validation:Optional
	ExistSecret string `json:"existSecret,omitempty"`
}
