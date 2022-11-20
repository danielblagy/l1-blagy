package main

import "fmt"

type MediaPlayer interface {
	Play(audioType string, filename string)
}

type AdvancedMediaPlayer interface {
	PlayVcl(filename string)
	PlayMp4(filename string)
}

// implements AdvancedMediaPlayer
type VclPlayer struct {
}

func (p VclPlayer) PlayVcl(filename string) {
	fmt.Println("Playing vcl file:", filename)
}
func (p VclPlayer) PlayMp4(filename string) {
	// not implemented
}

// implements AdvancedMediaPlayer
type Mp4Player struct {
}

func (p Mp4Player) PlayVcl(filename string) {
	// not implemented
}
func (p Mp4Player) PlayMp4(filename string) {
	fmt.Println("Playing mp4 file:", filename)
}

// implements MediaPlayer
type MediaAdapter struct {
	player AdvancedMediaPlayer
}

func (adapter *MediaAdapter) Play(audioType string, filename string) {
	if audioType == "vcl" {
		adapter.player = VclPlayer{}
		adapter.player.PlayVcl(filename)
	} else if audioType == "mp4" {
		adapter.player = Mp4Player{}
		adapter.player.PlayMp4(filename)
	}
}

/*
	We want to make AudioPlayer to play other formats as well.
	To attain this, we have created an adapter class MediaAdapter which implements
	the MediaPlayer interface and uses AdvancedMediaPlayer objects to play the required format.

	AudioPlayer uses the adapter class MediaAdapter passing it the desired audio type
	without knowing the actual class which can play the desired format.
*/
// implements MediaPlayer and uses MediaAdapter to support mp4 and vcl file formats
type AudioPlayer struct {
	mediaAdapter MediaAdapter
}

func (ap *AudioPlayer) Play(audioType string, filename string) {
	// AudioPlayer can play mp3 files by default
	if audioType == "mp3" {
		fmt.Println("Playing mp3 file:", filename)
	} else if audioType == "vcl" || audioType == "mp4" {
		// mediaAdapter provides support to play other media formats
		ap.mediaAdapter.Play(audioType, filename)
	} else {
		fmt.Println("Couldn't play file:", filename, ",", audioType, "format is not supported")
	}
}

func main() {
	player := AudioPlayer{}
	player.Play("mp3", "track1.mp3") // is supported by default in AudioPlayer
	player.Play("mp4", "track2.mp4") // is supported with the help of MediaAdapter
	player.Play("vcl", "track3.vcl") // is supported with the help of MediaAdapter
	player.Play("avi", "track4.avi") // is not supported
}
