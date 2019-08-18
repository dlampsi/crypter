package crypter

import "testing"

func Test_GenerateRandString(t *testing.T) {
	f := func(length int, expect int) {
		t.Helper()
		st := GenerateRandString(length)
		if len(st) != expect {
			t.Fatalf("unexpected len of generated password; want: %d, get: %d", expect, len(st))
		}
	}

	f(0, 0)
	f(1, 1)
	f(32, 32)
	f(-400, 0)
}

func Test_EncryptDecrypt(t *testing.T) {
	f := func(secret string, salt string, ok bool) {
		t.Helper()
		ec, err := Encrypt([]byte(secret), []byte(salt))
		if err != nil {
			if ok {
				t.Fatalf("unexpected error from Encrypt: %s", err.Error())
			}
			return
		}
		if !ok {
			t.Fatal("expected error from Encrypt")
		}

		dc, err := Decrypt(ec, []byte(salt))
		if err != nil {
			t.Fatal(err)
		}
		if string(dc) != string(secret) {
			t.Fatalf("unexpected secred returned; whant: %s, get: %s", string(secret), string(dc))
		}
	}

	f("My Super Secret Code Stuff", GenerateRandString(32), true)
}
