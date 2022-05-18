package conv

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"strings"
)

type OldArgusConf struct {
	AccessID           string `yaml:"accessID,omitempty"`
	AccessKey          string `yaml:"accessKey,omitempty"`
	Account            string `yaml:"account,omitempty"`
	ClusterName        string `yaml:"clusterName,omitempty"`
	LogLevel           string `yaml:"logLevel,omitempty"`
	DeleteDevices      bool   `yaml:"deleteDevices,omitempty"`
	DisableAlerting    bool   `yaml:"disableAlerting,omitempty"`
	EnableRBAC         bool   `yaml:"enableRBAC,omitempty"`
	ClusterGroupID     int    `yaml:"clusterGroupID,omitempty"`
	ResourceGroupID    int    `yaml:"resourceGroupID,omitempty"`
	EtcdDiscoveryToken string `yaml:"etcdDiscoveryToken,omitempty"`
	ImageRepository    string `yaml:"imageRepository,omitempty"`
	ImageTag           string `yaml:"imageTag,omitempty"`
	ImagePullPolicy    string `yaml:"imagePullPolicy,omitempty"`
	ProxyURL           string `yaml:"proxyURL,omitempty"`
	ProxyUser          string `yaml:"proxyUser,omitempty"`
	ProxyPass          string `yaml:"proxyPass,omitempty"`
	NodeSelector       struct {
	} `yaml:"nodeSelector,omitempty"`
	Affinity struct {
	} `yaml:"affinity,omitempty"`
	PriorityClassName string        `yaml:"priorityClassName,omitempty"`
	Tolerations       []interface{} `yaml:"tolerations,omitempty"`
	Labels            struct {
	} `yaml:"labels,omitempty"`
	Annotations struct {
	} `yaml:"annotations,omitempty"`
	IgnoreSsl             bool `yaml:"ignore_ssl,omitempty"`
	RegisterGenericFilter bool `yaml:"registerGenericFilter,omitempty"`
	AppIntervals          struct {
		PeriodicSyncInterval   string `yaml:"periodic_sync_interval,omitempty"`
		PeriodicDeleteInterval string `yaml:"periodic_delete_interval,omitempty"`
		CacheSyncInterval      string `yaml:"cache_sync_interval,omitempty"`
	} `yaml:"app_intervals,omitempty"`
	DeviceGroupProps struct {
		Cluster []struct {
			Name     string `yaml:"name,omitempty"`
			Value    string `yaml:"value,omitempty"`
			Override bool   `yaml:"override,omitempty"`
		} `yaml:"cluster,omitempty"`
		Pods []struct {
			Name     string `yaml:"name,omitempty"`
			Value    string `yaml:"value,omitempty"`
			Override bool   `yaml:"override,omitempty"`
		} `yaml:"pods,omitempty"`
		Services []struct {
			Name     string `yaml:"name,omitempty"`
			Value    string `yaml:"value,omitempty"`
			Override bool   `yaml:"override,omitempty"`
		} `yaml:"services,omitempty"`
		Deployments []struct {
			Name     string `yaml:"name,omitempty"`
			Value    string `yaml:"value,omitempty"`
			Override bool   `yaml:"override,omitempty"`
		} `yaml:"deployments,omitempty"`
		Nodes []struct {
			Name     string `yaml:"name,omitempty"`
			Value    string `yaml:"value,omitempty"`
			Override bool   `yaml:"override,omitempty"`
		} `yaml:"nodes,omitempty"`
		Etcd []struct {
			Name     string `yaml:"name,omitempty"`
			Value    string `yaml:"value,omitempty"`
			Override bool   `yaml:"override,omitempty"`
		} `yaml:"etcd,omitempty"`
		Hpas []struct {
			Name     string `yaml:"name,omitempty"`
			Value    string `yaml:"value,omitempty"`
			Override bool   `yaml:"override,omitempty"`
		} `yaml:"hpas,omitempty"`
	} `yaml:"device_group_props,omitempty"`
	Filters     map[string]any `yaml:"filters,omitempty"`
	Openmetrics struct {
		Port int `yaml:"port,omitempty"`
	} `yaml:"openmetrics,omitempty"`
	Collector struct {
		Replicas          int    `yaml:"replicas,omitempty"`
		Size              string `yaml:"size,omitempty"`
		ImageRepository   string `yaml:"imageRepository,omitempty"`
		ImageTag          string `yaml:"imageTag,omitempty"`
		ImagePullPolicy   string `yaml:"imagePullPolicy,omitempty"`
		SecretName        string `yaml:"secretName,omitempty"`
		GroupID           int    `yaml:"groupID,omitempty"`
		EscalationChainID int    `yaml:"escalationChainID,omitempty"`
		CollectorVersion  int    `yaml:"collectorVersion,omitempty"`
		UseEA             bool   `yaml:"useEA,omitempty"`
		ProxyURL          string `yaml:"proxyURL,omitempty"`
		ProxyUser         string `yaml:"proxyUser,omitempty"`
		ProxyPass         string `yaml:"proxyPass,omitempty"`
		Annotations       struct {
		} `yaml:"annotations,omitempty"`
		Labels struct {
		} `yaml:"labels,omitempty"`
		Statefulsetspec struct {
			Template struct {
				Spec struct {
				} `yaml:"spec,omitempty"`
			} `yaml:"template,omitempty"`
		} `yaml:"statefulsetspec,omitempty"`
	} `yaml:"collector,omitempty"`
	DisableResourceMonitoring []string `yaml:"disableResourceMonitoring,omitempty"`
	DisableResourceAlerting   []string `yaml:"disableResourceAlerting,omitempty"`
	Replicas                  int      `yaml:"replicas,omitempty"`
	KubeStateMetrics          struct {
		Enabled     bool `yaml:"enabled,omitempty"`
		Replicas    int  `yaml:"replicas,omitempty"`
		SelfMonitor struct {
			Enabled       bool `yaml:"enabled,omitempty"`
			TelemetryPort int  `yaml:"telemetryPort,omitempty"`
		} `yaml:"selfMonitor,omitempty"`
		Collectors []string `yaml:"collectors,omitempty"`
	} `yaml:"kube-state-metrics,omitempty"`
	Resources struct {
	} `yaml:"resources,omitempty"`
}

