# Translation Plugin

This is a translation plugin for [OpenAgents](https://openagents.com/) built with Extism Go PDK that uses the Microsoft Translator Text API via RapidAPI . The plugin takes a text input and translates it to a specified language.

## Features

- Translates text from one language (auto detect) to another using the Microsoft Translator Text API.
- Supported Language :

| Code   | Name                  |
|--------|-----------------------|
| af     | Afrikaans             |
| am     | Amharic               |
| ar     | Arabic                |
| as     | Assamese              |
| az     | Azerbaijani           |
| ba     | Bashkir               |
| bg     | Bulgarian             |
| bho    | Bhojpuri              |
| bn     | Bangla                |
| bo     | Tibetan               |
| brx    | Bodo                  |
| bs     | Bosnian               |
| ca     | Catalan               |
| cs     | Czech                 |
| cy     | Welsh                 |
| da     | Danish                |
| de     | German                |
| doi    | Dogri                 |
| dsb    | Lower Sorbian         |
| dv     | Divehi                |
| el     | Greek                 |
| en     | English               |
| es     | Spanish               |
| et     | Estonian              |
| eu     | Basque                |
| fa     | Persian               |
| fi     | Finnish               |
| fil    | Filipino              |
| fj     | Fijian                |
| fo     | Faroese               |
| fr     | French                |
| fr-CA  | French (Canada)       |
| ga     | Irish                 |
| gl     | Galician              |
| gom    | Konkani               |
| gu     | Gujarati              |
| ha     | Hausa                 |
| he     | Hebrew                |
| hi     | Hindi                 |
| hne    | Chhattisgarhi         |
| hr     | Croatian              |
| hsb    | Upper Sorbian         |
| ht     | Haitian Creole        |
| hu     | Hungarian             |
| hy     | Armenian              |
| id     | Indonesian            |
| ig     | Igbo                  |
| ikt    | Inuinnaqtun           |
| is     | Icelandic             |
| it     | Italian               |
| iu     | Inuktitut             |
| iu-Latn| Inuktitut (Latin)     |
| ja     | Japanese              |
| ka     | Georgian              |
| kk     | Kazakh                |
| km     | Khmer                 |
| kmr    | Kurdish (Northern)    |
| kn     | Kannada               |
| ko     | Korean                |
| ks     | Kashmiri              |
| ku     | Kurdish (Central)     |
| ky     | Kyrgyz                |
| ln     | Lingala               |
| lo     | Lao                   |
| lt     | Lithuanian            |
| lug    | Ganda                 |
| lv     | Latvian               |
| lzh    | Chinese (Literary)    |
| mai    | Maithili              |
| mg     | Malagasy              |
| mi     | Māori                 |
| mk     | Macedonian            |
| ml     | Malayalam             |
| mn-Cyrl| Mongolian (Cyrillic)  |
| mn-Mong| Mongolian (Traditional)|
| mni    | Manipuri              |
| mr     | Marathi               |
| ms     | Malay                 |
| mt     | Maltese               |
| mww    | Hmong Daw             |
| my     | Myanmar (Burmese)     |
| nb     | Norwegian             |
| ne     | Nepali                |
| nl     | Dutch                 |
| nso    | Sesotho sa Leboa      |
| nya    | Nyanja                |
| or     | Odia                  |
| otq    | Querétaro Otomi       |
| pa     | Punjabi               |
| pl     | Polish                |
| prs    | Dari                  |
| ps     | Pashto                |
| pt     | Portuguese (Brazil)   |
| pt-PT  | Portuguese (Portugal) |
| ro     | Romanian              |
| ru     | Russian               |
| run    | Rundi                 |
| rw     | Kinyarwanda           |
| sd     | Sindhi                |
| si     | Sinhala               |
| sk     | Slovak                |
| sl     | Slovenian             |
| sm     | Samoan                |
| sn     | Shona                 |
| so     | Somali                |
| sq     | Albanian              |
| sr-Cyrl| Serbian (Cyrillic)    |
| sr-Latn| Serbian (Latin)       |
| st     | Sesotho               |
| sv     | Swedish               |
| sw     | Swahili               |
| ta     | Tamil                 |
| te     | Telugu                |
| th     | Thai                  |
| ti     | Tigrinya              |
| tk     | Turkmen               |
| tlh-Latn| Klingon (Latin)      |
| tlh-Piqd| Klingon (pIqaD)      |
| tn     | Setswana              |
| to     | Tongan                |
| tr     | Turkish               |
| tt     | Tatar                 |
| ty     | Tahitian              |
| ug     | Uyghur                |
| uk     | Ukrainian             |
| ur     | Urdu                  |
| uz     | Uzbek (Latin)         |
| vi     | Vietnamese            |
| xh     | Xhosa                 |
| yo     | Yoruba                |
| yua    | Yucatec Maya          |
| yue    | Cantonese (Traditional)|
| zh-Hans| Chinese Simplified    |
| zh-Hant| Chinese Traditional  |
| zu     | Zulu                  |


## Prerequisites

- Go 1.16 or higher
- A RapidAPI account
- Microsoft Translator Text API key from RapidAPI

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/AndrewLWZZ/Translation-Plugin.git
    cd Translation-Plugin
    ```

2. Install dependencies:
    ```sh
    go get -u github.com/extism/go-pdk
    ```

## Usage

1. Set your RapidAPI key in the `API_KEY` variable in the `main.go` file:
    ```go
    var API_KEY string = "YOUR_API_KEY"
    ```

2. Build the plugin:
    ```sh
    tinygo build -o plugin.wasm -target wasi main.go
    ```

3. Run the plugin using Extism-CLI:
    ```sh
    extism call plugin.wasm run --input='{"text": "YOUR_TEXT?", "to": "LANGUAGE_CODE"}' --wasi --allow-host='*.rapidapi.com'
    ```

## Example

![image](https://github.com/AndrewLWZZ/Translation-Plugin/assets/170183093/290a4493-73c6-456f-856a-7b1d84d7c291)


## API Reference

The plugin uses the [Microsoft Translator Text API](https://rapidapi.com/microsoft-azure-org-microsoft-cognitive-services/api/microsoft-translator-text). Refer to the API documentation for details on supported languages and additional features.

## Acknowledgements

- Extism Go PDK: https://github.com/extism/go-pdk
- RapidAPI: https://rapidapi.com/
