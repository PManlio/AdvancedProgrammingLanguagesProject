import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  private serverUrl = `localhost:8085/login`
  private headers = new HttpHeaders({
    'Accept': 'application/json',
  });

  constructor(private http: HttpClient) { }

  public authenticate(email, password: string) {
    let basic = btoa(`${email}:${password}`);
    let headers = this.headers.append('Authorization', `Basic ${basic}`);

    return this.http.post(`${this.serverUrl}`, null, { headers }).toPromise().then(v => console.log(v));
  }
}
