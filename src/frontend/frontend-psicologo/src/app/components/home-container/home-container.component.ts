import { Component, OnInit } from '@angular/core';
import { JwtService } from 'src/app/services/jwt.service';

@Component({
  selector: 'app-home-container',
  templateUrl: './home-container.component.html',
  styleUrls: ['./home-container.component.css']
})
export class HomeContainerComponent implements OnInit {

  public title: string = "Frontend Psicologo"
  public subtitle: string = "Progetto Universitario - Manlio Puglisi"

  public isJwtPresent: boolean;

  constructor(private jwt: JwtService) { }

  ngOnInit(): void {
    this.isJwtPresent = this.jwt.isJWTPresent();
    this.jwt.watchJWT().subscribe((isStored: boolean) => { this.isJwtPresent = isStored });
    if (this.isJwtPresent) { this.jwt.isJWTvalid(); }
  }

}
