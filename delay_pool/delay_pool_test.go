package delay_pool

import (
	"testing"
	"time"
)

func TestAllocateDelayPoolItem(t *testing.T) {
	var dp DelayPool
	dp.Reset([]interface{}{1, 2, 3, 4, 5}, 5, time.Second)
	s := 0
	st := time.Now()

	for i := 0; i < 6; i++ {
		v, ok := dp.GetValue()

		if !ok {
			break
		}

		s += v.(int)
	}

	et := time.Now()

	if s != 15 {
		t.Errorf("%#v", s)
	}

	if d := et.Sub(st); d < time.Second {
		t.Errorf("%#v", d)
	}

	_, ok := dp.GetValue()

	if ok {
		t.Error()
	}
}
