package main

import (
	"testing"
)

func TestCheckSingleWebsite(t *testing.T) {
	var checker httpChecker = httpChecker{}

	status, err := checker.Check("http://www.google.com")

	if err != nil || status == false {
		t.Errorf("Expected status = true, but got false")
	}
}

func TestWebsites(t *testing.T) {
	arrayOfWebsites := []struct {
		nameOfTest string
		website    string
		status     bool
	}{
		{nameOfTest: "Checking http://www.google.com", website: "http://www.google.com", status: true},
		{"Checking http://www.facebook.com", "http://www.facebook.com", true},
		{"Checking 1234", "1234", false},
		// It is an API created on my local server. Checked and verified the working by below unit test.
		{"Checking http://127.0.0.1:8080/welcome", "http://127.0.0.1:8080/welcome", false},
	}

	for _, websiteTest := range arrayOfWebsites {
		t.Run(websiteTest.nameOfTest, func(t *testing.T) {
			var checker httpChecker

			status, _ := checker.Check(websiteTest.website)

			if status != websiteTest.status {
				t.Errorf("Expecting %v for %s, but got %v", websiteTest.status, websiteTest.website, status)
			}
		})
	}
}
