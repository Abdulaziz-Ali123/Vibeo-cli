# Vibeo-cli
Vibeo-cli is an easy-of-use/one-stop-shop type of tool that combines preexisting tools and combines them to make audio-to-text generation easy and kept local. This command line interface tool allows you to transcribe the audio of a `.wav` either  already obtained or from a YouTube link. Vibeo-cli also can convert audio and video to `.wav` files for easy use.
### Transparency
Vibeo-cli acquires YouTube video audio using `youtube-dl` and converts files through `FFmpeg`.
The audio-to-text generation is done using Python's `speech_recognition` module.
## Set Up

### Pre-requisites
* `Python` For download instructions full link is at the bottom or click [here](https://python.org/)
* `youtube-dl` For download instructions full link is at the bottom or click [here](https://github.com/ytdl-org/youtube-dl)
* `FFmpeg` For download instructions full link is at the bottom or click [here](https://ffmpeg.org/)

To set up, after downloading or cloning the repo: \
In the terminal/command prompt change directories to the clone/downloaded repo run
```
$ make --makefile=make.mk package
```
## Usage
### File conversion
For file conversions, Vibeo can take in a video or audio and turn it into `.wav`. For example: \
```
./Vibeo-cli convert -f example/directory/video.mp4 -o audio_file
```
### Transcription
For file conversions, Vibeo can take in a file path or YouTube link and turn it into `.wav` then pass it on to be transcribed. For example: 
 <br /> <br />  <br />  ***__For files__***
```
./Vibeo-cli transcribe -f example/directory/video.mp4 -o output_name
```
 <br /> <br />   ***__For Links__***
```
./Vibeo-cli transcribe -l https:\\youtube.com/watch?v=videothatexists -o output_name
```

***__Note__: For further documentation use:***

```
./Vibeo-cli --help
```

 <br/> <br/> <br/> 

### Links
python: https://python.org/ \
youtube-dl: https://github.com/ytdl-org/youtube-dl \
ffmpeg: https://ffmpeg.org/


