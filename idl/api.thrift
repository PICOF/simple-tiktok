namespace go tiktokapi

struct VideoInfo {
    1: i64 id
    2: UserInfo author
    3: string play_url
    4: string cover_url
    5: i64 favorite_count
    6: i64 comment_count
    7: bool is_favorite
    8: string title
}

struct FeedRequest{
    1: optional string latest_time
    2: optional string token
}

struct FeedResponse{
    1: i32 status_code
    2: optional string status_msg
    3: optional i64 next_time
    4: optional list<VideoInfo> video_list
}

service FeedService{
    FeedResponse GetVideoList(1: FeedRequest request) (api.get="/douyin/feed/")
}

struct UserInfo{
    1: i64 id
    2: string name
    3: i64 follow_count
    4: i64 follower_count
    5: bool is_follow
}

struct RegisterRequest {
    1: string username (api.query="username")
    2: string password (api.query="password")
}

struct RegisterResponse {
    1: i64 status_code
    2: string status_msg
    3: i64 user_id
    4: string token
}

struct LoginRequest {
    1: string username (api.query="username")
    2: string password (api.query="password")
}

struct LoginResponse {
    1: i64 status_code
    2: optional string status_msg
    3: optional i64 user_id
    4: optional string token
}

struct UserInfoRequest {
    1: string user_id (api.query="user_id")
    2: string token (api.query="token")
}

struct UserInfoResponse {
    1: i64 status_code
    2: optional string status_msg
    3: optional UserInfo user
}

service UserService {
    RegisterResponse UserRegister(1: RegisterRequest request)(api.post="/douyin/user/register/")
    LoginResponse UserLogin(1: LoginRequest request)(api.post="/douyin/user/login/")
    UserInfoResponse GetUserInfo(1: UserInfoRequest request)(api.get="/douyin/user/")
}

struct PublishRequest {
    1: binary data (api.form="date")
    2: string token (api.form="token")
    3: string title (api.form="title")
}

struct PublishResponse {
    1: i64 status_code
    2: optional string status_msg
}

struct PublishListRequest {
    1: string token (api.query="token")
    2: string user_id (api.query="user_id")
}

struct PublishListResponse {
    1: i64 status_code
    2: optional string status_msg
    4: optional list<VideoInfo> video_list
}

service PublishService {
    PublishResponse PublishAction(1: PublishRequest request) (api.post="/douyin/publish/action/")
    PublishListResponse GetPublishList(1: PublishListRequest request) (api.get="/douyin/publish/list/")
}

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
    4: optional list<VideoInfo> video_list
}

service FavoriteService {
    LikeResponse LikeAction(1: LikeRequest request) (api.post="/douyin/favorite/action/")
    LikeListResponse GetLikeList(1: LikeListRequest request) (api.get="/douyin/favorite/list/")
}

struct CommentInfo {
    1: i64 id
    2: UserInfo user
    3: string content
    4: string create_date
}

struct CommentRequest {
    1: string token (api.query="token")
    2: string video_id (api.query="video_id")
    3: string action_type (api.query="action_type")
    4: optional string comment_text (api.query="comment_text")
    5: optional string comment_id (api.query="comment_id")
}

struct CommentResponse {
    1: i64 status_code
    2: optional string status_msg
    3: optional CommentInfo comment
}

struct CommentListRequest {
    1: string token (api.query="token")
    2: string video_id (api.query="video_id")
}

struct CommentListResponse {
    1: i64 status_code
    2: optional string status_msg
    4: optional list<CommentInfo> comment_list
}

service CommentService {
    CommentResponse CommentAction(1: CommentRequest request) (api.post="/douyin/comment/action/")
    CommentListResponse GetCommentList(1: CommentListRequest request) (api.get="/douyin/comment/list/")
}

struct FollowRequest {
    1: string token (api.query="token")
    2: string to_user_id (api.query="to_user_id")
    3: string action_type (api.query="action_type")
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
    3: optional list<UserInfo> user_list
}

service RelationService {
    FollowResponse FollowAction(1: FollowRequest request) (api.post="/douyin/relation/action/")
    RelationListResponse GetFollowList(1: RelationListRequest request) (api.get="/douyin/relation/follow/list/")
    RelationListResponse GetFollowerList(1: RelationListRequest request) (api.get="/douyin/relation/follower/list/")
    RelationListResponse GetFriendList(1: RelationListRequest request) (api.get="/douyin/relation/friend/list/")
}

struct MessageInfo {
    1: i64 id
    2: string content
    3: string create_time
}

struct MessageRequest {
    1: string token (api.query="token")
    2: string to_user_id (api.query="to_user_id")
    3: string action_type (api.query="action_type")
    4: string content (api.query="content")
}

struct MessageResponse {
    1: i64 status_code
    2: string status_msg
}

struct ChatRecordRequest {
    1: string token (api.query="token")
    2: string to_user_id (api.query="to_user_id")
}

struct ChatRecordResponse {
    1: i64 status_code
    2: optional string status_msg
    3: optional list<MessageInfo> message_list
}

service MessageService {
    MessageResponse SendMessage(1: MessageRequest request) (api.post="/douyin/message/action/")
    ChatRecordResponse GetChatRecord(1: ChatRecordRequest request) (api.get="/douyin/message/chat/")
}
