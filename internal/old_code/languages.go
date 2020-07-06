package main

// languages.go contains code for foreign/alien language word generation

import (
	"fmt"
	"strings"
)

// generateZhodaniWord generates a single Zhodani word and returns a string containing the word and the structure
func generateZhodaniWord() string {
	syll := D6()

	var finalWord string
	var structure string
	useAlternate := false

	// For each syllable
	for m := 0; m < syll; m++ {

		if !useAlternate {
			// Basic structure table
			roll := Dice(36)

			if roll <= 3 {
				finalWord += getZhodVowel()
				structure += "[V[]"
				useAlternate = true
			} else if roll <= 6 {
				finalWord += getZhodInitCons() + getZhodVowel()
				structure += "[CV[]"
				useAlternate = true
			} else if roll <= 15 {
				finalWord += getZhodVowel() + getZhodFinalCons()
				structure += "[VC[]"
				useAlternate = false
			} else {
				finalWord += getZhodInitCons() + getZhodVowel() + getZhodFinalCons()
				structure += "[CVC[]"
				useAlternate = false
			}

		} else {
			// Alternate structure table
			roll := Dice(36)

			if roll <= 6 {
				finalWord += getZhodVowel()
				structure += "[a:V[]"
				useAlternate = true
			} else if roll <= 12 {
				finalWord += getZhodInitCons() + getZhodVowel()
				structure += "[a:CV[]"
				useAlternate = true
			} else if roll <= 18 {
				finalWord += getZhodVowel() + getZhodFinalCons()
				structure += "[a:VC[]"
				useAlternate = false
			} else {
				finalWord += getZhodInitCons() + getZhodVowel() + getZhodFinalCons()
				structure += "[a:CVC[]"
				useAlternate = false
			}
		}
	}
	return strings.Title(finalWord) + " (" + structure + ")"
}

// generateZhodaniWords generates a bunch of Zhodani words.
func generateZhodaniWords() {
	logPane.Log("Generating Zhodani words")

	mainPageFocus()
	// Do this a number of times.
	for n := 0; n < config.PageOptionSize; n++ {
		// Generate a word
		fmt.Fprintln(mainPage, generateZhodaniWord())
	}
	fmt.Fprintln(mainPage)
}

// getZhodInitCons returns an Zhodani initial consonant based on the frequency table.
func getZhodInitCons() string {
	roll := Dice(127)

	if roll <= 3 {
		return "b"
	} else if roll <= 5 {
		return "bl"
	} else if roll <= 8 {
		return "br"
	} else if roll <= 11 {
		return "ch"
	} else if roll <= 18 {
		return "cht"
	} else if roll <= 24 {
		return "d"
	} else if roll <= 28 {
		return "dl"
	} else if roll <= 31 {
		return "dr"
	} else if roll <= 34 {
		return "f"
	} else if roll <= 36 {
		return "fl"
	} else if roll <= 38 {
		return "fr"
	} else if roll <= 42 {
		return "j"
	} else if roll <= 45 {
		return "jd"
	} else if roll <= 48 {
		return "k"
	} else if roll <= 49 {
		return "kl"
	} else if roll <= 50 {
		return "kr"
	} else if roll <= 52 {
		return "l"
	} else if roll <= 53 {
		return "m"
	} else if roll <= 58 {
		return "n"
	} else if roll <= 62 {
		return "p"
	} else if roll <= 66 {
		return "pl"
	} else if roll <= 68 {
		return "pr"
	} else if roll <= 69 {
		return "q"
	} else if roll <= 70 {
		return "ql"
	} else if roll <= 71 {
		return "qr"
	} else if roll <= 74 {
		return "r"
	} else if roll <= 78 {
		return "s"
	} else if roll <= 82 {
		return "sh"
	} else if roll <= 86 {
		return "sht"
	} else if roll <= 90 {
		return "st"
	} else if roll <= 93 {
		return "t"
	} else if roll <= 99 {
		return "tl"
	} else if roll <= 101 {
		return "ts"
	} else if roll <= 104 {
		return "v"
	} else if roll <= 105 {
		return "vl"
	} else if roll <= 106 {
		return "vr"
	} else if roll <= 108 {
		return "y"
	} else if roll <= 111 {
		return "z"
	} else if roll <= 117 {
		return "zd"
	} else if roll <= 121 {
		return "zh"
	} else {
		return "zhd"
	}
}

