package env

import (
	"os"
)

const (
	YANDEX_KEY           = "YANDEX_KEY"
	AUDIO_PATH           = "AUDIO_PATH"
	DEFAULT_LANG         = "DEFAULT_LANG"
	DEFAULT_SPEED        = "DEFAULT_SPEED"
	DEFAULT_EMOTION      = "DEFAULT_EMOTION"
	DEFAULT_SPEAKER      = "DEFAULT_SPEAKER"
	DEFAULT_QUALITY      = "DEFAULT_QUALITY"
	DEFAULT_AUDIO_FORMAT = "DEFAULT_AUDIO_FORMAT"
)

func GetYandexKey() string {
	return os.Getenv(YANDEX_KEY)
}

func GetDefaultLang() string {
	return os.Getenv(DEFAULT_LANG)
}

func GetDefaultSpeed() string {
	return os.Getenv(DEFAULT_SPEED)
}

func GetDefaultEmotion() string {
	return os.Getenv(DEFAULT_EMOTION)
}

func GetDefaultSpeaker() string {
	return os.Getenv(DEFAULT_SPEAKER)
}

func GetDefaultQuality() string {
	return os.Getenv(DEFAULT_QUALITY)
}

func GetDefaultAudioFormat() string {
	return os.Getenv(DEFAULT_AUDIO_FORMAT)
}

func GetPathForAudio() string {
	return os.Getenv(AUDIO_PATH)
}
