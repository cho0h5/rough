# rough
간편하게 cpp프로젝트를 만들고 실행하기 위해서 만들었습니다!
알고리즘 문제를 풀거나 간단한 소스코드를 빠르게 만들고 실행할 때 유용합니다.
## 설치
```bash
$ go install github.com/cho0h5/rough@latest
```
then, add this to ~/.bashrc
```
PATH=$PATH:~/go/bin
```
### 사용법
```bash
$ rough new directory_name  # 폴더를 만들고 main.cpp파일을 만들어 넣습니다.
$ cd directory_name
$ vim main.cpp
$ rough run # 실행합니다.
```
