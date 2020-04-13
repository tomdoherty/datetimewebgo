all: date time web

clean:
	rm -f date time web

date:
	go build -a -ldflags '-extldflags "-static"' cmd/date/date.go
	docker build -f build/Dockerfile.date -t tomdo/date:latest .
	docker push tomdo/date:latest

time:
	go build -a -ldflags '-extldflags "-static"' cmd/time/time.go
	docker build -f build/Dockerfile.time -t tomdo/time:latest .
	docker push tomdo/time:latest

web:
	go build -a -ldflags '-extldflags "-static"' cmd/web/web.go
	docker build -f build/Dockerfile.web -t tomdo/web:latest .
	docker push tomdo/web:latest
