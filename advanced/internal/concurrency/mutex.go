package concurrency

import "sync"

type CharCounter struct {
	counter int
	mu      sync.Mutex
}

func (cc *CharCounter) Count(char rune, source string) int {
	numberOfRoutines := len(source) / 10
	sections := toSections(source, numberOfRoutines)

	waitGroup := &sync.WaitGroup{}
	for _, section := range sections {
		waitGroup.Add(1)
		cc.count(char, section, waitGroup)
	}
	waitGroup.Wait()
	return cc.counter
}

func (cc *CharCounter) count(char rune, source string, waitGroup *sync.WaitGroup) {
	for _, ch := range source {
		if ch == char {
			cc.incr()
		}
	}
	waitGroup.Done()
}

func (cc *CharCounter) incr() {
	defer cc.mu.Unlock()
	cc.mu.Lock()
	cc.counter += 1
}
