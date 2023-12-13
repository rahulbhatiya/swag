package model

type ArtifactoryReleaseBundleSummary struct {
	Name    string
	Version string
	Created string
	Status  string
	Type    string
}

// BundleVersionStatus is an alias type for bundles statuses defined below
type BundleVersionStatus string

func (s *BundleVersionStatus) String() string { return string(*s) }

// ArtifactoryReleaseBundles is a set of bundles in an Artifactory response
type ArtifactoryReleaseBundles struct {
	Bundles map[string][]ArtifactoryReleaseBundleVersionStatus
}

// ArtifactoryReleaseBundleVersionStatus descripes a release bundle version in Artifactory response
type ArtifactoryReleaseBundleVersionStatus struct {
	Version string              `json:"version"`
	Created string              `json:"created"`
	Status  BundleVersionStatus `json:"status"`
}
