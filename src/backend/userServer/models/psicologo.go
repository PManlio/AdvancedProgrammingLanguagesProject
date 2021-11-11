package models

type Psicologo struct {
	Utente   Utente   `json:"utente"`
	Pazienti []string `json:"pazienti"`
}

type PsicoInfo struct {
	Nome, Cognome, Email, Citta, Cellulare, Genere string
}

/* func (p *Psicologo) GetInfoPsicologo() *PsicoInfo {
	info := new(PsicoInfo)
	info.Nome = p.Utente.Nome
	info.Cognome = p.Utente.Cognome
	info.Email = p.Utente.Email
	info.Cellulare = p.Utente.Cellulare
	info.Genere = p.Utente.Genere

	return info
} */
