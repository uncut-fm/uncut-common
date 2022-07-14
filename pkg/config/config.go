package config

const (
	CommonSecretEnvVar     = "COMMON_SECRET_ID"
	ManagementSecretEnvVar = "MANAGEMENT_SECRET_ID"
	BackofficeSecretEnvVar = "BACKOFFICE_SECRET_ID"
	AuthSecretEnvVar       = "AUTH_SECRET_ID"
	TranscoderSecretEnvVar = "TRANSCODER_SECRET_ID"
	Web3EventsSecretEnvVar = "WEB3_EVENTS_SECRET_ID"
)

type Configuration struct {
	Common     Common            `yaml:"common"`
	Management ManagementConfigs `yaml:"management"`
	Backoffice BackofficeConfigs `yaml:"backoffice"`
	Auth       AuthConfigs       `yaml:"auth"`
	Transcoder TranscoderConfigs `yaml:"transcoder"`
	Web3Events Web3EventsConfigs `yaml:"web3_events"`
}

type Common struct {
	ProjectID   string      `yaml:"project_id"`
	Environment string      `yaml:"environment"`
	JWT         JWTConfigs  `yaml:"jwt"`
	ServiceURLs ServiceUrls `yaml:"service_urls"`
}

type ManagementConfigs struct {
	PlatformDB           DBConfigs      `yaml:"platform_db"`
	Server               ServerConfigs  `yaml:"server"`
	Sentry               SentryConfigs  `yaml:"sentry"`
	PaperAPIKey          string         `yaml:"paper_api_key"`
	BlockchainName       string         `yaml:"blockchain_name"`
	StoreContractAddress string         `yaml:"store_contract_address"`
	Currencies           Web3Currencies `yaml:"currencies"`
	BlockchainRpcUrl     string         `yaml:"blockchain_rpc_url"`
	SegmentWriteKey      string         `yaml:"segment_write_key"`
	AuthAdminToken       string         `yaml:"auth_admin_token"`
	SendgridAPIKey       string         `yaml:"sendgrid_api_key"`
	Web3AdminToken       string         `yaml:"web3_admin_token"`
	GcpStorageBucket     string         `yaml:"gcp_storage_bucket"`
	Twitter              TwitterConfigs `yaml:"twitter"`
	ReservedSlugs        []string       `yaml:"reserved_slugs"`
	BackofficeAdminToken string         `yaml:"backoffice_admin_token"`
}

type BackofficeConfigs struct {
	PlatformDB             DBConfigs      `yaml:"platform_db"`
	Server                 ServerConfigs  `yaml:"server"`
	Twitter                TwitterConfigs `yaml:"twitter"`
	GcpStorageBucket       string         `yaml:"gcp_storage_bucket"`
	AdminToken             string         `yaml:"admin_token"`
	AuthAdminToken         string         `yaml:"auth_admin_token"`
	AudioCutterPubsubTopic string         `yaml:"audio_cutter_pubsub_topic"`
	Sentry                 SentryConfigs  `yaml:"sentry"`
	StoreContractAddress   string         `yaml:"store_contract_address"`
	Web3AdminToken         string         `yaml:"web3_admin_token"`
	SegmentWriteKey        string         `yaml:"segment_write_key"`
	SendgridAPIKey         string         `yaml:"sendgrid_api_key"`
	Currencies             Web3Currencies `yaml:"currencies"`
	BlockchainRpcUrl       string         `yaml:"blockchain_rpc_url"`
	BlockchainWsRpcUrl     string         `yaml:"blockchain_ws_rpc_url"`
	BlockchainName         string         `yaml:"blockchain_name"`
	FaucetWallet           FaucetWallet   `yaml:"faucet_wallet"`
	PIRTokenIDs            []int          `yaml:"pir_token_ids"`
}

type FaucetWallet struct {
	WalletAddress   string  `yaml:"wallet_address"`
	MaticAlertPoint float64 `yaml:"matic_alert_point"`
}

