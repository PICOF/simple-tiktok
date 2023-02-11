namespace go relation

include "user.thrift"

struct FollowRequest {
    1: string token (api.query="token")
    2: string to_user_id
    3: string action_type
}

struct FollowResponse {
    1: i64 status_code
    2: string status_msg
}

struct RelationListRequest {
    1: string user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct RelationListResponse {
    1: string status_code
    2: optional string status_msg
    3: optional list<user.UserInfo> user_list
}

service RelationService {
    FollowResponse FollowAction(1: FollowRequest request) (api.post="/douyin/relation/action/")
    RelationListResponse GetFollowList(1: RelationListRequest request) (api.get="/douyin/relation/follow/list/")
    RelationListResponse GetFollowerList(1: RelationListRequest request) (api.get="/douyin/relation/follower/list/")
    RelationListResponse GetFriendList(1: RelationListRequest request) (api.get="/douyin/relation/friend/list/")
}