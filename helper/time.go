package helper

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	DefaultTimeFmt = "2006年01月02日 15:04:05"
)

func FmtTimeByTimeStamp(seconds int64, nanos int32, format string) string {
	tm := time.Unix(seconds, int64(nanos))
	return tm.Format(format)
}

func FmtTimeByProtoTime(t *timestamppb.Timestamp, format string) string {
	return FmtTimeByTimeStamp(t.Seconds, t.Nanos, format)
}
