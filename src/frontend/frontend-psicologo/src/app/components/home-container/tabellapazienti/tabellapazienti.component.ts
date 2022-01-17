import { Component, OnInit } from '@angular/core';
import { Paziente } from 'src/app/interfaces/paziente';
import { PsicologoService } from 'src/app/services/psicologo.service';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-tabellapazienti',
  templateUrl: './tabellapazienti.component.html',
  styleUrls: ['./tabellapazienti.component.css']
})
export class TabellapazientiComponent implements OnInit {

  public pazienti: Paziente[]
  private codFisc: string;

  constructor(private psicologoservice: PsicologoService, private userinfo: UserInfoService) {
    this.userinfo.codFisc.subscribe(v => {
      this.codFisc = v;
      this.psicologoservice.getPazientiOfPsicologo(this.codFisc).subscribe(pazienti => this.pazienti = pazienti)
    })
  }

  ngOnInit(): void { }

  loggaPazienti() {
    console.log(this.pazienti)
  }
}
