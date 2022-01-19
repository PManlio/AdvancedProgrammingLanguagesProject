import { Component, Input, OnInit } from '@angular/core';
import { PazientimanagerService } from 'src/app/services/pazientimanager.service';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-rimuovipaziente-btn',
  templateUrl: './rimuovipaziente-btn.component.html',
  styleUrls: ['./rimuovipaziente-btn.component.css']
})
export class RimuovipazienteBtnComponent implements OnInit {

  constructor(private pazienteManager: PazientimanagerService, private userinfo: UserInfoService) { 
    this.userinfo.codFisc.subscribe(cod => this.codFiscPsicologo = cod);
  }

  @Input() nomeUtente: string;
  @Input() codFiscPaziente: string;
  private codFiscPsicologo: string;

  ngOnInit(): void {
  }

  public rimuoviPaziente() {
    // console.log(this.codFiscPsicologo, this.codFiscPaziente);
    this.pazienteManager.removePazienteByCodFisc(this.codFiscPaziente, this.codFiscPsicologo).subscribe(val => alert(val));
    window.location.reload();
  }

}
