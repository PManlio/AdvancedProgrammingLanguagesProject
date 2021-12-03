import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  private serverUrl = `http://localhost:8085/login`

  private headers = new HttpHeaders({
    'Content-Type': 'application/x-www-form-urlencoded',
    'Accept': 'application/json',
  });

  constructor(private http: HttpClient) { }

  public authenticate(email: string, password: string) {
    let basic = btoa(`${email}:${password}`);
    let headers = this.headers.append('Authorization', `Basic ${basic}`);

    return this.http.post(`${this.serverUrl}`, null, { headers }).toPromise().then(v => console.log(v));
  }
}
