package main

import "fmt"

// Definisikan interface untuk Kendaraan
type Kendaraan interface {
	TampilInfo()
	GantiBan(banBaru string)
	GantiBunyi(bunyiBaru string)
}

// Definisikan struct untuk merepresentasikan Mobil
type Mobil struct {
	Merk  string
	Ban   string
	Bunyi string
}

func (m *Mobil) TampilInfo() {
	fmt.Printf("Merk: %s\n", m.Merk)
	fmt.Printf("Ban: %s\n", m.Ban)
	fmt.Printf("Bunyi: %s\n", m.Bunyi)
}

func (m *Mobil) GantiBan(banBaru string) {
	m.Ban = banBaru
	fmt.Printf("Ban telah diganti menjadi %s\n", m.Ban)
}

func (m *Mobil) GantiBunyi(bunyiBaru string) {
	m.Bunyi = bunyiBaru
	fmt.Printf("Bunyi telah diganti menjadi %s\n", m.Bunyi)
}

func main() {

	mobilSaya := &Mobil{
		Merk:  "Toyota",
		Ban:   "ban karet",
		Bunyi: "deep",
	}

	mobilSaya.GantiBan("Ban Kayu")

	mobilSaya.GantiBunyi("Knock")
	mobilSaya.TampilInfo()
}