// getZhodVowel returns a Zhodani vowel based on the frequency table.
func getZhodVowel() string {
	roll := Dice(31)

	if roll <= 7 {
		return "a"
	} else if roll <= 15 {
		return "e"
	} else if roll <= 20 {
		return "i"
	} else if roll <= 24 {
		return "ia"
	} else if roll <= 28 {
		return "ie"
	} else if roll <= 30 {
		return "o"
	} else {
		return "r"
	}

}

// getZhodFinalCons returns a Zhodani final consonant based on the frequency table.
func getZhodFinalCons() string {
	roll := Dice(122)

	if roll <= 1 {
		return "b"
	} else if roll <= 5 {
		return "bl"
	} else if roll <= 9 {
		return "br"
	} else if roll <= 12 {
		return "ch"
	} else if roll <= 14 {
		return "d"
	} else if roll <= 18 {
		return "dl"
	} else if roll <= 22 {
		return "dr"
	} else if roll <= 25 {
		return "f"
	} else if roll <= 28 {
		return "fl"
	} else if roll <= 31 {
		return "fr"
	} else if roll <= 33 {
		return "j"
	} else if roll <= 34 {
		return "k"
	} else if roll <= 36 {
		return "kl"
	} else if roll <= 37 {
		return "kr"
	} else if roll <= 44 {
		return "l"
	} else if roll <= 45 {
		return "m"
	} else if roll <= 46 {
		return "n"
	} else if roll <= 50 {
		return "nch"
	} else if roll <= 53 {
		return "nj"
	} else if roll <= 56 {
		return "ns"
	} else if roll <= 60 {
		return "nsh"
	} else if roll <= 62 {
		return "nt"
	} else if roll <= 64 {
		return "nts"
	} else if roll <= 67 {
		return "nz"
	} else if roll <= 71 {
		return "nzh"
	} else if roll <= 72 {
		return "p"
	} else if roll <= 76 {
		return "pl"
	} else if roll <= 80 {
		return "pr"
	} else if roll <= 81 {
		return "q"
	} else if roll <= 82 {
		return "ql"
	} else if roll <= 83 {
		return "qr"
	} else if roll <= 86 {
		return "r"
	} else if roll <= 90 {
		return "sh"
	} else if roll <= 92 {
		return "t"
	} else if roll <= 96 {
		return "ts"
	} else if roll <= 101 {
		return "tl"
	} else if roll <= 104 {
		return "v"
	} else if roll <= 106 {
		return "vl"
	} else if roll <= 109 {
		return "vr"
	} else if roll <= 114 {
		return "z"
	} else if roll <= 118 {
		return "zh"
	} else {
		return "'"
	}

}

// generateVilaniWords prints a bunch of generated Vilani words based on the published frequency tables.
func generateVilaniWords() {

	logPane.Log("Generating Vilani words")
	mainPageFocus()

	// Do this a number of times.
	for n := 0; n < config.PageOptionSize; n++ {
		fmt.Fprintln(mainPage, generateVilaniWord())
	}
	fmt.Fprintln(mainPage)

}

// generateVilaniWord generates a single Vilani word
func generateVilaniWord() string {
	// Generate a word
	syll := D6()

	var finalWord string
	var structure string
	useAlternate := false

	// For each syllable
	for m := 0; m < syll; m++ {

		if !useAlternate {
			// Basic structure table
			roll := Dice(36)

			if roll <= 6 {
				finalWord += getVilVowel()
				structure += "[V[]"
				useAlternate = true
			} else if roll <= 21 {
				finalWord += getVilInitCons() + getVilVowel()
				structure += "[CV[]"
				useAlternate = true
			} else if roll <= 29 {
				finalWord += getVilVowel() + getVilFinalCons()
				structure += "[VC[]"
				useAlternate = false
			} else {
				finalWord += getVilInitCons() + getVilVowel() + getVilFinalCons()
				structure += "[CVC[]"
				useAlternate = false
			}

		} else {
			// Alternate structure table
			roll := Dice(36)

			if roll <= 21 {
				finalWord += getVilInitCons() + getVilVowel()
				structure += "[a:CV[]"
				useAlternate = false
			} else {
				finalWord += getVilInitCons() + getVilVowel() + getVilFinalCons()
				structure += "[a:CVC[]"
				useAlternate = true
			}
		}
	}
	return strings.Title(finalWord) + " (" + structure + ")"

}

