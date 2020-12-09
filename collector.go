package stream

type ICollector interface {
	Feed(interface{})
}

type ArrayCollector struct {
	data []interface{}
}

func (collector * ArrayCollector) Feed(item interface{})  {
	collector.data = append(collector.data, item)
}

func (collector * ArrayCollector) ToArray() []interface{} {
	return collector.data
}