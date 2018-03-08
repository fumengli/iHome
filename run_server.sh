#!/bin/bash
# 启动redis
#redis-server ./conf/redis.conf 
# 启动fastdfs [tracker]
#fdfs_trackerd ./conf/tracker.conf
fdfs_trackerd /home/fml/workspace/go/src/iHome/conf/tracker.conf
# 启动fastdfs [storage]
#fdfs_storaged ./conf/storage.conf
fdfs_storaged /home/fml/workspace/go/src/iHome/conf/storage.conf
