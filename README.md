Yandex TTS processor
---
Library is used to get, zip and cache TTS responses.

#### Installation
```
import (
    "github.com/titanium-codes/yandex-tts-processor/audiocacher"
)
```

#### Usage
```
// returns path to file, that can be easily served or used in any other way
pathToFile,err := audiocacher.GetTtsForText("Your awesome text")
```

For example for query "Your awesome text" it will create file in path you defined in environmental variable and will look like
$YOUR_PATH$\Your awesome text.zip and this zip will contain 1 file audio.<Audio format you defined in env vars, or mp3 by default>
For all next same queries it will reuse this zip

Main features:
* Downloads , caches, zips all tts files in specific folder that is defined in environment
* Creates directory for files if is not exists
* Logging of all activity
* Can be simply dockerized and used from box
* Fully customizable


#### Env vars

Essential:
* YANDEX_KEY - essential
* AUDIO_PATH - essential

Optional (If not defined will use default value):
* DEFAULT_LANG - default language for tts (*Default* is russian(ru-RU), can be en-US for english, uk-UK for ukrainian or tr-TR for turkish)
* DEFAULT_SPEED - default speed of speech (*Default* 1.0, can be from 0.1 till 3.0)
* DEFAULT_EMOTION - default emotion (*Default* is neutral, can be neutral, evil or good)
* DEFAULT_SPEAKER - default speaker voice (*Default* is oksana, can be jane, oksana, alyss ,omazh ,zahar ,ermil)
* DEFAULT_QUALITY - default quality of audio, will work only for wav audio format (*Default* is hi, can be hi or lo)
* DEFAULT_AUDIO_FORMAT - default audio format (*Default* is mp3, can be wav, opus)

Additional information for optional parameters can be found [here](https://tech.yandex.ru/speechkit/cloud/doc/guide/concepts/tts-http-request-docpage/)


