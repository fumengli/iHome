#!/bin/bash
#fdfs_storaged ./conf/storage.conf
fdfs_storaged /home/fml/workspace/go/src/iHome/conf/storage.conf restart
# 停止fastdfs [tracker]
#fdfs_trackerd ./conf/tracker.conf
fdfs_trackerd /home/fml/workspace/go/src/iHome/conf/tracker.conf restart
