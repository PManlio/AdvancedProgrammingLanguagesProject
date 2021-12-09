import { Component, OnInit } from '@angular/core';
import { JwtService } from './services/jwt.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  public title = 'frontend-apl';

  public isPresentJwt: boolean = this.jwt.isJWTPresent();

  constructor(private jwt: JwtService) { }

  ngOnInit(): void {
    // console.log(localStorage.getItem('token'))
    this.jwt.watchJWT().subscribe(
      (isStored: boolean) => {
        this.isPresentJwt = isStored;
      }
    );

    if (this.isPresentJwt) {
      this.jwt.isJWTvalid();
    }
  }

}
