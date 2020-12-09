package stream

type StageType int8

const (
	StageFilter = 0
	StageMap    = 1
)

type Stage struct {
	callback func(interface{}) interface{}
	stageType StageType
}

func MakeStage(callback func(interface{}) interface{}, stageType StageType) *Stage {
	return &Stage{
		callback: callback,
		stageType: stageType,
	}
}