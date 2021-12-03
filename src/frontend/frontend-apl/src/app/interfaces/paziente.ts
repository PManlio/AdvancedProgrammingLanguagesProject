export interface Paziente {
    nome: string;
    cognome: string;
    email: string;
    citta: string;
    patientOf?: string[];
}
