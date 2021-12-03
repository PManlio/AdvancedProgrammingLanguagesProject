import { Component, Input, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';

import { LoginService } from 'src/app/services/login.service';

@Component({
  selector: 'app-main-login',
  templateUrl: './main-login.component.html',
  styleUrls: ['./main-login.component.css']
})
export class MainLoginComponent implements OnInit {

  @Input() formValues: NgForm;

  constructor(private loginService: LoginService) { }

  ngOnInit(): void {
  }

  public login() {
    console.log(this.formValues.value.email, this.formValues.value.password);
    
    this.loginService.authenticate(this.formValues.value.email, this.formValues.value.password);
  }

}
