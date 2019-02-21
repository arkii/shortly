build:
	GOOS=linux go build -o ./app .
install:
	cp -f dist/shortly.service /lib/systemd/system/ && \
	    cp -f dist/shortly.conf /etc/rsyslog.d/ && \
	    systemctl daemon-reload && systemctl enable shortly && service shortly start && service rsyslog restart
clean:
	rm -f ./app
