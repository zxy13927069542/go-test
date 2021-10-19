### Delve

- ```
  Goland Debug报错:
  	Version of Delve is too old for this version of Go (maximum supported version 1.12, suppress this error with --check-go-version=false)
  	即Delve版本过旧或者压根没有Deve
  ```

  - 解决措施：下载一个高版本或新的Delve,并在goland中配置其路径

    1. ```
       go get -u github.com/go-delve/delve/cmd/dlv	//下载dlv到本地
       ```

    2. 用everything搜索delve.exe找到其路径，并克隆到Go安装路径的bin目录下

    3. 打开 `Hele->Edit Customer Properties`,若提示文件不存在，点击创建。然后在新加一行 

       ```
       dlv.path=Go的安装路径/bin/dlv.exe
       ```

        重启就可以了

