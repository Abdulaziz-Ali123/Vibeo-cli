

# Default target to build the application
build:
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


