namespace go user

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
    2: string password
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