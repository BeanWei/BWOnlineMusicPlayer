package controllers

import (
	"fmt"
	"github.com/thedevsaddam/gojsonq"
	"strconv"
	"strings"

	ms "github.com/BeanWei/MusicSpider"
)

/*
基于 github.com/BeanWei/MusicSpider 下的 format.go, 根据此项目的需求进行第二次开发
*/

/* 辅助工具函数 */

func safeParse(jsonStr, path string) string {
	strValue := ""
	result := gojsonq.New().JSONString(jsonStr).Find(path)
	if result == nil || result == "" {
		return strValue
	}
	textType := fmt.Sprintf("%T", result)
	switch textType {
	case "string":
		strValue = result.(string)
	case "int":
		strValue = strconv.Itoa(result.(int))
	case "int64":
		strValue = strconv.FormatInt(result.(int64), 10)
	case "float32", "float64":
		strValue = fmt.Sprintf("%.0f", result)
	default:
		fmt.Println("此类型暂不支持转换为Str类型: ", textType)
	}
	return strValue
}

func getNeteasePicURL(id string) string {
	//http://p3.music.126.net/[encrypted_song_id]/[song_dfsId].jpg
	if id == "" {
		return ""
	}
	encrypted_song_id := ms.NeteaseEncryptId(id)
	return fmt.Sprintf("http://p3.music.126.net/%s/%s.jpg", encrypted_song_id, id)
}

/*===================Search Songs API Response Format==================================*/
func searchFormat(site, jsonStr string) []interface{} {
	shslist := []interface{}{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return nil
	}
	switch site {
	case "netease":
		for i := 0; i >= 0; i++ {
			shs := map[string]interface{}{}
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].id", i))
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].name", i))
			shs["singer_id"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].artists.[0].id", i))
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].artists.[0].name", i))
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].album.id", i))
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("result.songs.[%d].album.name", i))
			shs["source"] = "netease"
			shs["url_id"] = shs["song_id"]
			shs["pic_id"] = shs["song_id"]
			shs["lyric_id"] = shs["song_id"]
			shs["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", shs["song_id"])
			shslist = append(shslist, shs)
		}
	case "tencent":
		for i := 0; i >= 0; i++ {
			shs := map[string]interface{}{}
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].mid", i))
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].title", i))
			shs["singer_id"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].singer.[0].id", i))
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].singer.[0].name", i))
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].album.mid", i))
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("data.song.list.[%d].album.name", i))
			shs["source"] = "tencent"
			shs["url_id"] = shs["song_id"]
			shs["pic_id"] = shs["song_id"]
			shs["lyric_id"] = shs["song_id"]
			shs["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", shs["song_id"])
			shslist = append(shslist, shs)
			fmt.Println("song list is: ", shslist, "endStr")
		}
	case "xiami":
		for i := 0; i >= 0; i++ {
			shs := map[string]interface{}{}
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].song_id", i))
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].song_name", i))
			shs["singer_id"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].artist_id", i))
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].artist_name", i))
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].album_id", i))
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("data.songs.[%d].album_name", i))
			shs["source"] = "xiami"
			shs["url_id"] = shs["song_id"]
			shs["pic_id"] = shs["song_id"]
			shs["lyric_id"] = shs["song_id"]
			shs["source_url"] = fmt.Sprintf("https://www.xiami.com/song/%s", shs["song_id"])
			shslist = append(shslist, shs)
		}
	case "kugou":
		for i := 0; i >= 0; i++ {
			shs := map[string]interface{}{}
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].hash", i))
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].songname", i))
			shs["singer_id"] = ""
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].singername", i))
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].album_id", i))
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].album_name", i))
			shs["source"] = "kugou"
			shs["url_id"] = shs["song_id"]
			shs["pic_id"] = shs["song_id"]
			shs["lyric_id"] = shs["song_id"]
			shs["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", shs["song_id"])
			shslist = append(shslist, shs)
		}
	case "baidu":
		for i := 0; i >= 0; i++ {
			shs := map[string]interface{}{}
			shs["song_id"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].song_id", i))
			if shs["song_id"] == "" {
				break
			}
			shs["song_name"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].title", i))
			shs["singer_id"] = ""
			shs["singer_name"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].author", i))
			shs["album_id"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].album_id", i))
			shs["album_name"] = safeParse(jsonStr, fmt.Sprintf("result.songinfo.song_list.[%d].album_title", i))
			shs["source"] = "baidu"
			shs["url_id"] = shs["song_id"]
			shs["pic_id"] = shs["song_id"]
			shs["lyric_id"] = shs["song_id"]
			shs["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", shs["song_id"])
			shslist = append(shslist, shs)
		}
	default:
	}
	return shslist
}

