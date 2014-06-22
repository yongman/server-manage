all:
	go build
t:
	go build
	echo "build done"
	./server-manage t
m:
	go build
	./server-manage g -m /home/users/yanming02/workspace/server-manage/host_mem.info
s:
	go build
	./server-manage g -s /home/users/yanming02/workspace/server-manage/host_redis.info
