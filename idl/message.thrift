namespace go message

struct MessageInfo {
    1: i64 id
    2: string content
    3: i64 create_time
}

struct MessageRequest {
    1: i64 user_id
    2: i64 to_user_id
    3: bool action_type
    4: string content
    5: i64 send_time
}

struct MessageResponse {
    1: i64 status_code
    2: string status_msg
}

struct ChatRecordRequest {
    1: i64 user_id
    2: i64 to_user_id
    3: i64 latest_time
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