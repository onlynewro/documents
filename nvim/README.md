nvim 설치 및 플러그인 설치
==========================
nvim 설치
---------
# docker 를 통한 리눅스 설치
1. docker image 중 debian/golang 을 선택
'''
#docker container run -it --name my-vim <docker-image-id> /bin/bash
#docker start <docker name>
#docker attach <docker name>
```
2. linux update
```
#apt update
#apt upgrade
```
3. neovim install and setting
```
#apt install neovim
#mkdir .config
#mkdir nvim
#nvim init.nvim
``` 
```
:set number
:set relativenumber
```
4. vim-plug install
```
#apt install curl
#apt install universal-ctags
```
tagbar plug install
5. vim coc install
apt install nodejs
apt install npm

```
~/.local/share/nvim/plugged/coc.nvim/npm install -g yarn
yarn install
cocInstall coc-python
apt install pip3
pip3 install jedi
cocInstall coc-go
apt install gopls

```
