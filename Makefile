test:
	go test

bench:
	go test -bench=. -count=1 -run=^$ -v

clean:
	find . -name "*_*_test_*.png" -delete