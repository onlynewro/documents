// 특정 작업을 수행하는 고루틴의 풀을 관리하는 패키지
package work

import "sync"

type Worker interface {
	Task()
}

// Worker 인터페이스를 실행하는 고루틴의 풀을 제공하기 위한 Pool 구조체
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// 고루틴의 작업을 실해하고 완료될 때까지 대기하는 풀 함수
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			// work 채널에서 Worker 인터페이스 값을 받는 한 계속해서 루프를 실행
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// 풀에 새로운 작업을 추가하는 작업
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// 모든 고루틴을 종료할때까지 대기하는 메서드
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
