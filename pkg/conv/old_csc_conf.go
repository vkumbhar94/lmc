package conv

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

type OldCscConf struct {
	AccessID           string            `yaml:"accessID,omitempty"`
	AccessKey          string            `yaml:"accessKey,omitempty"`
	Account            string            `yaml:"account,omitempty"`
	Debug              bool              `yaml:"debug,omitempty"`
	EnableRBAC         bool              `yaml:"enableRBAC,omitempty"`
	EtcdDiscoveryToken string            `yaml:"etcdDiscoveryToken,omitempty"`
	ImageRepository    string            `yaml:"imageRepository,omitempty"`
	ImageTag           string            `yaml:"imageTag,omitempty"`
	ImagePullPolicy    string            `yaml:"imagePullPolicy,omitempty"`
	ProxyURL           string            `yaml:"proxyURL,omitempty"`
	ProxyUser          string            `yaml:"proxyUser,omitempty"`
	ProxyPass          string            `yaml:"proxyPass,omitempty"`
	NodeSelector       map[string]any    `yaml:"nodeSelector,omitempty"`
	Affinity           map[string]any    `yaml:"affinity,omitempty"`
	PriorityClassName  string            `yaml:"priorityClassName,omitempty"`
	Tolerations        []map[string]any  `yaml:"tolerations,omitempty"`
	Labels             map[string]string `yaml:"labels,omitempty"`
	Annotations        map[string]string `yaml:"annotations,omitempty"`
	IgnoreSsl          bool              `yaml:"ignore_ssl,omitempty"`
}

func (oldConf *OldCscConf) ToNewCscConf() *NewCscConf {
	newCscConf := &NewCscConf{}
	newCscConf.Account = oldConf.Account
	newCscConf.AccessID = oldConf.AccessID
	newCscConf.AccessKey = oldConf.AccessKey

	if oldConf.Debug {
		newCscConf.Log.Level = "debug"
	}
	if oldConf.ImageRepository != "logicmonitor/collectorset-controller" {
		newCscConf.Image.Repository = oldConf.ImageRepository
	}
	newCscConf.Proxy.URL = oldConf.ProxyURL
	newCscConf.Proxy.User = oldConf.ProxyUser
	newCscConf.Proxy.Pass = oldConf.ProxyPass

	newCscConf.NodeSelector = oldConf.NodeSelector
	newCscConf.Affinity = oldConf.Affinity
	newCscConf.PriorityClassName = oldConf.PriorityClassName
	newCscConf.Tolerations = oldConf.Tolerations
	newCscConf.Labels = oldConf.Labels
	newCscConf.Annotations = oldConf.Annotations
	newCscConf.IgnoreSSL = oldConf.IgnoreSsl

	return newCscConf
}

func LoadCscConf(values string) error {
	conf := &OldCscConf{}
	err := yaml.Unmarshal([]byte(values), conf)
	if err != nil {
		return err
	}
	marshal, err := yaml.Marshal(*conf)
	if err != nil {
		return err
	}
	fmt.Println(string(marshal))

	newArgusConf := conf.ToNewCscConf()
	bytes, err := yaml.Marshal(newArgusConf)
	if err != nil {
		return err
	}
	fmt.Println("New CSC Config:")
	fmt.Println(string(bytes))
	return nil
}
