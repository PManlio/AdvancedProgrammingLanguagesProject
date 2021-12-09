import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
import { myEnv } from 'src/environments/myEnv';

@Injectable({
  providedIn: 'root'
})
export class JwtService {

  private isJWTstored = new Subject<boolean>();

  constructor(private http: HttpClient) { }

  public storeJWT(jwt: string): void {
    localStorage.setItem('token', jwt);
  }

  public removeJWT(): void {
    localStorage.removeItem('token');
    window.location.reload()
  }

  // controlla se il token c'è
  public isJWTPresent(): boolean {
    return localStorage.getItem('token') ? true : false;
  }

  // per update del token
  public watchJWT(): Observable<any> {
    this.isJWTstored.subscribe();
    return this.isJWTstored.asObservable();
  }

  private envHeaders = new HttpHeaders(myEnv.headers)

  private extractData(res: any) {
    let body = res;
    console.log(body);
    return body;
  }
  public isJWTvalid() {
    let headersWithToken = this.envHeaders.append('Authorization', `Bearer ${localStorage.getItem('token')}`);
    return this.http.get(myEnv.userServerUrl, { headers: headersWithToken })
      .toPromise()
      .then(response => console.log(response))
      .catch(err => console.log(err));
  }
}
