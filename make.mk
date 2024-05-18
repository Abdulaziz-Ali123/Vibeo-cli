
#check if python is installed
check_python:
	@python3 -c "import sys; \
	exit_code = 0 if sys.version_info >= (3, 0) else 1; \
	sys.exit(exit_code)" || (echo "Python 3 is required but not found. Please install Python 3."; exit 1)


# Default target to build the application
build: check_python
	go build -o Vibeo-Cli main.go

# Target to clean up built files
clean:
	rm -f Vibeo-Cli

# Target to package the application with the virtual environment
package: build
	python -m venv env && \
	source env/bin/activate && \
	pip install speechrecognition
	cp cmd/scripts/audio-to-text.py env


