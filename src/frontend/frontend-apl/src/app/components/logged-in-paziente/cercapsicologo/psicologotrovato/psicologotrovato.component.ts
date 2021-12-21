import { Component, Input, OnInit } from '@angular/core';
import { Psicologo } from 'src/app/interfaces/psicologo';
import { PazienteServiceService } from 'src/app/services/paziente-service.service';
import { UserInfoService } from 'src/app/services/user-info.service';

@Component({
  selector: 'app-psicologotrovato',
  templateUrl: './psicologotrovato.component.html',
  styleUrls: ['./psicologotrovato.component.css']
})
export class PsicologotrovatoComponent implements OnInit {

  @Input() public psicologo: Psicologo;
  @Input() public isPresent: boolean = false;
  private codFisc: string;

  constructor(private pazienteService: PazienteServiceService, private userInfo: UserInfoService) { 
    this.userInfo.codFisc.subscribe(codFisc => { this.codFisc = codFisc; })
  }

  ngOnInit(): void { }

  public addPsicologo() {
    this.pazienteService.addPsicologoByEmail(this.psicologo.utente.Email, this.codFisc)
  }

  public removePsicologo() {
    this.pazienteService.removePsicologoByEmail(this.psicologo.utente.Email, this.codFisc)
  }

}
