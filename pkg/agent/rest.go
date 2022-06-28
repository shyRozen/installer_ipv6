package agent

import (
	"context"
	"net"
	"net/url"

	"github.com/go-openapi/strfmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/openshift/assisted-service/client"
	"github.com/openshift/assisted-service/client/installer"

	"github.com/openshift/installer/pkg/asset/agent/manifests"
	assetstore "github.com/openshift/installer/pkg/asset/store"
)

// NodeZeroRestClient is a struct to interact with the Agent Rest API that is on node zero.
type NodeZeroRestClient struct {
	Client     *client.AssistedInstall
	ctx        context.Context
	config     client.Config
	NodeZeroIP string
}

// NewNodeZeroRestClient Initialize a new rest client to interact with the Agent Rest API on node zero.
func NewNodeZeroRestClient(ctx context.Context, assetDir string) (*NodeZeroRestClient, error) {
	restClient := &NodeZeroRestClient{}

	assetStore, err := assetstore.NewStore(assetDir)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create asset store")
	}
	nmState := &manifests.NMStateConfig{}
	if err := assetStore.Fetch(nmState); err != nil {
		return nil, errors.Wrapf(err, "failed to fetch %s", nmState.Name())
	}

	NodeZeroIP, err := manifests.GetNodeZeroIP(nmState.Config)
	if err != nil {
		return nil, err
	}
	config := client.Config{}
	config.URL = &url.URL{
		Scheme: "http",
		Host:   net.JoinHostPort(NodeZeroIP, "8090"),
		Path:   client.DefaultBasePath,
	}
	client := client.New(config)

	restClient.Client = client
	restClient.ctx = ctx
	restClient.config = config
	restClient.NodeZeroIP = NodeZeroIP

	return restClient, nil
}

// IsRestAPILive Determine if the Agent Rest API on node zero has initialized
func (rest *NodeZeroRestClient) IsRestAPILive() (bool, error) {
	// GET /v2/infraenvs
	listInfraEnvsParams := installer.NewListInfraEnvsParams()
	_, err := rest.Client.Installer.ListInfraEnvs(rest.ctx, listInfraEnvsParams)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetRestAPIServiceBaseURL Return the url of the Agent Rest API on node zero
func (rest *NodeZeroRestClient) GetRestAPIServiceBaseURL() *url.URL {
	return rest.config.URL
}

// getClusterID Return the cluster ID assigned by the Agent Rest API
func (rest *NodeZeroRestClient) getClusterID() (*strfmt.UUID, error) {
	// GET /v2/clusters and return first result
	listClusterParams := installer.NewV2ListClustersParams()
	clusterResult, err := rest.Client.Installer.V2ListClusters(rest.ctx, listClusterParams)
	clusterList := clusterResult.Payload
	if err != nil {
		return nil, err
	}
	if len(clusterList) == 1 {
		clusterID := clusterList[0].ID
		return clusterID, nil
	} else if len(clusterList) == 0 {
		logrus.Debug("cluster is not registered in rest API")
		return nil, nil
	} else {
		logrus.Debugf("number of clusters found: %d", len(clusterList))
		return nil, errors.New("found too many clusters")
	}
}

// getClusterID Return the infraEnv ID associated with the cluster in the Agent Rest API
func (rest *NodeZeroRestClient) getClusterInfraEnvID() (*strfmt.UUID, error) {
	// GET /v2/infraenvs and return first result
	listInfraEnvParams := installer.NewListInfraEnvsParams()
	infraEnvResult, err := rest.Client.Installer.ListInfraEnvs(rest.ctx, listInfraEnvParams)
	infraEnvList := infraEnvResult.Payload
	if err != nil {
		return nil, err
	}
	if len(infraEnvList) == 1 {
		clusterInfraEnvID := infraEnvList[0].ID
		return clusterInfraEnvID, nil
	} else if len(infraEnvList) == 0 {
		logrus.Debug("infraenv is not registered in rest API")
		return nil, nil
	} else {
		logrus.Debugf("number of infraenvs found: %d", len(infraEnvList))
		return nil, errors.New("found too many infraenvs")
	}
}