// getVilInitCons generates a Vilani initial consonant based on the frequency table.
func getVilInitCons() string {
	roll := Dice(216)

	if roll <= 39 {
		return "k"
	} else if roll <= 78 {
		return "g"
	} else if roll <= 99 {
		return "m"
	} else if roll <= 120 {
		return "d"
	} else if roll <= 141 {
		return "l"
	} else if roll <= 162 {
		return "sh"
	} else if roll <= 180 {
		return "kh"
	} else if roll <= 190 {
		return "n"
	} else if roll <= 200 {
		return "s"
	} else if roll <= 204 {
		return "p"
	} else if roll <= 208 {
		return "b"
	} else if roll <= 212 {
		return "z"
	} else {
		return "r"
	}
}

// getVilVowel gets a Vilani vowel based on the frequency table.
func getVilVowel() string {
	roll := Dice(216)

	if roll <= 67 {
		return "a"
	} else if roll <= 84 {
		return "e"
	} else if roll <= 143 {
		return "i"
	} else if roll <= 184 {
		return "u"
	} else if roll <= 192 {
		return "aa"
	} else if roll <= 208 {
		return "ii"
	} else {
		return "uu"
	}
}

// getVilFinalCons gets a Vilani final consonant based on the frequency table.
func getVilFinalCons() string {
	roll := Dice(216)

	if roll <= 76 {
		return "r"
	} else if roll <= 102 {
		return "n"
	} else if roll <= 139 {
		return "m"
	} else if roll <= 165 {
		return "sh"
	} else if roll <= 180 {
		return "g"
	} else if roll <= 191 {
		return "s"
	} else if roll <= 204 {
		return "d"
	} else if roll <= 210 {
		return "p"
	} else {
		return "k"
	}

}

// generateVargrWords generates a bunch of Vargr words.
func generateVargrWords() {

	logPane.Log("Generating Vagr words")
	mainPageFocus()
	// Do this a number of times.
	for n := 0; n < config.PageOptionSize; n++ {
		fmt.Fprintln(mainPage, generateVargrWord())
	}
	fmt.Fprintln(mainPage)
}

// generateVargrWord generates a single Vargr word based on the published frequency tables.
func generateVargrWord() string {
	// Generate a word
	syll := D6()

	var finalWord string
	var structure string
	useAlternate := false

	// For each syllable
	for m := 0; m < syll; m++ {

		if !useAlternate {
			// Basic structure table
			roll := Dice(36)

			if roll <= 6 {
				finalWord += getVargrVowel()
				structure += "[V[]"
				useAlternate = true
			} else if roll <= 18 {
				finalWord += getVargrVowel() + getVargrFinalCons()
				structure += "[VC[]"
				useAlternate = false
			} else if roll <= 22 {
				finalWord += getVargrInitCons() + getVargrVowel()
				structure += "[CV[]"
				useAlternate = true
			} else {
				finalWord += getVargrInitCons() + getVargrVowel() + getVargrFinalCons()
				structure += "[CVC[]"
				useAlternate = false
			}
		} else {
			// Alternate structure table
			roll := Dice(36)

			if roll <= 18 {
				finalWord += getVargrInitCons() + getVargrVowel()
				structure += "[a:CV[]"
				useAlternate = true
			} else {
				finalWord += getVargrInitCons() + getVargrVowel() + getVargrFinalCons()
				structure += "[a:CVC[]"
				useAlternate = false
			}
		}
	}
	return strings.Title(finalWord) + " (" + structure + ")"

}

// getVargrInitCons generates a Vargr initial consonant based on the frequency table.
func getVargrInitCons() string {
	roll := Dice(26)

	if roll <= 5 {
		return "d"
	} else if roll <= 10 {
		return "dh"
	} else if roll <= 13 {
		return "dz"
	} else if roll <= 17 {
		return "f"
	} else if roll <= 27 {
		return "g"
	} else if roll <= 33 {
		return "gh"
	} else if roll <= 35 {
		return "gn"
	} else if roll <= 39 {
		return "gv"
	} else if roll <= 43 {
		return "gz"
	} else if roll <= 53 {
		return "k"
	} else if roll <= 56 {
		return "kf"
	} else if roll <= 62 {
		return "kh"
	} else if roll <= 65 {
		return "kn"
	} else if roll <= 68 {
		return "ks"
	} else if roll <= 72 {
		return "l"
	} else if roll <= 76 {
		return "ll"
	} else if roll <= 78 {
		return "n"
	} else if roll <= 80 {
		return "ng"
	} else if roll <= 85 {
		return "r"
	} else if roll <= 89 {
		return "rr"
	} else if roll <= 94 {
		return "s"
	} else if roll <= 98 {
		return "t"
	} else if roll <= 102 {
		return "th"
	} else if roll <= 104 {
		return "ts"
	} else if roll <= 109 {
		return "v"
	} else {
		return "z"
	}
}

