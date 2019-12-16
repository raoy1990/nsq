package nsqd

import "sync/atomic"

type dummyBackendQueue struct {
	readChan chan []byte
	depth    int64
}

func newDummyBackendQueue() BackendQueue {
	return &dummyBackendQueue{readChan: make(chan []byte), depth: 0}
}

func (d *dummyBackendQueue) Put([]byte) error {
	atomic.AddInt64(&d.depth, 1)
	return nil
}

func (d *dummyBackendQueue) ReadChan() chan []byte {
	return d.readChan
}

func (d *dummyBackendQueue) Close() error {
	return nil
}

func (d *dummyBackendQueue) Delete() error {
	return nil
}

func (d *dummyBackendQueue) Depth() int64 {
	return atomic.LoadInt64(&d.depth)
}

func (d *dummyBackendQueue) Empty() error {
	return nil
}
