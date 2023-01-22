package cdn

import (
	"github.com/gin-gonic/gin"
)

func CDNRouterInit(r *gin.Engine) {
	// List of CDN Node APIs
	r.GET("/api/info", getInfo)
	r.GET("/api/cdn/status", getCDNStatus)
	r.POST("/api/cdn/file/add", addFile)
	r.GET("/api/cdn/file/get/:indexuuid", getFile)
	r.GET("/api/cdn/file/delete", deleteFile)
}

// CDN API Functions
// getInfo View the basic information of the current node
func getInfo(c *gin.Context) {
	// ...
}

// getCDNStatus View system usage, disk space usage, running time, etc.
func getCDNStatus(c *gin.Context) {
	// ...
}

// addFile Add a file to the CDN node (and index, user binding)
func addFile(c *gin.Context) {
	// ...
}

// getFile Download file
func getFile(c *gin.Context) {
	// ...
}

// deleteFile Delete Files
func deleteFile(c *gin.Context) {
	// ...
}