// getVargrVowel gets a Vargr vowel based on the frequency table.
func getVargrVowel() string {
	roll := Dice(26)

	if roll <= 5 {
		return "a"
	} else if roll <= 9 {
		return "ae"
	} else if roll <= 11 {
		return "e"
	} else if roll <= 12 {
		return "i"
	} else if roll <= 16 {
		return "o"
	} else if roll <= 18 {
		return "oe"
	} else if roll <= 20 {
		return "ou"
	} else if roll <= 23 {
		return "u"
	} else {
		return "ue"
	}
}

// getVargrFinalCons gets a Vargr final consonant based on the frequency table.
func getVargrFinalCons() string {
	roll := Dice(43)

	if roll <= 1 {
		return "dh"
	} else if roll <= 2 {
		return "dz"
	} else if roll <= 5 {
		return "g"
	} else if roll <= 7 {
		return "gh"
	} else if roll <= 8 {
		return "ghz"
	} else if roll <= 9 {
		return "gz"
	} else if roll <= 11 {
		return "k"
	} else if roll <= 13 {
		return "kh"
	} else if roll <= 14 {
		return "khs"
	} else if roll <= 15 {
		return "ks"
	} else if roll <= 17 {
		return "l"
	} else if roll <= 18 {
		return "ll"
	} else if roll <= 23 {
		return "n"
	} else if roll <= 28 {
		return "ng"
	} else if roll <= 31 {
		return "r"
	} else if roll <= 34 {
		return "rr"
	} else if roll <= 35 {
		return "rrg"
	} else if roll <= 36 {
		return "rrgh"
	} else if roll <= 37 {
		return "rs"
	} else if roll <= 38 {
		return "rz"
	} else if roll <= 39 {
		return "s"
	} else if roll <= 40 {
		return "th"
	} else if roll <= 41 {
		return "ts"
	} else {
		return "z"
	}
}

// generateAslanWord generates a single Aslan word based on the published frequency tables.
func generateAslanWord() string {
	// Generate a word
	syll := D6()

	var finalWord string
	var structure string
	useAlternate := false

	// For each syllable
	for m := 0; m < syll; m++ {

		if !useAlternate {
			// Basic structure table
			roll := Dice(36)

			if roll <= 13 {
				finalWord += getAslanVowel()
				structure += "[V[]"
				useAlternate = false
			} else if roll <= 22 {
				finalWord += getAslanInitCons() + getAslanVowel()
				structure += "[CV[]"
				useAlternate = false
			} else if roll <= 30 {
				finalWord += getAslanVowel() + getAslanFinalCons()
				structure += "[VC[]"
				useAlternate = true
			} else {
				finalWord += getAslanInitCons() + getAslanVowel() + getAslanFinalCons()
				structure += "[CVC[]"
				useAlternate = true
			}
		} else {
			// Alternate structure table
			roll := Dice(36)

			if roll <= 15 {
				finalWord += getAslanVowel()
				structure += "[a:V[]"
				useAlternate = false
			} else {
				finalWord += getAslanVowel() + getAslanFinalCons()
				structure += "[a:VC[]"
			}
		}
	}
	return strings.Title(finalWord) + " (" + structure + ")"
}

// generateAslanWords generates a bunch of Aslan words.
func generateAslanWords() {
	logPane.Log("Generating Aslan words")

	mainPageFocus()
	// Do this a number of times.
	for n := 0; n < config.PageOptionSize; n++ {
		fmt.Fprintln(mainPage, generateAslanWord())
	}
	fmt.Fprintln(mainPage)
}

