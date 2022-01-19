import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { myEnv } from 'src/environments/myEnv';

@Injectable({
  providedIn: 'root'
})
export class PazientimanagerService {

  private serverUrl = `${myEnv.userServerUrl}/psicologo`;
  private headers = new HttpHeaders(myEnv.headers);
  private headersWithToken;

  constructor(private httpClient: HttpClient) { 
    this.headersWithToken = this.headers.append('Authorization', `Bearer ${localStorage.getItem('PStoken')}`);
  }

  public addPazienteByEmail(emailPaziente: string, codFiscPsicologo: string) {
    let body = {
      email: emailPaziente,
      codFisc: codFiscPsicologo
    }

    return this.httpClient.put(`${this.serverUrl}/addpazientebyemail`, body, {headers: this.headersWithToken})
  }

  public removePazienteByCodFisc(codiceFiscalePaziente: string, codiceFiscalePsicologo: string) {
    let body = {
      codFiscPsicologo: codiceFiscalePsicologo,
      codFiscPaziente: codiceFiscalePaziente
    }

    return this.httpClient.put(`${this.serverUrl}/removepaziente`, body, {headers: this.headersWithToken})
  }
}
