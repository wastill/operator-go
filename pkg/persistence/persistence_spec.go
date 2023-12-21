package persistence

import corev1 "k8s.io/api/core/v1"

type PersistenceSpec struct {
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=true
	Enable bool `json:"enable,omitempty"`
	// +kubebuilder:validation:Optional
	ExistingClaim *string `json:"existingClaim,omitempty"`
	// +kubebuilder:validation:Optional
	Annotations map[string]string `json:"annotations,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="10Gi"
	StorageSize string `json:"storageSize,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default:="ReadWriteOnce"
	AccessMode string `json:"accessMode,omitempty"`
	// +kubebuilder:validation:Optional
	StorageClassName *string `json:"storageClassName,omitempty"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:default=Filesystem
	VolumeMode *corev1.PersistentVolumeMode `json:"volumeMode,omitempty"`
}

func (p *PersistenceSpec) Existing() bool {
	return p.ExistingClaim != nil
}
