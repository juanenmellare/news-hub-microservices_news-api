package jobs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestNewCronJob(t *testing.T) {
	cronJob := NewCronJob("FooJob", "* * * * *", func() {})

	assert.Implements(t, (*CronJob)(nil), cronJob)
}

func Test_cronJob_GetCron(t *testing.T) {
	cron := "* * * * *"
	cronJob := NewCronJob("FooJob", cron, func() {})

	assert.Equal(t, cron, cronJob.GetCron())
}

func Test_cronJob_GetName(t *testing.T) {
	name := "FooJob"
	cronJob := NewCronJob(name, "* * * * *", func() {})

	assert.Equal(t, name, cronJob.GetName())
}

func Test_cronJob_Run(t *testing.T) {
	c := make(chan string, 1)
	expectedValue := "foo"
	function := func() {
		c <- expectedValue
	}
	NewCronJob("FooJob", "* * * * *", function).Run()
	value := ""
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		value = <-c
		wg.Done()
	}()
	wg.Wait()

	assert.Equal(t, "foo", value)
}

func Test_cronJob_Run_do_error_panic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			str := fmt.Sprintf("the test should not panic: %v", r)
			t.Errorf(str)
		}
	}()

	NewCronJob("FooJob", "foo", func() {
		panic("foo")
	}).Run()
}
