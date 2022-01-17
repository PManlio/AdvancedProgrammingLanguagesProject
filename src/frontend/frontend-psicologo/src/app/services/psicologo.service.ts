import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { myEnv } from 'src/environments/myEnv';
import { Paziente } from '../interfaces/paziente';

@Injectable({
  providedIn: 'root'
})
export class PsicologoService {

  // private paziente: Paziente;
  private serverUrl = `${myEnv.userServerUrl}/psicologo`

  private headers = new HttpHeaders(myEnv.headers);
  private headersWithToken;


  // private codFiscPsicologo: string = localStorage.getItem('PScodFisc')

  constructor(private httpClient: HttpClient) {
    this.headersWithToken = this.headers.append('Authorization', `Bearer ${localStorage.getItem('PStoken')}`);
  }

  public getPazientiOfPsicologo(codFisc: string) {
    let body = {
      codfisc: codFisc
    }
    return this.httpClient.post<Paziente[]>(`${this.serverUrl}/getpazienti`, body, {headers: this.headersWithToken})
  }
}
