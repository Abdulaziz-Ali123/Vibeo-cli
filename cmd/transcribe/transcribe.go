/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package transcribe

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

// TranscribeCmd represents the transcribe command
var TranscribeCmd = &cobra.Command{
	Use:   "transcribe",
	Short: "Transcribes audio given youtube link or file",
	Long: `Transcribes audio given youtube link or file. For example:

Vibeo-cli transcribe -f example/direcrory/audio.mp3 -o transcribed_file
or
Vibeo-cli transcribe -l https://youtube.com/example-video -o transcribed_file`,
	Run: func(cmd *cobra.Command, args []string) {
		File, _ := cmd.Flags().GetString("file")
		Link, _ := cmd.Flags().GetString("link")
		Output, _ := cmd.Flags().GetString("output")

		if File != "None" {
			exec_cmd := exec.Command("env/bin/python", "env/audio-to-text.py", File, Output)
			_, err := exec_cmd.CombinedOutput()
			if err != nil {
				log.Printf("Could not retrive audio file %v", err)
			} else {
				fmt.Println("Audio Transcribed as:" + Output + "txt")
			}

		} else if Link != "None" {
			exec_cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", Link, "-o", Output+"_Audio.mp3")
			fmt.Println("Retriving Audio...")
			_, err := exec_cmd.CombinedOutput()
			if err != nil {
				log.Printf("Could not retrive audio file %v", err)
			} else {
				exec_cmd := exec.Command("ffmpeg", "-i", Output+"_Audio.mp3", Output+"_Audio.wav")

				_, err := exec_cmd.CombinedOutput()
				if err != nil {
					log.Printf("Convertion failed %v ", err)
				} else {
					fmt.Println("Audio Retrieved.")

					fmt.Println("Initiating transcription...")
					exec_cmd = exec.Command("env/bin/python", "env/audio-to-text.py", Output+"_Audio.wav", Output)
					_, err = exec_cmd.CombinedOutput()
					if err != nil {
						log.Printf("Could not retrive audio file %v", err)
					} else {
						fmt.Println("Audio Transcribed as:" + Output + ".txt")
						fmt.Println("Removeing audio files...")
						exec_cmd := exec.Command("rm", Output+"_Audio.mp3", Output+"_Audio.wav")

						_, err := exec_cmd.CombinedOutput()
						if err != nil {
							log.Printf("Convertion failed %v ", err)
						} else {
							fmt.Println("Audio Files have been removed.")
						}

					}
				}
			}
		} else {
			cmd.Help()
		}

	},
}

func init() {
	TranscribeCmd.Flags().StringP("file", "f", "None", "Path to file desired for conversion")
	TranscribeCmd.Flags().StringP("link", "l", "None", "youtube link to file desired for conversion")
	TranscribeCmd.Flags().StringP("output", "o", "Transcribed", "Name of the converted file")
}
