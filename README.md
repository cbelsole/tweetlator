# Tweetlator

## Description

Tweetlator is a translator service using Azure's translator api.

## API

| Method |    Route   |  Errors  |                        Description                        |
|:------:|:----------:|:--------:|:---------------------------------------------------------:|
| GET    | /          | 500      | Returns a 200 when healthy                                |
| GET    | /health    | 500      | Returns a 200 when healthy                                |
| POST   | /translate | 400, 500 | Given a valid translate request, returns translated words |

### Translate Request

```json
{
  "from": "string",
  "to": "string",
  "tweet": "string"
}
```

### Translate Response

```json
{
  "tweet": "string"
}
```

### Available Languages

|                         |                         |                         |
|:-----------------------:|:-----------------------:|:-----------------------:|
| Afrikaans               | Haitian Creole          | Querétaro Otomi         |
| Arabic                  | Hebrew                  | Romanian                |
| Bangla                  | Hindi                   | Russian                 |
| Bosnian (Latin)         | Hmong Daw               | Samoan                  |
| Bulgarian               | Hungarian               | Serbian (Cyrillic)      |
| Cantonese (Traditional) | Indonesian              | Serbian (Latin)         |
| Catalan                 | Italian                 | Slovak                  |
| Chinese Simplified      | Japanese                | Slovenian               |
| Chinese Traditional     | Kiswahili               | Spanish                 |
| Croatian                | Klingon                 | Swedish                 |
| Czech                   | Klingon (pIqaD)         | Tahitian                |
| Danish                  | Korean                  | Thai                    |
| Dutch                   | Latvian                 | Tongan                  |
| English                 | Lithuanian              | Turkish                 |
| Estonian                | Malagasy                | Ukrainian               |
| Fijian                  | Malay                   | Urdu                    |
| Filipino                | Maltese                 | Vietnamese              |
| Finnish                 | Norwegian Bokmål        | Welsh                   |
| French                  | Persian                 | Yucatec Maya            |
| German                  | Polish                  |                         |
| Greek                   | Portuguese              |                         |
