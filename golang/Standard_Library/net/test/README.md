# 샘플 작성 

## 구성
1. 명령행 파싱 (공통)
2. 프로그램 유형
    1. 단일 프로세스 단일 스레드
    2. 단일 프로세스 멀티 스레드
    3. 단일 서버 단일 클라이언트
    4. 단일 서버 멀티 클라이언트 
    5. 단일 서버 멀티 프로세스 멀티 클라이언트
3. 로그
4. 테스트

## 고루틴
- 고루틴은 비동기적 함수 실행이다.
- 고루틴실행 후 프로그램이 먼저 종료되면 고루틴은 다 실행되지 않더라도 실행이 멈춘다.
- 모든 고루틴이 실행의 보장을 위해서는 몇가지 방법이 필요함
