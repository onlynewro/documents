# docker 컨테이너에서 nvidia(CUDA) 가속 사용하기

## 준비
- docker, docker-compose, nvidia CUDA를 호스트에 설치
- nvidia-docker가 docker command에 통합되어 docker run --gpus로 변경
<code><pre>
$ nvidia-smi 호스트 실행
</pre></code>
<code><pre>
$docker run --gpus all -it -rm nvidia/cuda:latest nvidia-smi
</pre></code>
- 두 결과가 같은지 확인
- CUDA 버젼 확인