// getAslanInitCons gets an Aslan initial consonant based on the frequency table.
func getAslanInitCons() string {
	roll := Dice(87)

	if roll <= 5 {
		return "f"
	} else if roll <= 9 {
		return "ft"
	} else if roll <= 16 {
		return "h"
	} else if roll <= 18 {
		return "hf"
	} else if roll <= 23 {
		return "hk"
	} else if roll <= 26 {
		return "hl"
	} else if roll <= 29 {
		return "hr"
	} else if roll <= 34 {
		return "ht"
	} else if roll <= 36 {
		return "hw"
	} else if roll <= 43 {
		return "k"
	} else if roll <= 49 {
		return "kh"
	} else if roll <= 53 {
		return "kht"
	} else if roll <= 57 {
		return "kt"
	} else if roll <= 59 {
		return "l"
	} else if roll <= 62 {
		return "r"
	} else if roll <= 66 {
		return "s"
	} else if roll <= 69 {
		return "st"
	} else if roll <= 77 {
		return "t"
	} else if roll <= 79 {
		return "tl"
	} else if roll <= 81 {
		return "tr"
	} else if roll <= 87 {
		return "w"
	} else if roll <= 87 {
		return ""
	} else if roll <= 87 {
		return ""
	} else if roll <= 87 {
		return ""
	} else if roll <= 87 {
		return ""
	} else {
		return ""
	}

}

// getAslanVowel gets an Aslan vowel based on the frequency table.
func getAslanVowel() string {
	roll := Dice(216)

	if roll <= 41 {
		return "a"
	} else if roll <= 52 {
		return "ai"
	} else if roll <= 60 {
		return "ao"
	} else if roll <= 64 {
		return "au"
	} else if roll <= 90 {
		return "e"
	} else if roll <= 114 {
		return "ea"
	} else if roll <= 127 {
		return "ei"
	} else if roll <= 143 {
		return "i"
	} else if roll <= 155 {
		return "iy"
	} else if roll <= 163 {
		return "o"
	} else if roll <= 167 {
		return "oa"
	} else if roll <= 175 {
		return "oi"
	} else if roll <= 180 {
		return "ou"
	} else if roll <= 184 {
		return "u"
	} else if roll <= 188 {
		return "ua"
	} else if roll <= 192 {
		return "ui"
	} else if roll <= 200 {
		return "ya"
	} else if roll <= 208 {
		return "ye"
	} else if roll <= 212 {
		return "yo"
	} else {
		return "yu"
	}
}

// getAslanFinalCons gets an Aslan final consonant based on the frequency table.
func getAslanFinalCons() string {
	roll := Dice(47)

	if roll <= 10 {
		return "h"
	} else if roll <= 14 {
		return "kh"
	} else if roll <= 21 {
		return "l"
	} else if roll <= 24 {
		return "lr"
	} else if roll <= 29 {
		return "r"
	} else if roll <= 33 {
		return "rl"
	} else if roll <= 38 {
		return "s"
	} else if roll <= 44 {
		return "w"
	} else {
		return "'"
	}
}

// generateDarrianWord generates a single Darrian word based on the published frequency tables.
func generateDarrianWord() string {
	// Generate a word
	syll := D6()

	var finalWord string
	var structure string
	useAlternate := false

	// For each syllable
	for m := 0; m < syll; m++ {

		if !useAlternate {
			// Basic structure table
			roll := Dice(36)

			if roll <= 27 {
				finalWord += getDarrianInitCons() + getDarrianVowel() + getDarrianFinalCons()
				structure += "[CVC[]"
				useAlternate = true
			} else {
				finalWord += getDarrianInitCons() + getDarrianVowel()
				structure += "[CV[]"
				useAlternate = false
			}
		} else {
			// Alternate structure table
			roll := Dice(36)

			if roll <= 27 {
				finalWord += getDarrianVowel() + getDarrianFinalCons()
				structure += "[a:VC[]"
				useAlternate = true
			} else {
				finalWord += getDarrianVowel()
				structure += "[a:V[]"
				useAlternate = false
			}
		}
	}
	return strings.Title(finalWord) + " (" + structure + ")"

}

// generateDarrianWords generates a bunch of Darrian words.
func generateDarrianWords() {

	logPane.Log("Generating Darrian words")
	mainPageFocus()

	// Do this a number of times.
	for n := 0; n < config.PageOptionSize; n++ {
		fmt.Fprintln(mainPage, generateDarrianWord())
	}
	fmt.Fprintln(mainPage)
}

