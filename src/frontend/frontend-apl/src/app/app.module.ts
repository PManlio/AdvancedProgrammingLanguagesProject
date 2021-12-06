import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { MainDivComponent } from './components/mainTitle/main-div/main-div.component';
import { MainLoginComponent } from './components/mainTitle/main-div/main-login/main-login.component';
import { MainRegisterComponent } from './components/mainTitle/main-div/main-register/main-register.component';
import { LogoutbtnComponent } from './components/logoutbtn/logoutbtn.component';
import { RegistrationModalComponent } from './components/mainTitle/main-div/registration-modal/registration-modal.component';

@NgModule({
  declarations: [
    AppComponent,
    MainDivComponent,
    MainLoginComponent,
    MainRegisterComponent,
    LogoutbtnComponent,
    RegistrationModalComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
