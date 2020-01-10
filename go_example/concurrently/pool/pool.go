// 사용자가 정의한 리소스의 집합을 관리하는 패키지
package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	m         sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

// 리소스를 획득하려 할 때 풀이 닫혀있는 경우에 발생한다.
var ErrPoolClosed = errors.New("풀이 닫혔습니다.")

// 리소스 관리 풀을 생성한다.
func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("풀의 크기가 너무 작습니다.")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

// 풀에서 리소스를 획득하는 메서드
func (p *Pool) Acquire() (io.Closer, error) {
	select {
	// 사용 가능한 리소스가 있는지 검사한다.
	case r, ok := <-p.resources:
		log.Println("리소스 획득:", "공유된 리소스")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	// 사용 가능한 리소스가 없는 경우 새로운 리소스를 생성한다.
	default:
		log.Println("리소스 획득:", "새로운 리소스")
		return p.factory()
	}
}

// 풀에 리소스를 반환하는 메서드
func (p *Pool) Release(r io.Closer) {
	//안전한 작업을 위해 잠금을 설정한다.
	p.m.Lock()
	defer p.m.Unlock()

	// 풀이 닫혔으면 리소스를 해제한다.
	if p.closed {
		r.Close()
		return
	}

	select {
	// 새로운 리소스를 큐에 추가한다.
	case p.resources <- r:
		log.Println("리소스 반환:", "리소스 큐에 반환")

	// 리소스 큐가 가득 찬 경우 리소스를 해제한다.
	default:
		log.Println("리소스 반환", "리소스 해제")
		r.Close()
	}
}

func (p *Pool) Close() {
	// 안전한 작업을 위해 잠금을 설정한다.
	p.m.Lock()
	defer p.m.Unlock()

	// 풀이 이미 닫혔으면 아무런 작업도 수행하지 않는다.
	if p.closed {
		return
	}

	// 풀을 닫힌 상태로 전환 한다.
	p.closed = true

	// 리소스를 해제하기에 앞서 채너을 먼저 닫는다.
	// 그렇지 않으면 데드락에 걸릴 수 있다.
	close(p.resources)

	// 리소스를 해제한다.
	for r := range p.resources {
		r.Close()
	}
}