// getDarrianInitCons gets an initial Darrian syllable consonant.
func getDarrianInitCons() string {
	roll := Dice(209)

	if roll <= 17 {
		return "b"
	} else if roll <= 39 {
		return "d"
	} else if roll <= 46 {
		return "g"
	} else if roll <= 58 {
		return "p"
	} else if roll <= 66 {
		return "t"
	} else if roll <= 73 {
		return "th"
	} else if roll <= 78 {
		return "k"
	} else if roll <= 88 {
		return "m"
	} else if roll <= 110 {
		return "n"
	} else if roll <= 132 {
		return "z"
	} else if roll <= 142 {
		return "l"
	} else if roll <= 156 {
		return "r"
	} else if roll <= 162 {
		return "y"
	} else if roll <= 166 {
		return "zb"
	} else if roll <= 171 {
		return "zd"
	} else if roll <= 173 {
		return "zg"
	} else if roll <= 176 {
		return "zl"
	} else if roll <= 181 {
		return "mb"
	} else if roll <= 186 {
		return "nd"
	} else if roll <= 189 {
		return "ngg"
	} else if roll <= 194 {
		return "ry"
	} else if roll <= 197 {
		return "ly"
	} else if roll <= 202 {
		return "lz"
	} else {
		return "ld"
	}
}

// getDarrianVowel gets a Darrian vowel.
func getDarrianVowel() string {
	roll := Dice(45)

	if roll <= 8 {
		return "a"
	} else if roll <= 16 {
		return "e"
	} else if roll <= 21 {
		return "eh"
	} else if roll <= 30 {
		return "i"
	} else if roll <= 38 {
		return "ih"
	} else if roll <= 43 {
		return "o"
	} else {
		return "u"
	}
}

// getDarrianFinalCons gets a final Darrian consonant.
func getDarrianFinalCons() string {
	roll := Dice(216)

	if roll <= 9 {
		return "bh"
	} else if roll <= 18 {
		return "dh"
	} else if roll <= 24 {
		return "gh"
	} else if roll <= 30 {
		return "p"
	} else if roll <= 36 {
		return "t"
	} else if roll <= 45 {
		return "k"
	} else if roll <= 74 {
		return "n"
	} else if roll <= 86 {
		return "ng"
	} else if roll <= 109 {
		return "l"
	} else if roll <= 138 {
		return "r"
	} else if roll <= 156 {
		return "s"
	} else if roll <= 171 {
		return "m"
	} else if roll <= 177 {
		return "mb"
	} else if roll <= 183 {
		return "nd"
	} else if roll <= 186 {
		return "ngg"
	} else if roll <= 192 {
		return "yr"
	} else if roll <= 195 {
		return "ly"
	} else if roll <= 198 {
		return "ny"
	} else if roll <= 201 {
		return "lbh"
	} else if roll <= 207 {
		return "lz"
	} else {
		return "ld"
	}
}

// generateKkreeWord generates a single K'kree word based on the published frequency tables.
func generateKkreeWord() string {
	syll := D6()

	var finalWord string
	var structure string
	useTable := 1

	// For each syllable
	for m := 0; m < syll; m++ {

		switch useTable {
		case 1:
			roll := Dice(13)
			switch {
			case roll == 1:
				finalWord += getKkreeVowel()
				structure += "[V[]"
				useTable = 3
			case roll <= 7:
				finalWord += getKkreeInitCons() + getKkreeVowel()
				structure += "[CV[]"
				useTable = 3
			case roll <= 9:
				finalWord += getKkreeVowel() + getKkreeFinalCons()
				structure += "[VC[]"
				useTable = 2
			default:
				finalWord += getKkreeInitCons() + getKkreeVowel() + getKkreeFinalCons()
				structure += "[CVC[]"
				m = syll
			}
		case 2:
			roll := Dice(3)
			switch {
			case roll == 1:
				finalWord += getKkreeVowel()
				structure += "[V[]"
				useTable = 3
			default:
				finalWord += getKkreeVowel() + getKkreeFinalCons()
				structure += "[VC[]"
				useTable = 2
			}
		default:
			roll := Dice(5)
			switch {
			case roll <= 3:
				finalWord += getKkreeInitCons() + getKkreeVowel()
				structure += "[CV[]"
				useTable = 3
			default:
				finalWord += getKkreeInitCons() + getKkreeVowel() + getKkreeFinalCons()
				structure += "[CVC[]"
				m = syll
			}
		}
	}
	return strings.Title(finalWord) + " (" + structure + ")"

}

