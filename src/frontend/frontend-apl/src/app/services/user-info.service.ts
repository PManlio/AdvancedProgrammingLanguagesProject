import { Injectable } from '@angular/core';
import { BehaviorSubject } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UserInfoService {

  public codFisc = new BehaviorSubject(this.localCodFisc);
  public nome = new BehaviorSubject(this.localNome);
  public cognome = new BehaviorSubject(this.localCognome);
  public email = new BehaviorSubject(this.localEmail);

  get localCodFisc(): string {
    return localStorage.getItem("codFisc");
  }
  set localCodFisc(newCodFisc) {
    this.codFisc.next(newCodFisc)
    localStorage.setItem("codFisc", newCodFisc);
  }

  get localNome(): string {
    return localStorage.getItem("nome");
  }
  set localNome(newName) {
    this.nome.next(newName)
    localStorage.setItem("nome", newName);
  }

  get localCognome(): string {
    return localStorage.getItem("cognome");
  }
  set localCognome(newCognome) {
    this.cognome.next(newCognome)
    localStorage.setItem("cognome", newCognome);
  }

  get localEmail(): string {
    return localStorage.getItem("email");
  }
  set localEmail(newEmail) {
    this.email.next(newEmail)
    localStorage.setItem("email", newEmail);
  }

  constructor() { }
}
