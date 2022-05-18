package conv

type NewCscConf struct {
	AccessID  string `yaml:"accessID,omitempty"`
	AccessKey string `yaml:"accessKey,omitempty"`
	Account   string `yaml:"account,omitempty"`
	Log       struct {
		Level string `yaml:"level,omitempty"`
	} `yaml:"log,omitempty"`
	NameOverride     string `yaml:"nameOverride,omitempty"`
	FullnameOverride string `yaml:"fullnameOverride,omitempty"`
	ServiceAccount   struct {
		Create bool `yaml:"create,omitempty"`
	} `yaml:"serviceAccount,omitempty"`
	Rbac struct {
		Create bool `yaml:"create,omitempty"`
	} `yaml:"rbac,omitempty"`
	Image struct {
		Registry   string `yaml:"registry,omitempty"`
		Repository string `yaml:"repository,omitempty"`
		PullPolicy string `yaml:"pullPolicy,omitempty"`
		Tag        string `yaml:"tag,omitempty"`
	} `yaml:"image,omitempty"`
	Proxy struct {
		URL  string `yaml:"url,omitempty"`
		User string `yaml:"user,omitempty"`
		Pass string `yaml:"pass,omitempty"`
	} `yaml:"proxy,omitempty"`
	NodeSelector      map[string]any    `yaml:"nodeSelector,omitempty"`
	Affinity          map[string]any    `yaml:"affinity,omitempty"`
	PriorityClassName string            `yaml:"priorityClassName,omitempty"`
	Tolerations       []map[string]any  `yaml:"tolerations,omitempty"`
	Labels            map[string]string `yaml:"labels,omitempty"`
	Annotations       map[string]string `yaml:"annotations,omitempty"`
	IgnoreSSL         bool              `yaml:"ignoreSSL,omitempty"`
	ImagePullSecrets  []string          `yaml:"imagePullSecrets,omitempty"`
}
