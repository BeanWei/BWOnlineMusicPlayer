package controllers

import (
	ms "github.com/BeanWei/MusicSpider"
	"github.com/gin-gonic/gin"
	"strconv"
)

func MusicApiHandler(c *gin.Context) {
	types := c.PostForm("types")
	switch types {
	case "search":
		count := c.PostForm("count")
		source := c.PostForm("source")
		pages := c.PostForm("pages")
		name := c.PostForm("name")
		args := map[string]int{"page": str2intSafe(pages), "limit": str2intSafe(count)}
		r := ms.Search("netease", name, args)
		c.JSON(200, gin.H{"data": searchFormat(source, r["result"])})
	case "url":
		id := c.PostForm("id")
		source := c.PostForm("source")
		c.JSON(200, gin.H{"url": ms.Downloadurl(source, id)["url"]})
	case "pic":
		id := c.PostForm("id")
		source := c.PostForm("source")
		songinfo := songFormat(source, ms.Song(source, id)["result"])
		c.JSON(200, gin.H{"url": songinfo["cover_url"]})
	case "playlist":
		lid := c.PostForm("lid")
		source := c.PostForm("source")
		r := ms.Playlist(source, lid)
		c.JSON(200, gin.H{"data": playlistFormat(source, r["result"])})
	case "song":
		id := c.PostForm("id")
		source := c.PostForm("source")
		r := ms.Song(source, id)
		c.JSON(200, gin.H{"data": songFormat(source, r["result"])})
	case "lyric":
		id := c.PostForm("id")
		source := c.PostForm("source")
		c.JSON(200, gin.H{"data": ms.Lyric(source, id)["lyric"]})
	case "userlist":
		offset := c.PostForm("offset")
		limit := c.PostForm("limit")
		uid := c.PostForm("uid")
		source := c.PostForm("source")
		args := map[string]int{"page": str2intSafe(offset), "limit": str2intSafe(limit)}
		r := ms.UserPlaylist(source, uid, args)
		c.JSON(200, gin.H{"data": userPlaylistFormat(source, r["result"])})
	default:
	}
}

func str2intSafe(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0
	} else {
		return i
	}
}
