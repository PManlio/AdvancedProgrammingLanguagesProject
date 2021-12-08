import { HttpClient, HttpErrorResponse, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { RegistrationInterface } from '../interfaces/registration-interface';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {

  private url: string = "http://localhost:8085/paziente/create"

  private headers = new HttpHeaders({
    'Content-Type': 'application/x-www-form-urlencoded',
    'Accept': 'application/json',
  });
  constructor(private http: HttpClient) { }

  public register(body: RegistrationInterface) {

    body.codFisc = btoa(body.codFisc)
    body.password = btoa(body.password)
    body.email = btoa(body.email)

    let JSONBody = {
      "utente": body,
      "patientOf": []
    }

    let headers = this.headers;
    this.http.post(this.url, JSONBody, { headers }).toPromise().then(v => {
      console.log(body);
      window.location.reload();
    }).catch((err: HttpErrorResponse) => { window.alert(`${err.status} ${err.message})`) })
  }
}
