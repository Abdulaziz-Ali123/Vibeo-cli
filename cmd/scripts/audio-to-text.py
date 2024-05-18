import sys
import speech_recognition as sr


def transcribe(path):
    r = sr.Recognizer()
    with sr.AudioFile(path) as source:
        audio_text = r.record(source)

        text = r.recognize_google(audio_text)
        #print('Converting audio transcripts into text ...')
        return text


audio_path = sys.argv[1]
output = sys.argv[2]+".txt"

audio_text = transcribe(audio_path)

file = open(output, "w")
file.write(audio_text)
file.close()