type AuthConfigs struct {
	AuthDB               DBConfigs      `yaml:"auth_db"`
	Server               ServerConfigs  `yaml:"server"`
	OauthProviders       OauthProviders `yaml:"oauth_providers"`
	AdminToken           string         `yaml:"admin_token"`
	BackofficeAdminToken string         `yaml:"backoffice_admin_token"`
	BaseURL              string         `yaml:"base_url"`
	MagicLinkKey         string         `yaml:"magic_link_key"`
	Sentry               SentryConfigs  `yaml:"sentry"`
	MaticFaucetLimit     FaucetLimit    `yaml:"matic_faucet_limit"`
	SegmentWriteKey      string         `yaml:"segment_write_key"`
	Twitter              TwitterConfigs `yaml:"twitter"`
}

type FaucetLimit struct {
	LimitPerUser            float32 `yaml:"limit_per_user"`
	RateLimitPeriodDuration string  `yaml:"rate_limit_period_duration"`
}

type TranscoderConfigs struct {
	GcpStorageBucket string            `yaml:"gcp_storage_bucket"`
	Server           ServerConfigs     `yaml:"server"`
	Headers          map[string]string `yaml:"headers"`
	Sentry           SentryConfigs     `yaml:"sentry"`
}

type Web3EventsConfigs struct {
	BlockchainRpcUrl     string         `yaml:"blockchain_rpc_url"`
	StoreContractAddress string         `yaml:"store_contract_address"`
	BackofficeAdminToken string         `yaml:"backoffice_admin_token"`
	Currencies           Web3Currencies `yaml:"currencies"`
	Sentry               SentryConfigs  `yaml:"sentry"`
}

type Web3Currencies struct {
	Weth  Web3CurrencyDetails `yaml:"weth"`
	Cdols Web3CurrencyDetails `yaml:"cdols"`
	Usdc  Web3CurrencyDetails `yaml:"usdc"`
	Matic Web3CurrencyDetails `yaml:"matic"`
}

func (w Web3Currencies) GetContractAddressBySymbol(symbol string) string {
	switch symbol {
	case w.Weth.Symbol:
		return w.Weth.ContractAddress
	case w.Cdols.Symbol:
		return w.Cdols.ContractAddress
	case w.Usdc.Symbol:
		return w.Cdols.ContractAddress
	default:
		return ""
	}
}

type Web3CurrencyDetails struct {
	Label           string `yaml:"label"`
	ContractAddress string `yaml:"contract_address"`
	Symbol          string `yaml:"symbol"`
	Image           string `yaml:"image"`
}

type ServerConfigs struct {
	Port string `yaml:"port"`
}

type DBConfigs struct {
	Host                  string `yaml:"host"`
	Port                  string `yaml:"port"`
	DBName                string `yaml:"db_name"`
	User                  string `yaml:"user"`
	Password              string `yaml:"password"`
	MaxOpenConnections    int    `yaml:"max_open_conns"`
	MaxIdleConnections    int    `yaml:"max_idle_conns"`
	ConnectionMaxLifetime string `yaml:"conn_max_lifetime"`
}

type TwitterConfigs struct {
	ConsumerKey    string `yaml:"consumer_key"`
	ConsumerSecret string `yaml:"consumer_secret"`
}

type GoogleOauthConfigs struct {
	ClientKey string `yaml:"client_key"`
	Secret    string `yaml:"secret"`
}

type RedisConfigs struct {
	Address  string `yaml:"address"`
	Port     string `yaml:"port"`
	DB       int    `yaml:"db"`
	Password string `yaml:"password"`
}

type JWTConfigs struct {
	SigningMethod   string `yaml:"signing_method"`
	AccessSecret    string `yaml:"access_secret"`
	AccessDuration  int    `yaml:"access_duration"`
	RefreshDuration int    `yaml:"refresh_duration"`
}

type ServiceUrls struct {
	ManagementURL       string `yaml:"management_url"`
	ManagementSchemaUrl string `yaml:"management_schema_url"`
	BackofficeSchemaUrl string `yaml:"backoffice_schema_url"`
	GatewaySchemaUrl    string `yaml:"gateway_schema_url"`
	AuthAPIUrl          string `yaml:"auth_api_url"`
	SearchAPIUrl        string `yaml:"search_api_url"`
	WebappUrl           string `yaml:"webapp_url"`
	Web3AdminUrl        string `yaml:"web3_admin_url"`
}

type OauthProviders struct {
	Twitter TwitterConfigs     `yaml:"twitter"`
	Google  GoogleOauthConfigs `yaml:"google"`
}

type SentryConfigs struct {
	DSN        string  `yaml:"dsn"`
	SampleRate float32 `yaml:"sample_rate"`
}
