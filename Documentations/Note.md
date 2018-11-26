api: localhost:8080/api/v1
meathod: post

/* 搜索 */
data: 
    {
        types: search
        count:   搜索结果一次加载多少条
        source:  来源
        pages:   当前页  
        name:    关键字   
    }

返回数据：
    [
        {
            song_id:             音乐ID        
            song_name:           音乐名称
            artist_name:         艺术家名字
            album_name:          专辑名字
            source:         音乐来源
            url_id:         链接ID
            pic_id:         封面ID
            lyric_id:       歌词ID
            pic:   null         专辑图片
            url:   null         mp3链接
        },
        ...   
    ] 


/* 直链 */
data:
    {
        types: url
        id:
        source:
    }

返回数据:
    {
        url: 
    }    


/* 封面图 */
data:
    {
        types: pic
        id:
        source:
    }

返回数据:
    {
        url: 
    }   


/* 歌单(默认站点为网易云) */
data:
    {
        types: playlist
        lid:
        source: "netease"
    }

返回数据:
data:    {
        source: 来源站点
        name: 列表名
        coverImgUrl:  列表封面
        creator_nickname: 列表创建者 
        creator_avatarUrl: 列表创建者头像
        brief_desc: 专辑简述
        tracks: [
            song_id:
            song_name:
            artist_name:
            album_name:
            url_id:
            pic_id: null  //封面ID
            lyric_id:
            album_coverImgUrl: null?
            url: null   // mp3链接
        ]
    }


/* 歌词 */
data:
    {
        types: lyric
        id:
        source:
    }

返回数据:
    {
        lyric: 
    }  


/* 用户播放列表(默认站点为网易云) */ 
data: 
    {
        types: userlist
        uid:
    }

返回数据:
data:{  
        code: 
        playlist: [
            {
                id:
                name:
                coverImgUrl:
                update_time:
                track_count:
                play_count:
                creator_nickname:
                creator_avatarUrl:
            },
            ...
        ]
    }    