/*===================Album Songs API Response Format==================================*/
func albumFormat(site, jsonStr string) map[string]interface{} {
	songslist, album := []interface{}{}, map[string]interface{}{}
	album["source"], album["album_name"], album["singer_name"], album["singer_id"], album["brief_desc"], album["publish_time"], album["cover_url"], album["total_songs"], album["songs"] = "", "", "", "", "", "", "", 0, nil

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return album
	}
	switch site {
	case "netease":
		album["source"] = "netease"
		album["album_name"] = safeParse(jsonStr, "album.name")
		album["singer_name"] = safeParse(jsonStr, "album.artist.[0].name")
		album["singer_id"] = safeParse(jsonStr, "album.artist.[0].id")
		album["brief_desc"] = safeParse(jsonStr, "album.briefDesc")
		album["publish_time"] = safeParse(jsonStr, "album.publishTime")
		album["cover_url"] = safeParse(jsonStr, "album.blurPicUrl")
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("songs.[%d].id", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("songs.[%d].name", i))
			songs["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	case "tencent":
		album["source"] = "tencent"
		album["album_name"] = safeParse(jsonStr, "data.getAlbumInfo.Falbum_name")
		album["singer_name"] = safeParse(jsonStr, "data.singerInfo.[0].Fsinger_name")
		album["singer_id"] = safeParse(jsonStr, "data.singerInfo.[0].Fsinger_mid")
		album["brief_desc"] = safeParse(jsonStr, "data.getAlbumDesc.Falbum_desc")
		album["publish_time"] = safeParse(jsonStr, "data.getAlbumInfo.Fpublic_time")
		album["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.getSongInfo.[%d].mid", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.getSongInfo.[%d].name", i))
			songs["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	case "xiami":
		album["source"] = "xiami"
		album["album_name"] = safeParse(jsonStr, "data.[0].album_name")
		album["singer_name"] = safeParse(jsonStr, "data.[0].artist_name")
		album["singer_id"] = safeParse(jsonStr, "data.[0].artist_id")
		album["brief_desc"] = ""
		album["publish_time"] = ""
		// TODO: 获取虾米音乐专辑高清封面
		album["cover_url"] = "https://pic.xiami.net/" + safeParse(jsonStr, "data.[0].album_logo")
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.[%d].songId", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("songs.[%d].songName", i))
			songs["source_url"] = fmt.Sprintf("https://www.xiami.com/song/%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	case "kugou":
		album["source"] = "kugou"
		album["album_name"] = ""
		album["singer_name"] = strings.Split(safeParse(jsonStr, "data.info.[0].filename"), " - ")[0]
		album["singer_id"] = ""
		album["brief_desc"] = ""
		album["publish_time"] = ""
		album["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].hash", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].filename", i))
			songs["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	case "baidu":
		album["source"] = "baidu"
		album["album_name"] = safeParse(jsonStr, "albumInfo.title")
		album["singer_name"] = safeParse(jsonStr, "albumInfo.author")
		album["singer_id"] = safeParse(jsonStr, "albumInfo.artist_id")
		album["brief_desc"] = safeParse(jsonStr, "albumInfo.info")
		album["publish_time"] = safeParse(jsonStr, "albumInfo.publishtime")
		album["cover_url"] = safeParse(jsonStr, "albumInfo.pic_s1000")
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("songlist.[%d].song_id", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("songlist.[%d].title", i))
			songs["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		album["total_songs"], album["songs"] = len(songslist), songslist
	default:
	}
	return album
}

/*===================Artist Songs API Response Format==================================*/
func artistFormat(site, jsonStr string) map[string]interface{} {
	songslist, artist := []interface{}{}, map[string]interface{}{}
	artist["source"], artist["singer_name"], artist["brief_desc"], artist["cover_url"], artist["total_songs"], artist["songs"] = "", "", "", "", 0, nil

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return artist
	}
	switch site {
	case "netease":
		artist["source"] = "netease"
		artist["singer_name"] = safeParse(jsonStr, "artist.name")
		artist["brief_desc"] = safeParse(jsonStr, "artist.briefDesc")
		artist["cover_url"] = safeParse(jsonStr, "artist.picUrl")
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("hotSongs.[%d].id", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("hotSongs.[%d].name", i))
			songs["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		artist["total_songs"], artist["songs"] = len(songslist), songslist
	case "tencent":
		artist["source"] = "tencent"
		artist["singer_name"] = safeParse(jsonStr, "data.singerInfo.[0].Fsinger_name")
		artist["brief_desc"] = safeParse(jsonStr, "data.getAlbumDesc.Falbum_desc")
		artist["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.list.[%d].musicData.mid", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.getSongInfo.[%d].musicData.name", i))
			songs["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", songs["song_id"])
			songslist = append(songslist, songs)
		}
		artist["total_songs"], artist["songs"] = len(songslist), songslist
	case "xiami":
		// TODO: 虾米艺人API接口待完善
		return artist
	case "kugou":
		artist["source"] = "kugou"
		artist["singer_name"] = strings.Split(safeParse(jsonStr, "data.info.[0].filename"), " - ")[0]
		artist["brief_desc"] = ""
		artist["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].hash", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].filename", i))
			songs["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		artist["total_songs"], artist["songs"] = len(songslist), songslist
	case "baidu":
		artist["source"] = "baidu"
		artist["singer_name"] = safeParse(jsonStr, "songlist.[0].author")
		artist["brief_desc"] = ""
		artist["cover_url"] = ""
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("songlist.[%d].song_id", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("songlist.[%d].title", i))
			songs["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		artist["total_songs"], artist["songs"] = len(songslist), songslist
	default:
	}
	return artist
}

/*===================Playlist Songs API Response Format==================================*/
func playlistFormat(site, jsonStr string) map[string]interface{} {
	songslist, playlist := []interface{}{}, map[string]interface{}{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return nil
	}
	switch site {
	case "netease":
		playlist["source"] = "netease"
		playlist["name"] = safeParse(jsonStr, "playlist.name")
		playlist["coverImgUrl"] = safeParse(jsonStr, "playlist.coverImgUrl")
		playlist["creator_nickname"] = safeParse(jsonStr, "playlist.creator.nickname")
		playlist["creator_avatarUrl"] = safeParse(jsonStr, "playlist.creator.avatarUrl")
		playlist["brief_desc"] = safeParse(jsonStr, "playlist.description")
		tagsStr := ""
		for i := 0; i >= 0; i++ {
			tag := safeParse(jsonStr, fmt.Sprintf("playlist.tags.[%d]", i))
			if tag == "" {
				break
			}
			tagsStr += tag + ","
		}
		playlist["tags"] = strings.TrimRight(tagsStr, ",")
		playlist["creat_time"] = safeParse(jsonStr, "playlist.createTime")
		playlist["play_count"] = safeParse(jsonStr, "playlist.playCount")
		playlist["subscribed_count"] = safeParse(jsonStr, "playlist.subscribedCount")
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].id", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].name", i))
			songs["artist_id"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].ar.[0].id", i))
			songs["artist_name"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].ar.[0].name", i))
			songs["album_id"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].al.id", i))
			songs["album_name"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].al.name", i))
			songs["url_id"] = songs["song_id"]
			songs["lyric_id"] = songs["song_id"]
			songs["pic_id"] = songs["song_id"]
			songs["song_coverImgUrl"] = getNeteasePicURL(songs["song_id"].(string))
			songs["album_coverImgUrl"] = safeParse(jsonStr, fmt.Sprintf("playlist.tracks.[%d].al.picUrl", i))
			songs["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		playlist["total_songs"], playlist["tracks"] = len(songslist), songslist
	case "tencent":
		playlist["source"] = "tencent"
		playlist["name"] = safeParse(jsonStr, "data.cdlist.dissname")
		playlist["coverImgUrl"] = safeParse(jsonStr, "data.cdlist.logo")
		playlist["creator_nickname"] = safeParse(jsonStr, "data.cdlist.nickname")
		playlist["creator_avatarUrl"] = safeParse(jsonStr, "data.cdlist.headurl")
		playlist["brief_desc"] = safeParse(jsonStr, "data.cdlist.desc")
		tagsStr := ""
		for i := 0; i >= 0; i++ {
			tag := safeParse(jsonStr, fmt.Sprintf("data.cdlist.tags.[%d].name", i))
			if tag == "" {
				break
			}
			tagsStr += tag + ","
		}
		playlist["tags"] = strings.TrimRight(tagsStr, ",")
		playlist["creat_time"] = safeParse(jsonStr, "data.cdlist.ctime")
		playlist["play_count"] = safeParse(jsonStr, "data.cdlist.visitnum")
		playlist["subscribed_count"] = ""
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].mid", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].name", i))
			songs["artist_id"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].singer.mid", i))
			songs["artist_name"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].singer.name", i))
			songs["album_id"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].album.id", i))
			songs["album_name"] = safeParse(jsonStr, fmt.Sprintf("data.cdlist.songlist.[%d].album.name", i))
			songs["url_id"] = songs["song_id"]
			songs["pic_id"] = songs["song_id"]
			songs["lyric_id"] = songs["song_id"]
			songs["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", songs["song_id"])
			songslist = append(songslist, songs)
		}
		playlist["total_songs"], playlist["tracks"] = len(songslist), songslist
	case "xiami":
		// TODO: 虾米歌单API接口待完善
		return nil
	case "kugou":
		playlist["source"] = "kugou"
		playlist["name"] = ""
		playlist["coverImgUrl"] = ""
		playlist["creator_nickname"] = ""
		playlist["creator_avatarUrl"] = ""
		playlist["brief_desc"] = ""
		playlist["tags"] = ""
		playlist["creat_time"] = ""
		playlist["play_count"] = ""
		playlist["subscribed_count"] = ""
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].hash", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].filename", i))
			songs["artist_id"] = ""
			songs["artist_name"] = strings.Split(songs["song_name"].(string), " - ")[0]
			songs["album_id"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].album_id", i))
			songs["album_name"] = safeParse(jsonStr, fmt.Sprintf("data.info.[%d].remark", i))
			songs["url_id"] = songs["song_id"]
			songs["pic_id"] = songs["song_id"]
			songs["lyric_id"] = songs["song_id"]
			songs["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		playlist["total_songs"], playlist["tracks"] = len(songslist), songslist
	case "baidu":
		playlist["source"] = "baidu"
		playlist["name"] = safeParse(jsonStr, "title")
		playlist["coverImgUrl"] = safeParse(jsonStr, "pic_w700")
		playlist["creator_nickname"] = ""
		playlist["creator_avatarUrl"] = ""
		playlist["brief_desc"] = safeParse(jsonStr, "desc")
		playlist["tags"] = safeParse(jsonStr, "tag")
		playlist["creat_time"] = ""
		playlist["play_count"] = safeParse(jsonStr, "listenum")
		playlist["subscribed_count"] = safeParse(jsonStr, "collectnum")
		for i := 0; i >= 0; i++ {
			songs := map[string]interface{}{}
			songs["song_id"] = safeParse(jsonStr, fmt.Sprintf("content.[%d].song_id", i))
			if songs["song_id"] == "" {
				break
			}
			songs["song_name"] = safeParse(jsonStr, fmt.Sprintf("content.[%d].title", i))
			songs["artist_id"] = ""
			songs["artist_name"] = safeParse(jsonStr, fmt.Sprintf("content.[%d].author", i))
			songs["album_id"] = safeParse(jsonStr, fmt.Sprintf("content.[%d].album_id", i))
			songs["album_name"] = safeParse(jsonStr, fmt.Sprintf("content.[%d].album_title", i))
			songs["url_id"] = songs["song_id"]
			songs["pic_id"] = songs["song_id"]
			songs["lyric_id"] = songs["song_id"]
			songs["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", songs["song_id"])
			songslist = append(songslist, songs)
		}
		playlist["total_songs"], playlist["tracks"] = len(songslist), songslist
	default:
	}
	return playlist
}

