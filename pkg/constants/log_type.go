package constants

type LogType = uint

const (
	LogType_MemberBackfill LogType = iota
	LogType_MemberJoin
	LogType_MemberLeave
	LogType_AssignRole
	LogType_RevokeRole
	LogType_UpdateNickname
	LogType_AdminKick
	LogType_AdminBan
)
