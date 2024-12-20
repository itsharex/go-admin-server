package dvroutes

import (
	"github.com/MarchGe/go-admin-server/app"
	"github.com/MarchGe/go-admin-server/app/admin/apis/devops"
	"github.com/MarchGe/go-admin-server/app/common/middleware/authz"
	"github.com/MarchGe/go-admin-server/app/common/middleware/recorder"
	"github.com/gin-gonic/gin"
)

func init() {
	app.RouterRegister(registerExplorerSftpRoutes)
}

func registerExplorerSftpRoutes(g *gin.RouterGroup) {
	a := devops.GetExplorerSftpApi()
	rg := g.Group("/devops/explorer/sftp")

	rg.DELETE("/entry", authz.RequiresPermissions("explorer_sftp:delete"), recorder.RecordOpLog("资源管理器资源（SFTP）"), a.DeleteEntry)
	rg.GET("/entries", authz.RequiresPermissions("explorer_sftp:entries"), a.GetEntries)
	rg.POST("/upload", authz.RequiresPermissions("explorer_sftp:upload"), recorder.RecordOpLog("资源管理器资源（SFTP）", "上传文件"), a.Upload)
	rg.GET("/download", authz.RequiresPermissions("explorer_sftp:download"), a.Download)
	rg.POST("/create", authz.RequiresPermissions("explorer_sftp:create"), recorder.RecordOpLog("资源管理器资源（SFTP）", "创建目录"), a.CreateDir)
	rg.POST("/rename", authz.RequiresPermissions("explorer_sftp:rename"), recorder.RecordOpLog("资源管理器资源（SFTP）", "重命名"), a.Rename)
}
