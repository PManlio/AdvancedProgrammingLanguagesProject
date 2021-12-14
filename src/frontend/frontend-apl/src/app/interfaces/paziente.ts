export interface Paziente {
    codFisc: string;
    nome: string;
    cognome: string;
    email: string;
    citta: string;
    patientOf?: string[];
}
