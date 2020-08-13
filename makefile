# Make file CD INTO https://bit.ly/2XSh3ZA | https://bit.ly/31I776g

Chapter01 = "01. file-reverse"
Chapter06 = "06. cli-streamer"

run-01:
	cd $(Chapter01) && make run

run-02:
	echo "there is no code to run, shared code directory"

run-03:
	cd "03. string-parsing" && make run

run-04:
	cd "04. parse-testing" && make run

run-04-cmd:
	cd "04. parse-testing" && make run-cmd
	
run-04-ps:
	cd "04. parse-testing" && make run-ps
	
run-04-chrome:
	cd "04. parse-testing" && make run-chrome

run-05:
	cd "05. console-app" && make run

run-06:
	cd $(Chapter06) && make run