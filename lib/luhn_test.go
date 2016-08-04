// Test for the Luhn package
package luhn

import "testing"

func TestValidPANs(t *testing.T) {
    // Test some valid PANs 
    valid_pans := []string{"4507990000000010", "5323601111111112", "376411234531007"}
    for _,pan := range valid_pans {
        if !validate(pan) {
            t.Error("Valid PAN validated as invalid", pan)
        }
    }
}

func TestInvalidPANs(t *testing.T) {
    // Test some valid PANs 
    valid_pans := []string{"4507990000000011", "5323601111111113", "376411234531004"}
    for _,pan := range valid_pans {
        if validate(pan) {
            t.Error("Invalid PAN validated as valid", pan)
        }
    }
}
