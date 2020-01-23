# interface 
- 서로 다른 타입들에 대한 처리
- 어떤 타입들도 담을 수 있는 바구니(인터페이스)
- 바구니에 담긴 타입에 따라 같은 이름 다른 처리 하는 메소드 정의
* handler
<pre><code>
    type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
    }

    type HandlerFunc func(ResponseWriter, *Request)

    func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)

    ServeHTTP calls f(w, r).
</code></pre>