/*===================Song Info API Response Format==================================*/
func songFormat(site, jsonStr string) map[string]interface{} {
	songinfo := map[string]interface{}{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return nil
	}
	switch site {
	case "netease":
		songinfo["source"] = "netease"
		songinfo["song_name"] = safeParse(jsonStr, "songs.[0].name")
		songinfo["singer_id"] = safeParse(jsonStr, "songs.[0].ar.[0].id")
		songinfo["singer_name"] = safeParse(jsonStr, "songs.[0].ar.[0].name")
		songinfo["album_id"] = safeParse(jsonStr, "songs.[0].al.id")
		songinfo["album_name"] = safeParse(jsonStr, "songs.[0].al.name")
		songinfo["publish_time"] = safeParse(jsonStr, "songs.[0].publishTime")
		//songinfo["song_coverImgUrl"]  = getNeteasePicURL(safeParse(jsonStr, "songs.[0].id"))
		songinfo["cover_url"] = safeParse(jsonStr, "songs.[0].al.picUrl")
		songinfo["source_url"] = fmt.Sprintf("https://music.163.com/#/song?id=%s", safeParse(jsonStr, "songs.[0].id"))

	case "tencent":
		songinfo["source"] = "tencent"
		songinfo["song_name"] = safeParse(jsonStr, "data.[0].name")
		songinfo["singer_id"] = safeParse(jsonStr, "data.singer.[0].id")
		songinfo["singer_name"] = safeParse(jsonStr, "data.singer.[0].name")
		songinfo["album_id"] = safeParse(jsonStr, "data.album.id")
		songinfo["album_name"] = safeParse(jsonStr, "data.album.name")
		songinfo["publish_time"] = safeParse(jsonStr, "data.[0].time_public")
		songinfo["cover_url"] = fmt.Sprintf("https://y.gtimg.cn/music/photo_new/T002R300x300M000%s.jpg?max_age=2592000", safeParse(jsonStr, "data.[0].mid"))
		songinfo["source_url"] = fmt.Sprintf("https://y.qq.com/n/yqq/song/%s.html", safeParse(jsonStr, "data.[0].mid"))
	case "xiami":
		songinfo["source"] = "xiami"
		songinfo["song_name"] = safeParse(jsonStr, "data.trackList.[0].songName")
		songinfo["singer_id"] = safeParse(jsonStr, "data.trackList.[0].artistId")
		songinfo["singer_name"] = safeParse(jsonStr, "data.trackList.[0].singers")
		songinfo["album_id"] = safeParse(jsonStr, "data.trackList.[0].album_id")
		songinfo["album_name"] = safeParse(jsonStr, "data.trackList.[0].album_name")
		songinfo["publish_time"] = safeParse(jsonStr, "data.trackList.[0].demoCreateTime")
		songinfo["cover_url"] = safeParse(jsonStr, "data.trackList.[0].pic")
		songinfo["source_url"] = fmt.Sprintf("https://www.xiami.com/song/%s", safeParse(jsonStr, "data.trackList.[0].songId"))
	case "kugou":
		songinfo["source"] = "kugou"
		songinfo["song_name"] = safeParse(jsonStr, "songName")
		songinfo["singer_id"] = safeParse(jsonStr, "singerId")
		songinfo["singer_name"] = safeParse(jsonStr, "choricSinger")
		songinfo["album_id"] = safeParse(jsonStr, "albumid")
		songinfo["album_name"] = ""
		songinfo["publish_time"] = ""
		songinfo["cover_url"] = safeParse(jsonStr, "imgUrl")
		songinfo["source_url"] = fmt.Sprintf("http://www.kugou.com/song/#hash=%s", safeParse(jsonStr, "hash"))
	case "baidu":
		songinfo["source"] = "baidu"
		songinfo["song_name"] = safeParse(jsonStr, "songinfo.title")
		songinfo["singer_id"] = safeParse(jsonStr, "songinfo.artist_id")
		songinfo["singer_name"] = safeParse(jsonStr, "songinfo.author")
		songinfo["album_id"] = safeParse(jsonStr, "songinfo.album_id")
		songinfo["album_name"] = safeParse(jsonStr, "songinfo.album_title")
		songinfo["publish_time"] = safeParse(jsonStr, "songinfo.publishtime")
		songinfo["cover_url"] = safeParse(jsonStr, "songinfo.pic_huge")
		songinfo["source_url"] = fmt.Sprintf("http://music.taihe.com/song/%s", safeParse(jsonStr, "songinfo.song_id"))
	default:
	}
	return songinfo
}

