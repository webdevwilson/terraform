package artifactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// LocalRepositoryConfiguration contains items present in local repository requests
type LocalRepositoryConfiguration struct {
	Key                          string   `json:"key,omitempty"`
	RClass                       string   `json:"rclass,omitempty"`
	PackageType                  string   `json:"packageType,omitempty"`
	Description                  string   `json:"description,omitempty"`
	Notes                        string   `json:"notes,omitempty"`
	IncludesPattern              string   `json:"includesPattern,omitempty"`
	ExcludesPattern              string   `json:"excludesPattern,omitempty"`
	RepoLayoutRef                string   `json:"repoLayoutRef,omitempty"`
	HandleReleases               bool     `json:"handleReleases,omitempty"`
	HandleSnapshots              bool     `json:"handleSnapshots,omitempty"`
	MaxUniqueSnapshots           int      `json:"maxUniqueSnapshots,omitempty"`
	DebianTrivialLayout          bool     `json:"debianTrivialLayout,omitempty"`
	ChecksumPolicyType           string   `json:"checksumPolicyType,omitempty"`
	MaxUniqueTags                int      `json:"maxUniqueTags,omitempty"`
	SnapshotVersionBehavior      string   `json:"snapshotVersionBehavior,omitempty"`
	SuppressPomConsistencyChecks bool     `json:"suppressPomConsistencyChecks,omitempty"`
	BlackedOut                   bool     `json:"blackedOut,omitempty"`
	PropertySets                 []string `json:"propertySets,omitempty"`
	ArchiveBrowsingEnabled       bool     `json:"archiveBrowsingEnabled,omitempty"`
	CalculateYumMetadata         bool     `json:"calculateYumMetadata,omitempty"`
	YumRootDepth                 int      `json:"yumRootDepth,omitempty"`
	DockerAPIVersion             string   `json:"dockerApiVersion,omitempty"`
}

// RemoteRepositoryConfiguration for configuring a remote repository
type RemoteRepositoryConfiguration struct {
	Key                               string   `json:"key,omitempty"`
	RClass                            string   `json:"rclass,omitempty"`
	PackageType                       string   `json:"packageType,omitempty"`
	URL                               string   `json:"url,omitempty"`
	Username                          string   `json:"username,omitempty"`
	Password                          string   `json:"password,omitempty"`
	Proxy                             string   `json:"proxy,omitempty"`
	Description                       string   `json:"description,omitempty"`
	Notes                             string   `json:"notes,omitempty"`
	IncludesPattern                   string   `json:"includesPattern,omitempty"`
	ExcludesPattern                   string   `json:"excludesPattern,omitempty"`
	RepoLayoutRef                     string   `json:"repoLayoutRef,omitempty"`
	RemoteRepoChecksumPolicyType      string   `json:"remoteRepoChecksumPolicyType,omitempty"`
	HandleReleases                    bool     `json:"handleReleases,omitempty"`
	HandleSnapshots                   bool     `json:"handleSnapshots,omitempty"`
	MaxUniqueSnapshots                int      `json:"maxUniqueSnapshots,omitempty"`
	SuppressPomConsistencyChecks      bool     `json:"suppressPomConsistencyChecks,omitempty"`
	HardFail                          bool     `json:"hardFail,omitempty"`
	Offline                           bool     `json:"offline,omitempty"`
	BlackedOut                        bool     `json:"blackedOut,omitempty"`
	StoreArtifactsLocally             bool     `json:"storeArtifactsLocally,omitempty"`
	SocketTimeoutMillis               int      `json:"socketTimeoutMillis,omitempty"`
	LocalAddress                      string   `json:"localAddress,omitempty"`
	RetrievalCachePeriodSeconds       int      `json:"retrievalCachePeriodSecs,omitempty"`
	FailedCachePeriodSeconds          int      `json:"failedRetrievalCachePeriodSecs,omitempty"`
	MissedCachePeriodSeconds          int      `json:"missedRetrievalCachePeriodSecs,omitempty"`
	UnusedArtifactsCleanupEnabled     bool     `json:"unusedArtifactsCleanupEnabled,omitempty"`
	UnusedArtifactsCleanupPeriodHours int      `json:"unusedArtifactsCleanupPeriodHours,omitempty"`
	FetchJarsEagerly                  bool     `json:"fetchJarsEagerly,omitempty"`
	FetchSourcesEagerly               bool     `json:"fetchSourcesEagerly,omitempty"`
	ShareConfiguration                bool     `json:"shareConfiguration,omitempty"`
	SynchronizeProperties             bool     `json:"synchronizeProperties,omitempty"`
	PropertySets                      []string `json:"propertySets,omitempty"`
	AllowAnyHostAuth                  bool     `json:"allowAnyHostAuth,omitempty"`
	EnableCookieManagement            bool     `json:"enableCookieManagement,omitempty"`
	BowerRegistryURL                  string   `json:"bowerRegistryUrl,omitempty"`
	VCSType                           string   `json:"vcsType,omitempty"`
	VCSGitProvider                    string   `json:"vcsGitProvider,omitempty"`
	VCSGitDownloadURL                 string   `json:"vcsGitDownloadUrl,omitempty"`
}

