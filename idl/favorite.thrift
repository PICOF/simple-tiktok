namespace go favorite

include "feed.thrift"

struct LikeRequest {
    1: string token (api.query="token")
    2: string video_id (api.query="video_id")
    3: string action_type (api.query="action_type")
}

struct LikeResponse {
    1: i64 status_code
    2: string status_msg
}

struct LikeListRequest {
    1: string user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct LikeListResponse {
    1: i64 status_code
    2: optional string status_msg
    4: optional list<feed.VideoInfo> video_list
}

service FavoriteService {
    LikeResponse LikeAction(1: LikeRequest request) (api.post="/douyin/favorite/action/")
    LikeListResponse GetLikeList(1: LikeRequest request) (api.get="/douyin/favorite/list/")
}