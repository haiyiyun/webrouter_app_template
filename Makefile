PWD=`pwd`
BIN=$(PWD)/bin
GO_DEBUG=",debug"
GO_PORT_PROJECT_NAME="10090"

#==========================================
install_project_name:
	GOBIN=$(BIN) go install ./cmd/project_name/...

run_project_name:
ifeq (bin/pid.project_name, $(wildcard bin/pid.project_name))
	@echo 'Process exists PID: ' && cat bin/pid.project_name
else
	@mkdir -p $(BIN)/logs
	cd $(BIN) && `GO_PORT=$(GO_PORT_PROJECT_NAME) GO_DEBUG=$(GO_DEBUG) nohup ./project_name > ./logs/nohup_project_name.log 2>&1 & echo $$! >./pid.project_name`
endif

status_project_name:
	@cd $(BIN) && pgrep -aF ./pid.project_name

kill_project_name:
	@cd $(BIN) && pkill -9 -eF ./pid.project_name && rm ./pid.project_name

log_project_name:
	@tail -f $(BIN)/logs/project_name.log -f $(BIN)/logs/nohup_project_name.log
#==========================================

install:
	make install_project_name
#==========================================

runall:
	make run_project_name
#==========================================

statusall:
	make status_project_name
#==========================================

killall:
	make kill_project_name
#==========================================

logall:
	make log_project_name
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
ifeq (bin/pid.project_name, $(wildcard bin/pid.project_name))
	@echo 'Process exists PID: ' && cat bin/pid.project_name
else
	rm -rf $(BIN)/
endif