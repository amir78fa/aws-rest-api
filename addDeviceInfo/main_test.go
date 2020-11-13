package main

import (
	"testing"
)

func TestHandler(t *testing.T) {
	resp, _ := handler(deviceInfo{
		ID:          "device1",
		DeviceModel: "XSA22",
		Name:        "device1",
		Note:        "note",
		Serial:      "23234134",
	})

	if resp.StatusCode != 500 {
		t.Errorf("%d", resp.StatusCode)
	}

	resp, _ = handler(deviceInfo{
		ID:          "device1",
		DeviceModel: "XSA22",
		Name:        "device1",
		Note:        "note",
	})

	if resp.StatusCode != 400 {
		t.Errorf("%d", resp.StatusCode)
	}

	resp, _ = handler(deviceInfo{
		ID:          "device1",
		DeviceModel: "XSA22",
		Name:        "device1",
		Note:        "note",
		Serial:      "23234134",
	})

	if resp.StatusCode != 500 {
		t.Errorf("%d", resp.StatusCode)
	}
}
