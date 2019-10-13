package object

type SyncType string

const (
	DTW SyncType = "DTW" // 日库到周库
	DTM SyncType = "DTM" // 日库到月库
	MTQ SyncType = "MTQ" // 月库到季库
	MTY SyncType = "MTY" // 月库到年库
)

func (this SyncType) ToString() string {
	return string(this)
}