// VirtualRepositoryConfiguration for
type VirtualRepositoryConfiguration struct {
	Key                                           string   `json:"key,omitempty"`
	RClass                                        string   `json:"rclass,omitempty"`
	PackageType                                   string   `json:"packageType,omitempty"`
	Description                                   string   `json:"description,omitempty"`
	Notes                                         string   `json:"notes,omitempty"`
	IncludesPattern                               string   `json:"includesPattern,omitempty"`
	ExcludesPattern                               string   `json:"excludesPattern,omitempty"`
	ArtifactoryRequestsCanRetrieveRemoteArtifacts bool     `json:"artifactoryRequestsCanRetrieveRemoteArtifacts,omitempty"`
	KeyPair                                       string   `json:"keyPair,omitempty"`
	PomRepositoryReferencesCleanupPolicy          string   `json:"pomRepositoryReferencesCleanupPolicy,omitempty"`
	DefaultDeploymentRepo                         string   `json:"defaultDeploymentRepo,omitempty"`
	Repositories                                  []string `json:"repositories,omitempty"`
}

type clientConfig struct {
	user string
	pass string
	url  string
}

var _ Client = clientConfig{}

// Client is used to call Artifactory REST APIs
type Client interface {
	GetRepository(key string, v interface{}) error
	CreateRepository(key string, v interface{}) error
	UpdateRepository(key string, v interface{}) error
	DeleteRepository(key string) error
	CheckRepositoryExists(key string) (bool, error)
}

// NewClient constructs a new artifactory client
func NewClient(username string, pass string, url string) Client {
	return clientConfig{
		username,
		pass,
		strings.TrimRight(url, "/"),
	}
}

// GetRepository fetches repository configuration from Artifactory
func (c clientConfig) GetRepository(key string, v interface{}) error {
	path := fmt.Sprintf("repositories/%s", key)
	_, err := c.execute("GET", path, nil, v)
	return err
}

func (c clientConfig) CreateRepository(key string, v interface{}) error {
	path := fmt.Sprintf("repositories/%s", key)
	_, err := c.execute("PUT", path, v, nil)
	return err
}

func (c clientConfig) UpdateRepository(key string, v interface{}) error {
	path := fmt.Sprintf("repositories/%s", key)
	_, err := c.execute("POST", path, v, nil)
	return err
}

// DeleteRepository deletes a repository from Artifactory
func (c clientConfig) DeleteRepository(key string) error {
	path := fmt.Sprintf("repositories/%s", key)
	_, err := c.execute("DELETE", path, nil, nil)
	return err
}

// CheckRepositoryExists returns boolean whether the repository exists
func (c clientConfig) CheckRepositoryExists(key string) (bool, error) {
	path := fmt.Sprintf("repositories/%s", key)
	status, err := c.execute("HEAD", path, nil, nil)
	if status == 400 {
		return false, nil
	} else if status == 200 {
		return true, nil
	} else {
		return false, err
	}
}

func (c clientConfig) execute(method string, endpoint string, payload interface{}, writeTo interface{}) (int, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/api/%s", c.url, endpoint)

	var jsonpayload *bytes.Buffer
	if payload == nil {
		jsonpayload = &bytes.Buffer{}
	} else {
		var jsonbuffer []byte
		jsonpayload = bytes.NewBuffer(jsonbuffer)
		enc := json.NewEncoder(jsonpayload)
		enc.Encode(payload)
	}

	req, err := http.NewRequest(method, url, jsonpayload)
	if err != nil {
		return -1, err
	}
	req.SetBasicAuth(c.user, c.pass)
	req.Header.Add("content-type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return resp.StatusCode, err
	}

	if resp.StatusCode != 200 {
		return resp.StatusCode, fmt.Errorf("%s %s returned status %d", method, url, resp.StatusCode)
	}

	// read the body of the request into the passed writeTo argument
	if writeTo != nil {
		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(writeTo)
		if err != nil {
			return 200, err
		}
	}

	return 200, nil
}
