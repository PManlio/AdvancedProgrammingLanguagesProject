use wrivtone_Manlio;

# drop table utente;
# drop table paziente;
# drop table psicologo;

create table if not exists utente(
    codFisc varchar(55) PRIMARY KEY,
    nome varchar(255),
    cognome varchar(255),
    email varchar(255),
    password varchar(255),
    citta varchar(255),
    cellulare varchar(255),
    genere varchar(255)
);

create table if not exists paziente(
    codFisc varchar(55) PRIMARY KEY,
    patientOf varchar(5000)
);

create table if not exists psicologo(
    codFisc varchar(55) PRIMARY KEY,
    pazienti varchar(5000)
);

select * from utente;
select * from psicologo;
select * from paziente;