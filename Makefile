.PHONY: ALL acceptance test wip

ALL: test

clean:
	rm -rf logs

logs:
	mkdir logs

test: logs
	(cd dronetown && godog run --tags ~@Pending) 2>&1 | tee logs/test_not_pending.log

wip: logs
	(cd dronetown && godog run --tags @wip) 2>&1 | tee logs/test_wip.log

watch:
	find . -type f -name '*.feature' -or -name '*.go' | entr -d -c bash -c 'make acceptance && make test && make wip'
