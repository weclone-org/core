package restfulapi

import "github.com/Weclone-org/core/config"

func RESTfulAPIInit() {
	r := setupRouter()
	err := r.Run(config.Conf.Listen)
	if err != nil {
		return
	}
}
