import { Component, OnInit } from '@angular/core';
import { Psicologo } from 'src/app/interfaces/psicologo';
import { PazienteServiceService } from 'src/app/services/paziente-service.service';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-psicologo-card',
  templateUrl: './psicologo-card.component.html',
  styleUrls: ['./psicologo-card.component.css']
})
export class PsicologoCardComponent implements OnInit {

  public psicologi: Psicologo[];
  public codFisc: string;
  public isPsicologiFull: boolean;

  constructor(private pazienteService: PazienteServiceService, private userInfo: UserInfoService) {
    this.userInfo.codFisc.subscribe((codFisc) => {this.codFisc = codFisc; this.getPsicologi()})
  }

  ngOnInit(): void { }

  public getPsicologi() {
    this.pazienteService.getAllPsicologiOfPatient(this.codFisc).subscribe(psicologi => {
      this.psicologi = psicologi;
      if (this.psicologi == null || this.psicologi.length == 0) {
        this.isPsicologiFull = false;
      } else this.isPsicologiFull = true;
    })

  }

  public rimuoviPsicologo(email: string) {
    this.pazienteService.removePsicologoByEmail(email, this.codFisc);
    window.location.reload();
  }


}