/*===================UserPlaylist Info API Response Format==================================*/
func userPlaylistFormat(site, jsonStr string) map[string]interface{} {
	play, playlist := []interface{}{}, map[string]interface{}{}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if site == "" || jsonStr == "" {
		return nil
	}
	switch site {
	case "netease":
		for i := 0; i >= 0; i++ {
			item := map[string]interface{}{}
			item["id"] = safeParse(jsonStr, fmt.Sprintf("playlist.[%d].id", i))
			item["name"] = safeParse(jsonStr, fmt.Sprintf("playlist.[%d].id", i))
			item["coverImgUrl"] = safeParse(jsonStr, fmt.Sprintf("playlist.[%d].coverImgUrl", i))
			item["update_time"] = safeParse(jsonStr, fmt.Sprintf("playlist.[%d].updateTime", i))
			item["track_count"] = safeParse(jsonStr, fmt.Sprintf("playlist.[%d].trackCount", i))
			item["play_count"] = safeParse(jsonStr, fmt.Sprintf("playlist.[%d].playCount", i))
			item["creator_nickname"] = safeParse(jsonStr, fmt.Sprintf("playlist.[%d].creator.nickname", i))
			item["creator_avatarUrl"] = safeParse(jsonStr, fmt.Sprintf("playlist.[%d].creator.avatarUrl", i))
			play = append(play, item)
		}
		playlist["code"], playlist["playlist"] = safeParse(jsonStr, "code"), play
	default:
		//TODO: 支持其他站点
		return nil
	}
	return playlist
}
