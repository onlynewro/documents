# 정리 할 것들
* 텍스트 파일읽기
    - 텍스트 파일을 한 줄씩 읽기
    - 텍스트 파일을 한 단어씩 읽기
    - 텍스트 파일을 한 문자씩 읽기
    - /dev/random 읽
* 파일에서 읽고 싶은 만큼 읽기
* 바이너리 포맷을 사용하는 이유
* csv 파일 읽기
* 파일 쓰기
* 디스크에 데이터를 읽거나 쓰기
* 파일의 접근 권한
* 유닉스 시그널 다루기
    - 두 가지 시그널 처리하기
    - 모든 종류의 시그널 처리하기
* 유닉스 파이프 다루기
* 디렉토리 트리 탐색하기
* eBPF 사용하기
* 시스템 콜 추적하기
* 유저 ID 와 그룹 ID
## 네트워크 프로그래밍
## go 자료구조

## 핸들러 

* type HandlerFunc
<pre><code>
    type HandlerFunc func(ResponseWriter, *Request)
</code></pre>
The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers. If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.

* func (HandlerFunc) ServeHTTP
<pre><code>
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
</code></pre>
ServeHTTP calls f(w, r).
