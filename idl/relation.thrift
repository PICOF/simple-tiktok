namespace go relation

include "user.thrift"

struct FollowRequest {
    1: i64 user_id
    2: i64 to_user_id
    3: bool action_type
}

struct FollowResponse {
    1: i64 status_code
    2: string status_msg
}

struct RelationListRequest {
    1: i64 user_id
    2: i64 query_id
}

struct RelationListResponse {
    1: i64 status_code
    2: optional string status_msg
    3: optional list<user.UserInfo> user_list
}

service RelationService {
    FollowResponse FollowAction(1: FollowRequest request)
    RelationListResponse GetFollowList(1: RelationListRequest request)
    RelationListResponse GetFollowerList(1: RelationListRequest request)
    RelationListResponse GetFriendList(1: RelationListRequest request)
}