import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { myEnv } from 'src/environments/myEnv';
import { Psicologo } from '../interfaces/psicologo';

@Injectable({
  providedIn: 'root'
})
export class PazienteServiceService {

  private url = myEnv.userServerUrl;
  private headers = new HttpHeaders(myEnv.headers);
  private headersWithToken;

  private userMail: string = localStorage.getItem("email");

  constructor(private http: HttpClient) {
    this.headersWithToken = this.headers.append('Authorization', `Bearer ${localStorage.getItem('token')}`);
  }

  // cerca psicologo via email
  public findPsicologoByEmail(mail: string) {
    return this.http.post<Psicologo>(`${this.url}/psicologo/getbyemail`, {"email":mail},{headers: this.headersWithToken});//.toPromise()/*.then(v => console.log(v))*/.catch(err => console.log(err));
  }

  public getAllPsicologiOfPatient(codFisc: string) {
    return this.http.post<Psicologo[]>(`${this.url}/paziente/getallpsicologiofpatient`, {"codFisc":codFisc}, { headers: this.headersWithToken })
  }

  // aggiunge psicologo via email
  public addPsicologoByEmail(psicologoMail: string, pazienteCodFisc: string) {
    let body = {
      "codFisc": pazienteCodFisc,
      "email": psicologoMail
    }
    console.log("STO CLICKANDO IL METODO REMOVE PSICOLOGO BY EMAIL", body)
    return this.http.put(`${this.url}/paziente/addpsicologobyemail`, body, { headers: this.headersWithToken }).toPromise().then(() => window.location.reload()).catch(err => alert(err));
  }

  // rimuove psicologo via email
  public removePsicologoByEmail(psicologoMail: string, pazienteCodFisc: string) {
    let body = {
      "codFisc": pazienteCodFisc,
      "email": psicologoMail
    }
    console.log("STO CLICKANDO IL METODO REMOVE PSICOLOGO BY EMAIL", body);
    return this.http.put(`${this.url}/paziente/removepsicologobyemail`, body, { headers: this.headersWithToken }).toPromise().then(() => window.location.reload()).catch(err => alert(err));
  }


  public getAllPsicologi() {
    return this.http.get(`${this.url}/psicologo/getallpsicologi`, {headers: this.headersWithToken}).toPromise().then(v => console.log(v)).catch(err => console.log(err));
  }
}
