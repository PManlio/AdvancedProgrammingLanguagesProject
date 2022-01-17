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
    return localStorage.getItem("PScodFisc");
  }
  set localCodFisc(newCodFisc) {
    this.codFisc.next(newCodFisc)
    localStorage.setItem("PScodFisc", newCodFisc);
  }

  get localNome(): string {
    return localStorage.getItem("PSnome");
  }
  set localNome(newName) {
    this.nome.next(newName)
    localStorage.setItem("PSnome", newName);
  }

  get localCognome(): string {
    return localStorage.getItem("PScognome");
  }
  set localCognome(newCognome) {
    this.cognome.next(newCognome)
    localStorage.setItem("PScognome", newCognome);
  }

  get localEmail(): string {
    return localStorage.getItem("PSemail");
  }
  set localEmail(newEmail) {
    this.email.next(newEmail)
    localStorage.setItem("PSemail", newEmail);
  }

  constructor() { }
}
