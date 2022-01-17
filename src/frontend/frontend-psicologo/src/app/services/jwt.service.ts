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
    localStorage.setItem('PStoken', jwt);
  }

  public removeJWT(): void {
    localStorage.removeItem('PStoken');

    // successivamente bisogna rimuovere anche le altre informazioni ottenute al login
    localStorage.removeItem('PScodFisc');
    localStorage.removeItem('PSnome');
    localStorage.removeItem('PScognome');
    localStorage.removeItem('PSemail');
    
    window.location.reload()
  }

  // controlla se il token c'è
  public isJWTPresent(): boolean {
    return localStorage.getItem('PStoken') ? true : false;
  }

  // per update del token
  public watchJWT(): Observable<any> {
    this.isJWTstored.subscribe();
    return this.isJWTstored.asObservable();
  }

  private envHeaders = new HttpHeaders(myEnv.headers)

  public isJWTvalid() {
    let headersWithToken = this.envHeaders.append('Authorization', `Bearer ${localStorage.getItem('PStoken')}`);
    return this.http.get(myEnv.userServerUrl, { headers: headersWithToken }).subscribe(res => { this.saveInformation(res) })
      // .toPromise()
      // .then(response => this.saveInformation(response))
      // .catch(err => console.log(err));
  }

  private saveInformation(obj: any) {
    this.userInfo.localCodFisc = obj.CodFisc;
    this.userInfo.localNome = obj.Nome;
    this.userInfo.localCognome = obj.Cognome;
    this.userInfo.localEmail = obj.Email;
  }
}
