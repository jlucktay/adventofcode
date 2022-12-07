// Package cookiemonster will monster your cookies!
package cookiemonster

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"os/user"
	"strings"

	// Imported to help chew on cookies
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/pbkdf2"
)

// Inspiration:
// http://n8henrie.com/2013/11/use-chromes-cookies-for-easier-downloading-with-python-requests/

// Chromium Mac os_crypt:  http://dacort.me/1ynPMgx
var (
	salt       = "saltysalt"
	iv         = "                "
	length     = 16
	iterations = 1003
)

// Cookie - Items for a cookie
type Cookie struct {
	Domain         string
	Key            string
	Value          string
	EncryptedValue []byte
}

func (c *Cookie) decryptedValue() string {
	if c.Value > "" {
		return c.Value
	}

	if len(c.EncryptedValue) > 0 {
		encryptedValue := c.EncryptedValue[3:]
		return decryptValue(encryptedValue)
	}

	return ""
}

func decryptValue(encryptedValue []byte) string {
	key := pbkdf2.Key([]byte(getPassword()), []byte(salt), iterations, length, sha1.New)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}

	decrypted := make([]byte, len(encryptedValue))
	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	cbc.CryptBlocks(decrypted, encryptedValue)

	plainText, err := aesStripPadding(decrypted)
	if err != nil {
		fmt.Println("Error decrypting:", err)

		return ""
	}

	return string(plainText)
}

// In the padding scheme the last <padding length> bytes
// have a value equal to the padding length, always in (1,16]
func aesStripPadding(data []byte) ([]byte, error) {
	if len(data)%length != 0 {
		return nil, fmt.Errorf("decrypted data block length is not a multiple of %d", length)
	}

	paddingLen := int(data[len(data)-1])

	if paddingLen > 16 {
		return nil, fmt.Errorf("invalid last block padding length: %d", paddingLen)
	}

	return data[:len(data)-paddingLen], nil
}

func getPassword() string {
	parts := strings.Fields("security find-generic-password -wga Chrome")
	cmd := parts[0]
	parts = parts[1:]

	out, err := exec.Command(cmd, parts...).Output()
	if err != nil {
		log.Fatal("error finding password ", err)
	}

	return strings.Trim(string(out), "\n")
}

func getCookies(domain string) (cookies []Cookie) {
	usr, _ := user.Current()
	cookiesFile := fmt.Sprintf("%s/Library/Application Support/Google/Chrome/Default/Cookies", usr.HomeDir)

	db, err := sql.Open("sqlite3", cookiesFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(
		"SELECT name, value, host_key, encrypted_value FROM cookies WHERE host_key like ?",
		fmt.Sprintf("%%%s%%", domain))
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name, value, hostKey string
		var encryptedValue []byte
		rows.Scan(&name, &value, &hostKey, &encryptedValue)
		cookies = append(cookies, Cookie{hostKey, name, value, encryptedValue})
	}

	return
}

// GetCookieWithKey will check all of the cookies from the given domain, and
// return the value if one has a key that matches exactly.
// If no such cookie is found, an empty string is returned instead.
func GetCookieWithKey(domain, key string) (cookie string) {
	for _, b := range getCookies(domain) {
		if b.Key == key {
			cookie = b.decryptedValue()
			return
		}
	}

	return
}
