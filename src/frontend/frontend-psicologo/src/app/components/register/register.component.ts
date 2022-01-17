import { Component, OnInit } from '@angular/core';
import { NgForm } from '@angular/forms';
import { RegistrationInterface } from 'src/app/interfaces/registration-interface';
import { RegisterService } from 'src/app/services/register.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  constructor(private register: RegisterService) { }

  ngOnInit(): void {
  }

  onSubmit(f: NgForm) {
    // console.log(f.value);  // { first: '', last: '' }
    // console.log(f.value.email) // o f.value.password
    // console.log(f.value.password, f.value.ripetiPassword)
    
    if((f.value.password == f.value.ripetiPassword) && f.valid) {
      // console.log(f.value.codFisc, f.value.password)
      // console.log(f.value)
      // console.log(f.valid);  // false se il form è valido
      let body: RegistrationInterface = JSON.parse(JSON.stringify(f.value));
      // console.log(body)
      this.register.register(body)
    } else {
      if(f.value.password != f.value.ripetiPassword){ alert("ripetizione password errata"); }
      else { alert("devi riempire tutti i campi"); }
    }
  }
}
