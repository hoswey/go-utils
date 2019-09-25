package concurrent

import (
	"hash/fnv"
	"sync"
)

type StripedMutex struct {
	mutexes         map[uint32]sync.Mutex
	concurrentLevel int
}

func NewStripedMutex(concurrentLevel int) *StripedMutex {

	m := StripedMutex{concurrentLevel: concurrentLevel}
	m.mutexes = make(map[uint32]sync.Mutex, m.concurrentLevel)
	for i := 0; i < m.concurrentLevel; i++ {
		m.mutexes[uint32(i)] = sync.Mutex{}
	}

	return &m
}

func (s *StripedMutex) GetLock(key string) sync.Mutex {

	hash32 := fnv.New32()
	_, err := hash32.Write([]byte(key))
	if err != nil {
		panic(err)
	}

	k := hash32.Sum32() % uint32(s.concurrentLevel)
	return s.mutexes[k]
}