// generateKkreeWords generates a bunch of K'kree words.
func generateKkreeWords() {

	logPane.Log("Generating K'kree words")
	mainPageFocus()
	// Do this a number of times.
	for n := 0; n < config.PageOptionSize; n++ {
		// Generate a word
		fmt.Fprintln(mainPage, generateKkreeWord())
	}
	fmt.Fprintln(mainPage)
}

// getKkreeVowel generates a K'kree vowel.
func getKkreeVowel() string {

	roll := Dice(60)

	if roll <= 19 {
		return "a"
	} else if roll <= 21 {
		return "aa"
	} else if roll <= 24 {
		return "e"
	} else if roll <= 28 {
		return "ee"
	} else if roll <= 34 {
		return "i"
	} else if roll <= 36 {
		return "ii"
	} else if roll <= 37 {
		return "o"
	} else if roll <= 39 {
		return "oo"
	} else if roll <= 45 {
		return "u"
	} else if roll <= 47 {
		return "uu"
	} else if roll <= 55 {
		return "'"
	} else if roll <= 58 {
		return "!"
	} else if roll <= 59 {
		return "!!"
	} else {
		return "!'"
	}
}

// getKkreeInitCons generates a Kkree initial consonant.
func getKkreeInitCons() string {
	roll := Dice(98)

	if roll <= 1 {
		return "b"
	} else if roll <= 4 {
		return "g"
	} else if roll <= 10 {
		return "gh"
	} else if roll <= 14 {
		return "gn"
	} else if roll <= 16 {
		return "gr"
	} else if roll <= 17 {
		return "gz"
	} else if roll <= 19 {
		return "hk"
	} else if roll <= 43 {
		return "k"
	} else if roll <= 53 {
		return "kr"
	} else if roll <= 54 {
		return "kt"
	} else if roll <= 59 {
		return "l"
	} else if roll <= 61 {
		return "mb"
	} else if roll <= 62 {
		return "mb"
	} else if roll <= 66 {
		return "n"
	} else if roll <= 67 {
		return "p"
	} else if roll <= 79 {
		return "r"
	} else if roll <= 82 {
		return "rr"
	} else if roll <= 89 {
		return "t"
	} else if roll <= 91 {
		return "tr"
	} else if roll <= 95 {
		return "x"
	} else if roll <= 96 {
		return "xk"
	} else if roll <= 97 {
		return "xr"
	} else {
		return "xt"
	}

}

// getKkreeFinalCons returns a K'kree final consonant.
func getKkreeFinalCons() string {
	roll := Dice(42)

	if roll <= 1 {
		return "b"
	} else if roll <= 3 {
		return "gh"
	} else if roll <= 4 {
		return "gh"
	} else if roll <= 5 {
		return "gr"
	} else if roll <= 11 {
		return "k"
	} else if roll <= 14 {
		return "kr"
	} else if roll <= 16 {
		return "l"
	} else if roll <= 17 {
		return "m"
	} else if roll <= 19 {
		return "n"
	} else if roll <= 22 {
		return "ngg"
	} else if roll <= 23 {
		return "p"
	} else if roll <= 31 {
		return "r"
	} else if roll <= 35 {
		return "rr"
	} else if roll <= 38 {
		return "t"
	} else if roll <= 41 {
		return "x"
	} else {
		return "xk"
	}

}

// generateDroyneWord generates a single Droyne word based on the published frequency tables.
func generateDroyneWord() string {
	// Generate a word
	syll := D6()

	var finalWord string
	var structure string
	useTable := 1

	// For each syllable
	for m := 0; m < syll; m++ {

		switch useTable {
		case 1:
			roll := Dice(36)
			switch {
			case roll <= 7:
				finalWord += getDroyneVowel()
				structure += "[V[]"
				useTable = 1
			case roll <= 18:
				finalWord += getDroyneInitCons() + getDroyneVowel()
				structure += "[CV[]"
				useTable = 1
			case roll <= 29:
				finalWord += getDroyneVowel() + getDroyneFinalCons()
				structure += "[VC[]"
				useTable = 2
			default:
				finalWord += getDroyneInitCons() + getDroyneVowel() + getDroyneFinalCons()
				structure += "[CVC[]"
				useTable = 2
			}
		default:
			roll := D6()
			switch {
			case roll == 1:
				finalWord += getDroyneVowel()
				structure += "[V[]"
				useTable = 1
			case roll == 2:
				finalWord += getDroyneInitCons() + getDroyneVowel()
				structure += "[CV[]"
				useTable = 1
			case roll == 3:
				finalWord += getDroyneVowel() + getDroyneFinalCons()
				structure += "[VC[]"
				useTable = 2
			default:
				finalWord += getDroyneInitCons() + getDroyneVowel() + getDroyneFinalCons()
				structure += "[CVC[]"
				useTable = 2
			}
		}
	}
	return strings.Title(finalWord) + " (" + structure + ")"
}

