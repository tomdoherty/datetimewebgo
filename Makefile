all: date time web

clean:
	rm -f date time web

date:
	go build -o date -a -ldflags '-extldflags "-static"' cmd/date/main.go
	docker build -f build/Dockerfile.date -t tomdo/date:latest .
	docker push tomdo/date:latest

time:
	go build -o time -a -ldflags '-extldflags "-static"' cmd/time/main.go
	docker build -f build/Dockerfile.time -t tomdo/time:latest .
	docker push tomdo/time:latest

web:
	go build -o web -a -ldflags '-extldflags "-static"' cmd/web/main.go
	docker build -f build/Dockerfile.web -t tomdo/web:latest .
	docker push tomdo/web:latest

test:
	( cd cmd/date && go test )
	( cd cmd/time && go test )
