
# Ride explorer

Ride explorer je veb aplikacija koja služi za podršku prevoza ljudi deljenjem vozila.

## Funkcionalnosti

### Neregistrovani korisnik

- Registracija kao vozač
- Registracija kao korisnik (iznajmljivač mesta u vozilu)
- Prijava

### Administrator

- Registracija drugih admina
- Odobravanje registracije vozača
- Pregled i pretraga vozača i korisnika
- Banovanje i brisanje korisnika
- Pregled profila korisnika

### Vozač

- Upravljanje profilom (izmena šifre i ostalih podataka)
- Kreiranje vožnje (početno i završno stajalište, cena, kilometraža, napomena)
- Pregled i prihvatanje zahteva korisnika za učestvovanje u vožnji
- Izmena vožnje (broj mesta, vreme polaska, obaveštenje učesnicima u vožnji o izmeni..)
- Brisanje vožnje (obaveštenje učesnicima u vožnji)
- Uvid u završene vožnje
- Ocenjivanje korisnika koji su učestvovali u vožnji uz komentar
- Pisanje žalbi za korisnike koji su učestvovali u vožnji
- Pregled profila korisnika

### Korisnik

- Upravljanje profilom (izmena šifre i ostalih podataka)
- Pretraga vožnji 
- Zahtev za učestovanje u vožnji
- Uvid u završene vožnje
- Ocenjivanje korisnika/vozača koji su učestvovali u istoj vožnji uz komentar 
- Pisanje žalbi za korisnike/vozača koji su učestvovali u istoj vožnji
- Pregled profila vozača i drugih korisnika


## Arhitektura 

### Mikroservisna arhitektura sa servisima: 

- Gateway servis - Go
- Servis za korisnike (moguća podela na sitnije servise) - Go
- Email servis - Go
- Servis za vožnje - Go
- Servis za pronalaženje vožnji - Rust (moguće spajanje sa servisom iznad, gde će novi servis biti pisan u Rust-u)
- Servis za obradu zahteva za učestvovanje u vožnji - Go
- Servis za žalbe - Go 
- Servis za ocene - Go
- Klijentska veb aplikacija - Angular (React ukoliko budem imao dovoljno vremena)

*Napomena: Moguće izmene pre implementacije*
### Baze podataka:

- SQL - PostgreSQL
