package cartel

import (
	"testing"
	"time"
)

func Test_Runs_One_Per_Second(t *testing.T) {

	p := NewPool(
		PoolOptions{
			Size:        1,
			PerDuration: 1,
			Duration:    time.Second,
		},
	)
	p.Do(TestTimeTask{})
	p.Do(TestTimeTask{})
	p.End()

	values := p.GetOutput()

	if actual := len(values); actual != 2 {
		t.Errorf("expected 2 items but got %v", actual)
	}

	firstTime := values[0].(time.Time)
	secondTime := values[1].(time.Time)

	d := secondTime.Sub(firstTime)
	if d.Seconds() < 1 {
		t.Errorf("expected 1 seconds but got %v - %v %v", d.Seconds(), firstTime, secondTime)
	}
}

type TestTimeTask struct {
}

func (ttt TestTimeTask) Execute() interface{} {
	return time.Now()
}
