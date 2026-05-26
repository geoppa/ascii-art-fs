package main

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

// Βοηθητική συνάρτηση για την εκτέλεση της main με συγκεκριμένα ορίσματα
// και την καταγραφή της εξόδου (stdout)
func runMainWithArgs(args []string) (string, error) {
	// Αποθήκευση των αρχικών os.Args και os.Stdout για επαναφορά μετά το τεστ
	oldArgs := os.Args
	oldStdout := os.Stdout
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldStdout
	}()

	// Ορισμός των νέων ορισμάτων (το args[0] είναι πάντα το όνομα του προγράμματος)
	os.Args = append([]string{"cmd"}, args...)

	// Δημιουργία pipe για την καταγραφή των fmt.Println
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Εκτέλεση της main
	main()

	// Κλείσιμο του writer και ανάγνωση των δεδομένων
	w.Close()
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func TestMainScenarios(t *testing.T) {
	// Πριν τρέξεις το τεστ, βεβαιώσου ότι υπάρχει ο φάκελος banners με τα αρχεία.
	// Για το τεστ, αν δεν υπάρχουν τα πραγματικά αρχεία, μπορείς να φτιάξεις mock.

	tests := []struct {
		name           string
		args           []string
		wantInOutput   []string // Τι περιμένουμε να δούμε στο output
		dontWantOutput bool     // Αν περιμένουμε κενή έξοδο
	}{
		{
			name:         "Λιγότερα ορίσματα από 2",
			args:         []string{"Hello"},
			wantInOutput: []string{"Usage: go run . <text> <banner>"},
		},
		{
			name:         "Λανθασμένο banner - Fallback σε standard",
			args:         []string{"Hello", "wrong_banner"},
			wantInOutput: []string{"Error: Invalid or missing banner"},
		},
		{
			name:         "Περισσότερα ορίσματα μετά το banner - Warning",
			args:         []string{"Hello", "standard", "extra1", "extra2"},
			wantInOutput: []string{"Warning: Extra arguments found after the banner. Ignoring them."},
		},
		{
			name:           "Κενό input κειμένου",
			args:           []string{"", "standard"},
			dontWantOutput: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := runMainWithArgs(tt.args)
			if err != nil {
				t.Fatalf("Αποτυχία εκτέλεσης της main: %v", err)
			}

			if tt.dontWantOutput && output != "" {
				t.Errorf("Περιμέναμε κενή έξοδο, αλλά πήραμε: %q", output)
			}

			for _, want := range tt.wantInOutput {
				if !strings.Contains(output, want) {
					t.Errorf("Το output δεν περιέχει το αναμενόμενο κείμενο %q. Output:\n%s", want, output)
				}
			}
		})
	}
}
