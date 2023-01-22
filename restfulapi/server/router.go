package server

import (
	"github.com/gin-gonic/gin"
)

func ServerRouterInit(r *gin.Engine) {
	// List of Server APIs
	r.GET("/api/info", getInfo)

	// List of Server APIs
	r.GET("/api/tag/add/:tagname", addTag)
	r.GET("/api/index/get/uuid/:indexuuid", getIndexByUUID)
	r.GET("/api/index/getlist/tag/:tagname", getIndexListByTag)
	r.GET("/api/index/getlist/short", getShortIndexList)
	r.POST("/api/index/update", updateIndex)

	// Client User API List
	r.POST("/api/user/login", loginUser)
	r.POST("/api/user/logout", logoutUser)
	r.GET("/api/user/profile", getUserProfile)
	r.POST("/api/user/profile/update", updateUserProfile)
	r.POST("/api/user/upload", uploadFile)
	r.GET("/api/user/file/public", getPublicFiles)
	r.GET("/api/user/file/private", getPrivateFiles)
	r.GET("/api/user/cdn/list", getCDNNode)
	r.GET("/api/user/root/list", getRootNode)
	r.GET("/api/user/server/list", getServerNode)
}

// Server API Functions
// getInfo View the basic information of the current node
func getInfo(c *gin.Context) {
	// ...
}

// addTag add tag
func addTag(c *gin.Context) {
	// ...
}

// getIndexByUUID
func getIndexByUUID(c *gin.Context) {
	// ...
}

// getIndexListByTag
func getIndexListByTag(c *gin.Context) {
	// ...
}

// getShortIndexList
func getShortIndexList(c *gin.Context) {
	// ...
}

// updateIndex
func updateIndex(c *gin.Context) {
	// ...
}

// loginUser
func loginUser(c *gin.Context) {
	// ...
}

// logoutUser
func logoutUser(c *gin.Context) {
	// ...
}

// getUserProfile
func getUserProfile(c *gin.Context) {
	// ...
}

// updateUserProfile
func updateUserProfile(c *gin.Context) {
	// ...
}

// uploadFile
func uploadFile(c *gin.Context) {
	// ...
}

// getPublicFiles
func getPublicFiles(c *gin.Context) {
	// ...
}

// getPrivateFiles
func getPrivateFiles(c *gin.Context) {
	// ...
}

// getCDNNode
func getCDNNode(c *gin.Context) {
	// ...
}

// getRootNode
func getRootNode(c *gin.Context) {
	// ...
}

// getServerNode
func getServerNode(c *gin.Context) {
	// ...
}