// generateDroyneWords generates a bunch of Droyne words.
func generateDroyneWords() {
	logPane.Log("Generating Droyne words")
	mainPageFocus()
	// Do this a number of times.
	for n := 0; n < config.PageOptionSize; n++ {
		fmt.Fprintln(mainPage, generateDroyneWord())
	}
	fmt.Fprintln(mainPage)
}

// getDroyneVowel gets a Droyne Vowel.
func getDroyneVowel() string {

	roll := Dice(58)

	if roll <= 7 {
		return "a"
	} else if roll <= 15 {
		return "ay"
	} else if roll <= 20 {
		return "e"
	} else if roll <= 24 {
		return "i"
	} else if roll <= 28 {
		return "o"
	} else if roll <= 30 {
		return "oy"
	} else if roll <= 31 {
		return "u"
	} else if roll <= 40 {
		return "ya"
	} else if roll <= 47 {
		return "yo"
	} else {
		return "yu"
	}
}

// getDroyneInitCons returns a Droyne initial consonant.
func getDroyneInitCons() string {
	roll := Dice(216)

	if roll <= 8 {
		return "b"
	} else if roll <= 12 {
		return "br"
	} else if roll <= 24 {
		return "d"
	} else if roll <= 29 {
		return "dr"
	} else if roll <= 42 {
		return "f"
	} else if roll <= 55 {
		return "h"
	} else if roll <= 68 {
		return "k"
	} else if roll <= 71 {
		return "kr"
	} else if roll <= 80 {
		return "l"
	} else if roll <= 94 {
		return "m"
	} else if roll <= 108 {
		return "n"
	} else if roll <= 120 {
		return "p"
	} else if roll <= 122 {
		return "pr"
	} else if roll <= 133 {
		return "r"
	} else if roll <= 157 {
		return "s"
	} else if roll <= 167 {
		return "ss"
	} else if roll <= 170 {
		return "st"
	} else if roll <= 180 {
		return "t"
	} else if roll <= 185 {
		return "th"
	} else if roll <= 189 {
		return "tr"
	} else if roll <= 198 {
		return "ts"
	} else if roll <= 207 {
		return "tw"
	} else {
		return "v"
	}

}

// getDroyneFinalCons returns a Droyne final consonant.
func getDroyneFinalCons() string {
	roll := Dice(216)

	if roll <= 6 {
		return "b"
	} else if roll <= 17 {
		return "d"
	} else if roll <= 22 {
		return "f"
	} else if roll <= 28 {
		return "h"
	} else if roll <= 36 {
		return "k"
	} else if roll <= 40 {
		return "l"
	} else if roll <= 42 {
		return "lb"
	} else if roll <= 49 {
		return "ld"
	} else if roll <= 53 {
		return "lk"
	} else if roll <= 56 {
		return "lm"
	} else if roll <= 57 {
		return "ln"
	} else if roll <= 58 {
		return "lp"
	} else if roll <= 60 {
		return "ls"
	} else if roll <= 62 {
		return "lt"
	} else if roll <= 73 {
		return "m"
	} else if roll <= 80 {
		return "n"
	} else if roll <= 92 {
		return "p"
	} else if roll <= 101 {
		return "r"
	} else if roll <= 104 {
		return "rd"
	} else if roll <= 106 {
		return "rf"
	} else if roll <= 111 {
		return "rk"
	} else if roll <= 115 {
		return "rm"
	} else if roll <= 118 {
		return "rn"
	} else if roll <= 119 {
		return "rp"
	} else if roll <= 123 {
		return "rs"
	} else if roll <= 128 {
		return "rt"
	} else if roll <= 130 {
		return "rv"
	} else if roll <= 153 {
		return "s"
	} else if roll <= 159 {
		return "sk"
	} else if roll <= 167 {
		return "ss"
	} else if roll <= 172 {
		return "st"
	} else if roll <= 184 {
		return "t"
	} else if roll <= 190 {
		return "th"
	} else if roll <= 200 {
		return "ts"
	} else if roll <= 204 {
		return "v"
	} else {
		return "x"
	}

}
