package audiocacher

import (
	"log"
	"net/http"
	"github.com/titanium-codes/yandex-tts-processor/env"
	"encoding/base64"
	"fmt"
	"os"
	"io"
	"html/template"
	"archive/zip"
	"github.com/kataras/go-errors"
	"io/ioutil"
	"strconv"
	"strings"
)

const (
	urlFormat          = "https://tts.voicetech.yandex.net/generate?key=%s&format=%s&speaker=%s&speed=%s&emotion=%s&lang=%s&quality=%s&text=%s"
	defaultLang        = "ru-Ru"
	defaultSpeed       = "1.0"
	defaultEmotion     = "neutral"
	defaultSpeaker     = "oksana"
	defaultQuality     = "hi"
	defaultAudioFormat = "mp3"
)

var (
	lang, speed, emotion, speaker, quality, audioFormat, key, path string
)

func init() {
	initFolderForAudio()
	initYandexKey()
	lang = getDefaultIfNotDefined(env.GetDefaultLang, defaultLang)
	speed = getDefaultIfNotDefined(env.GetDefaultSpeed, defaultSpeed)
	emotion = getDefaultIfNotDefined(env.GetDefaultEmotion, defaultEmotion)
	speaker = getDefaultIfNotDefined(env.GetDefaultSpeaker, defaultSpeaker)
	quality = getDefaultIfNotDefined(env.GetDefaultQuality, defaultQuality)
	audioFormat = getDefaultIfNotDefined(env.GetDefaultAudioFormat, defaultAudioFormat)
}

func getDefaultIfNotDefined(getter func() string, defaultValue string) string {
	envString := getter()
	if len(envString) == 0 {
		return defaultValue
	}
	return envString
}

/**
Gets TTS for text query.
In case of this tts already exists -> will return existing, otherwise will go to yandex server and return tts from there.
Zips all audio files.
 */
func GetTtsForText(text string) (string, error) {
	log.Println("Getting tts for text =", text)
	zipName := base64Encode(text)
	pathForZip := path + zipName
	if _, err := os.Stat(pathForZip); os.IsNotExist(err) {
		if err := downloadAndSaveZip(text); err != nil {
			return "", err
		}
		return zipName, nil
	}
	log.Println("Got tts for", text, "from cache")
	return zipName, nil
}

func downloadAndSaveZip(text string) error {
	requestUrl := formUrl(text)
	fileName := "audio." + audioFormat
	zipName := base64Encode(text)
	log.Println(requestUrl)
	resp, err := http.Get(requestUrl)
	output, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Println("Error while creating", text, "-", err)
		return err
	}
	if err != nil {
		log.Println("Error:", "Get Audio failed with error =", err)
	}
	defer resp.Body.Close()
	//getting body from response
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		log.Println("Error with response body = ", string(body), "Response status code = ", strconv.Itoa(resp.StatusCode))
		if err != nil {
			return err
		}
		return errors.New("Request failed with status " + strconv.Itoa(resp.StatusCode) + ",message = " + string(body))
	}
	log.Println("Successfully got tts for query", text)
	n, err := io.Copy(output, resp.Body)
	if err != nil {
		log.Println("Error while downloading", requestUrl, "-", err)
		return err
	}
	log.Println(n, "bytes downloaded.")
	output.Close()
	return zipFiles(zipName, []string{fileName})
}

// ZipFiles compresses one or many files into a single zip archive file
func zipFiles(filename string, files []string) error {
	newfile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newfile.Close()

	zipWriter := zip.NewWriter(newfile)
	defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {

		zipfile, err := os.OpenFile(file, os.O_CREATE, 0755)
		if err != nil {
			return err
		}

		// Get the file information
		info, err := zipfile.Stat()
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Change to deflate to gain better compression
		// see http://golang.org/pkg/archive/zip/#pkg-constants
		header.Method = zip.Deflate

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, zipfile)
		if err != nil {
			return err
		}
		zipfile.Close()
		//delete file after usage
		err = os.Remove(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func formUrl(text string) string {
	text = template.URLQueryEscaper(text)
	return fmt.Sprintf(urlFormat, key, audioFormat, speaker, speed, emotion, lang, quality, text)
}

func initYandexKey() {
	key = env.GetYandexKey()
	if len(key) == 0 {
		panic(errors.New("Should define yandex key"))
	}
}

func initFolderForAudio() {
	path := env.GetPathForAudio()
	if len(path) == 0 {
		panic(errors.New("Should define folder for caching tts"))
	}
	if err := os.Chdir(path); err != nil {
		log.Println("Created folder with path : ", path)
		if err := os.Mkdir(path, 0755); err != nil {
			panic(err)
		}
	} else {
		log.Println("Using existing folder at path : ", path)
	}
}

func base64Encode(str string) string {
	return strings.Replace(base64.StdEncoding.EncodeToString([]byte(str)), "/", "|", -1)
}
