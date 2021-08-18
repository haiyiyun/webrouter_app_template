PWD=`pwd`
BIN=$(PWD)/bin
GO_DEBUG=",debug"
GO_PORT_APP1="10090"

#==========================================
install_app1:
	GOBIN=$(BIN) go install ./cmd/app1/...

run_app1:
ifeq (bin/pid.app1, $(wildcard bin/pid.app1))
	@echo 'Process exists PID: ' && cat bin/pid.app1
else
	@mkdir -p $(BIN)/logs
	cd $(BIN) && `GO_PORT=$(GO_PORT_APP1) GO_DEBUG=$(GO_DEBUG) nohup ./app1 > ./logs/nohup_app1.log 2>&1 & echo $$! >./pid.app1`
endif

status_app1:
	@cd $(BIN) && pgrep -aF ./pid.app1

kill_app1:
	@cd $(BIN) && pkill -9 -eF ./pid.app1 && rm ./pid.app1

log_app1:
	@tail -f $(BIN)/logs/app1.log -f $(BIN)/logs/nohup_app1.log
#==========================================

install:
	make install_app1
#==========================================

runall:
	make run_app1
#==========================================

statusall:
	make status_app1
#==========================================

killall:
	make kill_app1
#==========================================

logall:
	make log_app1
#==========================================

status:
	make statusall | grep -e ^[0-9]
#==========================================

rerunall:
	make killall
	make runall
#==========================================

reinstall:
	make killall
	make install
	make runall
#==========================================
clean:
ifeq (bin/pid.app1, $(wildcard bin/pid.app1))
	@echo 'Process exists PID: ' && cat bin/pid.app1
else
	rm -rf $(BIN)/
endif