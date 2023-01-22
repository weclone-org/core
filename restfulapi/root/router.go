package root

import (
	"github.com/gin-gonic/gin"
)

func RootRouterInit(r *gin.Engine) {
	//List of Root Node APIs
	r.GET("/api/info", getInfo)
	r.GET("/api/root/network/info", getNetworkInfo)
	r.GET("/api/root/node-list", getNodeList)
	r.GET("/api/root/node-info/:nodeuuid", getNodeInfo)
	r.GET("/api/root/login/:serveruuid", login)
	r.GET("/api/root/logout/:serveruuid", logout)
	r.GET("/api/root/report/:serveruuid", report)
	r.POST("/api/root/register", register)
	r.POST("/api/root/verify", verify)

	//索引API列表
	r.GET("/api/tag/add/:tagname", addTag)
	r.GET("/api/index/get/uuid/:indexuuid", getIndexByUUID)
	r.GET("/api/index/getlist/tag/:tagname", getIndexListByTag)
	r.GET("/api/index/getlist/short", getShortIndexList)
	r.POST("/api/index/update", updateIndex)
}

// Root API Functions
// getInfo View the basic information of the current node
func getInfo(c *gin.Context) {
	// ...
}

// getNetworkInfo Get information about the network
func getNetworkInfo(c *gin.Context) {
	// ...
}

// getNodeList Current node list (including server, root, cdn)
func getNodeList(c *gin.Context) {
	// ...
}

// getNodeInfo Query the information of a node (basic information, online status)
func getNodeInfo(c *gin.Context) {
	// ...
}

// login Node
func login(c *gin.Context) {
	// ...
}

// logout Node
func logout(c *gin.Context) {
	// ...
}

// report node
func report(c *gin.Context) {
	// ...
}

// register node
func register(c *gin.Context) {
	// ...
}

// verify that the network is trustworthy
func verify(c *gin.Context) {
	// ...
}

// Index API Functions
// addTag
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
