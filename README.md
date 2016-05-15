#Audicut

This will be a web application that will be able to take in an audio file, along with some meta data, and create a new edited version of the file.

Sample input:

```json
SoundClip: {
  "snippets": [
    {
      "start":"1:00",
      "length":"1:00"
    },
    {
      "start":"3:15",
      "length":"0:45"
    }
  ],
  "format": "MP3",
  "outFileName": "foo",
  "inFile": "<binary audio file>"
}
```
