all:
	go build
t:
	go build
	echo "build done"
	./server-manage t
