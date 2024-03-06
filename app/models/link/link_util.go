package link

import (
	"go_learn/pkg/app"
	"go_learn/pkg/cache"
	"go_learn/pkg/database"
	"go_learn/pkg/helpers"
	"go_learn/pkg/paginator"
	"time"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (link Link) {
	database.DB.Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	database.DB.Where("? = ?", field, value).First(&link)
	return
}

func All() (links []Link) {
	database.DB.Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

func AllCached() (links []Link) {
	// 设置缓存 key
	cacheKey := "links:all"
	// 设置过期时间
	expireTime := 120 * time.Minute
	//取数据
	cache.GetObject(cacheKey, &links)
	if helpers.Empty(links) {
		// 查数据库
		links = All()
		if !helpers.Empty(links) {
			// 设置缓存
			cache.Set(cacheKey, links, expireTime)
		}
	}
	return
}
