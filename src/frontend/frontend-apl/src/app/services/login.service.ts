import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { JwtService } from './jwt.service';
import { JwtInterface } from '../interfaces/jwtInterface';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  private serverUrl = `http://localhost:8085/login`

  private headers = new HttpHeaders({
    'Content-Type': 'application/x-www-form-urlencoded',
    'Accept': 'application/json',
  });

  private jwtInterface: JwtInterface;

  constructor(private http: HttpClient, private jwt: JwtService) { }

  public authenticate(email: string, password: string) {
    let basic = btoa(`${email}:${password}`);
    let headers = this.headers.append('Authorization', `Basic ${basic}`);

    return this.http.post(`${this.serverUrl}`, null, { headers })
      .toPromise()
      .then(tkn => {
        this.jwtInterface = JSON.parse(JSON.stringify(tkn));
        this.jwt.storeJWT(this.jwtInterface.token);
        window.location.reload();
      })
      .catch(err => console.error(err));
  }
}
