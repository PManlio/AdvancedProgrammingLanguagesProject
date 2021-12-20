import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { Psicologo } from 'src/app/interfaces/psicologo';
import { PazienteServiceService } from 'src/app/services/paziente-service.service';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-cercapsicologo',
  templateUrl: './cercapsicologo.component.html',
  styleUrls: ['./cercapsicologo.component.css']
})
export class CercapsicologoComponent implements OnInit {

  public codFisc: string;
  public isPresent: boolean;

  constructor(private pazienteService: PazienteServiceService, private userInfo: UserInfoService) {
    this.userInfo.codFisc.subscribe(() => { this.codFisc = this.userInfo.localCodFisc })
  }

  public psicologo: Psicologo;
  public isFound: boolean = false;

  public getPsicologoByEmail(email: string) {
    this.pazienteService.findPsicologoByEmail(email).subscribe((v) => {
      this.psicologo = JSON.parse(JSON.stringify(v));
      
      this.isFound = true;

      let paziente = this.psicologo.pazienti.find(paziente => paziente == this.codFisc);
      if (paziente) {
        this.isPresent = true;
      } else {
        this.isPresent = false;
      }
    });
  }

  ngOnInit(): void {
  }

}
