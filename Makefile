all: date time web

clean:
	rm -f date time web

date:
	go build -a -ldflags '-extldflags "-static"' date.go
	docker build -f Dockerfile.date -t tomdo/date:latest .
	docker push tomdo/date:latest

time:
	go build -a -ldflags '-extldflags "-static"' time.go
	docker build -f Dockerfile.time -t tomdo/time:latest .
	docker push tomdo/time:latest

web:
	go build -a -ldflags '-extldflags "-static"' web.go
	docker build -f Dockerfile.web -t tomdo/web:latest .
	docker push tomdo/web:latest
