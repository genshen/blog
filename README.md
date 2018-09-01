# blog
my personal blog with Golang and material design  
[demo](https://gensh.me):https://gensh.me  

## dependency tools
We use go modules to build project and manage dependencies.
Make sure you have go 1.11 installed or vgo installed.
For more details, see [https://github.com/golang/go/wiki/Modules](https://github.com/golang/go/wiki/Modules).

## Build
```bash
$ go build
```

## config database
Mongodb is used for blog storage,to configure Mongodb,add those lines to **[conf/database.conf](/conf/database.conf)** file:
```
db_type = mongodb
db_name = blog
db_debug = false
db_config = localhost:27017

db_auth = true
db_auth_user = genshen
db_auth_pwd = genshen_blog
```

## config qiniu cloud
to store your files(e.g. images) by using qiniu cloud,add those lines to  **[conf/keys.conf](/conf/keys.conf)** file:
```
qiniu_access_key = xxxxxx
qiniu_secret_key = xxxxx
qiniu_bucket_name = YOUR_BUCKET_NAME
qiniu_token_expires = 360
qiniu_config_domain = http://YOUR_OUTER_LINK.bkt.clouddn.com/
qiniu_config_upload_path = http://up-z1.qiniu.com
```
notice: if the file dose exists,you should create it.
## config github authentication
to configure github authentication for user's comments,add those lines to **[conf/keys.conf](/conf/keys.conf)** file:  
```
github_auth_url = https://github.com/login/oauth/access_token
github_client_id = xxxx
github_client_secret = xxx
```
notice: if the file dose exists,you should create it.

## for developers
##todo