package main

import "fmt"

type Mobil struct {
	Merk             string
	BanDepanKiri     string
	BanDepanKanan    string
	BanBelakangKiri  string
	BanBelakangKanan string
	BunyiPintuKiri   string
	BunyiPintuKanan  string
}

func (m *Mobil) TampilInfo() {
	fmt.Printf("!!!!!!!!!!!!!!!!!!!\n")
	fmt.Printf("Merk: %s\n", m.Merk)
	fmt.Printf("Ban Depan Kiri: %s\n", m.BanDepanKiri)
	fmt.Printf("Ban Depan Kanan: %s\n", m.BanDepanKanan)
	fmt.Printf("Ban Belakang Kiri: %s\n", m.BanBelakangKiri)
	fmt.Printf("Ban Belakang Kanan: %s\n", m.BanBelakangKanan)
	fmt.Printf("Bunyi: %s\n", m.BunyiPintuKiri)
	fmt.Printf("Bunyi: %s\n", m.BunyiPintuKanan)
}

func (m *Mobil) GantiBanDepanKiri(banBaruDepanKiri string) {
	m.BanDepanKiri = banBaruDepanKiri
	fmt.Printf("Ban Depan kiri telah diganti menjadi %s\n", m.BanDepanKiri)
}

func (m *Mobil) GantiBanDepanKanan(banBaruDepanKanan string) {
	m.BanDepanKanan = banBaruDepanKanan
	fmt.Printf("Ban Depan kanan telah diganti menjadi %s\n", m.BanDepanKanan)
}

func (m *Mobil) GantiBanBelakangKiri(banBaruBelakangKiri string) {
	m.BanBelakangKiri = banBaruBelakangKiri
	fmt.Printf("Ban Belakang kiri telah diganti menjadi %s\n", m.BanBelakangKiri)
}

func (m *Mobil) GantiBanBelakangKanan(banBaruBelakangKanan string) {
	m.BanBelakangKanan = banBaruBelakangKanan
	fmt.Printf("Ban Belakang kanan telah diganti menjadi %s\n", m.BanBelakangKanan)
}

func (m *Mobil) GantiBunyiKiri(bunyiBaruPintuKiri string) {
	m.BunyiPintuKiri = bunyiBaruPintuKiri
	fmt.Printf("Bunyi Pintu Kiri telah diganti menjadi %s\n", m.BunyiPintuKiri)
}

func (m *Mobil) GantiBunyiKanan(bunyiBaruPintuKanan string) {
	m.BunyiPintuKanan = bunyiBaruPintuKanan
	fmt.Printf("Bunyi Pintu kanan telah diganti menjadi %s\n", m.BunyiPintuKanan)
}

func main() {
	mobilSaya := Mobil{
		Merk:             "Toyota",
		BanDepanKiri:     "Ban Karet",
		BanDepanKanan:    "Ban Karet",
		BanBelakangKiri:  "Ban Karet",
		BanBelakangKanan: "Ban Karet",
		BunyiPintuKiri:   "beep",
		BunyiPintuKanan:  "Knock",
	}

	mobilSaya.GantiBanDepanKiri("Ban Karet")

	mobilSaya.GantiBanDepanKanan("Ban Besi")

	mobilSaya.GantiBanBelakangKiri("Ban Kayu")

	mobilSaya.GantiBanBelakangKanan("Ban Kayu")

	mobilSaya.GantiBunyiKiri("Knock")

	mobilSaya.GantiBunyiKanan("beep")

	mobilSaya.TampilInfo()
}
