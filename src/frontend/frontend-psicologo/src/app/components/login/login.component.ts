import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { LoginService } from 'src/app/services/login.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(private login: LoginService) { }

  ngOnInit(): void {
  }

  onSubmit(f: NgForm) {
    // console.log(f.value);  // { first: '', last: '' }
    // console.log(f.value.email) // o f.value.password
    // console.log(f.value)
    // console.log(f.valid);  // false
    if(f.valid) {
      this.login.authenticate(f.value.email, f.value.password)
    } else {
      alert("email o password mancanti")
    }
  }

}
