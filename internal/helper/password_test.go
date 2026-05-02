package helper

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "rahasia123"

	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Diharapkan tidak ada error, tapi mendapat: %v", err)
	}

	if hashedPassword == "" {
		t.Fatal("Diharapkan hashed password tidak kosong")
	}

	if hashedPassword == password {
		t.Fatal("Diharapkan hashed password berbeda dengan password asli")
	}
}

func TestCheckPassword(t *testing.T) {
	password := "rahasia123"
	wrongPassword := "salah123"

	// 1. Buat hash terlebih dahulu
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Setup gagal saat melakukan hashing: %v", err)
	}

	// 2. Test dengan password yang benar (harus true)
	isValid := CheckPassword(password, hashedPassword)
	if !isValid {
		t.Errorf("Diharapkan true untuk password yang benar, tapi mendapat false")
	}

	// 3. Test dengan password yang salah (harus false)
	isInvalid := CheckPassword(wrongPassword, hashedPassword)
	if isInvalid {
		t.Errorf("Diharapkan false untuk password yang salah, tapi mendapat true")
	}

	// 4. Test simulasi parameter terbalik seperti isu sebelumnya (harus false)
	isReversed := CheckPassword(hashedPassword, password)
	if isReversed {
		t.Errorf("Diharapkan false jika urutan parameter (password dan hash) terbalik")
	}
}
