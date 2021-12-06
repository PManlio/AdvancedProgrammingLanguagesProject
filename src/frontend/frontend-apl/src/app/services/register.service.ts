import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class RegisterService {

  private url: string = "http://localhost:8085/paziente/create"

  constructor(private http: HttpClient) { }

  public register() {
    let mybody:string;
    this.http.post(this.url, mybody, null).toPromise().then( v => { window.location.reload(); })
  }
}
