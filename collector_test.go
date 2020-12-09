package stream

import (
	"fmt"
	"strings"
	"testing"
)

func TestArrayCollector_ToArray(t *testing.T) {

	var (
		err error
		before = []string{"     Kemana   ", "Apa  ", "  siapa  ", "  dimana "}
		expected = []string{"Aku Kemana", "Aku siapa", "Aku dimana"}
		inStream *Stream
		collectorOne *ArrayCollector
		collectorTwo *ArrayCollector
	)

	collectorOne = &ArrayCollector{}
	collectorTwo = &ArrayCollector{}

	inStream = &Stream{
		source: func(feedFunc func(interface{})) error {
			for _, text := range before {
				feedFunc(text)
			}
			return nil
		},
	}

	err = inStream.
		Map(func(item interface{}) interface{} {
			return strings.Trim(item.(string), " ")
		}).
		Filter(func(item interface{}) interface{} {
			return strings.Compare(item.(string), "Apa") != 0
		}).
		Map(func(item interface{}) interface{} {
			return fmt.Sprintf("Aku %s", item)
		}).
		Collect(collectorOne, collectorTwo).
		Run()

	if err != nil {
		t.Error(err)
	}

	t.Log("Data in Collector One:")
	for i, dt := range collectorOne.ToArray() {
		if expected[i] != dt.(string) {
			t.Errorf("Expected %v but %v", expected[i], dt)
		}
		t.Log(dt)
	}

	t.Log("Data in Collector Two:")
	for i, dt := range collectorTwo.ToArray() {
		if expected[i] != dt.(string) {
			t.Errorf("Expected %v but %v", expected[i], dt)
		}
		t.Log(dt)
	}
}