func (oldConf *OldArgusConf) ToNewArgusConf() *NewArgusConf {
	newConf := &NewArgusConf{}
	newConf.AccessKey = oldConf.AccessKey
	newConf.AccessID = oldConf.AccessID
	newConf.Account = oldConf.Account

	newConf.ClusterName = oldConf.ClusterName

	if oldConf.ClusterGroupID > 1 {
		newConf.ClusterTreeParentID = oldConf.ClusterGroupID
	}

	if oldConf.ImageRepository != "logicmonitor/argus" {
		newConf.Image.Repository = oldConf.ImageRepository
	}

	newConf.NodeSelector = oldConf.NodeSelector
	newConf.Affinity = oldConf.Affinity
	newConf.PriorityClassName = oldConf.PriorityClassName
	newConf.Tolerations = oldConf.Tolerations
	newConf.Resources = oldConf.Resources
	newConf.Labels = oldConf.Labels
	newConf.Annotations = oldConf.Annotations
	if oldConf.ResourceGroupID > 1 {
		newConf.ResourceContainerID = oldConf.ResourceGroupID
	}

	if strings.ToLower(oldConf.LogLevel) != "info" {
		newConf.Log.Level = oldConf.LogLevel
	}

	newConf.Proxy.URL = oldConf.ProxyURL
	newConf.Proxy.User = oldConf.ProxyUser
	newConf.Proxy.Pass = oldConf.ProxyPass

	newConf.EtcdDiscoveryToken = oldConf.EtcdDiscoveryToken
	newConf.IgnoreSSL = oldConf.IgnoreSsl
	newConf.Monitoring.Disable = oldConf.DisableResourceMonitoring
	for _, prop := range oldConf.DeviceGroupProps.Cluster {
		if prop.Name == "lmlogs.k8sevent.enable" {
			if prop.Value == "true" {
				newConf.Lm.Lmlogs.K8Sevent.Enable = true
			}
		} else if prop.Name == "kubernetes.resourcedeleteafterduration" {
			newConf.Lm.Resource.GlobalDeleteAfterDuration = prop.Value
		} else {
			newConf.Lm.ResourceGroup.ExtraProps.Cluster = append(newConf.Lm.ResourceGroup.ExtraProps.Cluster, prop)
		}
	}
	for _, prop := range oldConf.DeviceGroupProps.Pods {
		if prop.Name == "lmlogs.k8spodlog.enable" && prop.Value == "true" {
			newConf.Lm.Lmlogs.K8Spodlog.Enable = true
		}
	}

	newConf.Lm.Resource.Alerting.Disable = oldConf.DisableResourceAlerting
	newConf.Lm.ResourceGroup.ExtraProps.Etcd = oldConf.DeviceGroupProps.Etcd
	newConf.Lm.ResourceGroup.ExtraProps.Nodes = oldConf.DeviceGroupProps.Nodes

	for k, v := range oldConf.Filters {
		if value, ok := v.(string); ok {
			if len(value) > 0 {
				arr := strings.Split(value, " || ")
				for _, r := range arr {
					//newConf.Filters = append(newConf.Filters, "type == \""+k+"\" && "+r)
					newConf.Filters = append(newConf.Filters, fmt.Sprintf("type == \"%s\" && %s", k, r))
				}
			}
		} else if value, ok := v.([]any); ok {
			if len(value) > 0 {
				for _, r := range value {
					newConf.Filters = append(newConf.Filters, fmt.Sprintf("type == \"%s\" && %s", k, r.(string)))
				}
			}
		}
	}

	newConf.Collector.Version = oldConf.Collector.CollectorVersion
	newConf.Collector.Size = oldConf.Collector.Size
	newConf.Collector.UseEA = oldConf.Collector.UseEA
	if oldConf.Collector.GroupID == -1 {
		newConf.Collector.Lm.GroupID = 0
	} else if oldConf.Collector.GroupID != 1 {
		newConf.Collector.Lm.GroupID = oldConf.Collector.GroupID
	}
	newConf.Collector.Lm.EscalationChainID = oldConf.Collector.EscalationChainID
	if oldConf.Collector.ImageRepository != "logicmonitor/collector" {
		newConf.Collector.Image.Repository = oldConf.Collector.ImageRepository
	}
	newConf.Collector.Proxy.URL = oldConf.Collector.ProxyURL
	newConf.Collector.Proxy.User = oldConf.Collector.ProxyUser
	newConf.Collector.Proxy.Pass = oldConf.Collector.ProxyPass

	if oldConf.Replicas > 1 {
		newConf.Collector.Replicas = oldConf.Replicas
	}
	newConf.Collector.Annotations = oldConf.Collector.Annotations
	newConf.Collector.Labels = oldConf.Collector.Labels
	newConf.Collector.StatefulsetSpec.Template.Spec = oldConf.Collector.Statefulsetspec.Template.Spec

	return newConf
}

func LoadArgusConf(values string) error {
	conf := &OldArgusConf{}
	err := yaml.Unmarshal([]byte(values), conf)
	if err != nil {
		return err
	}
	marshal, err := yaml.Marshal(*conf)
	if err != nil {
		return err
	}
	fmt.Println(string(marshal))

	newArgusConf := conf.ToNewArgusConf()
	bytes, err := yaml.Marshal(newArgusConf)
	if err != nil {
		return err
	}
	fmt.Println("New Argus Config:")
	fmt.Println(string(bytes))
	return nil
}
