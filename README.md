
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
- Prosta pretraga vožnji - unosi se početno stajalište i destinacija, a rezultat pretrage su sve vožnje koje vode od početnog stajališta do destinacije.
- Složena pretraga vožnji na osnovu kriterijuma - pronalazi sve, odnosno, ograničen broj podvožnji koje vode od početne destinacije do odredišta na osnovu kriterijuma (ograničen broj podvožnji postoji iz razloga što može da se desi da postoji ogroman broj podvožnji koje vode od početne destinacije do odredišta). 

  Jedinični rezultat pretrage je skup podvožnji koje vode od početne destinacije do odredišta. Tih jediničnih rezultata biće recimo naboljih 5 po kriterijumu. Osnovni kriterijum koji planiram da implementiram je ukupna dužina puta svih podvožnji, ako budem imao vremena dodaću i kriterijum za vreme i za broj stanica. 
  - Primer: <br />
    Početna destinacija: Beograd <br />
    Odredište: Novi Sad <br />
    Rezultati: |Beograd, Novi Sad| ; |Beograd, Stara Pazova, Novi Sad| ; |Beograd, Beška, Novi Sad| <br />
    Gde su rezultati pretrage 3 najbolja rezultata po kriterijumu dužine puta
  
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

### Napomena za diplomski
Ukoliko bude bilo potrebno neko proširenje za diplomski otvoren sam za sugestije i predloge
