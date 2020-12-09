package stream

type IStream interface {
	Map(mapFunc func(interface{}) interface{}) IStream
	Filter(filterFunc func(interface{}) interface{}) IStream
	Collect(collector ...ICollector) IStream
	Run() error
}

type Stream struct
{
	stages []*Stage
	collectors []ICollector

	source func(func(interface{})) error
}

func (stream *Stream) Map(mapFunc func(interface{}) interface{}) IStream {
	stream.stages = append(stream.stages, MakeStage(mapFunc, StageMap))
	return stream
}

func (stream *Stream) Filter(filterFunc func(interface{}) interface{}) IStream {
	stream.stages = append(stream.stages, MakeStage(filterFunc, StageFilter))
	return stream
}

func (stream *Stream) Collect(collector ...ICollector) IStream {
	stream.collectors = append(stream.collectors, collector...)
	return stream
}

func (stream *Stream) Run() error {
	return stream.source(stream.feed)
}


func (stream *Stream) feed(data interface{}) {
	var (
		current = data
	)

	for _, stage := range stream.stages {
		if stage.stageType == StageFilter {
			if !stage.callback(current).(bool) {
				return
			}
		}
		if stage.stageType == StageMap {
			current = stage.callback(current)
		}
	}

	for _, collector := range stream.collectors {
		collector.Feed(current)
	}
}