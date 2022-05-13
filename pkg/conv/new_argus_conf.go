package conv

type NewArgusConf struct {
	AccessID            string `yaml:"accessID,omitempty"`
	AccessKey           string `yaml:"accessKey,omitempty"`
	Account             string `yaml:"account,omitempty"`
	ClusterName         string `yaml:"clusterName,omitempty"`
	ClusterTreeParentID int    `yaml:"clusterTreeParentID,omitempty"`
	Image               struct {
		Registry   string `yaml:"registry,omitempty"`
		Repository string `yaml:"repository,omitempty"`
		PullPolicy string `yaml:"pullPolicy,omitempty"`
		Tag        string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	NodeSelector struct {
	} `yaml:"nodeSelector,omitempty"`
	Affinity struct {
	} `yaml:"affinity,omitempty"`
	PriorityClassName string        `yaml:"priorityClassName,omitempty"`
	Tolerations       []interface{} `yaml:"tolerations,omitempty"`
	Resources         struct {
	} `yaml:"resources,omitempty"`
	Labels struct {
	} `yaml:"labels,omitempty"`
	Annotations struct {
	} `yaml:"annotations,omitempty"`
	Replicas            int `yaml:"replicas,omitempty"`
	ResourceContainerID int `yaml:"resourceContainerID,omitempty"`
	Log                 struct {
		Level string `yaml:"level,omitempty"`
	} `yaml:"log,omitempty"`
	Collectorsetcontroller struct {
		Address string `yaml:"address,omitempty"`
		Port    int    `yaml:"port,omitempty"`
	} `yaml:"collectorsetcontroller,omitempty"`
	Proxy struct {
		URL  string `yaml:"url,omitempty"`
		User string `yaml:"user,omitempty"`
		Pass string `yaml:"pass,omitempty"`
	} `yaml:"proxy,omitempty"`
	EtcdDiscoveryToken string `yaml:"etcdDiscoveryToken,omitempty"`
	IgnoreSSL          bool   `yaml:"ignoreSSL,omitempty"`
	Daemons            struct {
		LmResourceSweeper struct {
			Interval string `yaml:"interval,omitempty"`
		} `yaml:"lmResourceSweeper,omitempty"`
		LmCacheSync struct {
			Interval string `yaml:"interval,omitempty"`
		} `yaml:"lmCacheSync,omitempty"`
		Worker struct {
			PoolSize int `yaml:"poolSize,omitempty"`
		} `yaml:"worker,omitempty"`
		Watcher struct {
			BulkSyncInterval string `yaml:"bulkSyncInterval,omitempty"`
			Runner           struct {
				PoolSize                       int `yaml:"poolSize,omitempty"`
				BackPressureQueueSizePerRunner int `yaml:"backPressureQueueSizePerRunner,omitempty"`
			} `yaml:"runner,omitempty"`
			SysIpsWaitTimeout string `yaml:"sysIpsWaitTimeout,omitempty"`
		} `yaml:"watcher,omitempty"`
	} `yaml:"daemons,omitempty"`
	Monitoring struct {
		Disable []string `yaml:"disable,omitempty"`
	} `yaml:"monitoring,omitempty"`
	Lm struct {
		Lmlogs struct {
			K8Sevent struct {
				Enable bool `yaml:"enable,omitempty"`
			} `yaml:"k8sevent,omitempty"`
			K8Spodlog struct {
				Enable bool `yaml:"enable,omitempty"`
			} `yaml:"k8spodlog,omitempty"`
		} `yaml:"lmlogs,omitempty"`
		Resource struct {
			GlobalDeleteAfterDuration string `yaml:"globalDeleteAfterDuration,omitempty"`
			Alerting                  struct {
				Disable []string `yaml:"disable,omitempty"`
			} `yaml:"alerting,omitempty"`
		} `yaml:"resource,omitempty"`
		ResourceGroup struct {
			ExtraProps struct {
				Cluster []struct {
					Name     string `yaml:"name,omitempty"`
					Value    string `yaml:"value,omitempty"`
					Override bool   `yaml:"override,omitempty"`
				} `yaml:"cluster,omitempty"`
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
			} `yaml:"extraProps,omitempty"`
		} `yaml:"resourceGroup,omitempty"`
	} `yaml:"lm,omitempty"`
	Filters     []string `yaml:"filters,omitempty"`
	SelfMonitor struct {
		Enable bool `yaml:"enable,omitempty"`
		Port   int  `yaml:"port,omitempty"`
	} `yaml:"selfMonitor,omitempty"`
	Debug struct {
		Profiling struct {
			Enable bool `yaml:"enable,omitempty"`
		} `yaml:"profiling,omitempty"`
	} `yaml:"debug,omitempty"`
	Collector struct {
		Replicas int    `yaml:"replicas,omitempty"`
		Version  int    `yaml:"version,omitempty"`
		Size     string `yaml:"size,omitempty"`
		UseEA    bool   `yaml:"useEA,omitempty"`
		Lm       struct {
			GroupID           int `yaml:"groupID,omitempty"`
			EscalationChainID int `yaml:"escalationChainID,omitempty"`
		} `yaml:"lm,omitempty"`
		Image struct {
			Registry   string `yaml:"registry,omitempty"`
			Repository string `yaml:"repository,omitempty"`
			Tag        string `yaml:"tag,omitempty"`
			PullPolicy string `yaml:"pullPolicy,omitempty"`
		} `yaml:"image,omitempty"`
		Proxy struct {
			URL  string `yaml:"url,omitempty"`
			User string `yaml:"user,omitempty"`
			Pass string `yaml:"pass,omitempty"`
		} `yaml:"proxy,omitempty"`
		Annotations struct {
		} `yaml:"annotations,omitempty"`
		Labels struct {
		} `yaml:"labels,omitempty"`
		StatefulsetSpec struct {
			Template struct {
				Spec struct {
				} `yaml:"spec,omitempty"`
			} `yaml:"template,omitempty"`
		} `yaml:"statefulsetSpec,omitempty"`
	} `yaml:"collector,omitempty"`
	KubeStateMetrics struct {
		Enabled     bool `yaml:"enabled,omitempty"`
		SelfMonitor struct {
			Enabled       bool `yaml:"enabled,omitempty"`
			TelemetryPort int  `yaml:"telemetryPort,omitempty"`
		} `yaml:"selfMonitor,omitempty"`
		Replicas   int      `yaml:"replicas,omitempty"`
		Collectors []string `yaml:"collectors,omitempty"`
	} `yaml:"kube-state-metrics,omitempty"`
	ImagePullSecrets []interface{} `yaml:"imagePullSecrets,omitempty"`
	Global           struct {
		AccessID  string `yaml:"accessID,omitempty"`
		AccessKey string `yaml:"accessKey,omitempty"`
		Account   string `yaml:"account,omitempty"`
		Proxy     struct {
			URL  string `yaml:"url,omitempty"`
			User string `yaml:"user,omitempty"`
			Pass string `yaml:"pass,omitempty"`
		} `yaml:"proxy,omitempty"`
		Image struct {
			Registry   string `yaml:"registry,omitempty"`
			PullPolicy string `yaml:"pullPolicy,omitempty"`
		} `yaml:"image,omitempty"`
		CollectorsetServiceNameSuffix string        `yaml:"collectorsetServiceNameSuffix,omitempty"`
		ImagePullSecrets              []interface{} `yaml:"imagePullSecrets,omitempty"`
	} `yaml:"global,omitempty"`
	NameOverride     string `yaml:"nameOverride,omitempty"`
	FullnameOverride string `yaml:"fullnameOverride,omitempty"`
	Rbac             struct {
		Create bool `yaml:"create,omitempty"`
	} `yaml:"rbac,omitempty"`
	ServiceAccount struct {
		Create bool `yaml:"create,omitempty"`
	} `yaml:"serviceAccount,omitempty"`
}
