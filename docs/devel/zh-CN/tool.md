## 常用工具

- [readme.so](https://readme.so/editor) 工具来协助生成 README 文件
-  [gitignore.io](https://www.toptal.com/developers/gitignore) 在线生成 .gitignore文件
  - [db2struct](https://github.com/Shelnutt2/db2struct) 根据数据库表生成 Model 文件
    ```shell
      go get github.com/Shelnutt2/db2struct/cmd/db2struct
      # 编译出命令行工具
      go build -o ~/go/bin/db2struct -v db2struct/cmd/db2struc/main.go
    ```
    然后在需要创建model的目录internal/pkg/model下执行以下命令 生成model
    ```shell
      db2struct --gorm --no-json -H 127.0.0.1 -d miniblog -t user --package model --struct UserM -u root -p '12345678' --target=user.go
      db2struct --gorm --no-json -H 127.0.0.1 -d miniblog -t post --package model --struct PostM -u root -p '12345678' --target=post.go
    ```
     
    db2struct 命令行参数说明如下：
    
  ```shell
        $ db2struct  -h
        Usage of db2struct:
        db2struct [-H] [-p] [-v] --package pkgName --struct structName --database databaseName --table tableName
        Options:
        -H, --host=         Host to check mariadb status of
        --mysql_port=3306   Specify a port to connect to
        -t, --table=        Table to build struct from
        -d, --database=nil  Database to for connection
        -u, --user=user     user to connect to database
        -v, --verbose       Enable verbose output
        --package=          name to set for package
        --struct=           name to set for struct
        --json              Add json annotations (default)
        --no-json           Disable json annotations
        --gorm              Add gorm annotations (tags)
        --guregu            Add guregu null types
        --target=           Save file path
        -p, --password=     Mysql password
        -h, --help          Show usage message
        --version           Show version
  ```
  导出Mysql
    
    ```
     mysqldump -h127.0.0.1 -uroot --databases miniblog -p'12345678' --add-drop-trigger --add-locks --no-data > configs/miniblog.sql
    ```
  
