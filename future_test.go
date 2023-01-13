package future

import (
	"sync"
	"testing"
	"time"
)

func TestFutureInt(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	p := NewPromise[int]()
	f1 := p.GetFuture()
	f2 := p.GetFuture()

	expect := 233

	go func() {
		defer wg.Done()
		if v := f1.Get(); v != expect {
			t.Errorf("promise value not match. expect: %v vs actual %v", expect, v)
		}
	}()

	go func() {
		defer wg.Done()
		if r := f2.WaitFor(time.Millisecond * 500); r != Timeout {
			t.Errorf("expect %v, but get %v", Timeout, r)
		}
		if r := f2.WaitFor(time.Millisecond * 500); r != Ready {
			t.Errorf("expect %v, but get %v", Ready, r)
		}
		if v := f2.Get(); v != expect {
			t.Errorf("promise value not match. expect: %v vs actual %v", expect, v)
		}
	}()

	go func() {
		defer wg.Done()
		if r := f2.WaitUntil(time.Now().Add(time.Millisecond * 500)); r != Timeout {
			t.Errorf("expect %v, but get %v", Timeout, r)
		}
		if r := f2.WaitUntil(time.Now().Add(time.Millisecond * 500)); r != Ready {
			t.Errorf("expect %v, but get %v", Ready, r)
		}
		if v := f2.Get(); v != expect {
			t.Errorf("promise value not match. expect: %v vs actual %v", expect, v)
		}
	}()

	time.Sleep(time.Millisecond * 800)
	p.SetValue(expect)
	wg.Wait()
}

func TestFutureStruct(t *testing.T) {
	type TestData struct {
		Name string
	}
	var wg sync.WaitGroup
	wg.Add(3)
	p := NewPromise[*TestData]()
	f1 := p.GetFuture()
	f2 := p.GetFuture()

	expect := "Test"

	go func() {
		defer wg.Done()
		if v := f1.Get(); v.Name != expect {
			t.Errorf("promise value not match. expect: %v vs actual %v", expect, v)
		}
	}()

	go func() {
		defer wg.Done()
		if r := f2.WaitFor(time.Millisecond * 500); r != Timeout {
			t.Errorf("expect %v, but get %v", Timeout, r)
		}
		if r := f2.WaitFor(time.Millisecond * 500); r != Ready {
			t.Errorf("expect %v, but get %v", Ready, r)
		}
		if v := f2.Get(); v.Name != expect {
			t.Errorf("promise value not match. expect: %v vs actual %v", expect, v)
		}
	}()

	go func() {
		defer wg.Done()
		if r := f2.WaitUntil(time.Now().Add(time.Millisecond * 500)); r != Timeout {
			t.Errorf("expect %v, but get %v", Timeout, r)
		}
		if r := f2.WaitUntil(time.Now().Add(time.Millisecond * 500)); r != Ready {
			t.Errorf("expect %v, but get %v", Ready, r)
		}
		if v := f2.Get(); v.Name != expect {
			t.Errorf("promise value not match. expect: %v vs actual %v", expect, v)
		}
	}()

	time.Sleep(time.Millisecond * 800)
	p.SetValue(&TestData{Name: expect})
	wg.Wait()
}
