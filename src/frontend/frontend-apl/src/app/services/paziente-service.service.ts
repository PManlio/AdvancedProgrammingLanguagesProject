import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { myEnv } from 'src/environments/myEnv';

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
    return this.http.post(`${this.url}/psicologo/getbyemail`, {"email":mail},{headers: this.headersWithToken}).toPromise().then(v => console.log(v)).catch(err => console.log(err));
  }
  // aggiunge psicologo via nome

  // aggiunge psicologo via email

  // rimuove psicologo via email

  public getAllPsicologi() {
    return this.http.get(`${this.url}/psicologo/getallpsicologi`, {headers: this.headersWithToken}).toPromise().then(v => console.log(v)).catch(err => console.log(err));
  }
}
