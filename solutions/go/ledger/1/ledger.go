package ledger

/*
Refactoring Log:
1. Removed the convoluted and buggy custom sorting algorithm, replacing it with the idiomatic sort.SliceStable to preserve the original order of identical entries.
2. Eliminated the unnecessary recursive call for empty entries; replaced with direct upfront validation of locale and currency.
3. Removed the highly inefficient and error-prone goroutine/channel-based parallelism for formatting entries. Replaced with a simple, sequential loop using strings.Builder for zero-allocation string concatenation.
4. Extracted date parsing, description truncation, and amount formatting into clean, single-responsibility helper functions.
5. Replaced manual string padding and concatenation for the header and entry lines with fmt.Sprintf using width specifiers (e.g., %-10s, %13s), which is safer, handles multi-byte characters (like €) correctly, and is vastly more readable.
6. Fixed the amount formatting logic to correctly handle thousands separators and locale-specific currency symbols/positions without relying on fragile string slicing.
7. Replaced empty error strings (errors.New("")) with descriptive error messages to aid in debugging.
8. Ensured all string manipulations are bounds-safe, preventing potential panics from malformed input dates or descriptions.
9. Preserved the original behavior of copying the input slice before sorting to prevent unintended side effects on the caller's data.
*/

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency != "USD" && currency != "EUR" {
		return "", errors.New("unsupported currency")
	}
	if locale != "en-US" && locale != "nl-NL" {
		return "", errors.New("unsupported locale")
	}

	var sb strings.Builder
	if locale == "en-US" {
		sb.WriteString(fmt.Sprintf("%-10s | %-25s | %-13s\n", "Date", "Description", "Change"))
	} else {
		sb.WriteString(fmt.Sprintf("%-10s | %-25s | %-13s\n", "Datum", "Omschrijving", "Verandering"))
	}

	if len(entries) == 0 {
		return sb.String(), nil
	}

	// Copy the slice to avoid mutating the caller's data, matching original behavior.
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	sort.SliceStable(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date != entriesCopy[j].Date {
			return entriesCopy[i].Date < entriesCopy[j].Date
		}
		if entriesCopy[i].Description != entriesCopy[j].Description {
			return entriesCopy[i].Description < entriesCopy[j].Description
		}
		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	for _, e := range entriesCopy {
		line, err := formatEntry(e, currency, locale)
		if err != nil {
			return "", err
		}
		sb.WriteString(line)
	}

	return sb.String(), nil
}

func formatEntry(e Entry, currency, locale string) (string, error) {
	if len(e.Date) != 10 {
		return "", errors.New("invalid date length")
	}
	if e.Date[4] != '-' || e.Date[7] != '-' {
		return "", errors.New("invalid date format")
	}

	y, m, d := e.Date[0:4], e.Date[5:7], e.Date[8:10]
	var dateStr string
	if locale == "en-US" {
		dateStr = fmt.Sprintf("%s/%s/%s", m, d, y)
	} else {
		dateStr = fmt.Sprintf("%s-%s-%s", d, m, y)
	}

	desc := e.Description
	if len(desc) > 25 {
		desc = desc[:22] + "..."
	}

	amountStr, err := formatAmount(e.Change, currency, locale)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%-10s | %-25s | %13s\n", dateStr, desc, amountStr), nil
}

func formatAmount(cents int, currency, locale string) (string, error) {
	var symbol string
	if currency == "USD" {
		symbol = "$"
	} else if currency == "EUR" {
		symbol = "€"
	} else {
		return "", errors.New("unsupported currency")
	}

	negative := cents < 0
	if negative {
		cents = -cents
	}

	// Format cents to ensure we always have at least 3 digits (e.g., 5 -> "005")
	s := fmt.Sprintf("%03d", cents)
	integerPart := s[:len(s)-2]
	decimalPart := s[len(s)-2:]

	// Add thousands separators
	var intWithSep strings.Builder
	for i, ch := range integerPart {
		if i > 0 && (len(integerPart)-i)%3 == 0 {
			if locale == "en-US" {
				intWithSep.WriteByte(',')
			} else {
				intWithSep.WriteByte('.')
			}
		}
		intWithSep.WriteRune(ch)
	}
	integerPart = intWithSep.String()

	var res string
	if locale == "en-US" {
		res = fmt.Sprintf("%s%s.%s", symbol, integerPart, decimalPart)
		if negative {
			res = "(" + res + ")"
		} else {
			res += " "
		}
	} else if locale == "nl-NL" {
		if negative {
			res = fmt.Sprintf("%s -%s,%s", symbol, integerPart, decimalPart)
		} else {
			res = fmt.Sprintf("%s %s,%s", symbol, integerPart, decimalPart)
		}
		res += " "
	} else {
		return "", errors.New("unsupported locale")
	}

	return res, nil
}