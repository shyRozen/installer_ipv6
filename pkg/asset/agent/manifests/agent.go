package manifests

import (
	aiv1beta1 "github.com/openshift/assisted-service/api/v1beta1"
	"github.com/openshift/assisted-service/models"
	"github.com/openshift/installer/pkg/asset"
	corev1 "k8s.io/api/core/v1"
)

const (
	// This could be change to "cluster-manifests" once all the agent code will be migrated to using
	// assets (and will stop reading from the hard-code "manifests" relative path)
	clusterManifestDir = "manifests"
)

var (
	_ asset.WritableAsset = (*AgentManifests)(nil)
)

// AgentManifests generates all the required manifests by the agent installer.
type AgentManifests struct {
	FileList []*asset.File

	PullSecret          *corev1.Secret
	InfraEnv            *aiv1beta1.InfraEnv
	StaticNetworkConfig []*models.HostStaticNetworkConfig
}

// Name returns a human friendly name.
func (m *AgentManifests) Name() string {
	return "Agent Manifests"
}

// Dependencies returns all of the dependencies directly needed the asset.
func (m *AgentManifests) Dependencies() []asset.Asset {
	return []asset.Asset{
		&AgentPullSecret{},
		&InfraEnv{},
		&NMStateConfig{},
	}
}

// Generate generates the respective manifest files.
func (m *AgentManifests) Generate(dependencies asset.Parents) error {
	for _, a := range []asset.WritableAsset{
		&AgentPullSecret{},
		&InfraEnv{},
		&NMStateConfig{},
	} {
		dependencies.Get(a)
		m.FileList = append(m.FileList, a.Files()...)

		switch v := a.(type) {
		case *AgentPullSecret:
			m.PullSecret = v.Config
		case *InfraEnv:
			m.InfraEnv = v.Config
		case *NMStateConfig:
			m.StaticNetworkConfig = append(m.StaticNetworkConfig, v.StaticNetworkConfig...)
		}
	}

	asset.SortFiles(m.FileList)

	return nil
}

// Files returns the files generated by the asset.
func (m *AgentManifests) Files() []*asset.File {
	return m.FileList
}

// Load returns the manifests asset from disk.
func (m *AgentManifests) Load(f asset.FileFetcher) (bool, error) {
	// yamlFileList, err := f.FetchByPattern(filepath.Join(clusterManifestDir, "*.yaml"))
	// if err != nil {
	// 	return false, errors.Wrap(err, "failed to load *.yaml files")
	// }
	// ymlFileList, err := f.FetchByPattern(filepath.Join(clusterManifestDir, "*.yml"))
	// if err != nil {
	// 	return false, errors.Wrap(err, "failed to load *.yml files")
	// }

	// fileList := append(yamlFileList, ymlFileList...)
	// if len(fileList) == 0 {
	// 	return false, nil
	// }

	// m.FileList = fileList
	// asset.SortFiles(m.FileList)

	return false, nil
}
