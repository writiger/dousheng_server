package snowfalke

import (
	zaplog "dousheng_server/deploy/log"
	"errors"
	"sync"
	"time"
)

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

var iDMaker *worker

func init() {
	var err error
	workerId, err := NewLeaseMaker().getLease()
	iDMaker, err = newWorker(workerId)
	if err != nil {
		zaplog.ZapLogger.Error("IDMaker Init Error")
		panic("IDMaker Init Error")
	}
}

type worker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

// NewUUID 通过雪花算法生成UUID
func NewUUID() int64 {
	return iDMaker.getId()
}

func newWorker(workerId int64) (*worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("err: worker ID excess of quantity")
	}
	// 生成一个新节点
	return &worker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

func (w *worker) getId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := (now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number)
	return ID
}
