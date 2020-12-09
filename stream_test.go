package stream

import (
	"testing"
)

func TestStream_Integration(t *testing.T) {
	var (
		err error
		inStream *Stream
		numbers []int

		counterOne int
		expectedOne []int

		counterTwo int
		expectedTwo []int

		counterThree int
		expectedThree []int
	)

	numbers = []int{2286,3200,4176,5120,6165,7206,875,9129,10109,11123,12111}

	counterOne = 0
	expectedOne = []int{2286,3200,4176,5120,6165,7206,9129}

	counterTwo = 0
	expectedTwo = []int{2,3,4,5,6,7,9}

	counterThree = 0
	expectedThree = []int{2,3,4}

	inStream = &Stream{
		source: func(feedFunc func(interface{})) error {
			var (
				number interface{}
			)

			for _, number = range numbers {
				feedFunc(number)
			}
			return nil
		},
	}

	err = inStream.
		Filter(func(item interface{}) interface{} {
			return item.(int) > 1000 && item.(int) < 10000
		}).
		Map(func(item interface{}) interface{} {

			if item.(int) != expectedOne[counterOne] {
				t.Errorf("Expected %v but got %v", expectedOne[counterOne], item)
			}
			counterOne++

			return item
		}).
		Map(func(item interface{}) interface{} {
			return item.(int) / 1000
		}).
		Map(func(item interface{}) interface{} {

			if item.(int) != expectedTwo[counterTwo] {
				t.Errorf("Expected %v but got %v", expectedTwo[counterTwo], item)
			}
			counterTwo++

			return item
		}).
		Filter(func(item interface{}) interface{} {
			return item.(int) < 5
		}).
		Map(func(item interface{}) interface{} {

			if item.(int) != expectedThree[counterThree] {
				t.Errorf("Expected %v but got %v", expectedThree[counterThree], item)
			}
			counterThree++

			return item
		}).
		Map(func(item interface{}) interface{} {
			t.Log(item)
			return nil
		}).
		Run()

	if err != nil {
		t.Error(err)
	}

	if counterOne != len(expectedOne) {
		t.Errorf("In Stage One, expected %v but %v", len(expectedOne), counterOne)
	}

	if counterTwo != len(expectedTwo) {
		t.Errorf("In Stage Two, expected %v but %v", len(expectedTwo), counterTwo)
	}

	if counterThree != len(expectedThree) {
		t.Errorf("In Stage Three, expected %v but %v", len(expectedThree), counterThree)
	}
}
