import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';
import { myEnv } from 'src/environments/myEnv';
import { UserInfoService } from './user-info.service';

@Injectable({
  providedIn: 'root'
})
export class JwtService {

  private isJWTstored = new Subject<boolean>();

  constructor(private http: HttpClient, private userInfo: UserInfoService) { }

  public storeJWT(jwt: string): void {
    localStorage.setItem('token', jwt);
  }

  public removeJWT(): void {
    localStorage.removeItem('token');

    // successivamente bisogna rimuovere anche le altre informazioni ottenute al login
    localStorage.removeItem('codFisc');
    localStorage.removeItem('nome');
    localStorage.removeItem('cognome');
    localStorage.removeItem('email');
    
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

  public isJWTvalid() {
    let headersWithToken = this.envHeaders.append('Authorization', `Bearer ${localStorage.getItem('token')}`);
    return this.http.get(myEnv.userServerUrl, { headers: headersWithToken })
      .toPromise()
      .then(response => this.saveInformation(response))
      .catch(err => console.log(err));
  }

  private saveInformation(obj: any) {
    this.userInfo.localCodFisc = obj.CodFisc;
    this.userInfo.localNome = obj.Nome;
    this.userInfo.localCognome = obj.Cognome;
    this.userInfo.localEmail = obj.Email;

    /* localStorage.setItem("codFisc", obj.CodFisc)
    localStorage.setItem("nome", obj.Nome)
    localStorage.setItem("cognome", obj.Cognome)
    localStorage.setItem("email", obj.Email) */
  }
}
