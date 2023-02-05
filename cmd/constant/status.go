package constant

type Status int

const (
	Success Status = iota
	Failed
	TokenInvalid
)

func (s Status) GetInfo() (int32, string) {
	switch s {
	case Success:
		return 0, "成功"
	case Failed:
		return 1, "失败"
	case TokenInvalid:
		return 2, "token 验证失败"
	default:
		return -1, "未知状态"
	}
}
