import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { PazientimanagerService } from 'src/app/services/pazientimanager.service';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-aggiungipazientemodal',
  templateUrl: './aggiungipazientemodal.component.html',
  styleUrls: ['./aggiungipazientemodal.component.css']
})
export class AggiungipazientemodalComponent implements OnInit {

  private codFiscPsicologo: string;
  constructor(private pazientiManager: PazientimanagerService, private userInfo: UserInfoService) {
    this.userInfo.codFisc.subscribe(cod => this.codFiscPsicologo = cod);
  }

  public paziente = [];

  ngOnInit() { }

  displayStyle = "none";

  openPopup() {
    this.displayStyle = "block";
  }
  closePopup() {
    this.displayStyle = "none";
  }

  public searchByEmail(f: NgForm) {
    if (f.valid) {
      // console.log(f.value.email)
      //this.pazientiManager.addPazienteByEmail(f.value.email, this.codFiscPsicologo)
      this.pazientiManager.addPazienteByEmail(f.value.email, this.codFiscPsicologo).subscribe(val => console.log(val))
      window.location.reload();
    }
  }
}
