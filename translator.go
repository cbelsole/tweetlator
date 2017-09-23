package tweetlator

import (
	"errors"
	"strings"

	"github.com/st3v/translator"
	"github.com/st3v/translator/microsoft"
)

var languages = []translator.Language{
	translator.Language{Name: "Bosnian (Latin)", Code: "bs-Latn"},
	translator.Language{Name: "Estonian", Code: "et"},
	translator.Language{Name: "Afrikaans", Code: "af"},
	translator.Language{Name: "Catalan", Code: "ca"},
	translator.Language{Name: "Persian", Code: "fa"},
	translator.Language{Name: "Chinese Simplified", Code: "zh-CHS"},
	translator.Language{Name: "Czech", Code: "cs"},
	translator.Language{Name: "Chinese Traditional", Code: "zh-CHT"},
	translator.Language{Name: "Russian", Code: "ru"},
	translator.Language{Name: "Latvian", Code: "lv"},
	translator.Language{Name: "Filipino", Code: "fil"},
	translator.Language{Name: "Samoan", Code: "sm"},
	translator.Language{Name: "Finnish", Code: "fi"},
	translator.Language{Name: "Bangla", Code: "bn"},
	translator.Language{Name: "Hebrew", Code: "he"},
	translator.Language{Name: "Lithuanian", Code: "lt"},
	translator.Language{Name: "Fijian", Code: "fj"},
	translator.Language{Name: "Dutch", Code: "nl"},
	translator.Language{Name: "English", Code: "en"},
	translator.Language{Name: "Arabic", Code: "ar"},
	translator.Language{Name: "Croatian", Code: "hr"},
	translator.Language{Name: "Slovak", Code: "sk"},
	translator.Language{Name: "Danish", Code: "da"},
	translator.Language{Name: "French", Code: "fr"},
	translator.Language{Name: "German", Code: "de"},
	translator.Language{Name: "Bulgarian", Code: "bg"},
	translator.Language{Name: "Maltese", Code: "mt"},
	translator.Language{Name: "Serbian (Latin)", Code: "sr-Latn"},
	translator.Language{Name: "Hmong Daw", Code: "mww"},
	translator.Language{Name: "Greek", Code: "el"},
	translator.Language{Name: "Swedish", Code: "sv"},
	translator.Language{Name: "Spanish", Code: "es"},
	translator.Language{Name: "Malay", Code: "ms"},
	translator.Language{Name: "Yucatec Maya", Code: "yua"},
	translator.Language{Name: "Malagasy", Code: "mg"},
	translator.Language{Name: "Norwegian Bokmål", Code: "no"},
	translator.Language{Name: "Indonesian", Code: "id"},
	translator.Language{Name: "Tongan", Code: "to"},
	translator.Language{Name: "Haitian Creole", Code: "ht"},
	translator.Language{Name: "Slovenian", Code: "sl"},
	translator.Language{Name: "Thai", Code: "th"},
	translator.Language{Name: "Tahitian", Code: "ty"},
	translator.Language{Name: "Querétaro Otomi", Code: "otq"},
	translator.Language{Name: "Italian", Code: "it"},
	translator.Language{Name: "Portuguese", Code: "pt"},
	translator.Language{Name: "Hindi", Code: "hi"},
	translator.Language{Name: "Polish", Code: "pl"},
	translator.Language{Name: "Serbian (Cyrillic)", Code: "sr-Cyrl"},
	translator.Language{Name: "Klingon (pIqaD)", Code: "tlh-Qaak"},
	translator.Language{Name: "Korean", Code: "ko"},
	translator.Language{Name: "Welsh", Code: "cy"},
	translator.Language{Name: "Turkish", Code: "tr"},
	translator.Language{Name: "Japanese", Code: "ja"},
	translator.Language{Name: "Ukrainian", Code: "uk"},
	translator.Language{Name: "Urdu", Code: "ur"},
	translator.Language{Name: "Romanian", Code: "ro"},
	translator.Language{Name: "Vietnamese", Code: "vi"},
	translator.Language{Name: "Hungarian", Code: "hu"},
	translator.Language{Name: "Klingon", Code: "tlh"},
	translator.Language{Name: "Cantonese (Traditional)", Code: "yue"},
	translator.Language{Name: "Kiswahili", Code: "sw"},
}

type translate struct {
	client translator.Translator
}

func NewTranslate(key string) *translate {
	return &translate{
		client: microsoft.NewTranslator(key),
	}
}

func (t *translate) translate(text, to, from string) (string, error) {
	toL, err := t.findLanguage(to)
	if err != nil {
		return "", err
	}

	fromL, err := t.findLanguage(from)
	if err != nil {
		return "", err
	}

	trans, err := t.client.Translate(text, fromL.Code, toL.Code)

	if err != nil {
		return "", err
	}

	return trans, nil
}

func (t *translate) findLanguage(lang string) (*translator.Language, error) {
	for _, language := range languages {
		if strings.ToLower(lang) == strings.ToLower(language.Name) {
			return &language, nil
		}
	}
	return nil, errors.New("language not found")
}

func (t *translate) health() error {
	if _, err := t.client.Languages(); err != nil {
		return err
	}
	return nil
}
