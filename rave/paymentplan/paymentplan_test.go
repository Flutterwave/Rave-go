package paymentplan

import (
	"Rave-go/rave"
	"testing"
)

var r = rave.Rave{
	false,
	"FLWPUBK-f54d8d24292e377a71620bd82a8bb17c-X",
	"FLWSECK-a18ca169cb007a93db4479aff683a387-X",
}

var p = PaymentPlan{r}

func TestPaymentPlan_Create(t *testing.T) {
	payloads := []PaymentPlanData{
		{
			Amount:   "2000",
			Name:     "Killer",
			Interval: "daily",
			Duration: "12",
			Seckey:   r.GetSecretKey(),
		},
		{
			Amount:   "400",
			Name:     "Killer Bean",
			Interval: "monthly",
			Duration: "12",
			Seckey:   r.GetSecretKey(),
		},
	}

	for _, payload := range payloads {
		err, response := p.Create(payload)
		if err != nil {
			t.Fatalf("An error occurred while testing single transfer: %v", err)
		}
		if response["status"] != "success" {
			t.Fatalf("Single transfer wasn't successful: %v", response)
		}
	}
}

func TestPaymentPlan_List(t *testing.T) {
	err, response := p.List()
	if err != nil {
		t.Fatalf("An error occurred while testing single transfer: %v", err)
	}
	if response["status"] != "success" {
		t.Fatalf("Single transfer wasn't successful: %v", response)
	}
}

func TestPaymentPlan_Fetch(t *testing.T) {
	planIds := [4]string{"1140", "1139", "1138", "1136"}
	for _, planId := range planIds {
		err, response := p.Fetch(planId)
		if err != nil {
			t.Fatalf("An error occurred while testing single transfer: %v", err)
		}
		if response["status"] != "success" {
			t.Fatalf("Single transfer wasn't successful: %v", response)
		}
	}
}

func TestPaymentPlan_Edit(t *testing.T) {
	payloads := []struct {
		Id     int
		Name   string
		Status string
	}{
		{
			Id:     1140,
			Name:   "Anita Becker",
			Status: "cancelled",
		},
		{
			Id:     1139,
			Name:   "Anita Goldie",
			Status: "active",
		},
	}

	for _, payload := range payloads {
		err, response := p.Edit(payload.Id, payload.Name, payload.Status)
		if err != nil {
			t.Fatalf("An error occurred while testing single transfer: %v", err)
		}
		if response["status"] != "success" {
			t.Fatalf("Single transfer wasn't successful: %v", response)
		}
	}
}

func TestPaymentPlan_Cancel(t *testing.T) {
	planIds := [2]int{978, 1133}
	for _, planId := range planIds {
		err, response := p.Cancel(planId)
		if err != nil {
			t.Fatalf("An error occurred while testing single transfer: %v", err)
		}
		if response["status"] != "success" {
			t.Fatalf("Single transfer wasn't successful: %v", response)
		}
	}
}
