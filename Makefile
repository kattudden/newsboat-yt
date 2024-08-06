install:
	go build -o newsboat-yt .
	mv newsboat-yt /usr/local/bin/newsboat-yt

clean:
	go clean
	rm -f newsboat-yt

uninstall:
	rm -f /usr/local/bin/newsboat-yt
