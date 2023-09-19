package main

import (
	"fmt"

	"github.com/eduardolat/goeasyi18n"
)

func main() {
	// 1. Create a new i18n instance
	i18n := goeasyi18n.NewI18n(goeasyi18n.Config{
		// You can set the fallback language (optional)
		// The default value is "en"
		FallbackLanguageName: "en",
	})

	// 2. Create your translations
	enTranslations := goeasyi18n.TranslateStrings{
		{
			Key:     "hello_message",
			Default: "Hello, welcome to Go Easy i18n!",
		},
	}

	esTranslations := goeasyi18n.TranslateStrings{
		{
			Key:     "hello_message",
			Default: "¡Hola, bienvenido a Go Easy i18n!",
		},
	}

	// 3. Add your languages with their translations
	// The name of the language can be anything you want
	// You can use simple strings like "en" or "es"
	// You you can use the full language name like "english" or "spanish"
	// You can even use the IETF Language Tag like "en-US" or "es-ES"
	i18n.AddLanguage("en", enTranslations)
	i18n.AddLanguage("es", esTranslations)

	// 4. You are done! 🎉 Just get that translations!
	t1 := i18n.Translate("en", "hello_message", &goeasyi18n.Options{})
	// Or you can use the T method (it's just an alias for Translate)
	// and you can pass nil if you don't need options
	t2 := i18n.T("es", "hello_message", nil)

	fmt.Println(t1)
	fmt.Println(t2)

	/*
		Prints:
		Hello, welcome to Go Easy i18n!
		¡Hola, bienvenido a Go Easy i18n!
	*/

	// 5. (Extra) You can check if a language exists in the i18n instance
	enExists := i18n.HasLanguage("en")
	xxExists := i18n.HasLanguage("xx")

	fmt.Printf("en exists: %v\n", enExists)
	fmt.Printf("xx exists: %v\n", xxExists)

	/*
		Prints:
		en exists: true
		xx exists: false
	*/

	// 6. (Extra) You can create a translate function for a specific language
	// to prevent passing the language name every time you want to translate
	// something
	translateEn := i18n.NewLangTranslateFunc("en")
	translateEs := i18n.NewLangTranslateFunc("es")

	t3 := translateEn("hello_message", nil)
	t4 := translateEs("hello_message", nil)

	fmt.Println(t3)
	fmt.Println(t4)

	/*
		Prints:
		Hello, welcome to Go Easy i18n!
		¡Hola, bienvenido a Go Easy i18n!
	*/
}
