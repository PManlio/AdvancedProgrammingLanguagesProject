export interface Psicologo {
    utente: {
        CodFisc: string;
        Cellulare: string;
        Citta: string;
        Nome: string;
        Cognome: string;
        Email: string;
        Genere: string;
        Password?: string;
    };
    pazienti?: string[];
}