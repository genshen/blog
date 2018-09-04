package utils

import (
	"github.com/BurntSushi/toml"
	"log"
)

type AuthO2Config struct {
	AuthUrl  string `toml:"auth_url"`
	ClientId string `toml:"client_id"`
	SecretId string `toml:"secret_id"`
}

var CustomConfig struct {
	Database struct {
		DbType   string `toml:"db_type"`
		DbName   string `toml:"db_name"`
		DbDebug  bool   `toml:"db_debug"`
		DbConfig string `toml:"db_config"`

		DbAuth     bool   `toml:"db_auth"` // database need auth if true
		DbAuthUser string `toml:"db_auth_user"`
		DbAuthPwd  string `toml:"db_auth_pwd"`
	} `toml:"database"`
	Auth struct {
		// authO2 config
		BeforeAuth string                  `toml:"before_auth"`
		Keys       map[string]AuthO2Config `toml:"keys"`
	} `toml:"auth"`
	Storage struct {
		EnableQiNiuCloud bool `toml:"enable_qiniu_cloud"`
		QiNiuConfig      struct {
			AccessKey    string `toml:"access_key"`
			SecretKey    string `toml:"secret_key"`
			BucketName   string `toml:"bucket_name"`
			TokenExpires uint32 `toml:"token_expires"`
			ConfigDomain string `toml:"config_domain"`
			UploadPath   string `toml:"upload_path"`
		} `toml:"qiniu_config"`
		LocalStorageDir       string `toml:"local_storage_dir"`
		LocalStorageUploadUrl string `toml:"local_storage_upload_url"`
		LocalStorageDomain    string `toml:"local_storage_domain"` // todo default "/images/:hash"
	} `toml:"storage"`
	Api struct {
		BlogPagesPrefix  string `toml:"blog_pages_prefix"`
		BlogApiPrefix    string `toml:"blog_api_prefix"`
		AdminPagesPrefix string `toml:"admin_pages_prefix"`
		AdminApiPrefix   string `toml:"admin_api_prefix"`
		// special admin pages
		Admin            string `toml:"admin"` // not used.
		AdminSignUpPath  string `toml:"admin_sign_up_path"`
		AdminSignInPath  string `toml:"admin_sign_in_path"`
		AdminSignOutPath string `toml:"admin_sign_out_path"`
		AdminHomePath    string `toml:"admin_home_page"` // real path will be admin_pages_prefix + admin_home_page
	} `toml:"api"`
}

func init() {
	if _, err := toml.DecodeFile("conf/config.toml", &CustomConfig); err != nil {
		log.Fatalln(err)
		return
	}
	// post process of api section in configure
	// sign-in sign-out sign-up and admin-home all start with {AdminPagesPrefix}
	CustomConfig.Api.AdminSignUpPath = CustomConfig.Api.AdminPagesPrefix + CustomConfig.Api.AdminSignUpPath
	CustomConfig.Api.AdminSignInPath = CustomConfig.Api.AdminPagesPrefix + CustomConfig.Api.AdminSignInPath
	CustomConfig.Api.AdminSignOutPath = CustomConfig.Api.AdminPagesPrefix + CustomConfig.Api.AdminSignOutPath
	CustomConfig.Api.AdminHomePath = CustomConfig.Api.AdminPagesPrefix + CustomConfig.Api.AdminHomePath

	// post process of storage section in config
	if !CustomConfig.Storage.EnableQiNiuCloud {
		CustomConfig.Storage.LocalStorageUploadUrl = CustomConfig.Api.AdminApiPrefix + CustomConfig.Storage.LocalStorageUploadUrl
	}
}
