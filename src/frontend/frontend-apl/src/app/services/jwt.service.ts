import { Injectable } from '@angular/core';
import { Observable, Subject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class JwtService {

  private isJWTstored = new Subject<boolean>();

  constructor() { }

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
}
