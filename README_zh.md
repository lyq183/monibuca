# 项目环境配置：

## 首先需要安装好golang环境；

## 拉取项目：
git clone -b v9 https://github.com/lyq183/monibuca

## 数据库：
使用项目目录configs 下的Monibuca_sql构建初步的数据库；

## 连接数据库和邮箱
在目录configs下的Web_config.go中配置

## 启动
执行main.go 
浏览器打开 监听端口 可以登陆用户
http://localhost:端口/admin：打开管理员后台


5. ffmpeg或者OBS推流到1935端口
6. 后台界面上提供直播预览、录制flv、rtsp拉流转发、日志跟踪等功能