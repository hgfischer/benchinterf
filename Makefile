test:
	go test -bench=. -benchmem | grep -v github | column -t

history:
	echo '```' > README.md
	go test -bench=. -benchmem | grep -v github | column -t >> README.md
	echo '```' >> README.md